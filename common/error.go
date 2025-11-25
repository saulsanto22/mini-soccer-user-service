package common

import (
	"errors"
	"fmt"

	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) (validationRespose []ValidationResponse) {
	var fieldError validator.ValidationErrors

	if errors.As(err, &fieldError) {
		for _, err := range fieldError {
			switch err.Tag() {
			case "required":
				validationRespose = append(validationRespose, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required!", err.Field()),
				})
			case "email":
				validationRespose = append(validationRespose, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not valid email address! ", err.Field()),
				})
			default:
				errValidator, ok := ErrValidator[err.Tag()]
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						validationRespose = append(validationRespose, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field()),
						})
					} else {
						validationRespose = append(validationRespose, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
						})
					}
				} else {
					validationRespose = append(validationRespose, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Param()),
					})
				}
			}
		}
	}
	return validationRespose
}

func WrapError(err error) error {
	logrus.Errorf("error: %v", err)
	return err
}
