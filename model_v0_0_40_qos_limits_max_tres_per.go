/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0040QosLimitsMaxTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMaxTresPer{}

// V0040QosLimitsMaxTresPer struct for V0040QosLimitsMaxTresPer
type V0040QosLimitsMaxTresPer struct {
	Account []V0040Tres `json:"account,omitempty"`
	Job []V0040Tres `json:"job,omitempty"`
	Node []V0040Tres `json:"node,omitempty"`
	User []V0040Tres `json:"user,omitempty"`
}

// NewV0040QosLimitsMaxTresPer instantiates a new V0040QosLimitsMaxTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMaxTresPer() *V0040QosLimitsMaxTresPer {
	this := V0040QosLimitsMaxTresPer{}
	return &this
}

// NewV0040QosLimitsMaxTresPerWithDefaults instantiates a new V0040QosLimitsMaxTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxTresPerWithDefaults() *V0040QosLimitsMaxTresPer {
	this := V0040QosLimitsMaxTresPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTresPer) GetAccount() []V0040Tres {
	if o == nil || IsNil(o.Account) {
		var ret []V0040Tres
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTresPer) GetAccountOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTresPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []V0040Tres and assigns it to the Account field.
func (o *V0040QosLimitsMaxTresPer) SetAccount(v []V0040Tres) {
	o.Account = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTresPer) GetJob() []V0040Tres {
	if o == nil || IsNil(o.Job) {
		var ret []V0040Tres
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTresPer) GetJobOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0040Tres and assigns it to the Job field.
func (o *V0040QosLimitsMaxTresPer) SetJob(v []V0040Tres) {
	o.Job = v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTresPer) GetNode() []V0040Tres {
	if o == nil || IsNil(o.Node) {
		var ret []V0040Tres
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTresPer) GetNodeOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTresPer) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given []V0040Tres and assigns it to the Node field.
func (o *V0040QosLimitsMaxTresPer) SetNode(v []V0040Tres) {
	o.Node = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTresPer) GetUser() []V0040Tres {
	if o == nil || IsNil(o.User) {
		var ret []V0040Tres
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTresPer) GetUserOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTresPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []V0040Tres and assigns it to the User field.
func (o *V0040QosLimitsMaxTresPer) SetUser(v []V0040Tres) {
	o.User = v
}

func (o V0040QosLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMaxTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMaxTresPer struct {
	value *V0040QosLimitsMaxTresPer
	isSet bool
}

func (v NullableV0040QosLimitsMaxTresPer) Get() *V0040QosLimitsMaxTresPer {
	return v.value
}

func (v *NullableV0040QosLimitsMaxTresPer) Set(val *V0040QosLimitsMaxTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMaxTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMaxTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMaxTresPer(val *V0040QosLimitsMaxTresPer) *NullableV0040QosLimitsMaxTresPer {
	return &NullableV0040QosLimitsMaxTresPer{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMaxTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


