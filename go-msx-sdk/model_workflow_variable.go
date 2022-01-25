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

// WorkflowVariable struct for WorkflowVariable
type WorkflowVariable struct {
	Id *string `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	BaseType NullableString `json:"base_type,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	DataType NullableString `json:"data_type,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	CreatedOn *string `json:"created_on,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedOn *string `json:"updated_on,omitempty"`
	UpdatedBy NullableString `json:"updated_by,omitempty"`
	Owner NullableString `json:"owner,omitempty"`
	UniqueName NullableString `json:"unique_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowVariable WorkflowVariable

// NewWorkflowVariable instantiates a new WorkflowVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowVariable() *WorkflowVariable {
	this := WorkflowVariable{}
	return &this
}

// NewWorkflowVariableWithDefaults instantiates a new WorkflowVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowVariableWithDefaults() *WorkflowVariable {
	this := WorkflowVariable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowVariable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowVariable) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowVariable) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *WorkflowVariable) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *WorkflowVariable) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *WorkflowVariable) UnsetType() {
	o.Type.Unset()
}

// GetBaseType returns the BaseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetBaseType() string {
	if o == nil || o.BaseType.Get() == nil {
		var ret string
		return ret
	}
	return *o.BaseType.Get()
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetBaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseType.Get(), o.BaseType.IsSet()
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowVariable) HasBaseType() bool {
	if o != nil && o.BaseType.IsSet() {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given NullableString and assigns it to the BaseType field.
func (o *WorkflowVariable) SetBaseType(v string) {
	o.BaseType.Set(&v)
}
// SetBaseTypeNil sets the value for BaseType to be an explicit nil
func (o *WorkflowVariable) SetBaseTypeNil() {
	o.BaseType.Set(nil)
}

// UnsetBaseType ensures that no value is present for BaseType, not even an explicit nil
func (o *WorkflowVariable) UnsetBaseType() {
	o.BaseType.Unset()
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowVariable) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowVariable) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowVariable) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *WorkflowVariable) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *WorkflowVariable) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *WorkflowVariable) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowVariable) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowVariable) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *WorkflowVariable) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetDataType() string {
	if o == nil || o.DataType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataType.Get()
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetDataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataType.Get(), o.DataType.IsSet()
}

// HasDataType returns a boolean if a field has been set.
func (o *WorkflowVariable) HasDataType() bool {
	if o != nil && o.DataType.IsSet() {
		return true
	}

	return false
}

// SetDataType gets a reference to the given NullableString and assigns it to the DataType field.
func (o *WorkflowVariable) SetDataType(v string) {
	o.DataType.Set(&v)
}
// SetDataTypeNil sets the value for DataType to be an explicit nil
func (o *WorkflowVariable) SetDataTypeNil() {
	o.DataType.Set(nil)
}

// UnsetDataType ensures that no value is present for DataType, not even an explicit nil
func (o *WorkflowVariable) UnsetDataType() {
	o.DataType.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *WorkflowVariable) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *WorkflowVariable) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *WorkflowVariable) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *WorkflowVariable) UnsetScope() {
	o.Scope.Unset()
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *WorkflowVariable) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *WorkflowVariable) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *WorkflowVariable) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkflowVariable) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkflowVariable) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *WorkflowVariable) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *WorkflowVariable) GetUpdatedOn() string {
	if o == nil || o.UpdatedOn == nil {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariable) GetUpdatedOnOk() (*string, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *WorkflowVariable) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *WorkflowVariable) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkflowVariable) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *WorkflowVariable) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *WorkflowVariable) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *WorkflowVariable) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowVariable) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *WorkflowVariable) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WorkflowVariable) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WorkflowVariable) UnsetOwner() {
	o.Owner.Unset()
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariable) GetUniqueName() string {
	if o == nil || o.UniqueName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UniqueName.Get()
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariable) GetUniqueNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UniqueName.Get(), o.UniqueName.IsSet()
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowVariable) HasUniqueName() bool {
	if o != nil && o.UniqueName.IsSet() {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given NullableString and assigns it to the UniqueName field.
func (o *WorkflowVariable) SetUniqueName(v string) {
	o.UniqueName.Set(&v)
}
// SetUniqueNameNil sets the value for UniqueName to be an explicit nil
func (o *WorkflowVariable) SetUniqueNameNil() {
	o.UniqueName.Set(nil)
}

// UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
func (o *WorkflowVariable) UnsetUniqueName() {
	o.UniqueName.Unset()
}

func (o WorkflowVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.BaseType.IsSet() {
		toSerialize["base_type"] = o.BaseType.Get()
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.DataType.IsSet() {
		toSerialize["data_type"] = o.DataType.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.UniqueName.IsSet() {
		toSerialize["unique_name"] = o.UniqueName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowVariable) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowVariable := _WorkflowVariable{}

	if err = json.Unmarshal(bytes, &varWorkflowVariable); err == nil {
		*o = WorkflowVariable(varWorkflowVariable)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "base_type")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "data_type")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_on")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "unique_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowVariable struct {
	value *WorkflowVariable
	isSet bool
}

func (v NullableWorkflowVariable) Get() *WorkflowVariable {
	return v.value
}

func (v *NullableWorkflowVariable) Set(val *WorkflowVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowVariable(val *WorkflowVariable) *NullableWorkflowVariable {
	return &NullableWorkflowVariable{value: val, isSet: true}
}

func (v NullableWorkflowVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


