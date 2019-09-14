package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

func main() {

	numberA := "100"
	numberB := "50"
	FNumberA := "15.05"
	FNumberB := "1.05"

	fmt.Println("----문자열 정수 연산----")
	fmt.Printf(" %s + %s = ", numberA, numberB)
	fmt.Println(calcAdd(numberA, numberB))
	fmt.Printf(" %s - %s = ", numberA, numberB)
	fmt.Println(calcSub(numberA, numberB))
	fmt.Printf(" %s * %s = ", numberA, numberB)
	fmt.Println(calcMul(numberA, numberB))
	fmt.Printf(" %s / %s = ", numberA, numberB)
	fmt.Println(calcDiv(numberA, numberB))
	fmt.Println("-----------------------")

	fmt.Println("----문자열 소수 연산----")
	fmt.Printf(" %s + %s = ", FNumberA, FNumberB)
	fmt.Println(calcAdd(FNumberA, FNumberB))
	fmt.Printf(" %s - %s = ", FNumberA, FNumberB)
	fmt.Println(calcSub(FNumberA, FNumberB))
	fmt.Printf(" %s * %s = ", FNumberA, FNumberB)
	fmt.Println(calcMul(FNumberA, FNumberB))
	fmt.Printf(" %s / %s = ", FNumberA, FNumberB)
	fmt.Println(calcDiv(FNumberA, FNumberB))
	fmt.Println("-----------------------")

}

func convertBigInt(a, b string) (*big.Int, *big.Int, error) {

	first, ok := new(big.Int).SetString(a, 10)
	second, ok := new(big.Int).SetString(b, 10)
	var error string

	if ok == false {
		error = "error "
	}
	return first, second, errors.New(error)
}

func decimalCompute(numberString string) (string, int, int) {

	var dotIndexPosition, underDotLength int

	if dotIndexPosition = strings.Index(numberString, "."); dotIndexPosition != -1 {
		underDotLength = len(numberString[dotIndexPosition+1:])
		numberString = strings.Replace(numberString, ".", "", -1)
	} else {
		underDotLength = 0
		temp := []rune(numberString)
		tempZero := []rune(".")
		tempZero = append(tempZero, '0')
		temp = append(temp, tempZero...)

		dotIndexPosition = strings.Index(string(temp), ".")
	}
	return numberString, dotIndexPosition, underDotLength
}

func calcAdd(a, b string) (string, error) {

	var fillZeroCount, dotIndex, countZero int

	originAdotIndex := strings.Index(a, ".")
	originBdotIndex := strings.Index(b, ".")

	/*
		1. a 문자열에서 " . " 의 유무를 확인한다.
		2. . 이 있을 경우 소수이고 소수점 아래로의 숫자길이를 구한다. ( a, b 를 정수화 하며 연산시 0을 붙여서 계산하기 위해)
		3. 해당 문자열에서 " . " 을 지워서 정수화 한다.

		4. 소수가 아닐경우(정수)  임시적으로 수의 끝에 .0을 넣어서( ex) 100.0 ) . 의 인덱스를 저장한다.

	*/
	numberStringA, dotIndexPositionA, underDotLengthA := decimalCompute(a)
	numberStringB, dotIndexPositionB, underDotLengthB := decimalCompute(b)
	a = numberStringA
	b = numberStringB
	/*
		1. a , b 의 문자열 둘 다 머리에 0 이 있다는 것을 확인 ( 둘 다 0.xx로 시작하는 소수인지 확인)
		2. 둘 다 소수라면 정수화 후 ". " , "0" 을 추가적으로 붙여줘야 하므로 0에 대한 값을 저장.
		3. 연산 후 0을 붙일 때 => a의 값이 b보다 작을 경우 b의 0을 따라간다.
		4. a값이 b보다 클 경우 a의 0의 개수를 저장한다. 같을 경우는 둘 중에 하나 저장.

	*/
	if strings.HasPrefix(a, "0") && strings.HasPrefix(b, "0") {

		countZeroA := strings.Count(a, "0")
		countZeroB := strings.Count(b, "0")

		if countZeroA > countZeroB {
			countZero = countZeroB
		} else if countZeroA < countZeroB {
			countZero = countZeroA
		} else {
			countZero = countZeroA
		}
	}

	/*
		1. a, b  각 문자열숫자 의 "." 아래로의 길이를 위에서 구했고 해당 변수를 사용해서 정수화된 숫자의 길이를 맞춘다.
		2. 0.1 + 0.02 => 01 + 002 => 010 + 002 와 같이 길이가 작은 숫자의 뒤에 차이만큼 0을 붙여 맞춘다.
	*/
	if underDotLengthA > underDotLengthB {
		fillZeroCount = underDotLengthA - underDotLengthB

		BString := []rune(b)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempB := make([]rune, 0, len(BString)+fillZeroCount)

		tempB = append(tempB, fillZero...)

		BString = append(BString, tempB...)

		b = string(BString)
	}

	if underDotLengthA < underDotLengthB {
		fillZeroCount = underDotLengthB - underDotLengthA
		AString := []rune(a)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempA := make([]rune, 0, len(AString)+fillZeroCount)

		tempA = append(tempA, fillZero...)

		AString = append(AString, tempA...)

		a = string(AString)
	}

	/*
		1. dotIndexPosition  은 a , b의 .의 위치를 나타낸다
		2. .의 인덱스의 위치 크기는 곧 해당 숫자의 크기를 나타낸다.
		  ex ) 1234.032 = 4번 인덱스 , 10.23 = 2번 인덱스
		3. 더 큰 수의 " . " 인덱스의 위치를 결과에 반영하도록 dotIndex에 저장한다.
	*/

	if dotIndexPositionA > dotIndexPositionB {
		dotIndex = dotIndexPositionA
	} else {
		dotIndex = dotIndexPositionB
	}

	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Add(first, second)

	/*
		처음에 parameter로 들어온 a와 b 둘 중에 하나가 소수일 때 해당 if문에서 정수화된 결과물의 후처리작업을 실행한다.
	*/
	if originAdotIndex != -1 || originBdotIndex != -1 {

		resultString := []rune(result.String())
		temp := make([]rune, 0, len(resultString)+1+countZero)

		if countZero != 0 {
			zeroStrings := []rune(strings.Repeat("0", countZero-1))
			temp = append(temp, '0', '.')
			temp = append(temp, zeroStrings...)
			resultString = append(temp, resultString...)
		} else {
			temp = append(temp, resultString[:dotIndex]...)
			temp = append(temp, '.')
			resultString = append(temp, resultString[dotIndex:]...)
		}

		return string(resultString), error
	}

	return result.String(), error
}

