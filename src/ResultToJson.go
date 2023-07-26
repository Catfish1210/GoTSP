package Gotsp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type toJson struct {
	Date     string `json:"Date"`
	Time     string `json:"Time"`
	Scramble string `json:"Scramble"`
}

func ResultToJson(scramble []string, duration time.Duration) error {
	var scrambleString string
	for i := 0; i < len(scramble); i++ {
		scrambleString += scramble[i]
		if i != len(scramble)-1 {
			scrambleString += " "
		}
	}

	result := toJson{
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Time:     fmt.Sprintf("%.3fs", duration.Seconds()),
		Scramble: scrambleString,
	}

	fileName := "leaderboard.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		jsonData := []toJson{result}
		fileData, err := json.MarshalIndent(jsonData, "", "    ")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fileName, fileData, 0644)
		if err != nil {
			return err
		}

	} else {
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}

		var jsonData []toJson
		err = json.Unmarshal(file, &jsonData)
		if err != nil {
			return err
		}

		jsonData = append(jsonData, result)

		// sort by time anFunc
		sort.Slice(jsonData, func(i, j int) bool {
			timeI, _ := strconv.ParseFloat(strings.TrimSuffix(jsonData[i].Time, "s"), 64)
			timeJ, _ := strconv.ParseFloat(strings.TrimSuffix(jsonData[j].Time, "s"), 64)
			return timeI < timeJ
		})

		fileData, err := json.MarshalIndent(jsonData, "", "    ")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fileName, fileData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
