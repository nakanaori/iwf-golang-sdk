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

// WorkflowSkipTimerRequest struct for WorkflowSkipTimerRequest
type WorkflowSkipTimerRequest struct {
	WorkflowId string `json:"workflowId"`
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
	WorkflowStateExecutionId string `json:"workflowStateExecutionId"`
	TimerCommandId *string `json:"timerCommandId,omitempty"`
	TimerCommandIndex *int32 `json:"timerCommandIndex,omitempty"`
}

// NewWorkflowSkipTimerRequest instantiates a new WorkflowSkipTimerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSkipTimerRequest(workflowId string, workflowStateExecutionId string) *WorkflowSkipTimerRequest {
	this := WorkflowSkipTimerRequest{}
	this.WorkflowId = workflowId
	this.WorkflowStateExecutionId = workflowStateExecutionId
	return &this
}

// NewWorkflowSkipTimerRequestWithDefaults instantiates a new WorkflowSkipTimerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSkipTimerRequestWithDefaults() *WorkflowSkipTimerRequest {
	this := WorkflowSkipTimerRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowSkipTimerRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSkipTimerRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowSkipTimerRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowSkipTimerRequest) GetWorkflowRunId() string {
	if o == nil || isNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSkipTimerRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowRunId) {
    return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowSkipTimerRequest) HasWorkflowRunId() bool {
	if o != nil && !isNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowSkipTimerRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetWorkflowStateExecutionId returns the WorkflowStateExecutionId field value
func (o *WorkflowSkipTimerRequest) GetWorkflowStateExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowStateExecutionId
}

// GetWorkflowStateExecutionIdOk returns a tuple with the WorkflowStateExecutionId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSkipTimerRequest) GetWorkflowStateExecutionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WorkflowStateExecutionId, true
}

// SetWorkflowStateExecutionId sets field value
func (o *WorkflowSkipTimerRequest) SetWorkflowStateExecutionId(v string) {
	o.WorkflowStateExecutionId = v
}

// GetTimerCommandId returns the TimerCommandId field value if set, zero value otherwise.
func (o *WorkflowSkipTimerRequest) GetTimerCommandId() string {
	if o == nil || isNil(o.TimerCommandId) {
		var ret string
		return ret
	}
	return *o.TimerCommandId
}

// GetTimerCommandIdOk returns a tuple with the TimerCommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSkipTimerRequest) GetTimerCommandIdOk() (*string, bool) {
	if o == nil || isNil(o.TimerCommandId) {
    return nil, false
	}
	return o.TimerCommandId, true
}

// HasTimerCommandId returns a boolean if a field has been set.
func (o *WorkflowSkipTimerRequest) HasTimerCommandId() bool {
	if o != nil && !isNil(o.TimerCommandId) {
		return true
	}

	return false
}

// SetTimerCommandId gets a reference to the given string and assigns it to the TimerCommandId field.
func (o *WorkflowSkipTimerRequest) SetTimerCommandId(v string) {
	o.TimerCommandId = &v
}

// GetTimerCommandIndex returns the TimerCommandIndex field value if set, zero value otherwise.
func (o *WorkflowSkipTimerRequest) GetTimerCommandIndex() int32 {
	if o == nil || isNil(o.TimerCommandIndex) {
		var ret int32
		return ret
	}
	return *o.TimerCommandIndex
}

// GetTimerCommandIndexOk returns a tuple with the TimerCommandIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSkipTimerRequest) GetTimerCommandIndexOk() (*int32, bool) {
	if o == nil || isNil(o.TimerCommandIndex) {
    return nil, false
	}
	return o.TimerCommandIndex, true
}

// HasTimerCommandIndex returns a boolean if a field has been set.
func (o *WorkflowSkipTimerRequest) HasTimerCommandIndex() bool {
	if o != nil && !isNil(o.TimerCommandIndex) {
		return true
	}

	return false
}

// SetTimerCommandIndex gets a reference to the given int32 and assigns it to the TimerCommandIndex field.
func (o *WorkflowSkipTimerRequest) SetTimerCommandIndex(v int32) {
	o.TimerCommandIndex = &v
}

func (o WorkflowSkipTimerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if !isNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if true {
		toSerialize["workflowStateExecutionId"] = o.WorkflowStateExecutionId
	}
	if !isNil(o.TimerCommandId) {
		toSerialize["timerCommandId"] = o.TimerCommandId
	}
	if !isNil(o.TimerCommandIndex) {
		toSerialize["timerCommandIndex"] = o.TimerCommandIndex
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSkipTimerRequest struct {
	value *WorkflowSkipTimerRequest
	isSet bool
}

func (v NullableWorkflowSkipTimerRequest) Get() *WorkflowSkipTimerRequest {
	return v.value
}

func (v *NullableWorkflowSkipTimerRequest) Set(val *WorkflowSkipTimerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSkipTimerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSkipTimerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSkipTimerRequest(val *WorkflowSkipTimerRequest) *NullableWorkflowSkipTimerRequest {
	return &NullableWorkflowSkipTimerRequest{value: val, isSet: true}
}

func (v NullableWorkflowSkipTimerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSkipTimerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


