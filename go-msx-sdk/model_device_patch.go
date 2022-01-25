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

// DevicePatch struct for DevicePatch
type DevicePatch struct {
	Name *string `json:"name,omitempty"`
	Model *string `json:"model,omitempty"`
	Type *string `json:"type,omitempty"`
	SubType NullableString `json:"subType,omitempty"`
	SerialKey NullableString `json:"serialKey,omitempty"`
	Version NullableString `json:"version,omitempty"`
	ComplianceState *DeviceComplianceState `json:"complianceState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePatch DevicePatch

// NewDevicePatch instantiates a new DevicePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePatch() *DevicePatch {
	this := DevicePatch{}
	return &this
}

// NewDevicePatchWithDefaults instantiates a new DevicePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePatchWithDefaults() *DevicePatch {
	this := DevicePatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicePatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicePatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicePatch) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DevicePatch) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePatch) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DevicePatch) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DevicePatch) SetModel(v string) {
	o.Model = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevicePatch) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePatch) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevicePatch) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevicePatch) SetType(v string) {
	o.Type = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DevicePatch) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DevicePatch) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *DevicePatch) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *DevicePatch) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *DevicePatch) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *DevicePatch) UnsetSubType() {
	o.SubType.Unset()
}

// GetSerialKey returns the SerialKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DevicePatch) GetSerialKey() string {
	if o == nil || o.SerialKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SerialKey.Get()
}

// GetSerialKeyOk returns a tuple with the SerialKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DevicePatch) GetSerialKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialKey.Get(), o.SerialKey.IsSet()
}

// HasSerialKey returns a boolean if a field has been set.
func (o *DevicePatch) HasSerialKey() bool {
	if o != nil && o.SerialKey.IsSet() {
		return true
	}

	return false
}

// SetSerialKey gets a reference to the given NullableString and assigns it to the SerialKey field.
func (o *DevicePatch) SetSerialKey(v string) {
	o.SerialKey.Set(&v)
}
// SetSerialKeyNil sets the value for SerialKey to be an explicit nil
func (o *DevicePatch) SetSerialKeyNil() {
	o.SerialKey.Set(nil)
}

// UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
func (o *DevicePatch) UnsetSerialKey() {
	o.SerialKey.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DevicePatch) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DevicePatch) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *DevicePatch) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *DevicePatch) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *DevicePatch) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *DevicePatch) UnsetVersion() {
	o.Version.Unset()
}

// GetComplianceState returns the ComplianceState field value if set, zero value otherwise.
func (o *DevicePatch) GetComplianceState() DeviceComplianceState {
	if o == nil || o.ComplianceState == nil {
		var ret DeviceComplianceState
		return ret
	}
	return *o.ComplianceState
}

// GetComplianceStateOk returns a tuple with the ComplianceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePatch) GetComplianceStateOk() (*DeviceComplianceState, bool) {
	if o == nil || o.ComplianceState == nil {
		return nil, false
	}
	return o.ComplianceState, true
}

// HasComplianceState returns a boolean if a field has been set.
func (o *DevicePatch) HasComplianceState() bool {
	if o != nil && o.ComplianceState != nil {
		return true
	}

	return false
}

// SetComplianceState gets a reference to the given DeviceComplianceState and assigns it to the ComplianceState field.
func (o *DevicePatch) SetComplianceState(v DeviceComplianceState) {
	o.ComplianceState = &v
}

func (o DevicePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SubType.IsSet() {
		toSerialize["subType"] = o.SubType.Get()
	}
	if o.SerialKey.IsSet() {
		toSerialize["serialKey"] = o.SerialKey.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.ComplianceState != nil {
		toSerialize["complianceState"] = o.ComplianceState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DevicePatch) UnmarshalJSON(bytes []byte) (err error) {
	varDevicePatch := _DevicePatch{}

	if err = json.Unmarshal(bytes, &varDevicePatch); err == nil {
		*o = DevicePatch(varDevicePatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "model")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subType")
		delete(additionalProperties, "serialKey")
		delete(additionalProperties, "version")
		delete(additionalProperties, "complianceState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePatch struct {
	value *DevicePatch
	isSet bool
}

func (v NullableDevicePatch) Get() *DevicePatch {
	return v.value
}

func (v *NullableDevicePatch) Set(val *DevicePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePatch(val *DevicePatch) *NullableDevicePatch {
	return &NullableDevicePatch{value: val, isSet: true}
}

func (v NullableDevicePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


