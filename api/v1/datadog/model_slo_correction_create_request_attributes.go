/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SLOCorrectionCreateRequestAttributes The attribute object associated with the SLO correction to be created
type SLOCorrectionCreateRequestAttributes struct {
	Category SLOCorrectionCategory `json:"category"`
	// Description of the correction being made.
	Description *string `json:"description,omitempty"`
	// Ending time of the correction in epoch seconds
	End int64 `json:"end"`
	// ID of the SLO that this correction will be applied to
	SloId string `json:"slo_id"`
	// Starting time of the correction in epoch seconds
	Start int64 `json:"start"`
	// The timezone to display in the UI for the correction times (defaults to \"UTC\")
	Timezone *string `json:"timezone,omitempty"`
}

// NewSLOCorrectionCreateRequestAttributes instantiates a new SLOCorrectionCreateRequestAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOCorrectionCreateRequestAttributes(category SLOCorrectionCategory, end int64, sloId string, start int64) *SLOCorrectionCreateRequestAttributes {
	this := SLOCorrectionCreateRequestAttributes{}
	this.Category = category
	this.End = end
	this.SloId = sloId
	this.Start = start
	return &this
}

// NewSLOCorrectionCreateRequestAttributesWithDefaults instantiates a new SLOCorrectionCreateRequestAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOCorrectionCreateRequestAttributesWithDefaults() *SLOCorrectionCreateRequestAttributes {
	this := SLOCorrectionCreateRequestAttributes{}
	return &this
}

// GetCategory returns the Category field value
func (o *SLOCorrectionCreateRequestAttributes) GetCategory() SLOCorrectionCategory {
	if o == nil {
		var ret SLOCorrectionCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetCategoryOk() (*SLOCorrectionCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *SLOCorrectionCreateRequestAttributes) SetCategory(v SLOCorrectionCategory) {
	o.Category = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SLOCorrectionCreateRequestAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SLOCorrectionCreateRequestAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SLOCorrectionCreateRequestAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnd returns the End field value
func (o *SLOCorrectionCreateRequestAttributes) GetEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *SLOCorrectionCreateRequestAttributes) SetEnd(v int64) {
	o.End = v
}

// GetSloId returns the SloId field value
func (o *SLOCorrectionCreateRequestAttributes) GetSloId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetSloIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloId, true
}

// SetSloId sets field value
func (o *SLOCorrectionCreateRequestAttributes) SetSloId(v string) {
	o.SloId = v
}

// GetStart returns the Start field value
func (o *SLOCorrectionCreateRequestAttributes) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *SLOCorrectionCreateRequestAttributes) SetStart(v int64) {
	o.Start = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SLOCorrectionCreateRequestAttributes) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequestAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SLOCorrectionCreateRequestAttributes) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SLOCorrectionCreateRequestAttributes) SetTimezone(v string) {
	o.Timezone = &v
}

func (o SLOCorrectionCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["slo_id"] = o.SloId
	}
	if true {
		toSerialize["start"] = o.Start
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

func (o *SLOCorrectionCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Category *SLOCorrectionCategory `json:"category"`
		End      *int64                 `json:"end"`
		SloId    *string                `json:"slo_id"`
		Start    *int64                 `json:"start"`
	}{}
	all := struct {
		Category    SLOCorrectionCategory `json:"category"`
		Description *string               `json:"description,omitempty"`
		End         int64                 `json:"end"`
		SloId       string                `json:"slo_id"`
		Start       int64                 `json:"start"`
		Timezone    *string               `json:"timezone,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Category == nil {
		return fmt.Errorf("Required field category missing")
	}
	if required.End == nil {
		return fmt.Errorf("Required field end missing")
	}
	if required.SloId == nil {
		return fmt.Errorf("Required field slo_id missing")
	}
	if required.Start == nil {
		return fmt.Errorf("Required field start missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Category = all.Category
	o.Description = all.Description
	o.End = all.End
	o.SloId = all.SloId
	o.Start = all.Start
	o.Timezone = all.Timezone
	return nil
}

type NullableSLOCorrectionCreateRequestAttributes struct {
	value *SLOCorrectionCreateRequestAttributes
	isSet bool
}

func (v NullableSLOCorrectionCreateRequestAttributes) Get() *SLOCorrectionCreateRequestAttributes {
	return v.value
}

func (v *NullableSLOCorrectionCreateRequestAttributes) Set(val *SLOCorrectionCreateRequestAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOCorrectionCreateRequestAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOCorrectionCreateRequestAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOCorrectionCreateRequestAttributes(val *SLOCorrectionCreateRequestAttributes) *NullableSLOCorrectionCreateRequestAttributes {
	return &NullableSLOCorrectionCreateRequestAttributes{value: val, isSet: true}
}

func (v NullableSLOCorrectionCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOCorrectionCreateRequestAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
