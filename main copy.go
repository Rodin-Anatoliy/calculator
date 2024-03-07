// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Println("Введите выражение вида \"a + b\"")

// 		var output string
// 		var result float64
// 		var isRoman bool
// 		isTypesEqual := true
// 		text, _ := reader.ReadString('\n')

// 		text = strings.Join(strings.Split(text, " "), "")
// 		text = strings.Join(strings.Split(text, "\r\n"), "")
		
// 		operation := getOperationSymbol(text)
		
// 		if strings.Count(text, operation) > 1 || operation == "" {
// 			fmt.Println("Паника, завершаю работу")
// 			return
// 		}

// 		textNumbers  := strings.Split(text, operation)

// 		if len(textNumbers) != 2 {
// 			fmt.Println("Паника, завершаю работу")
// 			return
// 		}

// 		for index, item := range textNumbers { 
// 			numberType := getNumberType(item)

// 			if numberType != numberTypes.arabic && numberType != numberTypes.roman {
// 				fmt.Println("Паника, завершаю работу")
// 				return
// 			} 

// 			if index != 0 {
// 				previousNumberType := getNumberType(textNumbers[index-1])
// 				isTypesEqual = isTypesEqual && (numberType == previousNumberType)
// 			}

// 			if !isTypesEqual {
// 				fmt.Println("Паника, завершаю работу")
// 				return
// 			}

// 			isRoman = numberType == numberTypes.roman

// 			number := getNumber(numberType, item)
// 			isValidNumber := numberValidation(number)

// 			if !isValidNumber {
// 				fmt.Println("Паника, завершаю работу")
// 				return
// 			}

// 			if index == 0 {
// 				result = float64(number)
// 			} else {
// 				switch {
// 				case operation == operationSymbol.plus:
// 					result = result + float64(number)
// 				case operation == operationSymbol.minus:
// 					result = result - float64(number)
// 				case operation == operationSymbol.multi:
// 					result = result * float64(number)
// 				case operation == operationSymbol.div:
// 					result = result / float64(number)
// 				}
// 			}

// 		}
		
// 		if isRoman && result < 0 {
// 			fmt.Println("Паника, завершаю работу")
// 			return
// 		}

// 		if isRoman {
// 			output = arabicToRoman(result)
// 		} else {
// 			output = strconv.FormatFloat(result, 'f', 6, 64)
// 		}

// 		fmt.Println(output)

// 	}
// }

// var operationSymbol = struct {
// 	plus string
// 	minus string
// 	multi string
// 	div string
// }{
// 	plus: "+",
// 	minus: "-",
// 	multi: "*",
// 	div: "/",
// } 

// func getOperationSymbol(text string) string {
// 	switch {
// 	case strings.Contains(text, operationSymbol.multi):
// 		return operationSymbol.multi
// 	case strings.Contains(text, operationSymbol.div):
// 		return operationSymbol.div
// 	case strings.Contains(text, operationSymbol.minus):
// 		return operationSymbol.minus
// 	case strings.Contains(text, operationSymbol.plus):
// 		return operationSymbol.plus
// 	default:
// 		return ""
// 	}
// }

// func getNumber(numberType string, text string) int {
// 	var number int
// 	if numberType == numberTypes.arabic {
// 		number, _ = strconv.Atoi(text)
// 	} else if numberType == numberTypes.roman {
// 		for index, item := range romanIntegers { 
// 			if text == item {
// 				number = index
// 			}
// 		}
// 	} else {
// 		number = 0
// 	}
// 	return number
// }

// var numberTypes = struct {
// 	arabic string
// 	roman string
// }{
// 	arabic: "arabic",
// 	roman: "roman",
// }

// var romanIntegers = map[int]string{
// 	1: "I",
// 	2: "II",
// 	3: "III",
// 	4: "IV",
// 	5: "V",
// 	6: "VI",
// 	7: "VII",
// 	8: "VIII",
// 	9: "IX",
// 	10: "X",
// 	50: "L",
// 	100: "C",
// 	500: "D",
// 	1000: "M",
// }

// var romanFractions = map[int]string{
// 	1: "·",
// 	2: "··",
// 	3: "···",
// 	4: "····",
// 	5: "·····",
// 	6: "S",
// 	7: "S·",
// 	8: "S··",
// 	9: "S···",
// 	10: "S····",
// 	11: "S·····",
// }

// func getNumberType(text string) string {
// 	var numberType string
// 	number, _ := strconv.Atoi(text)
// 	if text == "0" || number != 0 {
// 		numberType = numberTypes.arabic
// 	} else {
// 		for _, item := range romanIntegers { 
// 			if text == item {
// 				numberType = numberTypes.roman
// 			}
// 		}
// 	}
// 	return numberType
// }

// func numberValidation(number int) bool {
// 	if number > 0 && number <= 10 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func arabicToRoman(result float64) string {

// 	integer := int(result)
// 	fraction := int(math.Round((result - float64(integer)) * 12))

// 	if fraction == 12 {
// 		integer = integer + 1
// 		fraction = 0
// 	}

// 	firstDigit := romanIntegers[integer/100]
// 	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[10], romanIntegers[1000])
// 	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[5], romanIntegers[500])
// 	firstDigit = strings.ReplaceAll(firstDigit, romanIntegers[1], romanIntegers[100])

// 	secondDigit := romanIntegers[integer%100/10]
// 	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[10], romanIntegers[100])
// 	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[5], romanIntegers[50])
// 	secondDigit = strings.ReplaceAll(secondDigit, romanIntegers[1], romanIntegers[10])

// 	thirdDigit := romanIntegers[integer%10]

// 	romanInteger := firstDigit + secondDigit + thirdDigit

// 	romanFraction := romanFractions[fraction]

// 	return romanInteger + romanFraction
// }