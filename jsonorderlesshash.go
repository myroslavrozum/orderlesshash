package orderlesshash

import (
	"encoding/json"
	"errors"
)

//JSONMapUnorderedSha ...
func (hashes *Hashes) JSONMapUnorderedSha(jsonString []byte) ([]byte, error) {
	var tmpEntity map[string]interface{}

	if !json.Valid(jsonString) {
		return nil, errors.New("Invalid JSON")
	}

	if err := json.Unmarshal(jsonString, &tmpEntity); err != nil {
		return nil, err
	}
	return hashes.anonymousUnorderedSha(tmpEntity)
}
