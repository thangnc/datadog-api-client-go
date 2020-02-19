/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// LogsDateRemapper As Datadog receives logs, it timestamps them using the value(s) from any of these default attributes:    - timestamp    - date    - _timestamp    - Timestamp    - eventTime    - published_date     If your logs put their dates in an attribute not in this list, use the log date Remapper Processor to define their date attribute as the official log timestamp.   The recognized date formats are: ISO8601, UNIX (the milliseconds EPOCH format), and RFC3164.    **Note:**      - If your logs don’t contain any of the default attributes and you haven’t defined your own date attribute, Datadog timestamps the logs with the date it received them.      - If multiple log date remapper processors can be applied to a given log, only the first one (according to the pipelines order) is taken into account.
type LogsDateRemapper struct {
	// Array of source attributes.
	Sources []string `json:"sources"`
	// Type of processor
	Type *string `json:"type,omitempty"`
	// Whether or not the processor is enabled
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor
	Name *string `json:"name,omitempty"`
}

// NewLogsDateRemapper instantiates a new LogsDateRemapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsDateRemapper(sources []string) *LogsDateRemapper {
	this := LogsDateRemapper{}
	this.Sources = sources
	var type_ string = "date-remapper"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// NewLogsDateRemapperWithDefaults instantiates a new LogsDateRemapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsDateRemapperWithDefaults() *LogsDateRemapper {
	this := LogsDateRemapper{}
	var type_ string = "date-remapper"
	this.Type = &type_
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	return &this
}

// GetSources returns the Sources field value
func (o *LogsDateRemapper) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sources
}

// SetSources sets field value
func (o *LogsDateRemapper) SetSources(v []string) {
	o.Sources = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogsDateRemapper) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogsDateRemapper) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogsDateRemapper) SetType(v string) {
	o.Type = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsDateRemapper) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetIsEnabledOk() (bool, bool) {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsDateRemapper) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsDateRemapper) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsDateRemapper) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsDateRemapper) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsDateRemapper) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsDateRemapper) SetName(v string) {
	o.Name = &v
}

func (o LogsDateRemapper) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

// AsLogsProcessor wraps this instance of LogsDateRemapper in LogsProcessor
func (s *LogsDateRemapper) AsLogsProcessor() LogsProcessor {
	return LogsProcessor{LogsProcessorInterface: s}
}

type NullableLogsDateRemapper struct {
	value *LogsDateRemapper
	isSet bool
}

func (v NullableLogsDateRemapper) Get() *LogsDateRemapper {
	return v.value
}

func (v NullableLogsDateRemapper) Set(val *LogsDateRemapper) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsDateRemapper) IsSet() bool {
	return v.isSet
}

func (v NullableLogsDateRemapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsDateRemapper(val *LogsDateRemapper) *NullableLogsDateRemapper {
	return &NullableLogsDateRemapper{value: val, isSet: true}
}

func (v NullableLogsDateRemapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsDateRemapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}