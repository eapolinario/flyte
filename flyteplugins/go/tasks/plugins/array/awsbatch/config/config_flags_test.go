// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeRaw_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_jobStoreConfig.jacheSize", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("jobStoreConfig.jacheSize", testValue)
			if vInt, err := cmdFlags.GetInt("jobStoreConfig.jacheSize"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.JobStoreConfig.CacheSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_jobStoreConfig.parallelizm", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("jobStoreConfig.parallelizm", testValue)
			if vInt, err := cmdFlags.GetInt("jobStoreConfig.parallelizm"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.JobStoreConfig.Parallelizm)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_jobStoreConfig.batchChunkSize", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("jobStoreConfig.batchChunkSize", testValue)
			if vInt, err := cmdFlags.GetInt("jobStoreConfig.batchChunkSize"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.JobStoreConfig.BatchChunkSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_jobStoreConfig.resyncPeriod", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.JobStoreConfig.ResyncPeriod.String()

			cmdFlags.Set("jobStoreConfig.resyncPeriod", testValue)
			if vString, err := cmdFlags.GetString("jobStoreConfig.resyncPeriod"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.JobStoreConfig.ResyncPeriod)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defCacheSize", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("defCacheSize", testValue)
			if vInt, err := cmdFlags.GetInt("defCacheSize"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.JobDefCacheSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_getRateLimiter.rate", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("getRateLimiter.rate", testValue)
			if vInt64, err := cmdFlags.GetInt64("getRateLimiter.rate"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.GetRateLimiter.Rate)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_getRateLimiter.burst", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("getRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("getRateLimiter.burst"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.GetRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defaultRateLimiter.rate", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("defaultRateLimiter.rate", testValue)
			if vInt64, err := cmdFlags.GetInt64("defaultRateLimiter.rate"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.DefaultRateLimiter.Rate)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defaultRateLimiter.burst", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("defaultRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("defaultRateLimiter.burst"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.DefaultRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxArrayJobSize", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("maxArrayJobSize", testValue)
			if vInt64, err := cmdFlags.GetInt64("maxArrayJobSize"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt64), &actual.MaxArrayJobSize)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_minRetries", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("minRetries", testValue)
			if vInt32, err := cmdFlags.GetInt32("minRetries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt32), &actual.MinRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxRetries", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("maxRetries", testValue)
			if vInt32, err := cmdFlags.GetInt32("maxRetries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt32), &actual.MaxRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defaultTimeout", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.DefaultTimeOut.String()

			cmdFlags.Set("defaultTimeout", testValue)
			if vString, err := cmdFlags.GetString("defaultTimeout"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.DefaultTimeOut)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxErrLength", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("maxErrLength", testValue)
			if vInt, err := cmdFlags.GetInt("maxErrLength"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.MaxErrorStringLength)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_roleAnnotationKey", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("roleAnnotationKey", testValue)
			if vString, err := cmdFlags.GetString("roleAnnotationKey"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.RoleAnnotationKey)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_outputAssembler.workers", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("outputAssembler.workers", testValue)
			if vInt, err := cmdFlags.GetInt("outputAssembler.workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.OutputAssembler.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_outputAssembler.maxRetries", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("outputAssembler.maxRetries", testValue)
			if vInt, err := cmdFlags.GetInt("outputAssembler.maxRetries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.OutputAssembler.MaxRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_outputAssembler.maxItems", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("outputAssembler.maxItems", testValue)
			if vInt, err := cmdFlags.GetInt("outputAssembler.maxItems"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.OutputAssembler.IndexCacheMaxItems)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_errorAssembler.workers", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("errorAssembler.workers", testValue)
			if vInt, err := cmdFlags.GetInt("errorAssembler.workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ErrorAssembler.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_errorAssembler.maxRetries", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("errorAssembler.maxRetries", testValue)
			if vInt, err := cmdFlags.GetInt("errorAssembler.maxRetries"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ErrorAssembler.MaxRetries)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_errorAssembler.maxItems", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("errorAssembler.maxItems", testValue)
			if vInt, err := cmdFlags.GetInt("errorAssembler.maxItems"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ErrorAssembler.IndexCacheMaxItems)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
