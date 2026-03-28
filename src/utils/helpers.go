package data_parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

type JSON struct {
	*json.RawMessage
}

func NewJSON(data []byte) (*JSON, error) {
	return &JSON{&data}, nil
}

func (j *JSON) Unmarshal(data []byte) error {
	*j.RawMessage = data
	return nil
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetIntEnv(key string, defaultValue int) int {
	value := GetEnv(key, "")
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err!= nil {
		log.Printf("Failed to parse environment variable %s: %v", key, err)
		return defaultValue
	}
	return i
}

func GetBoolEnv(key string, defaultValue bool) bool {
	value := GetEnv(key, "")
	if value == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(value)
	if err!= nil {
		log.Printf("Failed to parse environment variable %s: %v", key, err)
		return defaultValue
	}
	return b
}

func IsStringSliceEmpty(s []string) bool {
	return len(s) == 0
}

func GetJoinedStringSlice(s []string) string {
	return strings.Join(s, ",")
}

func GetDateFromISO8601(dateString string) (Date, error) {
	t, err := time.Parse(time.RFC3339, dateString)
	if err!= nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func GetDateStringFromTime(t time.Time) string {
	return t.Format("2006-01-02")
}