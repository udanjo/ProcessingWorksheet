package middleware

import (
	"fmt"
	"strconv"
	"strings"
)

func Handler(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}

func ConvertFloat(st string) float64 {
	if st != "" {
		p, err := strconv.ParseFloat(strings.Replace(st, ",", ".", -1), 64)
		Handler(err)
		return p
	} else {
		return 0
	}
}
