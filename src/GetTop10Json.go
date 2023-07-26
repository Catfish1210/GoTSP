package Gotsp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type fromJson struct {
	Date     string `json:"Date"`
	Time     string `json:"Time"`
	Scramble string `json:"Scramble"`
}

func GetTop10() [][]string {
	fileName := "leaderboard.json"
	var jsonData []fromJson
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Printf("ERROR: leaderboard file not found")
	} else {
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("ERROR: unable to read from file")
		}
		err = json.Unmarshal(file, &jsonData)
	}
	var top10 [][]string
	for i, v := range jsonData {
		if i < 10 {
			position := fmt.Sprint(i + 1)
			time := v.Time
			date := (v.Date)[:10]
			scramble := v.Scramble
			toTop10 := []string{position, time, date, scramble}
			top10 = append(top10, toTop10)
		}
	}
	return top10
}
