package events

import (
	"bytes"
	_ "embed" // load of the watch topic template
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/Axway/agent-sdk/pkg/agent/resource"
	"github.com/Axway/agent-sdk/pkg/config"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	cv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
	mv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

//go:embed assets/watch-topic-template.json
var agentTemplate string

var agentTypesMap = map[config.AgentType]string{
	config.DiscoveryAgent:    "discoveryagents",
	config.TraceabilityAgent: "traceabilityagents",
	config.GovernanceAgent:   "governanceagents",
}

type watchTopicFeatures interface {
	IsMarketplaceSubsEnabled() bool
	GetAgentType() config.AgentType
}

const (
	desc = "Watch Topic used by a %s agent for resources in the %s environment."
	// WatchTopicFilterTypeCreated filter type name
	WatchTopicFilterTypeCreated = "created"
	// WatchTopicFilterTypeUpdated filter type name
	WatchTopicFilterTypeUpdated = "updated"
	// WatchTopicFilterTypeDeleted filter type name
	WatchTopicFilterTypeDeleted = "deleted"
)

var (
	created          = []string{WatchTopicFilterTypeCreated}
	updated          = []string{WatchTopicFilterTypeUpdated}
	deleted          = []string{WatchTopicFilterTypeDeleted}
	createdOrUpdated = append(created, updated...)
	all              = append(createdOrUpdated, deleted...)
)

// getOrCreateWatchTopic attempts to retrieve a watch topic from central, or create one if it does not exist.
func getOrCreateWatchTopic(name, scope string, client APIClient, features watchTopicFeatures) (*mv1.WatchTopic, error) {
	wt := mv1.NewWatchTopic("")
	ri, err := client.GetResource(fmt.Sprintf("%s/%s", wt.GetKindLink(), name))

	if err == nil {
		err = wt.FromInstance(ri)
		if err != nil {
			return nil, err
		}
	}

	var agentResourceGroupKind v1.GroupKind
	var tmplValuesFunc func(string, string, v1.GroupKind, watchTopicFeatures) WatchTopicValues

	switch features.GetAgentType() {
	case config.DiscoveryAgent:
		agentResourceGroupKind = mv1.DiscoveryAgentGVK().GroupKind
		tmplValuesFunc = NewDiscoveryWatchTopic
	case config.TraceabilityAgent:
		agentResourceGroupKind = mv1.TraceabilityAgentGVK().GroupKind
		tmplValuesFunc = NewTraceWatchTopic
	case config.GovernanceAgent:
		agentResourceGroupKind = mv1.GovernanceAgentGVK().GroupKind
		tmplValuesFunc = NewGovernanceAgentWatchTopic
	default:
		return nil, resource.ErrUnsupportedAgentType
	}

	newWT, err := parseWatchTopicTemplate(tmplValuesFunc(name, scope, agentResourceGroupKind, features))
	if err != nil {
		return nil, err
	}

	// if the existing wt has no name then it does not exist yet
	if wt.Name == "" {
		return createOrUpdateWatchTopic(newWT, client)
	}

	// compare the generated WT and the existing WT for changes
	if shouldPushUpdate(wt, newWT) {
		// update the spec in the existing watch topic
		wt.Spec = newWT.Spec
		return createOrUpdateWatchTopic(wt, client)
	}

	return wt, nil
}

func shouldPushUpdate(cur, new *mv1.WatchTopic) bool {
	filtersDiff := func(a, b []mv1.WatchTopicSpecFilters) bool {
		for _, aFilter := range a {
			found := false
			for _, bFilter := range b {
				if filtersEqual(aFilter, bFilter) {
					found = true
					break
				}
			}
			if !found {
				// update required
				return true
			}
		}
		return false
	}

	if filtersDiff(cur.Spec.Filters, new.Spec.Filters) {
		return true
	}
	return filtersDiff(new.Spec.Filters, cur.Spec.Filters)
}

func filtersEqual(a, b mv1.WatchTopicSpecFilters) (equal bool) {
	if a.Group != b.Group ||
		a.Kind != b.Kind ||
		a.Name != b.Name ||
		a.Scope == nil && b.Scope != nil ||
		a.Scope != nil && b.Scope == nil {
		return
	}

	if a.Scope != nil && b.Scope != nil {
		if a.Scope.Kind != b.Scope.Kind ||
			a.Scope.Name != b.Scope.Name {
			return
		}
	}

	typesDiff := func(aTypes, bTypes []string) bool {
		for _, aType := range aTypes {
			found := false
			for _, bType := range bTypes {
				if aType == bType {
					found = true
					break
				}
			}
			if !found {
				return true
			}
		}
		return false
	}

	if typesDiff(a.Type, b.Type) {
		return false
	}
	return !typesDiff(b.Type, a.Type)
}

// executeTemplate parses a WatchTopic from a template
func parseWatchTopicTemplate(values WatchTopicValues) (*mv1.WatchTopic, error) {
	tmpl, err := template.New("watch-topic-tmpl").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(agentTemplate)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, values)
	if err != nil {
		return nil, err
	}

	wt := mv1.NewWatchTopic("")
	err = json.Unmarshal(buf.Bytes(), wt)
	return wt, err
}

