package client

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func StructToQueryParams(data interface{}) (string, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("input data must be a struct")
	}

	var queryParams = make(url.Values)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldName := field.Tag.Get("query")
		if fieldName == "" {
			fieldName = strings.ToLower(field.Name)
		}
		fieldValue := fmt.Sprintf("%v", v.Field(i).Interface())
		queryParams.Add(fieldName, fieldValue)
	}

	return queryParams.Encode(), nil
}
