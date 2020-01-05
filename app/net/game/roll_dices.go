package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func handlerRolls2(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	faces, _ := strconv.Atoi(vals.Get("faces"))
	dices, _ := strconv.Atoi(vals.Get("dices"))
	rolls, _ := strconv.Atoi(vals.Get("rolls"))

	fmt.Fprintf(w, rollDices(faces, dices, rolls))
	fmt.Printf("rolls: dices=%d faces=%d rolls=%d \n", dices, faces, rolls)
}

func handlerOneRollOneDice(w http.ResponseWriter, r *http.Request) {
	faces := 6
	dices := 1
	rolls := 1
	fmt.Fprintf(w, rollDices(faces, dices, rolls))
	fmt.Printf("rolls: dices=%d faces=%d rolls=%d \n", dices, faces, rolls)
}

func main() {
	http.HandleFunc("/api/roll", handlerRolls2)
	http.HandleFunc("/api", handlerOneRollOneDice)
	port := "8888"
	fmt.Println("started game at :", 8888)
	http.ListenAndServe(":"+port, nil)
}

func rollDices(faces, dices, rolls int) string {
	var result strings.Builder
	pointsList := make([]string, 0)
	oneRollPoints := make([]int, 0)
	//var points int
	for i := 0; i < rolls; i++ {
		for j := 0; j < dices; j++ {
			points := rand.Intn(faces) + 1
			oneRollPoints = append(oneRollPoints, points)
		}
		result.Reset()
		for _, point := range oneRollPoints {
			result.WriteString(strconv.Itoa(point))
			result.WriteString(" ")
		}
		oneRollPoints = nil
		pointsList = append(pointsList, result.String())
	}
	result.Reset()
	for _, point := range pointsList {
		result.WriteString(point + "\n")
	}
	return result.String()
}
