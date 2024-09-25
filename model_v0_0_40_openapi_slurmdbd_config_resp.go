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

// checks if the V0040OpenapiSlurmdbdConfigResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiSlurmdbdConfigResp{}

// V0040OpenapiSlurmdbdConfigResp struct for V0040OpenapiSlurmdbdConfigResp
type V0040OpenapiSlurmdbdConfigResp struct {
	Clusters []V0040ClusterRec `json:"clusters,omitempty"`
	Tres []V0040Tres `json:"tres,omitempty"`
	Accounts []V0040Account `json:"accounts,omitempty"`
	Users []V0040User `json:"users,omitempty"`
	Qos []V0040Qos `json:"qos,omitempty"`
	Wckeys []V0040Wckey `json:"wckeys,omitempty"`
	Associations []V0040Assoc `json:"associations,omitempty"`
	Instances []V0040Instance `json:"instances,omitempty"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

// NewV0040OpenapiSlurmdbdConfigResp instantiates a new V0040OpenapiSlurmdbdConfigResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiSlurmdbdConfigResp() *V0040OpenapiSlurmdbdConfigResp {
	this := V0040OpenapiSlurmdbdConfigResp{}
	return &this
}

// NewV0040OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0040OpenapiSlurmdbdConfigResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiSlurmdbdConfigRespWithDefaults() *V0040OpenapiSlurmdbdConfigResp {
	this := V0040OpenapiSlurmdbdConfigResp{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetClusters() []V0040ClusterRec {
	if o == nil || IsNil(o.Clusters) {
		var ret []V0040ClusterRec
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetClustersOk() ([]V0040ClusterRec, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V0040ClusterRec and assigns it to the Clusters field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetClusters(v []V0040ClusterRec) {
	o.Clusters = v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetTres() []V0040Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0040Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetTresOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0040Tres and assigns it to the Tres field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetTres(v []V0040Tres) {
	o.Tres = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetAccounts() []V0040Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []V0040Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetAccountsOk() ([]V0040Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []V0040Account and assigns it to the Accounts field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetAccounts(v []V0040Account) {
	o.Accounts = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetUsers() []V0040User {
	if o == nil || IsNil(o.Users) {
		var ret []V0040User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetUsersOk() ([]V0040User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V0040User and assigns it to the Users field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetUsers(v []V0040User) {
	o.Users = v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetQos() []V0040Qos {
	if o == nil || IsNil(o.Qos) {
		var ret []V0040Qos
		return ret
	}
	return o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetQosOk() ([]V0040Qos, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given []V0040Qos and assigns it to the Qos field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetQos(v []V0040Qos) {
	o.Qos = v
}

// GetWckeys returns the Wckeys field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetWckeys() []V0040Wckey {
	if o == nil || IsNil(o.Wckeys) {
		var ret []V0040Wckey
		return ret
	}
	return o.Wckeys
}

// GetWckeysOk returns a tuple with the Wckeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetWckeysOk() ([]V0040Wckey, bool) {
	if o == nil || IsNil(o.Wckeys) {
		return nil, false
	}
	return o.Wckeys, true
}

// HasWckeys returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasWckeys() bool {
	if o != nil && !IsNil(o.Wckeys) {
		return true
	}

	return false
}

// SetWckeys gets a reference to the given []V0040Wckey and assigns it to the Wckeys field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetWckeys(v []V0040Wckey) {
	o.Wckeys = v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetAssociations() []V0040Assoc {
	if o == nil || IsNil(o.Associations) {
		var ret []V0040Assoc
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetAssociationsOk() ([]V0040Assoc, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []V0040Assoc and assigns it to the Associations field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetAssociations(v []V0040Assoc) {
	o.Associations = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetInstances() []V0040Instance {
	if o == nil || IsNil(o.Instances) {
		var ret []V0040Instance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetInstancesOk() ([]V0040Instance, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []V0040Instance and assigns it to the Instances field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetInstances(v []V0040Instance) {
	o.Instances = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiSlurmdbdConfigResp) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiSlurmdbdConfigResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiSlurmdbdConfigResp) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiSlurmdbdConfigResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiSlurmdbdConfigResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Wckeys) {
		toSerialize["wckeys"] = o.Wckeys
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableV0040OpenapiSlurmdbdConfigResp struct {
	value *V0040OpenapiSlurmdbdConfigResp
	isSet bool
}

func (v NullableV0040OpenapiSlurmdbdConfigResp) Get() *V0040OpenapiSlurmdbdConfigResp {
	return v.value
}

func (v *NullableV0040OpenapiSlurmdbdConfigResp) Set(val *V0040OpenapiSlurmdbdConfigResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiSlurmdbdConfigResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiSlurmdbdConfigResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiSlurmdbdConfigResp(val *V0040OpenapiSlurmdbdConfigResp) *NullableV0040OpenapiSlurmdbdConfigResp {
	return &NullableV0040OpenapiSlurmdbdConfigResp{value: val, isSet: true}
}

func (v NullableV0040OpenapiSlurmdbdConfigResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiSlurmdbdConfigResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


