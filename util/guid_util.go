package util

import "github.com/google/uuid"

func GenerateGuid() string {
	return uuid.New().String()
}
