package BusinessLogicLayer

import "encoding/json"

func Convert[T any](input T) (*T, error) {
	byte, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	var output T
	json.Unmarshal(byte, &output)

	return &output, nil
}
