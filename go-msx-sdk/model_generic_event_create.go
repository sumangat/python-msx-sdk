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

// GenericEventCreate struct for GenericEventCreate
type GenericEventCreate struct {
	Severity *GenericEventSeverity `json:"severity,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
	// Service that generated the event
	Service *string `json:"service,omitempty"`
	// Space separated list of keywords to be used for search
	Keywords *string `json:"keywords,omitempty"`
	Details *map[string]string `json:"details,omitempty"`
	Trace *GenericEventTrace `json:"trace,omitempty"`
	Security *GenericEventSecurity `json:"security,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericEventCreate GenericEventCreate

// NewGenericEventCreate instantiates a new GenericEventCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericEventCreate() *GenericEventCreate {
	this := GenericEventCreate{}
	return &this
}

// NewGenericEventCreateWithDefaults instantiates a new GenericEventCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericEventCreateWithDefaults() *GenericEventCreate {
	this := GenericEventCreate{}
	return &this
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *GenericEventCreate) GetSeverity() GenericEventSeverity {
	if o == nil || o.Severity == nil {
		var ret GenericEventSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetSeverityOk() (*GenericEventSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *GenericEventCreate) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given GenericEventSeverity and assigns it to the Severity field.
func (o *GenericEventCreate) SetSeverity(v GenericEventSeverity) {
	o.Severity = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *GenericEventCreate) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *GenericEventCreate) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *GenericEventCreate) SetSubtype(v string) {
	o.Subtype = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *GenericEventCreate) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *GenericEventCreate) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *GenericEventCreate) SetService(v string) {
	o.Service = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *GenericEventCreate) GetKeywords() string {
	if o == nil || o.Keywords == nil {
		var ret string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetKeywordsOk() (*string, bool) {
	if o == nil || o.Keywords == nil {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *GenericEventCreate) HasKeywords() bool {
	if o != nil && o.Keywords != nil {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given string and assigns it to the Keywords field.
func (o *GenericEventCreate) SetKeywords(v string) {
	o.Keywords = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GenericEventCreate) GetDetails() map[string]string {
	if o == nil || o.Details == nil {
		var ret map[string]string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetDetailsOk() (*map[string]string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GenericEventCreate) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]string and assigns it to the Details field.
func (o *GenericEventCreate) SetDetails(v map[string]string) {
	o.Details = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *GenericEventCreate) GetTrace() GenericEventTrace {
	if o == nil || o.Trace == nil {
		var ret GenericEventTrace
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetTraceOk() (*GenericEventTrace, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *GenericEventCreate) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given GenericEventTrace and assigns it to the Trace field.
func (o *GenericEventCreate) SetTrace(v GenericEventTrace) {
	o.Trace = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *GenericEventCreate) GetSecurity() GenericEventSecurity {
	if o == nil || o.Security == nil {
		var ret GenericEventSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetSecurityOk() (*GenericEventSecurity, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *GenericEventCreate) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given GenericEventSecurity and assigns it to the Security field.
func (o *GenericEventCreate) SetSecurity(v GenericEventSecurity) {
	o.Security = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenericEventCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenericEventCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenericEventCreate) SetDescription(v string) {
	o.Description = &v
}

func (o GenericEventCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Keywords != nil {
		toSerialize["keywords"] = o.Keywords
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericEventCreate) UnmarshalJSON(bytes []byte) (err error) {
	varGenericEventCreate := _GenericEventCreate{}

	if err = json.Unmarshal(bytes, &varGenericEventCreate); err == nil {
		*o = GenericEventCreate(varGenericEventCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "severity")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "service")
		delete(additionalProperties, "keywords")
		delete(additionalProperties, "details")
		delete(additionalProperties, "trace")
		delete(additionalProperties, "security")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericEventCreate struct {
	value *GenericEventCreate
	isSet bool
}

func (v NullableGenericEventCreate) Get() *GenericEventCreate {
	return v.value
}

func (v *NullableGenericEventCreate) Set(val *GenericEventCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericEventCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericEventCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericEventCreate(val *GenericEventCreate) *NullableGenericEventCreate {
	return &NullableGenericEventCreate{value: val, isSet: true}
}

func (v NullableGenericEventCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericEventCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


