/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: LEDGER_VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ListTransactions200Response struct for ListTransactions200Response
type ListTransactions200Response struct {
	Cursor ListTransactions200ResponseCursor `json:"cursor"`
}

// NewListTransactions200Response instantiates a new ListTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactions200Response(cursor ListTransactions200ResponseCursor) *ListTransactions200Response {
	this := ListTransactions200Response{}
	this.Cursor = cursor
	return &this
}

// NewListTransactions200ResponseWithDefaults instantiates a new ListTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactions200ResponseWithDefaults() *ListTransactions200Response {
	this := ListTransactions200Response{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListTransactions200Response) GetCursor() ListTransactions200ResponseCursor {
	if o == nil {
		var ret ListTransactions200ResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListTransactions200Response) GetCursorOk() (*ListTransactions200ResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListTransactions200Response) SetCursor(v ListTransactions200ResponseCursor) {
	o.Cursor = v
}

func (o ListTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactions200Response struct {
	value *ListTransactions200Response
	isSet bool
}

func (v NullableListTransactions200Response) Get() *ListTransactions200Response {
	return v.value
}

func (v *NullableListTransactions200Response) Set(val *ListTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactions200Response(val *ListTransactions200Response) *NullableListTransactions200Response {
	return &NullableListTransactions200Response{value: val, isSet: true}
}

func (v NullableListTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


