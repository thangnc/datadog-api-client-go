package datadogV2

import "time"

type LogRestrictionQueryAttributes struct {
	// Creation time of the role.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Time of last role modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the role. The name is neither unique nor a stable identifier of the role.
	RestrictionQuery *string `json:"restriction_query,omitempty"`
}
