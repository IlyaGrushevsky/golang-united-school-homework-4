package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var sum int = 0
	var count int = 0
	if _, e := strconv.ParseInt(input, 10, 64); e != nil {
		err := fmt.Errorf("Пустой ввод %w", errorEmptyInput)
		fmt.Println(err.Error())
		return "", err
	}
	input1 := strings.Split(strings.TrimSpace(strings.Replace(input, "-", "+-", -1)), "+")
	for _, i := range input1 {
		a, e := strconv.Atoi(i)
		sum += a
		count++
		err := fmt.Errorf("Некорректный ввод %w", e)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}
	if count != 2 {
		err := fmt.Errorf("Пустой ввод %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}
	output = strconv.Itoa(sum)

	return output, nil
}
