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

// AwsLogsListResponse struct for AwsLogsListResponse
type AwsLogsListResponse struct {
	// Your AWS Account ID without dashes.
	AccountId *string `json:"account_id,omitempty"`
	// List of ARNs configured in your Datadog account.
	Lambdas *[]AwsLogsListResponseLambdas `json:"lambdas,omitempty"`
	// Array of services IDs.
	Services *[]string `json:"services,omitempty"`
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AwsLogsListResponse) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListResponse) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AwsLogsListResponse) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AwsLogsListResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetLambdas returns the Lambdas field value if set, zero value otherwise.
func (o *AwsLogsListResponse) GetLambdas() []AwsLogsListResponseLambdas {
	if o == nil || o.Lambdas == nil {
		var ret []AwsLogsListResponseLambdas
		return ret
	}
	return *o.Lambdas
}

// GetLambdasOk returns a tuple with the Lambdas field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListResponse) GetLambdasOk() ([]AwsLogsListResponseLambdas, bool) {
	if o == nil || o.Lambdas == nil {
		var ret []AwsLogsListResponseLambdas
		return ret, false
	}
	return *o.Lambdas, true
}

// HasLambdas returns a boolean if a field has been set.
func (o *AwsLogsListResponse) HasLambdas() bool {
	if o != nil && o.Lambdas != nil {
		return true
	}

	return false
}

// SetLambdas gets a reference to the given []AwsLogsListResponseLambdas and assigns it to the Lambdas field.
func (o *AwsLogsListResponse) SetLambdas(v []AwsLogsListResponseLambdas) {
	o.Lambdas = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AwsLogsListResponse) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsListResponse) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		var ret []string
		return ret, false
	}
	return *o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AwsLogsListResponse) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *AwsLogsListResponse) SetServices(v []string) {
	o.Services = &v
}

type NullableAwsLogsListResponse struct {
	Value        AwsLogsListResponse
	ExplicitNull bool
}

func (v NullableAwsLogsListResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsLogsListResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
