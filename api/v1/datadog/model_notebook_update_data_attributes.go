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

// NotebookUpdateDataAttributes The data attributes of a notebook.
type NotebookUpdateDataAttributes struct {
	// List of cells to display in the notebook.
	Cells []NotebookUpdateCell `json:"cells"`
	// The name of the notebook.
	Name   string             `json:"name"`
	Status *NotebookStatus    `json:"status,omitempty"`
	Time   NotebookGlobalTime `json:"time"`
}

// NewNotebookUpdateDataAttributes instantiates a new NotebookUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookUpdateDataAttributes(cells []NotebookUpdateCell, name string, time NotebookGlobalTime) *NotebookUpdateDataAttributes {
	this := NotebookUpdateDataAttributes{}
	this.Cells = cells
	this.Name = name
	var status NotebookStatus = NOTEBOOKSTATUS_PUBLISHED
	this.Status = &status
	this.Time = time
	return &this
}

// NewNotebookUpdateDataAttributesWithDefaults instantiates a new NotebookUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookUpdateDataAttributesWithDefaults() *NotebookUpdateDataAttributes {
	this := NotebookUpdateDataAttributes{}
	var status NotebookStatus = NOTEBOOKSTATUS_PUBLISHED
	this.Status = &status
	return &this
}

// GetCells returns the Cells field value
func (o *NotebookUpdateDataAttributes) GetCells() []NotebookUpdateCell {
	if o == nil {
		var ret []NotebookUpdateCell
		return ret
	}

	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateDataAttributes) GetCellsOk() (*[]NotebookUpdateCell, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cells, true
}

// SetCells sets field value
func (o *NotebookUpdateDataAttributes) SetCells(v []NotebookUpdateCell) {
	o.Cells = v
}

// GetName returns the Name field value
func (o *NotebookUpdateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotebookUpdateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotebookUpdateDataAttributes) GetStatus() NotebookStatus {
	if o == nil || o.Status == nil {
		var ret NotebookStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookUpdateDataAttributes) GetStatusOk() (*NotebookStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotebookUpdateDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NotebookStatus and assigns it to the Status field.
func (o *NotebookUpdateDataAttributes) SetStatus(v NotebookStatus) {
	o.Status = &v
}

// GetTime returns the Time field value
func (o *NotebookUpdateDataAttributes) GetTime() NotebookGlobalTime {
	if o == nil {
		var ret NotebookGlobalTime
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *NotebookUpdateDataAttributes) SetTime(v NotebookGlobalTime) {
	o.Time = v
}

func (o NotebookUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cells"] = o.Cells
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookUpdateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Cells *[]NotebookUpdateCell `json:"cells"`
		Name  *string               `json:"name"`
		Time  *NotebookGlobalTime   `json:"time"`
	}{}
	all := struct {
		Cells  []NotebookUpdateCell `json:"cells"`
		Name   string               `json:"name"`
		Status *NotebookStatus      `json:"status,omitempty"`
		Time   NotebookGlobalTime   `json:"time"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Cells == nil {
		return fmt.Errorf("Required field cells missing")
	}
	if required.Name == nil {
		return fmt.Errorf("Required field name missing")
	}
	if required.Time == nil {
		return fmt.Errorf("Required field time missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Cells = all.Cells
	o.Name = all.Name
	o.Status = all.Status
	o.Time = all.Time
	return nil
}

type NullableNotebookUpdateDataAttributes struct {
	value *NotebookUpdateDataAttributes
	isSet bool
}

func (v NullableNotebookUpdateDataAttributes) Get() *NotebookUpdateDataAttributes {
	return v.value
}

func (v *NullableNotebookUpdateDataAttributes) Set(val *NotebookUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookUpdateDataAttributes(val *NotebookUpdateDataAttributes) *NullableNotebookUpdateDataAttributes {
	return &NullableNotebookUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableNotebookUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
