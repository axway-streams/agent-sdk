package events

import (
	"fmt"
	"testing"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	mv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestCreateWatchTopic(t *testing.T) {
	tests := []struct {
		name   string
		ri     *apiv1.ResourceInstance
		hasErr bool
		err    error
	}{
		{
			name:   "Should call create and return a WatchTopic",
			hasErr: false,
			err:    nil,
			ri: &apiv1.ResourceInstance{
				ResourceMeta: apiv1.ResourceMeta{
					Name: "wt-name",
				},
			},
		},
		{
			name:   "Should return an error when calling create",
			hasErr: true,
			err:    fmt.Errorf("error"),
			ri: &apiv1.ResourceInstance{
				ResourceMeta: apiv1.ResourceMeta{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rc := &mockAPIClient{
				ri:        tc.ri,
				createErr: tc.err,
			}

			wt := mv1.NewWatchTopic("")
			err := wt.FromInstance(tc.ri)
			assert.Nil(t, err)

			wt, err = createOrUpdateWatchTopic(wt, rc)
			if tc.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.ri.Name, wt.Name)
			}
		})
	}

}

type mockWatchTopicFeatures struct {
	isMPSEnabled bool
	agentType    config.AgentType
}

func (m *mockWatchTopicFeatures) IsMarketplaceSubsEnabled() bool {
	return m.isMPSEnabled
}

func (m *mockWatchTopicFeatures) GetAgentType() config.AgentType {
	return m.agentType
}

func Test_parseWatchTopic(t *testing.T) {
	tests := []struct {
		name         string
		isMPSEnabled bool
	}{
		{
			name: "Should create a watch topic without marketplace subs enabled",
		},
		{
			name:         "Should create a watch topic with marketplace subs enabled",
			isMPSEnabled: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			features := &mockWatchTopicFeatures{isMPSEnabled: tc.isMPSEnabled}

			wt, err := parseWatchTopicTemplate(NewDiscoveryWatchTopic("name", "scope", mv1.DiscoveryAgentGVK().GroupKind, features))
			assert.Nil(t, err)
			assert.NotNil(t, wt)

			wt, err = parseWatchTopicTemplate(NewTraceWatchTopic("name", "scope", mv1.TraceabilityAgentGVK().GroupKind, features))
			assert.Nil(t, err)
			assert.NotNil(t, wt)

			wt, err = parseWatchTopicTemplate(NewGovernanceAgentWatchTopic("name", "scope", mv1.GovernanceAgentGVK().GroupKind, features))
			assert.Nil(t, err)
			assert.NotNil(t, wt)
		})
	}
}

