/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_slurm_24_05

import (
	"encoding/json"
)

// checks if the SlurmV0041DeleteJobsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041DeleteJobsRequest{}

// SlurmV0041DeleteJobsRequest struct for SlurmV0041DeleteJobsRequest
type SlurmV0041DeleteJobsRequest struct {
	// Filter jobs to a specific account
	Account *string `json:"account,omitempty"`
	// Filter jobs according to flags
	Flags []string `json:"flags,omitempty"`
	// Filter jobs to a specific name
	JobName *string `json:"job_name,omitempty"`
	// List of jobs to signal
	Jobs []string `json:"jobs,omitempty"`
	// Filter jobs to a specific partition
	Partition *string `json:"partition,omitempty"`
	// Filter jobs to a specific QOS
	Qos *string `json:"qos,omitempty"`
	// Filter jobs to a specific reservation
	Reservation *string `json:"reservation,omitempty"`
	// Signal to send to jobs
	Signal *string `json:"signal,omitempty"`
	// Filter jobs to a specific state
	JobState []string `json:"job_state,omitempty"`
	// Filter jobs to a specific numeric user id
	UserId *string `json:"user_id,omitempty"`
	// Filter jobs to a specific user name
	UserName *string `json:"user_name,omitempty"`
	// Filter jobs to a specific wckey
	Wckey *string `json:"wckey,omitempty"`
	// Filter jobs to a set of nodes
	Nodes []string `json:"nodes,omitempty"`
}

// NewSlurmV0041DeleteJobsRequest instantiates a new SlurmV0041DeleteJobsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041DeleteJobsRequest() *SlurmV0041DeleteJobsRequest {
	this := SlurmV0041DeleteJobsRequest{}
	return &this
}

// NewSlurmV0041DeleteJobsRequestWithDefaults instantiates a new SlurmV0041DeleteJobsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041DeleteJobsRequestWithDefaults() *SlurmV0041DeleteJobsRequest {
	this := SlurmV0041DeleteJobsRequest{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *SlurmV0041DeleteJobsRequest) SetAccount(v string) {
	o.Account = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *SlurmV0041DeleteJobsRequest) SetFlags(v []string) {
	o.Flags = v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetJobName() string {
	if o == nil || IsNil(o.JobName) {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetJobNameOk() (*string, bool) {
	if o == nil || IsNil(o.JobName) {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasJobName() bool {
	if o != nil && !IsNil(o.JobName) {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *SlurmV0041DeleteJobsRequest) SetJobName(v string) {
	o.JobName = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetJobs() []string {
	if o == nil || IsNil(o.Jobs) {
		var ret []string
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetJobsOk() ([]string, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []string and assigns it to the Jobs field.
func (o *SlurmV0041DeleteJobsRequest) SetJobs(v []string) {
	o.Jobs = v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetPartition() string {
	if o == nil || IsNil(o.Partition) {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetPartitionOk() (*string, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *SlurmV0041DeleteJobsRequest) SetPartition(v string) {
	o.Partition = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetQos() string {
	if o == nil || IsNil(o.Qos) {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetQosOk() (*string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *SlurmV0041DeleteJobsRequest) SetQos(v string) {
	o.Qos = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetReservation() string {
	if o == nil || IsNil(o.Reservation) {
		var ret string
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetReservationOk() (*string, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given string and assigns it to the Reservation field.
func (o *SlurmV0041DeleteJobsRequest) SetReservation(v string) {
	o.Reservation = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetSignal() string {
	if o == nil || IsNil(o.Signal) {
		var ret string
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetSignalOk() (*string, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given string and assigns it to the Signal field.
func (o *SlurmV0041DeleteJobsRequest) SetSignal(v string) {
	o.Signal = &v
}

// GetJobState returns the JobState field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetJobState() []string {
	if o == nil || IsNil(o.JobState) {
		var ret []string
		return ret
	}
	return o.JobState
}

// GetJobStateOk returns a tuple with the JobState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetJobStateOk() ([]string, bool) {
	if o == nil || IsNil(o.JobState) {
		return nil, false
	}
	return o.JobState, true
}

// HasJobState returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasJobState() bool {
	if o != nil && !IsNil(o.JobState) {
		return true
	}

	return false
}

// SetJobState gets a reference to the given []string and assigns it to the JobState field.
func (o *SlurmV0041DeleteJobsRequest) SetJobState(v []string) {
	o.JobState = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SlurmV0041DeleteJobsRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SlurmV0041DeleteJobsRequest) SetUserName(v string) {
	o.UserName = &v
}

// GetWckey returns the Wckey field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetWckey() string {
	if o == nil || IsNil(o.Wckey) {
		var ret string
		return ret
	}
	return *o.Wckey
}

// GetWckeyOk returns a tuple with the Wckey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetWckeyOk() (*string, bool) {
	if o == nil || IsNil(o.Wckey) {
		return nil, false
	}
	return o.Wckey, true
}

// HasWckey returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasWckey() bool {
	if o != nil && !IsNil(o.Wckey) {
		return true
	}

	return false
}

// SetWckey gets a reference to the given string and assigns it to the Wckey field.
func (o *SlurmV0041DeleteJobsRequest) SetWckey(v string) {
	o.Wckey = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *SlurmV0041DeleteJobsRequest) GetNodes() []string {
	if o == nil || IsNil(o.Nodes) {
		var ret []string
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041DeleteJobsRequest) GetNodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *SlurmV0041DeleteJobsRequest) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []string and assigns it to the Nodes field.
func (o *SlurmV0041DeleteJobsRequest) SetNodes(v []string) {
	o.Nodes = v
}

func (o SlurmV0041DeleteJobsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041DeleteJobsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.JobName) {
		toSerialize["job_name"] = o.JobName
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	if !IsNil(o.JobState) {
		toSerialize["job_state"] = o.JobState
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.Wckey) {
		toSerialize["wckey"] = o.Wckey
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	return toSerialize, nil
}

type NullableSlurmV0041DeleteJobsRequest struct {
	value *SlurmV0041DeleteJobsRequest
	isSet bool
}

func (v NullableSlurmV0041DeleteJobsRequest) Get() *SlurmV0041DeleteJobsRequest {
	return v.value
}

func (v *NullableSlurmV0041DeleteJobsRequest) Set(val *SlurmV0041DeleteJobsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041DeleteJobsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041DeleteJobsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041DeleteJobsRequest(val *SlurmV0041DeleteJobsRequest) *NullableSlurmV0041DeleteJobsRequest {
	return &NullableSlurmV0041DeleteJobsRequest{value: val, isSet: true}
}

func (v NullableSlurmV0041DeleteJobsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041DeleteJobsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


