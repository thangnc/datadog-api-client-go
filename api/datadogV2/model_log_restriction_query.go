package datadogV2

type LogRestrictionQuery struct {
	// Attributes of the role.
	Attributes *LogRestrictionQueryAttributes `json:"attributes,omitempty"`
	// The unique identifier of the role.
	Id *string `json:"id,omitempty"`
	// Relationships of the role object returned by the API.
	Relationships *RelationshipToRole `json:"relationships,omitempty"`
	// Roles type.
	Type LogRestrictionQueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}
