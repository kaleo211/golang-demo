package arithmetic

import "errors"

func Double(input interface{}) (interface{}, error) {
	switch input.(type) {
	case string:
		return input.(string) + input.(string), nil

	case int:
		return input.(int) * 2, nil

	default:
		return nil, errors.New("not supported type")
	}
}
