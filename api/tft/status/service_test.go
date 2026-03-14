package status

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedStatus = ServiceStatus{
		ID:   "BR1",
		Name: "Brazil",
		Locales: []string{
			"pt_BR",
		},
		Maintenances: []Status{},
		Incidents: []Status{
			{
				ID:                1,
				MaintenanceStatus: "",
				IncidentSeverity:  "info",
				Titles: []Content{
					{
						Locale:  "pt_BR",
						Content: "Test",
					},
				},
				Updates: []Update{
					{
						ID:      15378,
						Author:  "Riot Games",
						Publish: true,
						PublishLocations: []PublishLocation{
							"game",
							"riotstatus",
						},
						Translations: []Content{
							{
								Locale:  "pt_BR",
								Content: "Test",
							},
						},
						CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				ArchiveAt: nil,
				Platforms: []Platform{
					"macos",
					"windows",
				},
			},
		},
	}

	statusJSON = `{
    "id": "BR1",
    "name": "Brazil",
    "locales": [
        "pt_BR"
    ],
    "maintenances": [],
    "incidents": [
        {
            "id": 1,
            "created_at": "2026-01-01T00:00:00.000000+00:00",
            "updated_at": "2026-01-01T00:00:00.000000+00:00",
            "archive_at": null,
            "titles": [
                {
                    "locale": "pt_BR",
                    "content": "Test"
                }
            ],
            "updates": [
                {
                    "id": 15378,
                    "created_at": "2026-01-01T00:00:00.000000+00:00",
                    "updated_at": "2026-01-01T00:00:00.000000+00:00",
                    "publish": true,
                    "author": "Riot Games",
                    "translations": [
                        {
                            "locale": "pt_BR",
                            "content": "Test"
                        }
                    ],
                    "publish_locations": [
                        "game",
                        "riotstatus"
                    ]
                }
            ],
            "platforms": [
                "macos",
                "windows"
            ],
            "maintenance_status": null,
            "incident_severity": "info"
        }
    ]
}`
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.PlatformBR1), "apiKey", internal.WithHTTP(mockDoer))

	return NewPlatformClient(baseClient), mockDoer
}

func TestGetStatus(t *testing.T) {
	tests := []struct {
		name string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult ServiceStatus

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			statusCode:   http.StatusOK,
			responseBody: statusJSON,

			expectedPath: "/tft/status/v1/platform-data",

			expectedResult: expectedStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetStatus(context.Background())

			if tt.wantErr {
				require.Error(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// Avoiding direct comparison of the returned due to timezones handling.
			// Marshal first so both are plain jsons.
			expectedJSON, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)
			assert.Equal(t, expectedJSON, jsonResp)
		})
	}
}
