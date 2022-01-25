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

// BillingEventsPageAllOf struct for BillingEventsPageAllOf
type BillingEventsPageAllOf struct {
	Contents *[]BillingEvent `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BillingEventsPageAllOf BillingEventsPageAllOf

// NewBillingEventsPageAllOf instantiates a new BillingEventsPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEventsPageAllOf() *BillingEventsPageAllOf {
	this := BillingEventsPageAllOf{}
	return &this
}

// NewBillingEventsPageAllOfWithDefaults instantiates a new BillingEventsPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEventsPageAllOfWithDefaults() *BillingEventsPageAllOf {
	this := BillingEventsPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *BillingEventsPageAllOf) GetContents() []BillingEvent {
	if o == nil || o.Contents == nil {
		var ret []BillingEvent
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventsPageAllOf) GetContentsOk() (*[]BillingEvent, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *BillingEventsPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []BillingEvent and assigns it to the Contents field.
func (o *BillingEventsPageAllOf) SetContents(v []BillingEvent) {
	o.Contents = &v
}

func (o BillingEventsPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BillingEventsPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEventsPageAllOf := _BillingEventsPageAllOf{}

	if err = json.Unmarshal(bytes, &varBillingEventsPageAllOf); err == nil {
		*o = BillingEventsPageAllOf(varBillingEventsPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingEventsPageAllOf struct {
	value *BillingEventsPageAllOf
	isSet bool
}

func (v NullableBillingEventsPageAllOf) Get() *BillingEventsPageAllOf {
	return v.value
}

func (v *NullableBillingEventsPageAllOf) Set(val *BillingEventsPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingEventsPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingEventsPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingEventsPageAllOf(val *BillingEventsPageAllOf) *NullableBillingEventsPageAllOf {
	return &NullableBillingEventsPageAllOf{value: val, isSet: true}
}

func (v NullableBillingEventsPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingEventsPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