func calcSub(a, b string) (string, error) {
	var fillZeroCount, dotIndex, countZero int

	originAdotIndex := strings.Index(a, ".")
	originBdotIndex := strings.Index(b, ".")

	/*
		1. a 문자열에서 " . " 의 유무를 확인한다.
		2. . 이 있을 경우 소수이고 소수점 아래로의 숫자길이를 구한다. ( a, b 를 정수화 하며 연산시 0을 붙여서 계산하기 위해)
		3. 해당 문자열에서 " . " 을 지워서 정수화 한다.

		4. 소수가 아닐경우(정수)  임시적으로 수의 끝에 .0을 넣어서( ex) 100.0 ) . 의 인덱스를 저장한다.
	*/

	numberStringA, dotIndexPositionA, underDotLengthA := decimalCompute(a)
	numberStringB, dotIndexPositionB, underDotLengthB := decimalCompute(b)
	a = numberStringA
	b = numberStringB
	/*
		1. a , b 의 문자열 둘 다 머리에 0 이 있다는 것을 확인 ( 둘 다 0.xx로 시작하는 소수인지 확인)
		2. 둘 다 소수라면 정수화 후 ". " , "0" 을 추가적으로 붙여줘야 하므로 0에 대한 값을 저장.
		3. 연산 후 0을 붙일 때 => a의 값이 b보다 작을 경우 b의 0을 따라간다.
		4. a값이 b보다 클 경우 a의 0의 개수를 저장한다. 같을 경우는 둘 중에 하나 저장.

	*/
	if strings.HasPrefix(a, "0") && strings.HasPrefix(b, "0") {

		countZeroA := strings.Count(a, "0")
		countZeroB := strings.Count(b, "0")

		if countZeroA > countZeroB {
			countZero = countZeroB
		} else if countZeroA < countZeroB {
			countZero = countZeroA
		} else {
			countZero = countZeroA
		}
	}

	/*
		1. a, b  각 문자열숫자 의 "." 아래로의 길이를 위에서 구했고 해당 변수를 사용해서 정수화된 숫자의 길이를 맞춘다.
		2. 0.1 + 0.02 => 01 + 002 => 010 + 002 와 같이 길이가 작은 숫자의 뒤에 차이만큼 0을 붙여 맞춘다.
	*/
	if underDotLengthA > underDotLengthB {
		fillZeroCount = underDotLengthA - underDotLengthB

		BString := []rune(b)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempB := make([]rune, 0, len(BString)+fillZeroCount)

		tempB = append(tempB, fillZero...)

		BString = append(BString, tempB...)

		b = string(BString)
	}

	if underDotLengthA < underDotLengthB {
		fillZeroCount = underDotLengthB - underDotLengthA
		AString := []rune(a)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempA := make([]rune, 0, len(AString)+fillZeroCount)

		tempA = append(tempA, fillZero...)

		AString = append(AString, tempA...)

		a = string(AString)
	}

	/*
		1. dotIndexPosition  은 a , b의 .의 위치를 나타낸다
		2. .의 인덱스의 위치 크기는 곧 해당 숫자의 크기를 나타낸다.
		  ex ) 1234.032 = 4번 인덱스 , 10.23 = 2번 인덱스
		3. 더 큰 수의 " . " 인덱스의 위치를 결과에 반영하도록 dotIndex에 저장한다.
	*/

	if dotIndexPositionA > dotIndexPositionB {
		dotIndex = dotIndexPositionA
	} else {
		dotIndex = dotIndexPositionB
	}

	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Sub(first, second)

	/*
		처음에 parameter로 들어온 a와 b 둘 중에 하나가 소수일 때 해당 if문에서 정수화된 결과물의 후처리작업을 실행한다.
	*/
	if originAdotIndex != -1 || originBdotIndex != -1 {

		resultString := []rune(result.String())
		temp := make([]rune, 0, len(resultString)+1+countZero)

		if countZero != 0 {
			zeroStrings := []rune(strings.Repeat("0", countZero-1))
			temp = append(temp, '0', '.')
			temp = append(temp, zeroStrings...)
			resultString = append(temp, resultString...)
		} else {
			temp = append(temp, resultString[:dotIndex]...)
			temp = append(temp, '.')
			resultString = append(temp, resultString[dotIndex:]...)
		}

		return string(resultString), error
	}

	return result.String(), error
}

