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

// ServiceUpdate struct for ServiceUpdate
type ServiceUpdate struct {
	Status *map[string]interface{} `json:"status,omitempty"`
	DefinitionAttributes *map[string]interface{} `json:"definitionAttributes,omitempty"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceUpdate ServiceUpdate

// NewServiceUpdate instantiates a new ServiceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUpdate() *ServiceUpdate {
	this := ServiceUpdate{}
	return &this
}

// NewServiceUpdateWithDefaults instantiates a new ServiceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUpdateWithDefaults() *ServiceUpdate {
	this := ServiceUpdate{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceUpdate) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUpdate) GetStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *ServiceUpdate) SetStatus(v map[string]interface{}) {
	o.Status = &v
}

// GetDefinitionAttributes returns the DefinitionAttributes field value if set, zero value otherwise.
func (o *ServiceUpdate) GetDefinitionAttributes() map[string]interface{} {
	if o == nil || o.DefinitionAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DefinitionAttributes
}

// GetDefinitionAttributesOk returns a tuple with the DefinitionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUpdate) GetDefinitionAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.DefinitionAttributes == nil {
		return nil, false
	}
	return o.DefinitionAttributes, true
}

// HasDefinitionAttributes returns a boolean if a field has been set.
func (o *ServiceUpdate) HasDefinitionAttributes() bool {
	if o != nil && o.DefinitionAttributes != nil {
		return true
	}

	return false
}

// SetDefinitionAttributes gets a reference to the given map[string]interface{} and assigns it to the DefinitionAttributes field.
func (o *ServiceUpdate) SetDefinitionAttributes(v map[string]interface{}) {
	o.DefinitionAttributes = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceUpdate) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUpdate) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceUpdate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ServiceUpdate) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

func (o ServiceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DefinitionAttributes != nil {
		toSerialize["definitionAttributes"] = o.DefinitionAttributes
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varServiceUpdate := _ServiceUpdate{}

	if err = json.Unmarshal(bytes, &varServiceUpdate); err == nil {
		*o = ServiceUpdate(varServiceUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "definitionAttributes")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceUpdate struct {
	value *ServiceUpdate
	isSet bool
}

func (v NullableServiceUpdate) Get() *ServiceUpdate {
	return v.value
}

func (v *NullableServiceUpdate) Set(val *ServiceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUpdate(val *ServiceUpdate) *NullableServiceUpdate {
	return &NullableServiceUpdate{value: val, isSet: true}
}

func (v NullableServiceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


