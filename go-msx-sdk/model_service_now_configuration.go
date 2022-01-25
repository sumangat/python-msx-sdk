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

// ServiceNowConfiguration struct for ServiceNowConfiguration
type ServiceNowConfiguration struct {
	Id *string `json:"id,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	CriticalEvent *bool `json:"criticalEvent,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Proxy NullableString `json:"proxy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceNowConfiguration ServiceNowConfiguration

// NewServiceNowConfiguration instantiates a new ServiceNowConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowConfiguration() *ServiceNowConfiguration {
	this := ServiceNowConfiguration{}
	return &this
}

// NewServiceNowConfigurationWithDefaults instantiates a new ServiceNowConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowConfigurationWithDefaults() *ServiceNowConfiguration {
	this := ServiceNowConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceNowConfiguration) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceNowConfiguration) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ServiceNowConfiguration) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetCriticalEvent returns the CriticalEvent field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetCriticalEvent() bool {
	if o == nil || o.CriticalEvent == nil {
		var ret bool
		return ret
	}
	return *o.CriticalEvent
}

// GetCriticalEventOk returns a tuple with the CriticalEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetCriticalEventOk() (*bool, bool) {
	if o == nil || o.CriticalEvent == nil {
		return nil, false
	}
	return o.CriticalEvent, true
}

// HasCriticalEvent returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasCriticalEvent() bool {
	if o != nil && o.CriticalEvent != nil {
		return true
	}

	return false
}

// SetCriticalEvent gets a reference to the given bool and assigns it to the CriticalEvent field.
func (o *ServiceNowConfiguration) SetCriticalEvent(v bool) {
	o.CriticalEvent = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ServiceNowConfiguration) SetDomain(v string) {
	o.Domain = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ServiceNowConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ServiceNowConfiguration) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfiguration) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ServiceNowConfiguration) SetUserName(v string) {
	o.UserName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceNowConfiguration) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceNowConfiguration) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ServiceNowConfiguration) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ServiceNowConfiguration) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ServiceNowConfiguration) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceNowConfiguration) GetProxy() string {
	if o == nil || o.Proxy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceNowConfiguration) GetProxyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *ServiceNowConfiguration) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *ServiceNowConfiguration) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *ServiceNowConfiguration) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *ServiceNowConfiguration) UnsetProxy() {
	o.Proxy.Unset()
}

func (o ServiceNowConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.CriticalEvent != nil {
		toSerialize["criticalEvent"] = o.CriticalEvent
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceNowConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varServiceNowConfiguration := _ServiceNowConfiguration{}

	if err = json.Unmarshal(bytes, &varServiceNowConfiguration); err == nil {
		*o = ServiceNowConfiguration(varServiceNowConfiguration)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "criticalEvent")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "password")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "proxy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceNowConfiguration struct {
	value *ServiceNowConfiguration
	isSet bool
}

func (v NullableServiceNowConfiguration) Get() *ServiceNowConfiguration {
	return v.value
}

func (v *NullableServiceNowConfiguration) Set(val *ServiceNowConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowConfiguration(val *ServiceNowConfiguration) *NullableServiceNowConfiguration {
	return &NullableServiceNowConfiguration{value: val, isSet: true}
}

func (v NullableServiceNowConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


