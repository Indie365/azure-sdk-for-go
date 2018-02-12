package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"reflect"
	"regexp"
	"strings"
)

// Constraint stores constraint name, target field name
// Rule and chain validations.
type constraint struct {
	// Target field name for validation.
	target string

	// Constraint name e.g. minLength, MaxLength, Pattern, etc.
	name string

	// Rule for constraint e.g. greater than 10, less than 5 etc.
	rule interface{}

	// Chain validations for struct type
	chain []constraint
}

// Validation stores parameter-wise validation.
type validation struct {
	targetValue interface{}
	constraints []constraint
}

// Constraint list
const (
	empty            = "Empty"
	null             = "Null"
	readOnly         = "ReadOnly"
	pattern          = "Pattern"
	maxLength        = "MaxLength"
	minLength        = "MinLength"
	maxItems         = "MaxItems"
	minItems         = "MinItems"
	multipleOf       = "MultipleOf"
	uniqueItems      = "UniqueItems"
	inclusiveMaximum = "InclusiveMaximum"
	exclusiveMaximum = "ExclusiveMaximum"
	exclusiveMinimum = "ExclusiveMinimum"
	inclusiveMinimum = "InclusiveMinimum"
)

// Validate method validates constraints on parameter
// passed in validation array.
func validate(m []validation) error {
	for _, item := range m {
		v := reflect.ValueOf(item.targetValue)
		for _, constraint := range item.constraints {
			var err error
			switch v.Kind() {
			case reflect.Ptr:
				err = validatePtr(v, constraint)
			case reflect.String:
				err = validateString(v, constraint)
			case reflect.Struct:
				err = validateStruct(v, constraint)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				err = validateInt(v, constraint)
			case reflect.Float32, reflect.Float64:
				err = validateFloat(v, constraint)
			case reflect.Array, reflect.Slice, reflect.Map:
				err = validateArrayMap(v, constraint)
			default:
				err = createError(v, constraint, fmt.Sprintf("unknown type %v", v.Kind()))
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateStruct(x reflect.Value, v constraint, name ...string) error {
	//Get field name from target name which is in format a.b.c
	s := strings.Split(v.target, ".")
	f := x.FieldByName(s[len(s)-1])
	if isZero(f) {
		return createError(x, v, fmt.Sprintf("field %q doesn't exist", v.target))
	}
	err := validate([]validation{
		{
			targetValue: getInterfaceValue(f),
			constraints: []constraint{v},
		},
	})
	return err
}

func validatePtr(x reflect.Value, v constraint) error {
	if v.name == readOnly {
		if !x.IsNil() {
			return createError(x.Elem(), v, "readonly parameter; must send as nil or empty in request")
		}
		return nil
	}
	if x.IsNil() {
		return checkNil(x, v)
	}
	if v.chain != nil {
		return validate([]validation{
			{
				targetValue: getInterfaceValue(x.Elem()),
				constraints: v.chain,
			},
		})
	}
	return nil
}

func validateInt(x reflect.Value, v constraint) error {
	i := x.Int()
	r, ok := v.rule.(int)
	if !ok {
		return createError(x, v, fmt.Sprintf("rule must be integer value for %v constraint; got: %v", v.name, v.rule))
	}
	switch v.name {
	case multipleOf:
		if i%int64(r) != 0 {
			return createError(x, v, fmt.Sprintf("value must be a multiple of %v", r))
		}
	case exclusiveMinimum:
		if i <= int64(r) {
			return createError(x, v, fmt.Sprintf("value must be greater than %v", r))
		}
	case exclusiveMaximum:
		if i >= int64(r) {
			return createError(x, v, fmt.Sprintf("value must be less than %v", r))
		}
	case inclusiveMinimum:
		if i < int64(r) {
			return createError(x, v, fmt.Sprintf("value must be greater than or equal to %v", r))
		}
	case inclusiveMaximum:
		if i > int64(r) {
			return createError(x, v, fmt.Sprintf("value must be less than or equal to %v", r))
		}
	default:
		return createError(x, v, fmt.Sprintf("constraint %v is not applicable for type integer", v.name))
	}
	return nil
}

func validateFloat(x reflect.Value, v constraint) error {
	f := x.Float()
	r, ok := v.rule.(float64)
	if !ok {
		return createError(x, v, fmt.Sprintf("rule must be float value for %v constraint; got: %v", v.name, v.rule))
	}
	switch v.name {
	case exclusiveMinimum:
		if f <= r {
			return createError(x, v, fmt.Sprintf("value must be greater than %v", r))
		}
	case exclusiveMaximum:
		if f >= r {
			return createError(x, v, fmt.Sprintf("value must be less than %v", r))
		}
	case inclusiveMinimum:
		if f < r {
			return createError(x, v, fmt.Sprintf("value must be greater than or equal to %v", r))
		}
	case inclusiveMaximum:
		if f > r {
			return createError(x, v, fmt.Sprintf("value must be less than or equal to %v", r))
		}
	default:
		return createError(x, v, fmt.Sprintf("constraint %s is not applicable for type float", v.name))
	}
	return nil
}

func validateString(x reflect.Value, v constraint) error {
	s := x.String()
	switch v.name {
	case empty:
		if len(s) == 0 {
			return checkEmpty(x, v)
		}
	case pattern:
		reg, err := regexp.Compile(v.rule.(string))
		if err != nil {
			return createError(x, v, err.Error())
		}
		if !reg.MatchString(s) {
			return createError(x, v, fmt.Sprintf("value doesn't match pattern %v", v.rule))
		}
	case maxLength:
		if _, ok := v.rule.(int); !ok {
			return createError(x, v, fmt.Sprintf("rule must be integer value for %v constraint; got: %v", v.name, v.rule))
		}
		if len(s) > v.rule.(int) {
			return createError(x, v, fmt.Sprintf("value length must be less than %v", v.rule))
		}
	case minLength:
		if _, ok := v.rule.(int); !ok {
			return createError(x, v, fmt.Sprintf("rule must be integer value for %v constraint; got: %v", v.name, v.rule))
		}
		if len(s) < v.rule.(int) {
			return createError(x, v, fmt.Sprintf("value length must be greater than %v", v.rule))
		}
	case readOnly:
		if len(s) > 0 {
			return createError(reflect.ValueOf(s), v, "readonly parameter; must send as nil or empty in request")
		}
	default:
		return createError(x, v, fmt.Sprintf("constraint %s is not applicable to string type", v.name))
	}
	if v.chain != nil {
		return validate([]validation{
			{
				targetValue: getInterfaceValue(x),
				constraints: v.chain,
			},
		})
	}
	return nil
}

func validateArrayMap(x reflect.Value, v constraint) error {
	switch v.name {
	case null:
		if x.IsNil() {
			return checkNil(x, v)
		}
	case empty:
		if x.IsNil() || x.Len() == 0 {
			return checkEmpty(x, v)
		}
	case maxItems:
		if _, ok := v.rule.(int); !ok {
			return createError(x, v, fmt.Sprintf("rule must be integer for %v constraint; got: %v", v.name, v.rule))
		}
		if x.Len() > v.rule.(int) {
			return createError(x, v, fmt.Sprintf("maximum item limit is %v; got: %v", v.rule, x.Len()))
		}
	case minItems:
		if _, ok := v.rule.(int); !ok {
			return createError(x, v, fmt.Sprintf("rule must be integer for %v constraint; got: %v", v.name, v.rule))
		}
		if x.Len() < v.rule.(int) {
			return createError(x, v, fmt.Sprintf("minimum item limit is %v; got: %v", v.rule, x.Len()))
		}
	case uniqueItems:
		if x.Kind() == reflect.Array || x.Kind() == reflect.Slice {
			if !checkForUniqueInArray(x) {
				return createError(x, v, fmt.Sprintf("all items in parameter %q must be unique; got:%v", v.target, x))
			}
		} else if x.Kind() == reflect.Map {
			if !checkForUniqueInMap(x) {
				return createError(x, v, fmt.Sprintf("all items in parameter %q must be unique; got:%v", v.target, x))
			}
		} else {
			return createError(x, v, fmt.Sprintf("type must be array, slice or map for constraint %v; got: %v", v.name, x.Kind()))
		}
	case readOnly:
		if x.Len() != 0 {
			return createError(x, v, "readonly parameter; must send as nil or empty in request")
		}
	case pattern:
		reg, err := regexp.Compile(v.rule.(string))
		if err != nil {
			return createError(x, v, err.Error())
		}
		keys := x.MapKeys()
		for _, k := range keys {
			if !reg.MatchString(k.String()) {
				return createError(k, v, fmt.Sprintf("map key doesn't match pattern %v", v.rule))
			}
		}
	default:
		return createError(x, v, fmt.Sprintf("constraint %v is not applicable to array, slice and map type", v.name))
	}
	if v.chain != nil {
		return validate([]validation{
			{
				targetValue: getInterfaceValue(x),
				constraints: v.chain,
			},
		})
	}
	return nil
}

func checkNil(x reflect.Value, v constraint) error {
	if _, ok := v.rule.(bool); !ok {
		return createError(x, v, fmt.Sprintf("rule must be bool value for %v constraint; got: %v", v.name, v.rule))
	}
	if v.rule.(bool) {
		return createError(x, v, "value can not be null; required parameter")
	}
	return nil
}

func checkEmpty(x reflect.Value, v constraint) error {
	if _, ok := v.rule.(bool); !ok {
		return createError(x, v, fmt.Sprintf("rule must be bool value for %v constraint; got: %v", v.name, v.rule))
	}
	if v.rule.(bool) {
		return createError(x, v, "value can not be null or empty; required parameter")
	}
	return nil
}

func checkForUniqueInArray(x reflect.Value) bool {
	if x == reflect.Zero(reflect.TypeOf(x)) || x.Len() == 0 {
		return false
	}
	arrOfInterface := make([]interface{}, x.Len())
	for i := 0; i < x.Len(); i++ {
		arrOfInterface[i] = x.Index(i).Interface()
	}
	m := make(map[interface{}]bool)
	for _, val := range arrOfInterface {
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}

func checkForUniqueInMap(x reflect.Value) bool {
	if x == reflect.Zero(reflect.TypeOf(x)) || x.Len() == 0 {
		return false
	}
	mapOfInterface := make(map[interface{}]interface{}, x.Len())
	keys := x.MapKeys()
	for _, k := range keys {
		mapOfInterface[k.Interface()] = x.MapIndex(k).Interface()
	}
	m := make(map[interface{}]bool)
	for _, val := range mapOfInterface {
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}

func getInterfaceValue(x reflect.Value) interface{} {
	if x.Kind() == reflect.Invalid {
		return nil
	}
	return x.Interface()
}

func isZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func createError(x reflect.Value, v constraint, message string) error {
	return pipeline.NewError(nil, fmt.Sprintf("validation failed: parameter=%s constraint=%s value=%#v details: %s",
		v.target, v.name, getInterfaceValue(x), message))
}
