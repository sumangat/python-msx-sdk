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

// DeviceCreate struct for DeviceCreate
type DeviceCreate struct {
	ServiceInstanceId NullableString `json:"serviceInstanceId,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	TenantId string `json:"tenantId"`
	ServiceType NullableString `json:"serviceType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Managed bool `json:"managed"`
	OnboardType string `json:"onboardType"`
	OnboardInformation map[string]interface{} `json:"onboardInformation,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Name string `json:"name"`
	Model string `json:"model"`
	Type string `json:"type"`
	SubType NullableString `json:"subType,omitempty"`
	SerialKey NullableString `json:"serialKey,omitempty"`
	Version NullableString `json:"version,omitempty"`
	ComplianceState *DeviceComplianceState `json:"complianceState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceCreate DeviceCreate

// NewDeviceCreate instantiates a new DeviceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreate(tenantId string, managed bool, onboardType string, name string, model string, type_ string) *DeviceCreate {
	this := DeviceCreate{}
	this.TenantId = tenantId
	this.Managed = managed
	this.OnboardType = onboardType
	this.Name = name
	this.Model = model
	this.Type = type_
	return &this
}

// NewDeviceCreateWithDefaults instantiates a new DeviceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateWithDefaults() *DeviceCreate {
	this := DeviceCreate{}
	var managed bool = false
	this.Managed = managed
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetServiceInstanceId() string {
	if o == nil || o.ServiceInstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId.Get()
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceInstanceId.Get(), o.ServiceInstanceId.IsSet()
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *DeviceCreate) HasServiceInstanceId() bool {
	if o != nil && o.ServiceInstanceId.IsSet() {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given NullableString and assigns it to the ServiceInstanceId field.
func (o *DeviceCreate) SetServiceInstanceId(v string) {
	o.ServiceInstanceId.Set(&v)
}
// SetServiceInstanceIdNil sets the value for ServiceInstanceId to be an explicit nil
func (o *DeviceCreate) SetServiceInstanceIdNil() {
	o.ServiceInstanceId.Set(nil)
}

// UnsetServiceInstanceId ensures that no value is present for ServiceInstanceId, not even an explicit nil
func (o *DeviceCreate) UnsetServiceInstanceId() {
	o.ServiceInstanceId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *DeviceCreate) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *DeviceCreate) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *DeviceCreate) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *DeviceCreate) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetTenantId returns the TenantId field value
func (o *DeviceCreate) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *DeviceCreate) SetTenantId(v string) {
	o.TenantId = v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetServiceType() string {
	if o == nil || o.ServiceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceType.Get()
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetServiceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceType.Get(), o.ServiceType.IsSet()
}

// HasServiceType returns a boolean if a field has been set.
func (o *DeviceCreate) HasServiceType() bool {
	if o != nil && o.ServiceType.IsSet() {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given NullableString and assigns it to the ServiceType field.
func (o *DeviceCreate) SetServiceType(v string) {
	o.ServiceType.Set(&v)
}
// SetServiceTypeNil sets the value for ServiceType to be an explicit nil
func (o *DeviceCreate) SetServiceTypeNil() {
	o.ServiceType.Set(nil)
}

// UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
func (o *DeviceCreate) UnsetServiceType() {
	o.ServiceType.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetTags() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceCreate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *DeviceCreate) SetTags(v map[string]string) {
	o.Tags = v
}

// GetManaged returns the Managed field value
func (o *DeviceCreate) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetManagedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *DeviceCreate) SetManaged(v bool) {
	o.Managed = v
}

// GetOnboardType returns the OnboardType field value
func (o *DeviceCreate) GetOnboardType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OnboardType
}

// GetOnboardTypeOk returns a tuple with the OnboardType field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetOnboardTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnboardType, true
}

// SetOnboardType sets field value
func (o *DeviceCreate) SetOnboardType(v string) {
	o.OnboardType = v
}

// GetOnboardInformation returns the OnboardInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetOnboardInformation() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.OnboardInformation
}

// GetOnboardInformationOk returns a tuple with the OnboardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetOnboardInformationOk() (*map[string]interface{}, bool) {
	if o == nil || o.OnboardInformation == nil {
		return nil, false
	}
	return &o.OnboardInformation, true
}

