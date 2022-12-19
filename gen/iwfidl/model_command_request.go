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

// CommandRequest struct for CommandRequest
type CommandRequest struct {
	DeciderTriggerType string `json:"deciderTriggerType"`
	TimerCommands []TimerCommand `json:"timerCommands,omitempty"`
	SignalCommands []SignalCommand `json:"signalCommands,omitempty"`
	InterStateChannelCommands []InterStateChannelCommand `json:"interStateChannelCommands,omitempty"`
}

// NewCommandRequest instantiates a new CommandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandRequest(deciderTriggerType string) *CommandRequest {
	this := CommandRequest{}
	this.DeciderTriggerType = deciderTriggerType
	return &this
}

// NewCommandRequestWithDefaults instantiates a new CommandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandRequestWithDefaults() *CommandRequest {
	this := CommandRequest{}
	return &this
}

// GetDeciderTriggerType returns the DeciderTriggerType field value
func (o *CommandRequest) GetDeciderTriggerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeciderTriggerType
}

// GetDeciderTriggerTypeOk returns a tuple with the DeciderTriggerType field value
// and a boolean to check if the value has been set.
func (o *CommandRequest) GetDeciderTriggerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeciderTriggerType, true
}

// SetDeciderTriggerType sets field value
func (o *CommandRequest) SetDeciderTriggerType(v string) {
	o.DeciderTriggerType = v
}

// GetTimerCommands returns the TimerCommands field value if set, zero value otherwise.
func (o *CommandRequest) GetTimerCommands() []TimerCommand {
	if o == nil || o.TimerCommands == nil {
		var ret []TimerCommand
		return ret
	}
	return o.TimerCommands
}

// GetTimerCommandsOk returns a tuple with the TimerCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandRequest) GetTimerCommandsOk() ([]TimerCommand, bool) {
	if o == nil || o.TimerCommands == nil {
		return nil, false
	}
	return o.TimerCommands, true
}

// HasTimerCommands returns a boolean if a field has been set.
func (o *CommandRequest) HasTimerCommands() bool {
	if o != nil && o.TimerCommands != nil {
		return true
	}

	return false
}

// SetTimerCommands gets a reference to the given []TimerCommand and assigns it to the TimerCommands field.
func (o *CommandRequest) SetTimerCommands(v []TimerCommand) {
	o.TimerCommands = v
}

// GetSignalCommands returns the SignalCommands field value if set, zero value otherwise.
func (o *CommandRequest) GetSignalCommands() []SignalCommand {
	if o == nil || o.SignalCommands == nil {
		var ret []SignalCommand
		return ret
	}
	return o.SignalCommands
}

// GetSignalCommandsOk returns a tuple with the SignalCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandRequest) GetSignalCommandsOk() ([]SignalCommand, bool) {
	if o == nil || o.SignalCommands == nil {
		return nil, false
	}
	return o.SignalCommands, true
}

// HasSignalCommands returns a boolean if a field has been set.
func (o *CommandRequest) HasSignalCommands() bool {
	if o != nil && o.SignalCommands != nil {
		return true
	}

	return false
}

// SetSignalCommands gets a reference to the given []SignalCommand and assigns it to the SignalCommands field.
func (o *CommandRequest) SetSignalCommands(v []SignalCommand) {
	o.SignalCommands = v
}

// GetInterStateChannelCommands returns the InterStateChannelCommands field value if set, zero value otherwise.
func (o *CommandRequest) GetInterStateChannelCommands() []InterStateChannelCommand {
	if o == nil || o.InterStateChannelCommands == nil {
		var ret []InterStateChannelCommand
		return ret
	}
	return o.InterStateChannelCommands
}

// GetInterStateChannelCommandsOk returns a tuple with the InterStateChannelCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandRequest) GetInterStateChannelCommandsOk() ([]InterStateChannelCommand, bool) {
	if o == nil || o.InterStateChannelCommands == nil {
		return nil, false
	}
	return o.InterStateChannelCommands, true
}

// HasInterStateChannelCommands returns a boolean if a field has been set.
func (o *CommandRequest) HasInterStateChannelCommands() bool {
	if o != nil && o.InterStateChannelCommands != nil {
		return true
	}

	return false
}

// SetInterStateChannelCommands gets a reference to the given []InterStateChannelCommand and assigns it to the InterStateChannelCommands field.
func (o *CommandRequest) SetInterStateChannelCommands(v []InterStateChannelCommand) {
	o.InterStateChannelCommands = v
}

func (o CommandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deciderTriggerType"] = o.DeciderTriggerType
	}
	if o.TimerCommands != nil {
		toSerialize["timerCommands"] = o.TimerCommands
	}
	if o.SignalCommands != nil {
		toSerialize["signalCommands"] = o.SignalCommands
	}
	if o.InterStateChannelCommands != nil {
		toSerialize["interStateChannelCommands"] = o.InterStateChannelCommands
	}
	return json.Marshal(toSerialize)
}

type NullableCommandRequest struct {
	value *CommandRequest
	isSet bool
}

func (v NullableCommandRequest) Get() *CommandRequest {
	return v.value
}

func (v *NullableCommandRequest) Set(val *CommandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandRequest(val *CommandRequest) *NullableCommandRequest {
	return &NullableCommandRequest{value: val, isSet: true}
}

func (v NullableCommandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

