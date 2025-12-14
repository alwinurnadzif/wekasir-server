package utils

import (
	"fmt"
	"reflect"
)

func SliceToMapByJsonTag[T any, K comparable](items []T, tagField string) (map[K]T, error) {
	result := make(map[K]T)

	for _, item := range items {
		val := reflect.ValueOf(item)
		typ := reflect.TypeOf(item)

		if typ.Kind() != reflect.Struct {
			return nil, fmt.Errorf("T must be a struct")
		}

		var key K
		found := false

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("json")

			if tag == tagField {
				fv := val.Field(i)

				// Interfacenya harus bisa cast ke K
				v, ok := fv.Interface().(K)
				if !ok {
					return nil, fmt.Errorf("field %s must be type %T", tagField, key)
				}

				key = v
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("no field with json tag '%s' found", tagField)
		}

		result[key] = item
	}

	return result, nil
}
