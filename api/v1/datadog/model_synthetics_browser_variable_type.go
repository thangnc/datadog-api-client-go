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

// SyntheticsBrowserVariableType the model 'SyntheticsBrowserVariableType'
type SyntheticsBrowserVariableType string

// List of SyntheticsBrowserVariableType
const (
	SYNTHETICSBROWSERVARIABLETYPE_ELEMENT SyntheticsBrowserVariableType = "element"
	SYNTHETICSBROWSERVARIABLETYPE_EMAIL   SyntheticsBrowserVariableType = "email"
	SYNTHETICSBROWSERVARIABLETYPE_GLOBAL  SyntheticsBrowserVariableType = "global"
	SYNTHETICSBROWSERVARIABLETYPE_TEXT    SyntheticsBrowserVariableType = "text"
)

// Ptr returns reference to SyntheticsBrowserVariableType value
func (v SyntheticsBrowserVariableType) Ptr() *SyntheticsBrowserVariableType {
	return &v
}

type NullableSyntheticsBrowserVariableType struct {
	value *SyntheticsBrowserVariableType
	isSet bool
}

func (v NullableSyntheticsBrowserVariableType) Get() *SyntheticsBrowserVariableType {
	return v.value
}

func (v NullableSyntheticsBrowserVariableType) Set(val *SyntheticsBrowserVariableType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserVariableType) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsBrowserVariableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserVariableType(val *SyntheticsBrowserVariableType) *NullableSyntheticsBrowserVariableType {
	return &NullableSyntheticsBrowserVariableType{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserVariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserVariableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}