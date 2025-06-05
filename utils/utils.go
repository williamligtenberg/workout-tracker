package utils

import (
	"log"
	"strconv"
)

func StringToInt(str string) (int, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("[ERROR] Conversion failed for string '%s': %v", str, err)
		return 0, err
	}
	return value, nil
}
