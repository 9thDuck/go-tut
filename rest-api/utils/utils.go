package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func GetEnvVariable(varName string) (value string, err error) {
	value, varPresent := os.LookupEnv(varName)

	if !varPresent {
		return value, fmt.Errorf("error: %v .env variable not found", varName)
	}

	return value, err
}

func VerifyEnvAndGetMisingVars(envVarList []string) []string {
	var missingEnvVars []string

	for _, varName := range envVarList {
		_, err := GetEnvVariable(varName)
		if err != nil {
			missingEnvVars = append(missingEnvVars, varName)
		}
	}

	return missingEnvVars
}

func GetJwtDuration() (time.Duration, error) {
	duration := time.Duration(0)
	jwtDurationInHoursStr, err := GetEnvVariable("JWT_DURATION_IN_HOURS")

	if err != nil {
		return duration, err
	}

	jwtDurationInHours, err := strconv.ParseInt(jwtDurationInHoursStr, 10, 64)

	if err != nil {
		return duration, err
	}

	duration = time.Hour * time.Duration(jwtDurationInHours)

	return duration, nil
}
