/**
 * @author:       wangxuebing
 * @fileName:     convertor.go
 * @date:         2023/5/5 0:09
 * @description:
 */

package convertor

import (
	"encoding/json"
	"fmt"
)

// InterfaceToSlice convert interface to slice
func InterfaceToSlice[T any](i interface{}) ([]T, error) {
	s := make([]T, 0)
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert iterface to slice error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &s); err != nil {
		return nil, fmt.Errorf("convert iterface to slice error. %s", err.Error())
	}
	return s, nil
}

// InterfaceToStruct convert interface to struct
func InterfaceToStruct[T any](i interface{}) (*T, error) {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert interface to struct error. %s", err.Error())
	}

	t := new(T)
	if err := json.Unmarshal(marshal, t); err != nil {
		return nil, fmt.Errorf("convert interface to struct error. %s", err.Error())
	}
	return t, nil
}

// InterfaceToMap convert interface to map
func InterfaceToMap[K comparable, V any](i interface{}) (map[K]V, error) {
	marshal, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("convert interface to map error. %s", err.Error())
	}
	m := make(map[K]V, 0)
	if err := json.Unmarshal(marshal, &m); err != nil {
		return nil, fmt.Errorf("convert interface to map error. %s", err.Error())
	}
	return m, nil
}

// MapToStruct convert map to struct
func MapToStruct[T any](m map[string]interface{}) (*T, error) {
	t := new(T)
	marshal, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("convert map to struct error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &t); err != nil {
		return nil, fmt.Errorf("convert map to struct error. %s", err.Error())
	}
	return t, nil
}

// StructToMap convert struct to map
func StructToMap[K comparable, V any](s interface{}) (map[K]V, error) {
	m := make(map[K]V, 0)
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("convert struct to map error. %s", err.Error())
	}
	if err := json.Unmarshal(marshal, &m); err != nil {
		return nil, fmt.Errorf("convert struct to map error. %s", err.Error())
	}
	return m, nil
}
