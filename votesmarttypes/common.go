package votesmarttypes

import (
	"encoding/json"
	"errors"
)

// MaybeList is a generic unmarshaler. Votesmart response structure varies
// depending on the length of the response. If there is a single item, it
// returns the object. If there are multiple, it may return a list. This handler
// converts those cases to lists.
type MaybeList[T any] []T

func (m *MaybeList[T]) UnmarshalJSON(in []byte) error {
	if string(in) == `""` {
		return nil
	}

	var single T
	if err := json.Unmarshal(in, &single); err == nil {
		*m = []T{single}
		return nil
	}

	var list []T
	if err := json.Unmarshal(in, &list); err == nil {
		*m = list
		return nil
	}

	return errors.New("unexpected list response format")
}
