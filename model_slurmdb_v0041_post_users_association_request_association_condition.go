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
	"bytes"
	"fmt"
)

// checks if the SlurmdbV0041PostUsersAssociationRequestAssociationCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041PostUsersAssociationRequestAssociationCondition{}

// SlurmdbV0041PostUsersAssociationRequestAssociationCondition Filters to select associations for users
type SlurmdbV0041PostUsersAssociationRequestAssociationCondition struct {
	// CSV accounts list
	Accounts []string `json:"accounts,omitempty"`
	Association *SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation `json:"association,omitempty"`
	// CSV clusters list
	Clusters []string `json:"clusters,omitempty"`
	// CSV partitions list
	Partitions []string `json:"partitions,omitempty"`
	// CSV users list
	Users []string `json:"users"`
	// CSV WCKeys list
	Wckeys []string `json:"wckeys,omitempty"`
}

type _SlurmdbV0041PostUsersAssociationRequestAssociationCondition SlurmdbV0041PostUsersAssociationRequestAssociationCondition

// NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition(users []string) *SlurmdbV0041PostUsersAssociationRequestAssociationCondition {
	this := SlurmdbV0041PostUsersAssociationRequestAssociationCondition{}
	this.Users = users
	return &this
}

// NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionWithDefaults instantiates a new SlurmdbV0041PostUsersAssociationRequestAssociationCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionWithDefaults() *SlurmdbV0041PostUsersAssociationRequestAssociationCondition {
	this := SlurmdbV0041PostUsersAssociationRequestAssociationCondition{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAccounts() []string {
	if o == nil || IsNil(o.Accounts) {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetAccounts(v []string) {
	o.Accounts = v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAssociation() SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation {
	if o == nil || IsNil(o.Association) {
		var ret SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetAssociationOk() (*SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation and assigns it to the Association field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetAssociation(v SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociation) {
	o.Association = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetClusters(v []string) {
	o.Clusters = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetPartitions() []string {
	if o == nil || IsNil(o.Partitions) {
		var ret []string
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetPartitionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []string and assigns it to the Partitions field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetPartitions(v []string) {
	o.Partitions = v
}

// GetUsers returns the Users field value
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetUsers(v []string) {
	o.Users = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetWckeys() []string {
	if o == nil || IsNil(o.Wckeys) {
		var ret []string
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) GetWckeysOk() ([]string, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []string and assigns it to the Wckeys field.
func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) SetWckeys(v []string) {
	o.Wckeys = v
}

func (o SlurmdbV0041PostUsersAssociationRequestAssociationCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041PostUsersAssociationRequestAssociationCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	toSerialize["users"] = o.Users
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	return toSerialize, nil
}

func (o *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSlurmdbV0041PostUsersAssociationRequestAssociationCondition := _SlurmdbV0041PostUsersAssociationRequestAssociationCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlurmdbV0041PostUsersAssociationRequestAssociationCondition)

	if err != nil {
		return err
	}

	*o = SlurmdbV0041PostUsersAssociationRequestAssociationCondition(varSlurmdbV0041PostUsersAssociationRequestAssociationCondition)

	return err
}

type NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition struct {
	value *SlurmdbV0041PostUsersAssociationRequestAssociationCondition
	isSet bool
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) Get() *SlurmdbV0041PostUsersAssociationRequestAssociationCondition {
	return v.value
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) Set(val *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition(val *SlurmdbV0041PostUsersAssociationRequestAssociationCondition) *NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition {
	return &NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition{value: val, isSet: true}
}

func (v NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041PostUsersAssociationRequestAssociationCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


