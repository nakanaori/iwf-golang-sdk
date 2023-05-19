/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the WorkflowRpcResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRpcResponse{}

// WorkflowRpcResponse struct for WorkflowRpcResponse
type WorkflowRpcResponse struct {
	Output *EncodedObject `json:"output,omitempty"`
}

// NewWorkflowRpcResponse instantiates a new WorkflowRpcResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRpcResponse() *WorkflowRpcResponse {
	this := WorkflowRpcResponse{}
	return &this
}

// NewWorkflowRpcResponseWithDefaults instantiates a new WorkflowRpcResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRpcResponseWithDefaults() *WorkflowRpcResponse {
	this := WorkflowRpcResponse{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowRpcResponse) GetOutput() EncodedObject {
	if o == nil || IsNil(o.Output) {
		var ret EncodedObject
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRpcResponse) GetOutputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowRpcResponse) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given EncodedObject and assigns it to the Output field.
func (o *WorkflowRpcResponse) SetOutput(v EncodedObject) {
	o.Output = &v
}

func (o WorkflowRpcResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRpcResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullableWorkflowRpcResponse struct {
	value *WorkflowRpcResponse
	isSet bool
}

func (v NullableWorkflowRpcResponse) Get() *WorkflowRpcResponse {
	return v.value
}

func (v *NullableWorkflowRpcResponse) Set(val *WorkflowRpcResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRpcResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRpcResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRpcResponse(val *WorkflowRpcResponse) *NullableWorkflowRpcResponse {
	return &NullableWorkflowRpcResponse{value: val, isSet: true}
}

func (v NullableWorkflowRpcResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRpcResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
