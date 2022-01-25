/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// VulnerabilitiesPageAllOf struct for VulnerabilitiesPageAllOf
type VulnerabilitiesPageAllOf struct {
	Contents []Vulnerability `json:"contents"`
	AdditionalProperties map[string]interface{}
}

type _VulnerabilitiesPageAllOf VulnerabilitiesPageAllOf

// NewVulnerabilitiesPageAllOf instantiates a new VulnerabilitiesPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerabilitiesPageAllOf(contents []Vulnerability) *VulnerabilitiesPageAllOf {
	this := VulnerabilitiesPageAllOf{}
	this.Contents = contents
	return &this
}

// NewVulnerabilitiesPageAllOfWithDefaults instantiates a new VulnerabilitiesPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerabilitiesPageAllOfWithDefaults() *VulnerabilitiesPageAllOf {
	this := VulnerabilitiesPageAllOf{}
	return &this
}

// GetContents returns the Contents field value
func (o *VulnerabilitiesPageAllOf) GetContents() []Vulnerability {
	if o == nil {
		var ret []Vulnerability
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesPageAllOf) GetContentsOk() (*[]Vulnerability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *VulnerabilitiesPageAllOf) SetContents(v []Vulnerability) {
	o.Contents = v
}

func (o VulnerabilitiesPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VulnerabilitiesPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVulnerabilitiesPageAllOf := _VulnerabilitiesPageAllOf{}

	if err = json.Unmarshal(bytes, &varVulnerabilitiesPageAllOf); err == nil {
		*o = VulnerabilitiesPageAllOf(varVulnerabilitiesPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVulnerabilitiesPageAllOf struct {
	value *VulnerabilitiesPageAllOf
	isSet bool
}

func (v NullableVulnerabilitiesPageAllOf) Get() *VulnerabilitiesPageAllOf {
	return v.value
}

func (v *NullableVulnerabilitiesPageAllOf) Set(val *VulnerabilitiesPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerabilitiesPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerabilitiesPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerabilitiesPageAllOf(val *VulnerabilitiesPageAllOf) *NullableVulnerabilitiesPageAllOf {
	return &NullableVulnerabilitiesPageAllOf{value: val, isSet: true}
}

func (v NullableVulnerabilitiesPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerabilitiesPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


