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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer struct {
	// GrpTRESRunMins
	Qos []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"qos,omitempty"`
	// MaxTRESMinsPerJob
	Job []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"job,omitempty"`
	// MaxTRESRunMinsPerAccount
	Account []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"account,omitempty"`
	// MaxTRESRunMinsPerUser
	User []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner `json:"user,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetQos() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Qos) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetQosOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Qos field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) SetQos(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Qos = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetJob() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Job) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetJobOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Job field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) SetJob(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Job = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetAccount() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.Account) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetAccountOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the Account field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) SetAccount(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.Account = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetUser() []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner {
	if o == nil || IsNil(o.User) {
		var ret []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) GetUserOk() ([]SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner and assigns it to the User field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) SetUser(v []SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner) {
	o.User = v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresMinutesPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