func TestGetOrCreateWatchTopic(t *testing.T) {
	tests := []struct {
		name      string
		client    *mockAPIClient
		hasErr    bool
		agentType config.AgentType
	}{
		{
			name:      "should retrieve a watch topic if it exists",
			hasErr:    false,
			agentType: config.DiscoveryAgent,
			client: &mockAPIClient{
				ri: &apiv1.ResourceInstance{
					ResourceMeta: apiv1.ResourceMeta{
						Name: "wt-name",
					},
				},
			},
		},
		{
			name:      "should create a watch topic for a trace agent if it does not exist",
			agentType: config.TraceabilityAgent,
			hasErr:    false,
			client: &mockAPIClient{
				getErr: fmt.Errorf("not found"),
				ri: &apiv1.ResourceInstance{
					ResourceMeta: apiv1.ResourceMeta{
						Name: "wt-name",
					},
				},
			},
		},
		{
			name:      "should create a watch topic for a discovery agent if it does not exist",
			agentType: config.DiscoveryAgent,
			hasErr:    false,
			client: &mockAPIClient{
				getErr: fmt.Errorf("not found"),
				ri: &apiv1.ResourceInstance{
					ResourceMeta: apiv1.ResourceMeta{
						Name: "wt-name",
					},
				},
			},
		},
		{
			name:      "should create a watch topic for a governance agent if it does not exist",
			agentType: config.GovernanceAgent,
			hasErr:    false,
			client: &mockAPIClient{
				getErr: fmt.Errorf("not found"),
				ri: &apiv1.ResourceInstance{
					ResourceMeta: apiv1.ResourceMeta{
						Name: "wt-name",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			name := "agent-name"
			features := &mockWatchTopicFeatures{isMPSEnabled: true, agentType: tc.agentType}

			wt, err := getOrCreateWatchTopic(name, "scope", tc.client, features)
			if tc.hasErr == true {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.client.ri.Name, wt.Name)
			}
		})
	}
}

type mockCacheGet struct {
	item interface{}
	err  error
}

func (m mockCacheGet) Get(_ string) (interface{}, error) {
	return m.item, m.err
}

func Test_shouldPushUpdate(t *testing.T) {
	type args struct {
		cur []mv1.WatchTopicSpecFilters
		new []mv1.WatchTopicSpecFilters
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should not push update",
			args: args{
				cur: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
				new: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
			},
			want: false,
		},
		{
			name: "should push update, second more",
			args: args{
				cur: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
				new: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind1",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
			},
			want: true,
		},
		{
			name: "should push update, first more",
			args: args{
				cur: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind1",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
				new: []mv1.WatchTopicSpecFilters{
					{
						Group: "group",
						Scope: nil,
						Kind:  "kind",
						Name:  "name",
						Type:  []string{"type1", "type2", "type3"},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createWatchTopic := func(filters []mv1.WatchTopicSpecFilters) *mv1.WatchTopic {
				wt := mv1.NewWatchTopic("")
				wt.Spec.Filters = filters
				return wt
			}

			if got := shouldPushUpdate(createWatchTopic(tt.args.cur), createWatchTopic(tt.args.new)); got != tt.want {
				t.Errorf("shouldPushUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filtersEqual(t *testing.T) {
	type args struct {
		a mv1.WatchTopicSpecFilters
		b mv1.WatchTopicSpecFilters
	}
	tests := []struct {
		name      string
		args      args
		wantEqual bool
	}{
		{
			name: "group diff",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group1",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "kind diff",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind1",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "name diff",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name1",
					Type:  []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff 1",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff 2",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: nil,
					Kind:  "kind",
					Name:  "name",
					Type:  []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff name",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name1",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff name",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind1",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff types 1",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2"},
				},
			},
			wantEqual: false,
		},
		{
			name: "scope diff types 2",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: false,
		},
		{
			name: "equal",
			args: args{
				a: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
				b: mv1.WatchTopicSpecFilters{
					Group: "group",
					Scope: &mv1.WatchTopicSpecScope{
						Kind: "kind",
						Name: "name",
					},
					Kind: "kind",
					Name: "name",
					Type: []string{"type1", "type2", "type3"},
				},
			},
			wantEqual: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEqual := filtersEqual(tt.args.a, tt.args.b); gotEqual != tt.wantEqual {
				t.Errorf("filtersEqual() = %v, want %v", gotEqual, tt.wantEqual)
			}
		})
	}
}

func Test_getWatchTopic(t *testing.T) {
	wt := &mv1.WatchTopic{}
	ri, _ := wt.AsInstance()
	httpClient := &mockAPIClient{
		ri: ri,
	}
	cfg := &config.CentralConfiguration{
		AgentType:     1,
		TenantID:      "12345",
		Environment:   "stream-test",
		EnvironmentID: "123",
		AgentName:     "discoveryagents",
		URL:           "http://abc.com",
		TLS:           &config.TLSConfiguration{},
	}

	wt, err := GetWatchTopic(cfg, httpClient)
	assert.NotNil(t, wt)
	assert.Nil(t, err)

	wt, err = GetWatchTopic(cfg, httpClient)
	assert.NotNil(t, wt)
	assert.Nil(t, err)
}

type mockAPIClient struct {
	ri        *apiv1.ResourceInstance
	getErr    error
	createErr error
	updateErr error
	deleteErr error
}

func (m mockAPIClient) GetResource(url string) (*apiv1.ResourceInstance, error) {
	return m.ri, m.getErr
}

func (m mockAPIClient) CreateResourceInstance(_ apiv1.Interface) (*apiv1.ResourceInstance, error) {
	return m.ri, m.createErr
}

func (m mockAPIClient) UpdateResourceInstance(_ apiv1.Interface) (*apiv1.ResourceInstance, error) {
	return m.ri, m.updateErr
}

func (m mockAPIClient) DeleteResourceInstance(_ apiv1.Interface) error {
	return m.deleteErr
}

func (m *mockAPIClient) GetAPIV1ResourceInstancesWithPageSize(_ map[string]string, _ string, _ int) ([]*apiv1.ResourceInstance, error) {
	return nil, nil
}
