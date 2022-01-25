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

// DeviceTemplateBatchAttachRequest struct for DeviceTemplateBatchAttachRequest
type DeviceTemplateBatchAttachRequest struct {
	DeviceIds *[]string `json:"deviceIds,omitempty"`
	TemplateDetails *[]DeviceTemplateDetails `json:"templateDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplateBatchAttachRequest DeviceTemplateBatchAttachRequest

// NewDeviceTemplateBatchAttachRequest instantiates a new DeviceTemplateBatchAttachRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateBatchAttachRequest() *DeviceTemplateBatchAttachRequest {
	this := DeviceTemplateBatchAttachRequest{}
	return &this
}

// NewDeviceTemplateBatchAttachRequestWithDefaults instantiates a new DeviceTemplateBatchAttachRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateBatchAttachRequestWithDefaults() *DeviceTemplateBatchAttachRequest {
	this := DeviceTemplateBatchAttachRequest{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *DeviceTemplateBatchAttachRequest) GetDeviceIds() []string {
	if o == nil || o.DeviceIds == nil {
		var ret []string
		return ret
	}
	return *o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateBatchAttachRequest) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil || o.DeviceIds == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *DeviceTemplateBatchAttachRequest) HasDeviceIds() bool {
	if o != nil && o.DeviceIds != nil {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *DeviceTemplateBatchAttachRequest) SetDeviceIds(v []string) {
	o.DeviceIds = &v
}

// GetTemplateDetails returns the TemplateDetails field value if set, zero value otherwise.
func (o *DeviceTemplateBatchAttachRequest) GetTemplateDetails() []DeviceTemplateDetails {
	if o == nil || o.TemplateDetails == nil {
		var ret []DeviceTemplateDetails
		return ret
	}
	return *o.TemplateDetails
}

// GetTemplateDetailsOk returns a tuple with the TemplateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateBatchAttachRequest) GetTemplateDetailsOk() (*[]DeviceTemplateDetails, bool) {
	if o == nil || o.TemplateDetails == nil {
		return nil, false
	}
	return o.TemplateDetails, true
}

// HasTemplateDetails returns a boolean if a field has been set.
func (o *DeviceTemplateBatchAttachRequest) HasTemplateDetails() bool {
	if o != nil && o.TemplateDetails != nil {
		return true
	}

	return false
}

// SetTemplateDetails gets a reference to the given []DeviceTemplateDetails and assigns it to the TemplateDetails field.
func (o *DeviceTemplateBatchAttachRequest) SetTemplateDetails(v []DeviceTemplateDetails) {
	o.TemplateDetails = &v
}

func (o DeviceTemplateBatchAttachRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceIds != nil {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if o.TemplateDetails != nil {
		toSerialize["templateDetails"] = o.TemplateDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceTemplateBatchAttachRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplateBatchAttachRequest := _DeviceTemplateBatchAttachRequest{}

	if err = json.Unmarshal(bytes, &varDeviceTemplateBatchAttachRequest); err == nil {
		*o = DeviceTemplateBatchAttachRequest(varDeviceTemplateBatchAttachRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceIds")
		delete(additionalProperties, "templateDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTemplateBatchAttachRequest struct {
	value *DeviceTemplateBatchAttachRequest
	isSet bool
}

func (v NullableDeviceTemplateBatchAttachRequest) Get() *DeviceTemplateBatchAttachRequest {
	return v.value
}

func (v *NullableDeviceTemplateBatchAttachRequest) Set(val *DeviceTemplateBatchAttachRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateBatchAttachRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateBatchAttachRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateBatchAttachRequest(val *DeviceTemplateBatchAttachRequest) *NullableDeviceTemplateBatchAttachRequest {
	return &NullableDeviceTemplateBatchAttachRequest{value: val, isSet: true}
}

func (v NullableDeviceTemplateBatchAttachRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateBatchAttachRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


