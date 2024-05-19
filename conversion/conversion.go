package conversion

import (
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, err
		}

		floats = append(floats, floatPrice)

	}
	return floats, nil
}
