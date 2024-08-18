package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloat64UserInput(label string) (input float64) {
	fmt.Print(label)
	fmt.Scan(&input)
	return input
}

func WriteFloat64BalanceToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func ReadFloat64BalanceFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0.0, errors.New("failed to find balance file. balance resetted to 0.00. please contact customer care")
	}
	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0.0, errors.New("failed to parse stored balance value. balance resetted to 0.00. please contact customer care")
	}
	return balance, nil
}
