package handlers

import (
	"errors"
	"reflect"
	"restapi/pkg/utils"
	"strings"
)

func CheckBlankFields(value interface{}) error {
	val := reflect.ValueOf(value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			// fmt.Println("field.Kind():", field.Kind())
			// fmt.Println("reflect.String:", reflect.String)
			// fmt.Println("field.String():", field.String())
			// http.Error(w, "All Fields are required", http.StatusBadRequest)

			return utils.ErrorHandler(errors.New("All fields are required"), "All fields are required")
		}
	}
	return nil
}

// step: 4)
func GetFieldNames(model interface{}) []string {
	val := reflect.TypeOf(model)
	fields := []string{}

	// Step: 5)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// Step: 13)
		fieldToAdd := strings.TrimSuffix(field.Tag.Get("json"), ",omitempty")
		// Step: 5)
		// fields = append(fields, field.Tag.Get("json")) // Get JSON tag|this wont work. so we need above code
		fields = append(fields, fieldToAdd) // Get JSON tag
	}
	return fields
}
