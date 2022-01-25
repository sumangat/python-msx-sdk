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

// WorkflowMetadataGitInfo struct for WorkflowMetadataGitInfo
type WorkflowMetadataGitInfo struct {
	TargetId NullableString `json:"target_id,omitempty"`
	CommitHash NullableString `json:"commit_hash,omitempty"`
	CommittedBy NullableString `json:"committed_by,omitempty"`
	CommitedOn NullableString `json:"commited_on,omitempty"`
	CommittedOn NullableString `json:"committed_on,omitempty"`
	CommitMessage NullableString `json:"commit_message,omitempty"`
	FileName NullableString `json:"file_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMetadataGitInfo WorkflowMetadataGitInfo

// NewWorkflowMetadataGitInfo instantiates a new WorkflowMetadataGitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMetadataGitInfo() *WorkflowMetadataGitInfo {
	this := WorkflowMetadataGitInfo{}
	return &this
}

// NewWorkflowMetadataGitInfoWithDefaults instantiates a new WorkflowMetadataGitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMetadataGitInfoWithDefaults() *WorkflowMetadataGitInfo {
	this := WorkflowMetadataGitInfo{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetTargetId() string {
	if o == nil || o.TargetId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetTargetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasTargetId() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableString and assigns it to the TargetId field.
func (o *WorkflowMetadataGitInfo) SetTargetId(v string) {
	o.TargetId.Set(&v)
}
// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetTargetId() {
	o.TargetId.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetCommitHash() string {
	if o == nil || o.CommitHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetCommitHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *WorkflowMetadataGitInfo) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetCommittedBy returns the CommittedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetCommittedBy() string {
	if o == nil || o.CommittedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommittedBy.Get()
}

// GetCommittedByOk returns a tuple with the CommittedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetCommittedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CommittedBy.Get(), o.CommittedBy.IsSet()
}

// HasCommittedBy returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasCommittedBy() bool {
	if o != nil && o.CommittedBy.IsSet() {
		return true
	}

	return false
}

// SetCommittedBy gets a reference to the given NullableString and assigns it to the CommittedBy field.
func (o *WorkflowMetadataGitInfo) SetCommittedBy(v string) {
	o.CommittedBy.Set(&v)
}
// SetCommittedByNil sets the value for CommittedBy to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetCommittedByNil() {
	o.CommittedBy.Set(nil)
}

// UnsetCommittedBy ensures that no value is present for CommittedBy, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetCommittedBy() {
	o.CommittedBy.Unset()
}

// GetCommitedOn returns the CommitedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetCommitedOn() string {
	if o == nil || o.CommitedOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitedOn.Get()
}

// GetCommitedOnOk returns a tuple with the CommitedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetCommitedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CommitedOn.Get(), o.CommitedOn.IsSet()
}

// HasCommitedOn returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasCommitedOn() bool {
	if o != nil && o.CommitedOn.IsSet() {
		return true
	}

	return false
}

// SetCommitedOn gets a reference to the given NullableString and assigns it to the CommitedOn field.
func (o *WorkflowMetadataGitInfo) SetCommitedOn(v string) {
	o.CommitedOn.Set(&v)
}
// SetCommitedOnNil sets the value for CommitedOn to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetCommitedOnNil() {
	o.CommitedOn.Set(nil)
}

// UnsetCommitedOn ensures that no value is present for CommitedOn, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetCommitedOn() {
	o.CommitedOn.Unset()
}

// GetCommittedOn returns the CommittedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetCommittedOn() string {
	if o == nil || o.CommittedOn.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommittedOn.Get()
}

// GetCommittedOnOk returns a tuple with the CommittedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetCommittedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CommittedOn.Get(), o.CommittedOn.IsSet()
}

// HasCommittedOn returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasCommittedOn() bool {
	if o != nil && o.CommittedOn.IsSet() {
		return true
	}

	return false
}

// SetCommittedOn gets a reference to the given NullableString and assigns it to the CommittedOn field.
func (o *WorkflowMetadataGitInfo) SetCommittedOn(v string) {
	o.CommittedOn.Set(&v)
}
// SetCommittedOnNil sets the value for CommittedOn to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetCommittedOnNil() {
	o.CommittedOn.Set(nil)
}

// UnsetCommittedOn ensures that no value is present for CommittedOn, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetCommittedOn() {
	o.CommittedOn.Unset()
}

// GetCommitMessage returns the CommitMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetCommitMessage() string {
	if o == nil || o.CommitMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommitMessage.Get()
}

// GetCommitMessageOk returns a tuple with the CommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetCommitMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CommitMessage.Get(), o.CommitMessage.IsSet()
}

// HasCommitMessage returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasCommitMessage() bool {
	if o != nil && o.CommitMessage.IsSet() {
		return true
	}

	return false
}

// SetCommitMessage gets a reference to the given NullableString and assigns it to the CommitMessage field.
func (o *WorkflowMetadataGitInfo) SetCommitMessage(v string) {
	o.CommitMessage.Set(&v)
}
// SetCommitMessageNil sets the value for CommitMessage to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetCommitMessageNil() {
	o.CommitMessage.Set(nil)
}

// UnsetCommitMessage ensures that no value is present for CommitMessage, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetCommitMessage() {
	o.CommitMessage.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMetadataGitInfo) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMetadataGitInfo) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *WorkflowMetadataGitInfo) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *WorkflowMetadataGitInfo) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *WorkflowMetadataGitInfo) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *WorkflowMetadataGitInfo) UnsetFileName() {
	o.FileName.Unset()
}

func (o WorkflowMetadataGitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetId.IsSet() {
		toSerialize["target_id"] = o.TargetId.Get()
	}
	if o.CommitHash.IsSet() {
		toSerialize["commit_hash"] = o.CommitHash.Get()
	}
	if o.CommittedBy.IsSet() {
		toSerialize["committed_by"] = o.CommittedBy.Get()
	}
	if o.CommitedOn.IsSet() {
		toSerialize["commited_on"] = o.CommitedOn.Get()
	}
	if o.CommittedOn.IsSet() {
		toSerialize["committed_on"] = o.CommittedOn.Get()
	}
	if o.CommitMessage.IsSet() {
		toSerialize["commit_message"] = o.CommitMessage.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["file_name"] = o.FileName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMetadataGitInfo) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMetadataGitInfo := _WorkflowMetadataGitInfo{}

	if err = json.Unmarshal(bytes, &varWorkflowMetadataGitInfo); err == nil {
		*o = WorkflowMetadataGitInfo(varWorkflowMetadataGitInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "target_id")
		delete(additionalProperties, "commit_hash")
		delete(additionalProperties, "committed_by")
		delete(additionalProperties, "commited_on")
		delete(additionalProperties, "committed_on")
		delete(additionalProperties, "commit_message")
		delete(additionalProperties, "file_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowMetadataGitInfo struct {
	value *WorkflowMetadataGitInfo
	isSet bool
}

func (v NullableWorkflowMetadataGitInfo) Get() *WorkflowMetadataGitInfo {
	return v.value
}

func (v *NullableWorkflowMetadataGitInfo) Set(val *WorkflowMetadataGitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMetadataGitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMetadataGitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMetadataGitInfo(val *WorkflowMetadataGitInfo) *NullableWorkflowMetadataGitInfo {
	return &NullableWorkflowMetadataGitInfo{value: val, isSet: true}
}

func (v NullableWorkflowMetadataGitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMetadataGitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


