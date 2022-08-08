/*
OneLogin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VerifyFactorSaml200Response struct for VerifyFactorSaml200Response
type VerifyFactorSaml200Response struct {
	// Provides the SAML assertion.
	Data *string `json:"data,omitempty"`
	// Plain text description describing the outcome of the response.
	Message *string `json:"message,omitempty"`
}

// NewVerifyFactorSaml200Response instantiates a new VerifyFactorSaml200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyFactorSaml200Response() *VerifyFactorSaml200Response {
	this := VerifyFactorSaml200Response{}
	return &this
}

// NewVerifyFactorSaml200ResponseWithDefaults instantiates a new VerifyFactorSaml200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyFactorSaml200ResponseWithDefaults() *VerifyFactorSaml200Response {
	this := VerifyFactorSaml200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VerifyFactorSaml200Response) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorSaml200Response) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VerifyFactorSaml200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *VerifyFactorSaml200Response) SetData(v string) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VerifyFactorSaml200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorSaml200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VerifyFactorSaml200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VerifyFactorSaml200Response) SetMessage(v string) {
	o.Message = &v
}

func (o VerifyFactorSaml200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyFactorSaml200Response struct {
	value *VerifyFactorSaml200Response
	isSet bool
}

func (v NullableVerifyFactorSaml200Response) Get() *VerifyFactorSaml200Response {
	return v.value
}

func (v *NullableVerifyFactorSaml200Response) Set(val *VerifyFactorSaml200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyFactorSaml200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyFactorSaml200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyFactorSaml200Response(val *VerifyFactorSaml200Response) *NullableVerifyFactorSaml200Response {
	return &NullableVerifyFactorSaml200Response{value: val, isSet: true}
}

func (v NullableVerifyFactorSaml200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyFactorSaml200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


