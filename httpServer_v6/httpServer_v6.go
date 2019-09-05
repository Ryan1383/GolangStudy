package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

// NameCard stuct
type NameCard struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

//포트 등으로 실행이 실패 했을때 알수있는 방법을 추가하세요

var nameCardMap map[string]NameCard

func main() {
	makeNameCards()
	handleHTTPRequest()
	log.Fatal(http.ListenAndServe(":80", nil))
	openBrowser()
}

func makeNameCards() {
	nameCard1 := NameCard{ID: "1", Name: "이인", Address: "수원", Phone: "010-2476-1383"}
	nameCard2 := NameCard{ID: "2", Name: "라이언", Address: "서울", Phone: "02-1444-4421"}
	nameCard3 := NameCard{ID: "3", Name: "제임스", Address: "충주", Phone: "041-322-2212"}
	nameCard4 := NameCard{ID: "4", Name: "트와이스", Address: "서울", Phone: "032-122-3212"}
	nameCard5 := NameCard{ID: "5", Name: "제이슨므라즈", Address: "미국", Phone: "091-352-7512"}
	nameCardMap = map[string]NameCard{
		"1": nameCard1,
		"2": nameCard2,
		"3": nameCard3,
		"4": nameCard4,
		"5": nameCard5,
	}
}

func handleHTTPRequest() {
	http.HandleFunc("/", handleRootView)
	http.HandleFunc("/data", getQueryData)
}

func handleRootView(w http.ResponseWriter, r *http.Request) {
	nameCardView, err := ioutil.ReadFile("./static/view/nameCardView.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(nameCardView))
}

// r.URL.Query를 통해 keys를 추출한다. => 추출할 때 정상적인 쿼리 데이터가 없으면 ok= false가 된다.
func getQueryData(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if ok == false || len(keys[0]) < 1 {
		log.Println("Url Param 'id ' is missing")
		return
	}

	key := keys[0]
	nameCard, match := nameCardMap[string(key)]

	if match == false {
		log.Printf("There is no match nameCard id='%s'", key)
		return
	}

	nameCardJSON, err := json.MarshalIndent(nameCard, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(nameCardJSON))
}

func openBrowser() {
	const url = "http://localhost:80"
	browser.OpenURL(url)
}
