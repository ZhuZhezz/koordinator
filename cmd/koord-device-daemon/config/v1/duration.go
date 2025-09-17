package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

// Duration wraps a time.Duration function with custom JSON marshaling/unmarshaling
type Duration time.Duration

// MarshalJSON marshals 'Duration' to its raw bytes representation
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

// UnmarshalJSON unmarshals raw bytes into a 'Duration' type.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return fmt.Errorf("invalid duration")
	}
}