func calcMul(a, b string) (string, error) {

	var dotIndexPositionA, dotIndexPositionB, dotIndex, countZero int

	originAdotIndex := strings.Index(a, ".")
	originBdotIndex := strings.Index(b, ".")

	if dotIndexPositionA = strings.Index(a, "."); dotIndexPositionA != -1 {
		a = strings.Replace(a, ".", "", -1)
	}

	if dotIndexPositionB = strings.Index(b, "."); dotIndexPositionB != -1 {
		b = strings.Replace(b, ".", "", -1)
	}

	if dotIndexPositionA*dotIndexPositionB-1 == 0 {
		dotIndex = 1
	} else {
		dotIndex = dotIndexPositionA * dotIndexPositionB
	}

	if strings.HasPrefix(a, "0") {
		countZero += strings.Count(a, "0")
	}

	if strings.HasPrefix(b, "0") {
		countZero += strings.Count(b, "0")
	}

	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Mul(first, second)

	if originAdotIndex != -1 && originBdotIndex != -1 {

		if dotIndex > -1 {
			resultString := []rune(result.String())
			temp := make([]rune, 0, len(resultString)+1+countZero)

			temp = append(temp, resultString[:dotIndex]...)

			if countZero > 1 {
				zeroStrings := []rune(strings.Repeat("0", countZero-1))
				temp = append(temp, '0', '.')
				temp = append(temp, zeroStrings...)

			} else {
				temp = append(temp, '.')
			}
			resultString = append(temp, resultString[dotIndex:]...)

			return string(resultString), error
		}
	}
	return result.String(), error
}

func calcDiv(a, b string) (string, error) {
	var dotIndexPositionA, dotIndexPositionB, underDotLengthA, underDotLengthB, fillZeroCount int

	if dotIndexPositionA = strings.Index(a, "."); dotIndexPositionA != -1 {
		underDotLengthA = len(a[dotIndexPositionA+1:])
		a = strings.Replace(a, ".", "", -1)
	}

	if dotIndexPositionB = strings.Index(b, "."); dotIndexPositionB != -1 {
		underDotLengthB = len(b[dotIndexPositionB+1:])
		b = strings.Replace(b, ".", "", -1)
	}

	if underDotLengthA > underDotLengthB {
		fillZeroCount = underDotLengthA - underDotLengthB

		BString := []rune(b)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempB := make([]rune, 0, len(BString)+fillZeroCount)

		tempB = append(tempB, fillZero...)

		BString = append(BString, tempB...)

		b = string(BString)
	}

	if underDotLengthA < underDotLengthB {
		fillZeroCount = underDotLengthB - underDotLengthA
		AString := []rune(a)
		fillZero := []rune(strings.Repeat("0", fillZeroCount))

		tempA := make([]rune, 0, len(AString)+fillZeroCount)

		tempA = append(tempA, fillZero...)

		AString = append(AString, tempA...)

		a = string(AString)
	}

	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Div(first, second)

	return result.String(), error
}
