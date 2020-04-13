// Code generated by go generate
// This file was generated by robots at 2018-12-06 14:29:58.516631 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// String is an optional string.
type String struct {
	value string
	present bool
}

// NewString creates an optional.String from a string.
func NewString(v string) String {
	return String{value: v, present: true}
}

// Set sets the string value.
func (s *String) Set(v string) {
	s.value = v
	s.present = true
}

// Get returns the string value or an error if not present.
func (s String) Get() (string, error) {
	if !s.Present() {
		var zero string
		return zero, errors.New("value not present")
	}
	return s.value, nil
}

// Present returns whether or not the value is present.
func (s String) Present() bool {
	return s.present
}

// OrElse returns the string value or a default value if the value is not present.
func (s String) OrElse(v string) string {
	if s.Present() {
		return s.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (s String) If(fn func(string)) {
	if s.Present() {
		fn(s.value)
	}
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Present() {
		return json.Marshal(s.value)
	}
	return json.Marshal(nil)
}

func (s *String) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.present = false
		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.Set(value)
	return nil
}
