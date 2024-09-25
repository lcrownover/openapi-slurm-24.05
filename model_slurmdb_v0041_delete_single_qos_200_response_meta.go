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

// checks if the SlurmdbV0041DeleteSingleQos200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041DeleteSingleQos200ResponseMeta{}

// SlurmdbV0041DeleteSingleQos200ResponseMeta Slurm meta values
type SlurmdbV0041DeleteSingleQos200ResponseMeta struct {
	Plugin *SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin `json:"plugin,omitempty"`
	Client *SlurmdbV0041DeleteSingleQos200ResponseMetaClient `json:"client,omitempty"`
	// CLI command (if applicable)
	Command []string `json:"command,omitempty"`
	Slurm *SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm `json:"slurm,omitempty"`
}

// NewSlurmdbV0041DeleteSingleQos200ResponseMeta instantiates a new SlurmdbV0041DeleteSingleQos200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041DeleteSingleQos200ResponseMeta() *SlurmdbV0041DeleteSingleQos200ResponseMeta {
	this := SlurmdbV0041DeleteSingleQos200ResponseMeta{}
	return &this
}

// NewSlurmdbV0041DeleteSingleQos200ResponseMetaWithDefaults instantiates a new SlurmdbV0041DeleteSingleQos200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041DeleteSingleQos200ResponseMetaWithDefaults() *SlurmdbV0041DeleteSingleQos200ResponseMeta {
	this := SlurmdbV0041DeleteSingleQos200ResponseMeta{}
	return &this
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetPlugin() SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin {
	if o == nil || IsNil(o.Plugin) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetPluginOk() (*SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin, bool) {
	if o == nil || IsNil(o.Plugin) {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) HasPlugin() bool {
	if o != nil && !IsNil(o.Plugin) {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin and assigns it to the Plugin field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) SetPlugin(v SlurmdbV0041DeleteSingleQos200ResponseMetaPlugin) {
	o.Plugin = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetClient() SlurmdbV0041DeleteSingleQos200ResponseMetaClient {
	if o == nil || IsNil(o.Client) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMetaClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetClientOk() (*SlurmdbV0041DeleteSingleQos200ResponseMetaClient, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMetaClient and assigns it to the Client field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) SetClient(v SlurmdbV0041DeleteSingleQos200ResponseMetaClient) {
	o.Client = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetCommand() []string {
	if o == nil || IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) SetCommand(v []string) {
	o.Command = v
}

// GetSlurm returns the Slurm field value if set, zero value otherwise.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetSlurm() SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm {
	if o == nil || IsNil(o.Slurm) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm
		return ret
	}
	return *o.Slurm
}

// GetSlurmOk returns a tuple with the Slurm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) GetSlurmOk() (*SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm, bool) {
	if o == nil || IsNil(o.Slurm) {
		return nil, false
	}
	return o.Slurm, true
}

// HasSlurm returns a boolean if a field has been set.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) HasSlurm() bool {
	if o != nil && !IsNil(o.Slurm) {
		return true
	}

	return false
}

// SetSlurm gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm and assigns it to the Slurm field.
func (o *SlurmdbV0041DeleteSingleQos200ResponseMeta) SetSlurm(v SlurmdbV0041DeleteSingleQos200ResponseMetaSlurm) {
	o.Slurm = &v
}

func (o SlurmdbV0041DeleteSingleQos200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041DeleteSingleQos200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plugin) {
		toSerialize["plugin"] = o.Plugin
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Slurm) {
		toSerialize["slurm"] = o.Slurm
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041DeleteSingleQos200ResponseMeta struct {
	value *SlurmdbV0041DeleteSingleQos200ResponseMeta
	isSet bool
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) Get() *SlurmdbV0041DeleteSingleQos200ResponseMeta {
	return v.value
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) Set(val *SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041DeleteSingleQos200ResponseMeta(val *SlurmdbV0041DeleteSingleQos200ResponseMeta) *NullableSlurmdbV0041DeleteSingleQos200ResponseMeta {
	return &NullableSlurmdbV0041DeleteSingleQos200ResponseMeta{value: val, isSet: true}
}

func (v NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041DeleteSingleQos200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


