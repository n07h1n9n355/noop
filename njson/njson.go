package njson

// Noop is a struct used as a container for all `json`
type Noop struct{}

// MarshalJSON returns an empty byte slice, nil
func (n Noop) MarshalJSON() ([]byte, error) { return []byte{}, nil }

// UnmarshalJSON returns nil
func (n Noop) UnmarshalJSON([]byte) error { return nil }
