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

// TemplateApplicationStatusPatch struct for TemplateApplicationStatusPatch
type TemplateApplicationStatusPatch struct {
	Status TemplateStatus `json:"status"`
	StatusDetails *string `json:"statusDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateApplicationStatusPatch TemplateApplicationStatusPatch

// NewTemplateApplicationStatusPatch instantiates a new TemplateApplicationStatusPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateApplicationStatusPatch(status TemplateStatus) *TemplateApplicationStatusPatch {
	this := TemplateApplicationStatusPatch{}
	this.Status = status
	return &this
}

// NewTemplateApplicationStatusPatchWithDefaults instantiates a new TemplateApplicationStatusPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateApplicationStatusPatchWithDefaults() *TemplateApplicationStatusPatch {
	this := TemplateApplicationStatusPatch{}
	return &this
}

// GetStatus returns the Status field value
func (o *TemplateApplicationStatusPatch) GetStatus() TemplateStatus {
	if o == nil {
		var ret TemplateStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateApplicationStatusPatch) GetStatusOk() (*TemplateStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateApplicationStatusPatch) SetStatus(v TemplateStatus) {
	o.Status = v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *TemplateApplicationStatusPatch) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplicationStatusPatch) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *TemplateApplicationStatusPatch) HasStatusDetails() bool {
	if o != nil && o.StatusDetails != nil {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *TemplateApplicationStatusPatch) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o TemplateApplicationStatusPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["statusDetails"] = o.StatusDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateApplicationStatusPatch) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateApplicationStatusPatch := _TemplateApplicationStatusPatch{}

	if err = json.Unmarshal(bytes, &varTemplateApplicationStatusPatch); err == nil {
		*o = TemplateApplicationStatusPatch(varTemplateApplicationStatusPatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateApplicationStatusPatch struct {
	value *TemplateApplicationStatusPatch
	isSet bool
}

func (v NullableTemplateApplicationStatusPatch) Get() *TemplateApplicationStatusPatch {
	return v.value
}

func (v *NullableTemplateApplicationStatusPatch) Set(val *TemplateApplicationStatusPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateApplicationStatusPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateApplicationStatusPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateApplicationStatusPatch(val *TemplateApplicationStatusPatch) *NullableTemplateApplicationStatusPatch {
	return &NullableTemplateApplicationStatusPatch{value: val, isSet: true}
}

func (v NullableTemplateApplicationStatusPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateApplicationStatusPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


