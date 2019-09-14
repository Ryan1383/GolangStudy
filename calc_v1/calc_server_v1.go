package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/browser"
)

var portNumber string

//CalcResult :계산 결과 구조체
type CalcResult struct {
	ResultValue string `json:"resultValue"`
}

func main() {
	handleHTTPRequest()
	searchAvailablePorts()
}

func handleHTTPRequest() {
	http.HandleFunc("/", handleRootView)
	http.HandleFunc("/calc", getCalcResult)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func handleRootView(w http.ResponseWriter, r *http.Request) {
	calcView, err := ioutil.ReadFile("./static/calc_view.html")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(calcView))
}

func searchAvailablePorts() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	portNumber = port
	writePortScript(portNumber)
	openBrowser(portNumber)

	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	panic(http.Serve(listener, nil))
}

func writePortScript(portNumber string) {
	portFile, err := os.Create("./static/port.js")

	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(portFile)
	buf.WriteString("var portNumber=" + portNumber + ";")
	buf.Flush()
	portFile.Close()
}

func openBrowser(portNumber string) {
	var url = "http://localhost:" + portNumber
	fmt.Println(url)
	browser.OpenURL(url)

}

func getCalcResult(w http.ResponseWriter, r *http.Request) {
	firstParam, ok := r.URL.Query()["firstString"]
	secondParam, ok := r.URL.Query()["secondString"]
	opParam, ok := r.URL.Query()["operator"]
	var result string

	if ok == false || len(firstParam[0]) < 1 {
		log.Println("Url Param 'firstString' is missing")
		return
	}
	if ok == false || len(secondParam[0]) < 1 {
		log.Println("Url Param 'secondString' is missing")
		return
	}
	if ok == false || len(opParam[0]) < 1 {
		log.Println("Url Param 'operator' is missing")
		return
	}

	firstString := firstParam[0]
	secondString := secondParam[0]
	operator := opParam[0]

	switch operator {
	case "ADD":
		result, _ = calcAdd(firstString, secondString)
		log.Println(result)
		break
	case "SUB":
		result, _ = calcSub(firstString, secondString)
		log.Println(result)
		break
	case "DIV":
		result, _ = calcDiv(firstString, secondString)
		log.Println(result)
		break
	case "MUL":
		result, _ = calcMul(firstString, secondString)
		log.Println(result)
		break

	}

	resultStruct := CalcResult{ResultValue: result}

	resultJSON, err := json.MarshalIndent(resultStruct, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(resultJSON))
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

func calcAdd(a, b string) (string, error) {
	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Add(first, second)

	return result.String(), error
}

func calcSub(a, b string) (string, error) {
	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Sub(first, second)

	return result.String(), error
}

func calcDiv(a, b string) (string, error) {
	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Div(first, second)

	return result.String(), error
}

func calcMul(a, b string) (string, error) {
	first, second, error := convertBigInt(a, b)
	result := new(big.Int)
	result.Mul(first, second)

	return result.String(), error
}
