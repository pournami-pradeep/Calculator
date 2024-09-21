package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"strconv"
)

type Numbers struct {
	array []float64
	// length int
}

type Operators struct {
	arrayOperators []byte
}

func (num *Numbers) PushNum(val float64) {
	num.array = append(num.array, val)
}
func (operator *Operators) PushOperator(val byte) {
	operator.arrayOperators = append(operator.arrayOperators, val)
}

func (num *Numbers) Pop() float64 {
	length := len(num.array)
	if length == 0 {
		return 0
	}
	lastItem := num.array[length-1]
	num.array = num.array[:length-1]
	return lastItem

}
func (operator *Operators) Length() int {
	return len(operator.arrayOperators)
}

func (operator *Operators) lastElement() byte {
	length := len(operator.arrayOperators)
	lastItem := operator.arrayOperators[length-1]
	return lastItem
}
func (operator *Operators) Pop() byte {
	length := len(operator.arrayOperators)
	if length == 0 {
		return 'o'
	}
	lastoperator := operator.arrayOperators[length-1]
	operator.arrayOperators = operator.arrayOperators[:length-1]
	return lastoperator

}
func (operator *Operators) isEmpty() bool {
	length := len(operator.arrayOperators)
	return length == 0

}

func main() {
	for {
		userInput()
	}
}

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		expression := scanner.Text()
		expression = strings.ReplaceAll(expression, " ", "")
		expression = postfixConversion(expression)

		Result(expression)
		// fmt.Printf("Input was: %q\n", expression)
	}
	// var expression string
	// fmt.Scanln(&expression)
	// final_expression := postfixConversion(expression)
	// Result(final_expression)

	// expression := "-2+3-1"
	// expression = postfixConversion(expression)

	// fmt.Println("Expression", expression)
}

// Function for converting an infix expression to postfix expression
func postfixConversion(expression string) string {
	operator_precedence := map[byte]int{
		'+': 0,
		'-': 0,
		'/': 1,
		'*': 1,
		'(': -1,
		'^': 2,
		')': 3,
	}
	operator_stack := Operators{}
	len_of_expression := len(expression)
	postfix_expression := ""
	i := 0
	for i < len_of_expression {

		if expression[i] == '-' {
			if operator_stack.isEmpty() {
				if i == 0 {
					operator_stack.PushOperator(expression[i])
					postfix_expression = postfix_expression + "0 "

				}
				if i > 0 {
					operator_stack.PushOperator(expression[i])

				}

				i = i + 1
				continue

			}
			if !operator_stack.isEmpty() {
				prev_exp := expression[i-1]
				if _, err := strconv.Atoi(string(prev_exp)); err != nil {
					postfix_expression = postfix_expression + "0 "
					operator_stack.PushOperator(expression[i])
					i = i + 1
					continue
				}

				if operator_precedence[operator_stack.lastElement()] < operator_precedence[expression[i]] {
					operator_stack.PushOperator(expression[i])

				}
				for !operator_stack.isEmpty() && operator_precedence[expression[i]] <= operator_precedence[operator_stack.lastElement()] {
					if operator_stack.lastElement() == '(' {
						operator_stack.Pop()
					}
					postfix_expression = postfix_expression + string(operator_stack.Pop()) + " "
				}
				operator_stack.PushOperator('-')
			}
			// i = i + 1
			// continue

		}
		if expression[i] == '+' || expression[i] == '*' || expression[i] == '/' || expression[i] == '^' {
			if operator_stack.isEmpty() {
				operator_stack.PushOperator(expression[i])
				// i = i + 1
				// continue
			}
			if operator_precedence[operator_stack.lastElement()] < operator_precedence[expression[i]] {
				operator_stack.PushOperator(expression[i])
				// i = i + 1
				// continue

			}
			for !operator_stack.isEmpty() && operator_precedence[operator_stack.lastElement()] >= operator_precedence[expression[i]] {
				if operator_stack.lastElement() == '(' {
					operator_stack.Pop()
				}
				postfix_expression = postfix_expression + string(operator_stack.Pop()) + " "
			}
			operator_stack.PushOperator(expression[i])
			i = i + 1
			continue

		}
		// }

		if expression[i] == '(' {
			operator_stack.PushOperator('(')
			i = i + 1
			continue

		}

		if expression[i] == ')' {
			for {
				if operator_stack.lastElement() == '(' {
					operator_stack.Pop()

					break

				} else {
					postfix_expression = postfix_expression + string(operator_stack.Pop()) + " "
				}
			}
			i += 1
			continue

		}
		if _, err := strconv.Atoi(string(expression[i])); err == nil {
			fmt.Println("ji")
			// break
			num := ""
			for i < len_of_expression {
				if _, ok := operator_precedence[expression[i]]; ok {
					break
				}

				num = num + string(expression[i])
				// fmt.Println(num,"num")
				i = i + 1
			}
			postfix_expression = postfix_expression + num + " "
			i = i - 1
		}

		i = i + 1
	}

	for !operator_stack.isEmpty() {
		postfix_expression = postfix_expression + string(operator_stack.Pop()) + " "
	}
	fmt.Println(operator_stack)
	fmt.Println("POSTFIX EXPRESSION", postfix_expression)
	return postfix_expression

}

func Result(final_expression string) float64 {
	j := 0
	number_stack := Numbers{}
	len_of_final_expression := len(final_expression)
	// fmt.Println(len_of_final_expression,"AA")
	for j < len_of_final_expression {
		curr_val := final_expression[j]

		if curr_val == '+' || curr_val == '-' || curr_val == '/' || curr_val == '*' || curr_val == '^' {
			val1 := float64(number_stack.Pop())
			val2 := float64(number_stack.Pop())
			if curr_val == '+' {
				number_stack.PushNum(val1 + val2)
			}
			if curr_val == '-' {
				number_stack.PushNum(val2 - val1)
			}
			if curr_val == '*' {
				number_stack.PushNum(val1 * val2)
			}
			if curr_val == '/' {
				number_stack.PushNum(val2 / val1)
			}
			if curr_val == '^' {
				number_stack.PushNum(math.Pow(val2, val1))
			}
		} else {
			curr_num := ""
			for j < len_of_final_expression && final_expression[j] != ' ' {
				// fmt.Println("Po")
				curr_num = curr_num + string(final_expression[j])
				j = j + 1
				// fmt.Println(curr_num)

			}
			// fmt.Println(curr_num,"CURR")
			if len(curr_num) > 0 {
				number, err := strconv.ParseFloat(curr_num, 32)
				if err != nil {
					fmt.Println("Invalid number!")
					break

				}
				number_stack.PushNum(number)
			}

		}
		j = j + 1

	}
	ans := number_stack.Pop()
	fmt.Println("Answer is ", ans)
	return ans

}
