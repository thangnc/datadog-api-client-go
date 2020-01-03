/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// MonitorState struct for MonitorState
type MonitorState struct {
	Groups       *map[string]MonitorStateGroup `json:"groups,omitempty"`
	MonitorId    *int64                        `json:"monitor_id,omitempty"`
	OverallState *MonitorOverallStates         `json:"overall_state,omitempty"`
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MonitorState) GetGroups() map[string]MonitorStateGroup {
	if o == nil || o.Groups == nil {
		var ret map[string]MonitorStateGroup
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorState) GetGroupsOk() (map[string]MonitorStateGroup, bool) {
	if o == nil || o.Groups == nil {
		var ret map[string]MonitorStateGroup
		return ret, false
	}
	return *o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MonitorState) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]MonitorStateGroup and assigns it to the Groups field.
func (o *MonitorState) SetGroups(v map[string]MonitorStateGroup) {
	o.Groups = &v
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *MonitorState) GetMonitorId() int64 {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorState) GetMonitorIdOk() (int64, bool) {
	if o == nil || o.MonitorId == nil {
		var ret int64
		return ret, false
	}
	return *o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *MonitorState) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given int64 and assigns it to the MonitorId field.
func (o *MonitorState) SetMonitorId(v int64) {
	o.MonitorId = &v
}

// GetOverallState returns the OverallState field value if set, zero value otherwise.
func (o *MonitorState) GetOverallState() MonitorOverallStates {
	if o == nil || o.OverallState == nil {
		var ret MonitorOverallStates
		return ret
	}
	return *o.OverallState
}

// GetOverallStateOk returns a tuple with the OverallState field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorState) GetOverallStateOk() (MonitorOverallStates, bool) {
	if o == nil || o.OverallState == nil {
		var ret MonitorOverallStates
		return ret, false
	}
	return *o.OverallState, true
}

// HasOverallState returns a boolean if a field has been set.
func (o *MonitorState) HasOverallState() bool {
	if o != nil && o.OverallState != nil {
		return true
	}

	return false
}

// SetOverallState gets a reference to the given MonitorOverallStates and assigns it to the OverallState field.
func (o *MonitorState) SetOverallState(v MonitorOverallStates) {
	o.OverallState = &v
}

type NullableMonitorState struct {
	Value        MonitorState
	ExplicitNull bool
}

func (v NullableMonitorState) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMonitorState) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
