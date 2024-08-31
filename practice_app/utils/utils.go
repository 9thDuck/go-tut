package utils

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func ReadFloat64DataFile(path string) (float64Slice []float64, err error) {
	byteData, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	strData := string(byteData)

	strSlice := strings.Split(strData, ",")

	for _, val := range strSlice {
		float64Val, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return nil, err
		}
		float64Slice = append(float64Slice, float64Val)
	}

	return float64Slice, err
}

func WriteJsonToFile(data any, filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}
