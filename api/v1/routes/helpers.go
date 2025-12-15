package dataparser

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DateStringToTime converts a date string in a specific format to a time.Time object
func DateStringToTime(dateString string, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}

// DateStringToTimeWithDefault converts a date string in a specific format to a time.Time object
// If the date string is empty, it returns the current time
func DateStringToTimeWithDefault(dateString string, format string) (time.Time, error) {
	if strings.TrimSpace(dateString) == "" {
		return time.Now(), nil
	}
	return time.Parse(format, dateString)
}

// IsUUID checks if a string is a valid UUID
func IsUUID(uuid string) bool {
	if len(uuid)!= 36 {
		return false
	}
	return strings.ReplaceAll(uuid, "-", "")!= ""
}

// ValidateEmail checks if a string is a valid email address
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// ValidatePassword checks if a string is a valid password
func ValidatePassword(password string) bool {
	if password == "" {
		return false
	}
	if len(password) < 8 {
		return false
	}
	return true
}

// TrimString trims a string and removes leading/trailing whitespaces
func TrimString(s string) string {
	return strings.TrimSpace(s)
}

// ToLower converts a string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// SplitString splits a string by a delimiter and returns a slice of strings
func SplitString(s string, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// JoinString joins a slice of strings with a delimiter and returns a string
func JoinString(slice []string, delimiter string) string {
	return strings.Join(slice, delimiter)
}

// Errorf returns an error with a formatted message
func Errorf(format string, args...interface{}) error {
	return fmt.Errorf(format, args...)
}

// Errorsf returns multiple errors with formatted messages
func Errorsf(errs...error) error {
	var errorMessages []string
	for _, err := range errs {
		if err!= nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}
	if len(errorMessages) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(errorMessages, "; "))
}

// IsNil checks if a value is nil
func IsNil(value interface{}) bool {
	return value == nil
}

// IsNotNil checks if a value is not nil
func IsNotNil(value interface{}) bool {
	return!IsNil(value)
}

// IsError checks if a value is an error
func IsError(value interface{}) bool {
	_, ok := value.(error)
	return ok
}

// IsNotNilError checks if a value is not an error
func IsNotNilError(value interface{}) bool {
	return!IsError(value)
}

// IsNilError checks if a value is nil and an error
func IsNilError(value interface{}) bool {
	return IsNil(value) && IsError(value)
}