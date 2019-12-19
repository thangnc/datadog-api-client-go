/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// Error413 struct for Error413
type Error413 struct {
	Errors []string `json:"errors"`
}

// GetErrors returns the Errors field value
func (o *Error413) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// SetErrors sets field value
func (o *Error413) SetErrors(v []string) {
	o.Errors = v
}

type NullableError413 struct {
	Value        Error413
	ExplicitNull bool
}

func (v NullableError413) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableError413) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}