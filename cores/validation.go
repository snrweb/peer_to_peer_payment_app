package cores

import (
	"errors"
	"net/mail"
)

func Validate(value interface{}, valueName string, rules []string) error {
	for _, rule := range rules {
		switch rule {
		case "empty":
			if len(value.(string)) == 0 {
				return errors.New(valueName + " is empty")
			}

		case "greaterThanZero":
			if value.(float64) <= 0 {
				return errors.New(valueName + " should be greater than zero")
			}

		case "isEmail":
			if _, err := mail.ParseAddress(value.(string)); err != nil {
				return errors.New(valueName + " address is invalid")
			}
		}
	}
	return nil
}
