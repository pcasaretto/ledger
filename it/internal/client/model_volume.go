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

// Volume struct for Volume
type Volume struct {
	Input float32 `json:"input"`
	Output float32 `json:"output"`
	Balance float32 `json:"balance"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(input float32, output float32, balance float32) *Volume {
	this := Volume{}
	this.Input = input
	this.Output = output
	this.Balance = balance
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetInput returns the Input field value
func (o *Volume) GetInput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Volume) GetInputOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *Volume) SetInput(v float32) {
	o.Input = v
}

// GetOutput returns the Output field value
func (o *Volume) GetOutput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *Volume) GetOutputOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *Volume) SetOutput(v float32) {
	o.Output = v
}

// GetBalance returns the Balance field value
func (o *Volume) GetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Volume) GetBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Volume) SetBalance(v float32) {
	o.Balance = v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["output"] = o.Output
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	return json.Marshal(toSerialize)
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


