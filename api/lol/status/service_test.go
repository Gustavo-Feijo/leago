package status

import (
	"context"
	"encoding/json"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	statusJson = `{
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
)

func newTestPlatformClient(statusCode int, responseBody string, httpErr error) *PlatformClient {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody, httpErr)
	baseClient := internal.NewHttpClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient)
}

func TestGetStatus(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult ServiceStatus
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "invalid json",
			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   statusJson,
			expectedResult: expectedStatus,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := newTestPlatformClient(tt.statusCode, tt.responseBody, tt.httpErr)
			resp, err := pc.GetStatus(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}

				return
			}

			require.Nil(t, err)
			require.NotNil(t, resp)

			// Avoiding direct comparison of the returned due to timezones handling.
			// Marshal first so both are plain jsons.
			expectedJson, _ := json.Marshal(tt.expectedResult)
			jsonResp, _ := json.Marshal(resp)

			assert.Equal(t, expectedJson, jsonResp)
		})
	}
}
