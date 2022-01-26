package agent

import (
	"fmt"
	"os"
	"time"

	"github.com/Axway/agent-sdk/pkg/agent/cache"

	"github.com/Axway/agent-sdk/pkg/apic"
	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

// QA EnvVars
const qaTeamCacheInterval = "QA_CENTRAL_TEAMCACHE_INTERVAL"

type centralTeamsCache struct {
	jobs.Job
	teamChannel chan string
	cache       cache.Manager
	client      apic.Client
}

func (j *centralTeamsCache) Ready() bool {
	return true
}

func (j *centralTeamsCache) Status() error {
	return nil
}

func (j *centralTeamsCache) Execute() error {
	platformTeams, err := j.client.GetTeam(map[string]string{})
	if err != nil {
		return err
	}

	if len(platformTeams) == 0 {
		return fmt.Errorf("error: no teams returned from central")
	}

	for _, team := range platformTeams {
		savedTeam := j.cache.GetTeamById(team.ID)
		if savedTeam == nil {
			j.cache.AddTeam(&team)
			log.Tracef("sending %s (%s) team to acl", savedTeam.Name, savedTeam.ID)
			j.teamChannel <- savedTeam.ID
		}
	}

	return nil
}

// registerTeamMapCacheJob -
func registerTeamMapCacheJob(teamChannel chan string, cache cache.Manager, client apic.Client) {
	job := &centralTeamsCache{
		teamChannel: teamChannel,
		cache:       cache,
		client:      client,
	}
	// execute the job on startup to populate the team cache
	job.Execute()
	jobs.RegisterIntervalJobWithName(job, getJobInterval(), "Team Cache")
}

func getJobInterval() time.Duration {
	interval := time.Hour
	// check for QA env vars
	if val := os.Getenv(qaTeamCacheInterval); val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			log.Tracef("Using %s (%s) rather than the default (%s) for non-QA", qaTeamCacheInterval, val, time.Hour)
			interval = duration
		} else {
			log.Tracef("Could not use %s (%s) it is not a proper duration", qaTeamCacheInterval, val)
		}
	}

	return interval
}

func GetTeamFromCache(teamName string) (string, bool) {
	id, found := agent.teamMap.Get(teamName)
	if teamName == "" {
		// get the default team
		id, found = agent.teamMap.GetBySecondaryKey(apic.DefaultTeamKey)
	}
	if found != nil {
		return "", false
	}
	return id.(string), true
}
