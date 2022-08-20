package datadogV2

import (
	"encoding/json"
	"fmt"
)

type LogRestrictionQueryType string

const (
	LOG_RESTRICTION_QUERY_TYPE_QUERIES LogRestrictionQueryType = "logs_restriction_queries"
)

var allowedLogRestrictionQueryTypeEnumValues = []LogRestrictionQueryType{
	LOG_RESTRICTION_QUERY_TYPE_QUERIES,
}

func (v *LogRestrictionQueryType) GetAllowedValues() []LogRestrictionQueryType {
	return allowedLogRestrictionQueryTypeEnumValues
}

func (v *LogRestrictionQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogRestrictionQueryType(value)
	return nil
}

func NewLogRestrictionQueryTypeFromValue(v string) (*LogRestrictionQueryType, error) {
	ev := LogRestrictionQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogRestrictionQueryType: valid values are %v", v, allowedLogRestrictionQueryTypeEnumValues)
}

func (v LogRestrictionQueryType) IsValid() bool {
	for _, existing := range allowedLogRestrictionQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v LogRestrictionQueryType) Ptr() *LogRestrictionQueryType {
	return &v
}

type NullableLogRestrictionQueryType struct {
	value *LogRestrictionQueryType
	isSet bool
}

func (v NullableLogRestrictionQueryType) Get() *LogRestrictionQueryType {
	return v.value
}

func (v *NullableLogRestrictionQueryType) Set(val *LogRestrictionQueryType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogRestrictionQueryType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogRestrictionQueryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogRestrictionQueryType(val *LogRestrictionQueryType) *NullableLogRestrictionQueryType {
	return &NullableLogRestrictionQueryType{value: val, isSet: true}
}

func (v NullableLogRestrictionQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogRestrictionQueryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
