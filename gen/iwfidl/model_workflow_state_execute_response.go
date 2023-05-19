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

// checks if the WorkflowStateExecuteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateExecuteResponse{}

// WorkflowStateExecuteResponse struct for WorkflowStateExecuteResponse
type WorkflowStateExecuteResponse struct {
	StateDecision              *StateDecision                `json:"stateDecision,omitempty"`
	UpsertSearchAttributes     []SearchAttribute             `json:"upsertSearchAttributes,omitempty"`
	UpsertDataObjects          []KeyValue                    `json:"upsertDataObjects,omitempty"`
	RecordEvents               []KeyValue                    `json:"recordEvents,omitempty"`
	UpsertStateLocals          []KeyValue                    `json:"upsertStateLocals,omitempty"`
	PublishToInterStateChannel []InterStateChannelPublishing `json:"publishToInterStateChannel,omitempty"`
}

// NewWorkflowStateExecuteResponse instantiates a new WorkflowStateExecuteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateExecuteResponse() *WorkflowStateExecuteResponse {
	this := WorkflowStateExecuteResponse{}
	return &this
}

// NewWorkflowStateExecuteResponseWithDefaults instantiates a new WorkflowStateExecuteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateExecuteResponseWithDefaults() *WorkflowStateExecuteResponse {
	this := WorkflowStateExecuteResponse{}
	return &this
}

// GetStateDecision returns the StateDecision field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetStateDecision() StateDecision {
	if o == nil || IsNil(o.StateDecision) {
		var ret StateDecision
		return ret
	}
	return *o.StateDecision
}

// GetStateDecisionOk returns a tuple with the StateDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetStateDecisionOk() (*StateDecision, bool) {
	if o == nil || IsNil(o.StateDecision) {
		return nil, false
	}
	return o.StateDecision, true
}

// HasStateDecision returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasStateDecision() bool {
	if o != nil && !IsNil(o.StateDecision) {
		return true
	}

	return false
}

// SetStateDecision gets a reference to the given StateDecision and assigns it to the StateDecision field.
func (o *WorkflowStateExecuteResponse) SetStateDecision(v StateDecision) {
	o.StateDecision = &v
}

// GetUpsertSearchAttributes returns the UpsertSearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetUpsertSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.UpsertSearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.UpsertSearchAttributes
}

// GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetUpsertSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.UpsertSearchAttributes) {
		return nil, false
	}
	return o.UpsertSearchAttributes, true
}

// HasUpsertSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasUpsertSearchAttributes() bool {
	if o != nil && !IsNil(o.UpsertSearchAttributes) {
		return true
	}

	return false
}

// SetUpsertSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the UpsertSearchAttributes field.
func (o *WorkflowStateExecuteResponse) SetUpsertSearchAttributes(v []SearchAttribute) {
	o.UpsertSearchAttributes = v
}

// GetUpsertDataObjects returns the UpsertDataObjects field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetUpsertDataObjects() []KeyValue {
	if o == nil || IsNil(o.UpsertDataObjects) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertDataObjects
}

// GetUpsertDataObjectsOk returns a tuple with the UpsertDataObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetUpsertDataObjectsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.UpsertDataObjects) {
		return nil, false
	}
	return o.UpsertDataObjects, true
}

// HasUpsertDataObjects returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasUpsertDataObjects() bool {
	if o != nil && !IsNil(o.UpsertDataObjects) {
		return true
	}

	return false
}

// SetUpsertDataObjects gets a reference to the given []KeyValue and assigns it to the UpsertDataObjects field.
func (o *WorkflowStateExecuteResponse) SetUpsertDataObjects(v []KeyValue) {
	o.UpsertDataObjects = v
}

// GetRecordEvents returns the RecordEvents field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetRecordEvents() []KeyValue {
	if o == nil || IsNil(o.RecordEvents) {
		var ret []KeyValue
		return ret
	}
	return o.RecordEvents
}

