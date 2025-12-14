package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var (
	ImagesContentTypes map[string]bool = map[string]bool{"image/jpg": true, "image/jpeg": true, "image/png": true}
	ExcelContentTypes  map[string]bool = map[string]bool{"application/vnd.openxmlformats-officedocument.spreadsheetml": true, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true, "application/vnd.ms-excel": true, "application/wps-office.xlsx": true}
)

type ValidationError struct {
	Index   *uint  `json:"index"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

var (
	ErrValidationDetail = errors.New("validation error details")
	ErrValidation       = errors.New("validation error")
	ErrEmpty            = errors.New("empty data")
)

func ValidatePayload(payload any) (*[]ValidationError, error) {
	var validationErrors []ValidationError

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitAfterN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v := reflect.ValueOf(payload)

	// Kalau pointer, ambil element aslinya
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		// Kalau payload berupa array/slice
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i).Interface()
			if err := validate.Struct(elem); err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					index := uint(i)
					validationErrors = append(validationErrors, ValidationError{
						Field:   err.Field(),
						Message: err.Tag(),
						Value:   err.Param(),
						Index:   &index,
					})
				}
			}
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Type().Field(i)
			tag := fieldType.Tag.Get("validate")

			// handle field biasa (bukan slice)
			if !strings.Contains(tag, "dive") {
				if err := validate.Var(field.Interface(), tag); err != nil {
					for _, verr := range err.(validator.ValidationErrors) {
						jsonTag := fieldType.Tag.Get("json")
						jsonName := strings.Split(jsonTag, ",")[0]
						if jsonName == "" || jsonName == "-" {
							jsonName = fieldType.Name
						}

						validationErrors = append(validationErrors, ValidationError{
							Field:   jsonName,
							Message: verr.Tag(),
							Value:   verr.Param(),
							Index:   nil,
						})
					}
				}
				continue
			}

			// handle field dengan "dive"
			switch field.Kind() {
			case reflect.Slice, reflect.Array:
				for j := 0; j < field.Len(); j++ {
					elem := field.Index(j).Interface()
					if err := validate.Struct(elem); err != nil {
						for _, verr := range err.(validator.ValidationErrors) {
							idx := uint(j)

							jsonTag := fieldType.Tag.Get("json")
							jsonName := strings.Split(jsonTag, ",")[0]
							if jsonName == "" || jsonName == "-" {
								jsonName = fieldType.Name
							}
							validationErrors = append(validationErrors, ValidationError{
								Field:   fmt.Sprintf("%s.%s", jsonName, verr.Field()),
								Message: verr.Tag(),
								Value:   verr.Param(),
								Index:   &idx,
							})
						}
					}
				}
			case reflect.Struct:
				// fallback: struct tunggal (anggap index 0)
				idx := uint(0)
				if err := validate.Struct(field.Interface()); err != nil {
					for _, verr := range err.(validator.ValidationErrors) {
						validationErrors = append(validationErrors, ValidationError{
							Field:   fmt.Sprintf("%s.%s", fieldType.Name, verr.Field()),
							Message: verr.Tag(),
							Value:   verr.Param(),
							Index:   &idx,
						})
					}
				}
			}
		}

	default:
		return nil, fmt.Errorf("unsupported type: %s", v.Kind())
	}

	if len(validationErrors) > 0 {
		return &validationErrors, ErrValidation
	}

	return &validationErrors, nil
}

func ValidateFileSize(file *multipart.FileHeader, maxFileSizeMb uint, field string) interface{} {
	maxFileSizeByte := maxFileSizeMb * 1024 * 1024

	errResponses := []ValidationError{}
	if file.Size > int64(maxFileSizeByte) {
		errResponse := ValidationError{
			Field:   field,
			Message: "ukuran file terlalu besar",
		}

		errResponses = append(errResponses, errResponse)
		return errResponses
	}
	return nil
}

func ValidateFileContentType(file *multipart.FileHeader, allowedContentTypes map[string]bool, field string) interface{} {
	contentType := file.Header.Get("Content-Type")

	errResponses := []ValidationError{}
	if !allowedContentTypes[contentType] {
		errResponse := ValidationError{
			Field:   field,
			Message: "file tidak didukung",
		}

		errResponses = append(errResponses, errResponse)
		return errResponses
	}

	return nil
}
