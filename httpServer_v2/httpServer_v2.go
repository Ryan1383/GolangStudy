package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NameCard stuct
type NameCard struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

const noSearchDataScreen = `
	<html>
		<body>
			<p>There is no data
		</body>
	</html>
`

var nameCardMap map[string]NameCard

func main() {
	nameCard1 := NameCard{ID: "1", Name: "이인", Address: "수원", Phone: "010-2476-1383"}
	nameCard2 := NameCard{ID: "2", Name: "라이언", Address: "서울", Phone: "02-1444-4421"}
	nameCard3 := NameCard{ID: "3", Name: "제임스", Address: "충주", Phone: "041-322-2212"}
	nameCard4 := NameCard{ID: "4", Name: "트와이스", Address: "서울", Phone: "032-122-3212"}
	nameCard5 := NameCard{ID: "5", Name: "컴퓨터", Address: "미국", Phone: "091-352-7512"}

	nameCardMap = map[string]NameCard{
		"1": nameCard1,
		"2": nameCard2,
		"3": nameCard3,
		"4": nameCard4,
		"5": nameCard5,
	}

	http.HandleFunc("/data", getQueryData)
	log.Fatal(http.ListenAndServe(":80", nil))

}

func getQueryData(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'id ' is missing")
		return
	}

	key := keys[0]
	nameCard, ok := nameCardMap[string(key)]

	if !ok {
		screen := fmt.Sprintf(noSearchDataScreen)
		fmt.Fprint(w, screen)
		return
	}

	nameCardJSON, err := json.MarshalIndent(nameCard, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(nameCardJSON))
}
