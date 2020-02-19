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

// ServiceLevelObjectiveListResponse struct for ServiceLevelObjectiveListResponse
type ServiceLevelObjectiveListResponse struct {
	// An array of service level objective objects.
	Data []ServiceLevelObjective `json:"data"`
	// An array of error messages. Each endpoint documents how/whether this field is used.
	Errors *[]string `json:"errors,omitempty"`
}

// NewServiceLevelObjectiveListResponse instantiates a new ServiceLevelObjectiveListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectiveListResponse(data []ServiceLevelObjective) *ServiceLevelObjectiveListResponse {
	this := ServiceLevelObjectiveListResponse{}
	this.Data = data
	return &this
}

// NewServiceLevelObjectiveListResponseWithDefaults instantiates a new ServiceLevelObjectiveListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectiveListResponseWithDefaults() *ServiceLevelObjectiveListResponse {
	this := ServiceLevelObjectiveListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ServiceLevelObjectiveListResponse) GetData() []ServiceLevelObjective {
	if o == nil {
		var ret []ServiceLevelObjective
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *ServiceLevelObjectiveListResponse) SetData(v []ServiceLevelObjective) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ServiceLevelObjectiveListResponse) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveListResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ServiceLevelObjectiveListResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ServiceLevelObjectiveListResponse) SetErrors(v []string) {
	o.Errors = &v
}

func (o ServiceLevelObjectiveListResponse) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLevelObjectiveListResponse struct {
	value *ServiceLevelObjectiveListResponse
	isSet bool
}

func (v NullableServiceLevelObjectiveListResponse) Get() *ServiceLevelObjectiveListResponse {
	return v.value
}

func (v NullableServiceLevelObjectiveListResponse) Set(val *ServiceLevelObjectiveListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectiveListResponse) IsSet() bool {
	return v.isSet
}

func (v NullableServiceLevelObjectiveListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectiveListResponse(val *ServiceLevelObjectiveListResponse) *NullableServiceLevelObjectiveListResponse {
	return &NullableServiceLevelObjectiveListResponse{value: val, isSet: true}
}

func (v NullableServiceLevelObjectiveListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectiveListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}