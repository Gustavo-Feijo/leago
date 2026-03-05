package status

import "time"

type (
	// Status of the LoL services.
	ServiceStatus struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Locales      []string `json:"locales"`
		Maintenances []Status `json:"maintenances"`
		Incidents    []Status `json:"incidents"`
	}

	Status struct {
		ID                int               `json:"id"`
		MaintenanceStatus MaintenanceStatus `json:"maintenance_status,omitempty"`
		IncidentSeverity  IncidentSeverity  `json:"incident_severity,omitempty"`
		Titles            []Content         `json:"titles"`
		Updates           []Update          `json:"updates"`
		CreatedAt         time.Time         `json:"created_at"`
		ArchiveAt         *time.Time        `json:"archive_at,omitempty"`
		UpdatedAt         time.Time         `json:"updated_at"`
		Platforms         []Platform        `json:"platforms"`
	}

	Content struct {
		Locale  string `json:"locale"`
		Content string `json:"content"`
	}

	Update struct {
		ID               int               `json:"id"`
		Author           string            `json:"author"`
		Publish          bool              `json:"publish"`
		PublishLocations []PublishLocation `json:"publish_locations"`
		Translations     []Content         `json:"translations"`
		CreatedAt        time.Time         `json:"created_at"`
		UpdatedAt        time.Time         `json:"updated_at"`
	}
)

type (
	PublishLocation   string
	MaintenanceStatus string
	IncidentSeverity  string
	Platform          string
)

const (
	MaintenanceSchedule   MaintenanceStatus = "scheduled"
	MaintenanceInProgress MaintenanceStatus = "in_progress"
	MaintenanceComplete   MaintenanceStatus = "complete"

	IncidentInfo     IncidentSeverity = "info"
	IncidentWarning  IncidentSeverity = "warning"
	IncidentCritical IncidentSeverity = "critical"

	PlatformWindows Platform = "windows"
	PlatformMacOS   Platform = "macos"
	PlatformAndroid Platform = "android"
	PlatformIOS     Platform = "ios"
	PlatformPS4     Platform = "ps4"
	PlatformXBONE   Platform = "xbone"
	PlatformSwitch  Platform = "switch"

	PublishRiotClient PublishLocation = "riotclient"
	PublishRiotStatus PublishLocation = "riotstatus"
	PublishGame       PublishLocation = "game"
)
