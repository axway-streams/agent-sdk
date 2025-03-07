package agent

import (
	"context"
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/errors"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/google/uuid"
)

const (
	periodic  = "periodic status change"
	immediate = "immediate status change"
)

// This type is used for values added to context
type ctxKey int

// The key used for the logger in the context
const (
	ctxLogger ctxKey = iota
)

var previousStatus string // The global previous status to be used by both update jobs
var updateStatusMutex *sync.Mutex

func init() {
	updateStatusMutex = &sync.Mutex{}
}

type agentStatusUpdate struct {
	jobs.Job
	previousActivityTime  time.Time
	currentActivityTime   time.Time
	immediateStatusChange bool
	typeOfStatusUpdate    string
	logger                log.FieldLogger
}

var periodicStatusUpdate *agentStatusUpdate
var immediateStatusUpdate *agentStatusUpdate

func (su *agentStatusUpdate) Ready() bool {
	ctx := context.WithValue(context.Background(), ctxLogger, su.logger)
	// Do not start until status will be running
	status := su.getCombinedStatus(ctx)
	if status != AgentRunning && su.immediateStatusChange {
		return false
	}

	su.logger.Trace("Periodic status update is ready")
	su.currentActivityTime = time.Now()
	su.previousActivityTime = su.currentActivityTime
	return true
}

func (su *agentStatusUpdate) Status() error {
	return nil
}

func (su *agentStatusUpdate) Execute() error {
	id, _ := uuid.NewUUID()
	log := su.logger.WithField("status-update-id", id)

	ctx := context.WithValue(context.Background(), ctxLogger, log)
	// only one status update should execute at a time
	log.Tracef("get status update lock %s", su.typeOfStatusUpdate)
	updateStatusMutex.Lock()
	defer func() {
		log.Tracef("return status update lock %s", su.typeOfStatusUpdate)
		updateStatusMutex.Unlock()
	}()

	// get the status from the health check and jobs
	status := su.getCombinedStatus(ctx)
	log.Tracef("Type of agent status update being checked : %s ", su.typeOfStatusUpdate)

	if su.typeOfStatusUpdate == periodic {
		// always update on the periodic status update, even if the status has not changed
		log.
			WithField("previous-status", previousStatus).
			WithField("new-status", status).
			Debugf("%s -- Last activity updated", su.typeOfStatusUpdate)
		UpdateStatusWithContext(ctx, status, previousStatus, "")
		su.previousActivityTime = su.currentActivityTime
	} else if previousStatus != status {
		// if the status has changed then report that on the immediate check
		log.
			WithField("previous-status", previousStatus).
			WithField("new-status", status).
			Debug("status is changing")
		UpdateStatusWithContext(ctx, status, previousStatus, "")
		su.previousActivityTime = su.currentActivityTime
	}

	previousStatus = status
	return nil
}

// StartAgentStatusUpdate - starts 2 separate jobs that runs the periodic status updates and immediate status updates
func StartAgentStatusUpdate() {
	logger := log.NewFieldLogger().
		WithPackage("sdk.agent").
		WithComponent("agentStatusUpdate")
	if err := runStatusUpdateCheck(); err != nil {
		logger.WithError(err).Error("not starting status update jobs")
		return
	}
	startPeriodicStatusUpdate(logger)
	startImmediateStatusUpdate(logger)
}

// startPeriodicStatusUpdate - start periodic status updates based on report activity frequency config
func startPeriodicStatusUpdate(logger log.FieldLogger) {
	interval := agent.cfg.GetReportActivityFrequency()
	periodicStatusUpdate = &agentStatusUpdate{
		typeOfStatusUpdate: periodic,
		logger:             logger.WithField("status-check", periodic),
	}
	_, err := jobs.RegisterIntervalJobWithName(periodicStatusUpdate, interval, "Status Update")

	if err != nil {
		logger.Error(errors.ErrStartingAgentStatusUpdate.FormatError(periodic))
	}
}

// startImmediateStatusUpdate - start job that will 'immediately' update status.  NOTE : By 'immediately', this means currently 10 seconds.
// The time interval for this job is hard coded.
func startImmediateStatusUpdate(logger log.FieldLogger) {
	interval := 10 * time.Second
	immediateStatusUpdate = &agentStatusUpdate{
		immediateStatusChange: true,
		typeOfStatusUpdate:    immediate,
		logger:                logger.WithField("status-check", immediate),
	}
	_, err := jobs.RegisterDetachedIntervalJobWithName(immediateStatusUpdate, interval, "Immediate Status Update")

	if err != nil {
		logger.Error(errors.ErrStartingAgentStatusUpdate.FormatError(immediate))
	}
}

func (su *agentStatusUpdate) getCombinedStatus(ctx context.Context) string {
	log := ctx.Value(ctxLogger).(log.FieldLogger)
	status := su.getJobPoolStatus(ctx)
	hcStatus := su.getHealthcheckStatus(ctx)
	if hcStatus != AgentRunning {
		log.
			WithField("pool-status", status).
			WithField("healthcheck-status", hcStatus).
			Info("agent not in running status")
		status = hcStatus
	}
	return status
}

// getJobPoolStatus
func (su *agentStatusUpdate) getJobPoolStatus(ctx context.Context) string {
	log := ctx.Value(ctxLogger).(log.FieldLogger)
	status := jobs.GetStatus()
	log.Tracef("status from pool: %s", status)

	// update the status only if not running
	if status == jobs.PoolStatusStopped.String() {
		return AgentUnhealthy
	}
	return AgentRunning
}

// getHealthcheckStatus
func (su *agentStatusUpdate) getHealthcheckStatus(ctx context.Context) string {
	log := ctx.Value(ctxLogger).(log.FieldLogger)
	hcStatus := hc.GetGlobalStatus()
	log.Tracef("status from healthcheck: %s", hcStatus)

	// update the status only if not running
	if hcStatus == string(hc.FAIL) {
		return AgentUnhealthy
	}
	return AgentRunning
}

// runStatusUpdateCheck - returns an error if agent name is blank
func runStatusUpdateCheck() error {
	if agent.cfg.GetAgentName() == "" {
		return errors.ErrStartingAgentStatusUpdate.FormatError(periodic)
	}
	return nil
}