// HasOnboardInformation returns a boolean if a field has been set.
func (o *DeviceCreate) HasOnboardInformation() bool {
	if o != nil && o.OnboardInformation != nil {
		return true
	}

	return false
}

// SetOnboardInformation gets a reference to the given map[string]interface{} and assigns it to the OnboardInformation field.
func (o *DeviceCreate) SetOnboardInformation(v map[string]interface{}) {
	o.OnboardInformation = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetAttributes() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeviceCreate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *DeviceCreate) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetName returns the Name field value
func (o *DeviceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceCreate) SetName(v string) {
	o.Name = v
}

// GetModel returns the Model field value
func (o *DeviceCreate) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceCreate) SetModel(v string) {
	o.Model = v
}

// GetType returns the Type field value
func (o *DeviceCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceCreate) SetType(v string) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *DeviceCreate) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *DeviceCreate) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *DeviceCreate) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *DeviceCreate) UnsetSubType() {
	o.SubType.Unset()
}

// GetSerialKey returns the SerialKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetSerialKey() string {
	if o == nil || o.SerialKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SerialKey.Get()
}

// GetSerialKeyOk returns a tuple with the SerialKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetSerialKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialKey.Get(), o.SerialKey.IsSet()
}

// HasSerialKey returns a boolean if a field has been set.
func (o *DeviceCreate) HasSerialKey() bool {
	if o != nil && o.SerialKey.IsSet() {
		return true
	}

	return false
}

// SetSerialKey gets a reference to the given NullableString and assigns it to the SerialKey field.
func (o *DeviceCreate) SetSerialKey(v string) {
	o.SerialKey.Set(&v)
}
// SetSerialKeyNil sets the value for SerialKey to be an explicit nil
func (o *DeviceCreate) SetSerialKeyNil() {
	o.SerialKey.Set(nil)
}

// UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
func (o *DeviceCreate) UnsetSerialKey() {
	o.SerialKey.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceCreate) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceCreate) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceCreate) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *DeviceCreate) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *DeviceCreate) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *DeviceCreate) UnsetVersion() {
	o.Version.Unset()
}

// GetComplianceState returns the ComplianceState field value if set, zero value otherwise.
func (o *DeviceCreate) GetComplianceState() DeviceComplianceState {
	if o == nil || o.ComplianceState == nil {
		var ret DeviceComplianceState
		return ret
	}
	return *o.ComplianceState
}

// GetComplianceStateOk returns a tuple with the ComplianceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreate) GetComplianceStateOk() (*DeviceComplianceState, bool) {
	if o == nil || o.ComplianceState == nil {
		return nil, false
	}
	return o.ComplianceState, true
}

// HasComplianceState returns a boolean if a field has been set.
func (o *DeviceCreate) HasComplianceState() bool {
	if o != nil && o.ComplianceState != nil {
		return true
	}

	return false
}

// SetComplianceState gets a reference to the given DeviceComplianceState and assigns it to the ComplianceState field.
func (o *DeviceCreate) SetComplianceState(v DeviceComplianceState) {
	o.ComplianceState = &v
}

func (o DeviceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceInstanceId.IsSet() {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.ServiceType.IsSet() {
		toSerialize["serviceType"] = o.ServiceType.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["managed"] = o.Managed
	}
	if true {
		toSerialize["onboardType"] = o.OnboardType
	}
	if o.OnboardInformation != nil {
		toSerialize["onboardInformation"] = o.OnboardInformation
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
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

func (o *DeviceCreate) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceCreate := _DeviceCreate{}

	if err = json.Unmarshal(bytes, &varDeviceCreate); err == nil {
		*o = DeviceCreate(varDeviceCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceInstanceId")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "onboardType")
		delete(additionalProperties, "onboardInformation")
		delete(additionalProperties, "attributes")
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

type NullableDeviceCreate struct {
	value *DeviceCreate
	isSet bool
}

func (v NullableDeviceCreate) Get() *DeviceCreate {
	return v.value
}

func (v *NullableDeviceCreate) Set(val *DeviceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreate(val *DeviceCreate) *NullableDeviceCreate {
	return &NullableDeviceCreate{value: val, isSet: true}
}

func (v NullableDeviceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


