package utils

import (
	"fmt"
	"strconv"
)

func ToInt(str string) (int, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error convert rate limit to number", err)
		return 0, err
	}
	return value, nil
}