// GetRecordEventsOk returns a tuple with the RecordEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetRecordEventsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.RecordEvents) {
		return nil, false
	}
	return o.RecordEvents, true
}

// HasRecordEvents returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasRecordEvents() bool {
	if o != nil && !IsNil(o.RecordEvents) {
		return true
	}

	return false
}

// SetRecordEvents gets a reference to the given []KeyValue and assigns it to the RecordEvents field.
func (o *WorkflowStateExecuteResponse) SetRecordEvents(v []KeyValue) {
	o.RecordEvents = v
}

// GetUpsertStateLocals returns the UpsertStateLocals field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetUpsertStateLocals() []KeyValue {
	if o == nil || IsNil(o.UpsertStateLocals) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertStateLocals
}

// GetUpsertStateLocalsOk returns a tuple with the UpsertStateLocals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetUpsertStateLocalsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.UpsertStateLocals) {
		return nil, false
	}
	return o.UpsertStateLocals, true
}

// HasUpsertStateLocals returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasUpsertStateLocals() bool {
	if o != nil && !IsNil(o.UpsertStateLocals) {
		return true
	}

	return false
}

// SetUpsertStateLocals gets a reference to the given []KeyValue and assigns it to the UpsertStateLocals field.
func (o *WorkflowStateExecuteResponse) SetUpsertStateLocals(v []KeyValue) {
	o.UpsertStateLocals = v
}

// GetPublishToInterStateChannel returns the PublishToInterStateChannel field value if set, zero value otherwise.
func (o *WorkflowStateExecuteResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing {
	if o == nil || IsNil(o.PublishToInterStateChannel) {
		var ret []InterStateChannelPublishing
		return ret
	}
	return o.PublishToInterStateChannel
}

// GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteResponse) GetPublishToInterStateChannelOk() ([]InterStateChannelPublishing, bool) {
	if o == nil || IsNil(o.PublishToInterStateChannel) {
		return nil, false
	}
	return o.PublishToInterStateChannel, true
}

// HasPublishToInterStateChannel returns a boolean if a field has been set.
func (o *WorkflowStateExecuteResponse) HasPublishToInterStateChannel() bool {
	if o != nil && !IsNil(o.PublishToInterStateChannel) {
		return true
	}

	return false
}

// SetPublishToInterStateChannel gets a reference to the given []InterStateChannelPublishing and assigns it to the PublishToInterStateChannel field.
func (o *WorkflowStateExecuteResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing) {
	o.PublishToInterStateChannel = v
}

func (o WorkflowStateExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateExecuteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StateDecision) {
		toSerialize["stateDecision"] = o.StateDecision
	}
	if !IsNil(o.UpsertSearchAttributes) {
		toSerialize["upsertSearchAttributes"] = o.UpsertSearchAttributes
	}
	if !IsNil(o.UpsertDataObjects) {
		toSerialize["upsertDataObjects"] = o.UpsertDataObjects
	}
	if !IsNil(o.RecordEvents) {
		toSerialize["recordEvents"] = o.RecordEvents
	}
	if !IsNil(o.UpsertStateLocals) {
		toSerialize["upsertStateLocals"] = o.UpsertStateLocals
	}
	if !IsNil(o.PublishToInterStateChannel) {
		toSerialize["publishToInterStateChannel"] = o.PublishToInterStateChannel
	}
	return toSerialize, nil
}

type NullableWorkflowStateExecuteResponse struct {
	value *WorkflowStateExecuteResponse
	isSet bool
}

func (v NullableWorkflowStateExecuteResponse) Get() *WorkflowStateExecuteResponse {
	return v.value
}

func (v *NullableWorkflowStateExecuteResponse) Set(val *WorkflowStateExecuteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateExecuteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateExecuteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateExecuteResponse(val *WorkflowStateExecuteResponse) *NullableWorkflowStateExecuteResponse {
	return &NullableWorkflowStateExecuteResponse{value: val, isSet: true}
}

func (v NullableWorkflowStateExecuteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateExecuteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
