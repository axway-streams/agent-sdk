package agent

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Axway/agent-sdk/pkg/agent/cache"

	"github.com/Axway/agent-sdk/pkg/apic/definitions"

	"github.com/stretchr/testify/assert"
)

func TestTeamCache(t *testing.T) {
	testCases := []struct {
		name                 string
		teams                []*definitions.PlatformTeam
		cached               []*definitions.PlatformTeam
		expectedTeamsInCache int
	}{
		{
			name:                 "Should save one team to the cache",
			expectedTeamsInCache: 1,
			cached:               []*definitions.PlatformTeam{},
			teams: []*definitions.PlatformTeam{
				{
					Name:    "TeamA",
					ID:      "1",
					Default: true,
				},
			},
		},
		{
			name:                 "Should save two teams to the cache, and remove a team that was added",
			expectedTeamsInCache: 2,
			cached: []*definitions.PlatformTeam{
				{
					Name:    "TeamA",
					ID:      "1",
					Default: true,
				},
			},
			teams: []*definitions.PlatformTeam{
				{
					Name:    "TeamB",
					ID:      "2",
					Default: false,
				},
				{
					Name:    "TeamC",
					ID:      "3",
					Default: false,
				},
			},
		},
		{
			name:                 "Should save 4 teams in the cache",
			expectedTeamsInCache: 4,
			cached: []*definitions.PlatformTeam{
				{
					Name:    "TeamA",
					ID:      "1",
					Default: true,
				},
				{
					Name:    "TeamB",
					ID:      "2",
					Default: false,
				},
				{
					Name:    "TeamC",
					ID:      "3",
					Default: false,
				},
			},
			teams: []*definitions.PlatformTeam{
				{
					Name:    "TeamA",
					ID:      "1",
					Default: true,
				},
				{
					Name:    "TeamB",
					ID:      "2",
					Default: false,
				},
				{
					Name:    "TeamC",
					ID:      "3",
					Default: false,
				},
				{
					Name:    "TeamD",
					ID:      "4",
					Default: false,
				},
			},
		},
	}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				switch {
				case strings.Contains(req.RequestURI, "/auth"):
					token := "{\"access_token\":\"somevalue\",\"expires_in\": 12235677}"
					resp.Write([]byte(token))
				case strings.Contains(req.RequestURI, "platformTeams"):
					data, _ := json.Marshal(test.teams)
					resp.Write(data)
				}
			}))
			defer s.Close()

			cfg := createCentralCfg(s.URL, "env")
			caches := cache.NewAgentCacheManager(cfg, false)

			for _, item := range test.cached {
				caches.AddTeam(item)
			}

			resetResources()
			agent.teamMap = nil
			err := Initialize(cfg)
			assert.Nil(t, err)
			assert.NotNil(t, agent)
			assert.NotNil(t, agent.apicClient)

			job := centralTeamsCache{}

			job.Execute()
			teams := agent.cacheManager.GetTeamCache().GetKeys()
			assert.Equal(t, test.expectedTeamsInCache, len(teams))
		})
	}
}
