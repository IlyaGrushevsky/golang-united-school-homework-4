package string_sum

import (
	//"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
/*var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)*/

func StringSum(input string) (output string, err error) {
	var sum int = 0
	input1 := strings.Split(strings.TrimSpace(strings.Replace(input, "-", "+-", -1)), "+")
	for _, i := range input1 {
		a, err := strconv.Atoi(i)
		sum += a
		if err != nil {
			fmt.Println(err)
		}
	}
	output = strconv.Itoa(sum)

	return output, nil
}
