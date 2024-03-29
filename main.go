package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение вида \"a + b\"")
		input, _ := reader.ReadString('\n')
		output:= calculator(input)
		fmt.Println(output)
	}

}

func calculator(input string) string {

	text := strings.ReplaceAll(input, " ", "")
	text = strings.TrimSuffix(text, "\r\n") // для windows
	text = strings.TrimSuffix(text, "\n") // для linux и mac os
	
	operation := getOperationSymbol(text)
	
	if strings.Count(text, operation) > 1 || operation == "" {
		panic("не является допустимой математической операцией")
	}

	operands  := strings.Split(text, operation)

	if len(operands) != 2 {
		panic("не является допустимой математической операцией")
	}

	var aStr, bStr = operands[0], operands[1]

	aType := getNumberType(aStr)
	aInt := getNumber(aType, aStr)

	bType := getNumberType(bStr)
	bInt := getNumber(bType, bStr)

	if aType != bType {
		panic("используются одновременно разные системы счисления")
	}

	operandType := aType

	if operandType != numberTypes.arabic && operandType != numberTypes.roman {
		panic("введены не корректные символы")
	}

	if !(numberValidation(aInt) && numberValidation(bInt)) {
		panic("число выходит за допустимый диапазон")
	}

	isRoman := operandType == numberTypes.roman
	result := calculate(operation, aInt, bInt)
	
	if isRoman && result < 1 {
		panic("недопустимый результат для римских цифр")
	}

	output := formatResult(result, isRoman)
	return output
}

var operationSymbol = struct {
	plus string
	minus string
	multi string
	div string
}{
	plus: "+",
	minus: "-",
	multi: "*",
	div: "/",
} 

func getOperationSymbol(text string) string {
	switch {
	case strings.Contains(text, operationSymbol.multi):
		return operationSymbol.multi
	case strings.Contains(text, operationSymbol.div):
		return operationSymbol.div
	case strings.Contains(text, operationSymbol.minus):
		return operationSymbol.minus
	case strings.Contains(text, operationSymbol.plus):
		return operationSymbol.plus
	default:
		return ""
	}
}

func getNumber(numberType string, text string) int {
	var number int
	if numberType == numberTypes.arabic {
		number, _ = strconv.Atoi(text)
	} else if numberType == numberTypes.roman {
		for index, item := range romanIntegers { 
			if text == item {
				number = index
			}
		}
	} else {
		number = 0
	}
	return number
}

var numberTypes = struct {
	arabic string
	roman string
}{
	arabic: "arabic",
	roman: "roman",
}

var romanIntegers = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}

func getNumberType(text string) string {
	var numberType string
	number, _ := strconv.Atoi(text)
	if text == "0" || number != 0 {
		numberType = numberTypes.arabic
	} else {
		for _, item := range romanIntegers { 
			if text == item {
				numberType = numberTypes.roman
			}
		}
	}
	return numberType
}

func numberValidation(number int) bool {
	if number > 0 && number <= 10 {
		return true
	} else {
		return false
	}
}

func arabicToRoman(result int) string {

	firstDigit := romanIntegers[result/100]
	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[10], romanIntegers[1000])
	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[5], romanIntegers[500])
	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[1], romanIntegers[100])

	secondDigit := romanIntegers[result%100/10]
	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[10], romanIntegers[100])
	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[5], romanIntegers[50])
	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[1], romanIntegers[10])

	thirdDigit := romanIntegers[result%10]

	return firstDigit + secondDigit + thirdDigit
}

func calculate(operator string, a, b int) int {
	result := 0
	switch operator {
	case operationSymbol.plus:
		result = a + b
	case operationSymbol.minus:
		result = a - b
	case operationSymbol.multi:
		result = a * b
	case operationSymbol.div:
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	}
	return result
}

func formatResult(result int, isRoman bool) string { 
	if isRoman { 
		return arabicToRoman(result) 
	} 
	return strconv.Itoa(result) 
}
