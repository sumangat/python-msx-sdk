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
	"time"
)

// TemplateStatusMeta struct for TemplateStatusMeta
type TemplateStatusMeta struct {
	Status *TemplateStatus `json:"status,omitempty"`
	StatusDetails NullableString `json:"statusDetails,omitempty"`
	CreatedOn NullableTime `json:"createdOn,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	ModifiedBy NullableString `json:"modifiedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateStatusMeta TemplateStatusMeta

// NewTemplateStatusMeta instantiates a new TemplateStatusMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateStatusMeta() *TemplateStatusMeta {
	this := TemplateStatusMeta{}
	return &this
}

// NewTemplateStatusMetaWithDefaults instantiates a new TemplateStatusMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateStatusMetaWithDefaults() *TemplateStatusMeta {
	this := TemplateStatusMeta{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TemplateStatusMeta) GetStatus() TemplateStatus {
	if o == nil || o.Status == nil {
		var ret TemplateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateStatusMeta) GetStatusOk() (*TemplateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TemplateStatus and assigns it to the Status field.
func (o *TemplateStatusMeta) SetStatus(v TemplateStatus) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateStatusMeta) GetStatusDetails() string {
	if o == nil || o.StatusDetails.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails.Get()
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateStatusMeta) GetStatusDetailsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDetails.Get(), o.StatusDetails.IsSet()
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasStatusDetails() bool {
	if o != nil && o.StatusDetails.IsSet() {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given NullableString and assigns it to the StatusDetails field.
func (o *TemplateStatusMeta) SetStatusDetails(v string) {
	o.StatusDetails.Set(&v)
}
// SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil
func (o *TemplateStatusMeta) SetStatusDetailsNil() {
	o.StatusDetails.Set(nil)
}

// UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
func (o *TemplateStatusMeta) UnsetStatusDetails() {
	o.StatusDetails.Unset()
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateStatusMeta) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn.Get()
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateStatusMeta) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedOn.Get(), o.CreatedOn.IsSet()
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasCreatedOn() bool {
	if o != nil && o.CreatedOn.IsSet() {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given NullableTime and assigns it to the CreatedOn field.
func (o *TemplateStatusMeta) SetCreatedOn(v time.Time) {
	o.CreatedOn.Set(&v)
}
// SetCreatedOnNil sets the value for CreatedOn to be an explicit nil
func (o *TemplateStatusMeta) SetCreatedOnNil() {
	o.CreatedOn.Set(nil)
}

// UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
func (o *TemplateStatusMeta) UnsetCreatedOn() {
	o.CreatedOn.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateStatusMeta) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateStatusMeta) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *TemplateStatusMeta) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *TemplateStatusMeta) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *TemplateStatusMeta) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateStatusMeta) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateStatusMeta) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *TemplateStatusMeta) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *TemplateStatusMeta) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *TemplateStatusMeta) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateStatusMeta) GetModifiedBy() string {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateStatusMeta) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *TemplateStatusMeta) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *TemplateStatusMeta) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}
// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *TemplateStatusMeta) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *TemplateStatusMeta) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

func (o TemplateStatusMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails.IsSet() {
		toSerialize["statusDetails"] = o.StatusDetails.Get()
	}
	if o.CreatedOn.IsSet() {
		toSerialize["createdOn"] = o.CreatedOn.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modifiedBy"] = o.ModifiedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateStatusMeta) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateStatusMeta := _TemplateStatusMeta{}

	if err = json.Unmarshal(bytes, &varTemplateStatusMeta); err == nil {
		*o = TemplateStatusMeta(varTemplateStatusMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetails")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modifiedOn")
		delete(additionalProperties, "modifiedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateStatusMeta struct {
	value *TemplateStatusMeta
	isSet bool
}

func (v NullableTemplateStatusMeta) Get() *TemplateStatusMeta {
	return v.value
}

func (v *NullableTemplateStatusMeta) Set(val *TemplateStatusMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateStatusMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateStatusMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateStatusMeta(val *TemplateStatusMeta) *NullableTemplateStatusMeta {
	return &NullableTemplateStatusMeta{value: val, isSet: true}
}

func (v NullableTemplateStatusMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateStatusMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