// createOrUpdateWatchTopic creates a WatchTopic
func createOrUpdateWatchTopic(wt *mv1.WatchTopic, rc APIClient) (*mv1.WatchTopic, error) {
	if wt.Metadata.ID != "" {
		err := rc.DeleteResourceInstance(wt)
		if err != nil {
			return nil, err
		}
	}

	ri, err := rc.CreateResourceInstance(wt)
	if err != nil {
		return nil, err
	}

	err = wt.FromInstance(ri)

	return wt, err
}

type kindValues struct {
	v1.GroupKind
	EventTypes []string
	ScopeKind  string // blank defaults to Environment in template
	ScopeName  string // blank generates no scope in template
}

// WatchTopicValues values to populate the watch topic template
type WatchTopicValues struct {
	Name        string
	Title       string
	Description string
	Kinds       []kindValues
}

// NewDiscoveryWatchTopic creates a WatchTopic template string.
// Using a template instead of unmarshalling into a struct to avoid sending a request with empty fields
func NewDiscoveryWatchTopic(name, scope string, agentResourceGroupKind v1.GroupKind, features watchTopicFeatures) WatchTopicValues {
	kinds := []kindValues{
		{GroupKind: agentResourceGroupKind, ScopeName: scope, EventTypes: updated},
		{GroupKind: cv1.CategoryGVK().GroupKind, EventTypes: all},
		{GroupKind: mv1.APIServiceGVK().GroupKind, ScopeName: scope, EventTypes: all},
		{GroupKind: mv1.APIServiceInstanceGVK().GroupKind, ScopeName: scope, EventTypes: all},
		{GroupKind: mv1.AccessControlListGVK().GroupKind, ScopeName: scope, EventTypes: all},
	}
	if features.IsMarketplaceSubsEnabled() {
		kinds = append(kinds, []kindValues{
			{GroupKind: mv1.CredentialGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.AccessRequestGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.ManagedApplicationGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.CredentialRequestDefinitionGVK().GroupKind, ScopeName: scope, EventTypes: all},
			{GroupKind: mv1.AccessRequestDefinitionGVK().GroupKind, ScopeName: scope, EventTypes: all},
		}...)
	}
	return WatchTopicValues{
		Name:        name,
		Title:       name,
		Description: fmt.Sprintf(desc, "discovery", scope),
		Kinds:       kinds,
	}
}

// NewTraceWatchTopic creates a WatchTopic template string
func NewTraceWatchTopic(name, scope string, agentResourceGroupKind v1.GroupKind, features watchTopicFeatures) WatchTopicValues {
	kinds := []kindValues{
		{GroupKind: agentResourceGroupKind, ScopeName: scope, EventTypes: updated},
		{GroupKind: mv1.APIServiceGVK().GroupKind, ScopeName: scope, EventTypes: all},
		{GroupKind: mv1.APIServiceInstanceGVK().GroupKind, ScopeName: scope, EventTypes: all},
	}
	if features.IsMarketplaceSubsEnabled() {
		kinds = append(kinds, []kindValues{
			{GroupKind: mv1.AccessRequestGVK().GroupKind, ScopeName: scope, EventTypes: all},
			{GroupKind: mv1.ManagedApplicationGVK().GroupKind, ScopeName: scope, EventTypes: all},
		}...)
	}
	return WatchTopicValues{
		Name:        name,
		Title:       name,
		Description: fmt.Sprintf(desc, "traceability", scope),
		Kinds:       kinds,
	}
}

// NewGovernanceAgentWatchTopic creates a WatchTopic template string
func NewGovernanceAgentWatchTopic(name, scope string, agentResourceGroupKind v1.GroupKind, features watchTopicFeatures) WatchTopicValues {
	kinds := []kindValues{
		{GroupKind: mv1.APIServiceGVK().GroupKind, ScopeName: scope, EventTypes: all},
		{GroupKind: mv1.APIServiceInstanceGVK().GroupKind, ScopeName: scope, EventTypes: all},
		{GroupKind: agentResourceGroupKind, ScopeName: scope, EventTypes: updated},
		{GroupKind: mv1.AmplifyRuntimeConfigGVK().GroupKind, ScopeName: scope, EventTypes: all},
	}
	if features.IsMarketplaceSubsEnabled() {
		kinds = append(kinds, []kindValues{
			{GroupKind: mv1.CredentialGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.AccessRequestGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.ManagedApplicationGVK().GroupKind, ScopeName: scope, EventTypes: createdOrUpdated},
			{GroupKind: mv1.CredentialRequestDefinitionGVK().GroupKind, ScopeName: scope, EventTypes: all},
			{GroupKind: mv1.AccessRequestDefinitionGVK().GroupKind, ScopeName: scope, EventTypes: all},
		}...)
	}
	return WatchTopicValues{
		Name:        name,
		Title:       name,
		Description: fmt.Sprintf(desc, "governance", scope),
		Kinds:       kinds,
	}
}

// GetWatchTopic retrieves a watch topic based on the agent config. Creates a watch topic if one does not exist.
func GetWatchTopic(cfg config.CentralConfig, client APIClient) (*mv1.WatchTopic, error) {
	env := cfg.GetEnvironmentName()

	wtName := getWatchTopicName(env, cfg.GetAgentType())
	wt, err := getOrCreateWatchTopic(wtName, env, client, cfg)
	if err != nil {
		return nil, err
	}

	return wt, err
}

func getWatchTopicName(envName string, agentType config.AgentType) string {
	return envName + getWatchTopicNameSuffix(agentType)
}

func getWatchTopicNameSuffix(agentType config.AgentType) string {
	return "-" + agentTypesMap[agentType]
}
