package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nsf/termbox-go"
)

type asciiContainer struct {
	line1 string
	line2 string
	line3 string
	line4 string
	line5 string
}

type toJson struct {
	Date     string `json:"Date"`
	Time     string `json:"Time"`
	Scramble string `json:"Scramble"`
}

func main() {
	clearScreen()
	ui()
}

func ui() {
	Uinput := ""
	prompt := &survey.Select{
		Message:       "Welcome to GoTSP",
		Options:       []string{"Speedcube", "Scramble", "View leaderboard", "Timer", "Quit"},
		FilterMessage: "s",
		VimMode:       true,
		Filter:        nil,
		Description:   nil,
	}

	survey.AskOne(prompt, &Uinput, nil, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = "!"
		icons.Question.Format = "red+b"
	}))

	if Uinput == "Speedcube" {
		fmt.Println("Speedcube time")
		clearScreen()
		fmt.Println("Apply the scramble with white on the top and green on the front:")
		scramble := GenerateScramble()
		fmt.Println(scramble)
		fmt.Println("\nPress enter if the scramble is applied..")
		fmt.Scanln()

		duration := Timer()
		resultToJson(scramble, duration)
		clearScreen()
		displayTimeASCII(duration)
		// func resultToJson(scramble []string, duration time.duration) {}

	} else if Uinput == "Scramble" {
		clearScreen()
		fmt.Println("Apply the scramble with white on the top and green on the front:")
		fmt.Println(GenerateScramble())
	} else if Uinput == "View leaderboard" {
		fmt.Println("Nothing much to see here")
	} else if Uinput == "Timer" {
		clearScreen()
		fmt.Println("Start/stop the timer with spacebar")
		fmt.Println("Press any key to continue to the timer..")
		fmt.Scanln()
		clearScreen()
		duration := Timer()
		if duration.Seconds() > 60 {
			minutes := int(duration.Minutes())
			seconds := int(duration.Seconds()) - (minutes * 60)
			fmt.Printf("Timer stopped, Elapsed time: %dmin %dsec\n", minutes, seconds)
		} else {
			fmt.Printf("Timer stopped, Elapsed time: %.3f seconds\n", duration.Seconds())
		}
		// return to menu or quit prompt
	}
}

func GenerateScramble() []string {
	var Scramble []string
	Moves := []string{"U", "D", "R", "L", "F", "B", "M"}
	seed := time.Now().UnixNano()
	RandomGenerator := rand.New(rand.NewSource(seed))

	scrambleLength := RandomGenerator.Intn(8) + 14
	var lastScrambleLetter string
	var lastScrambleDouble int
	for i := 0; i < scrambleLength; i++ {
		var currentScramble string
		scrambleMove := RandomGenerator.Intn(7)
		scrambleMoveLetter := Moves[scrambleMove]
		scrambleDouble := RandomGenerator.Intn(2)
		scrambleApostrophe := RandomGenerator.Intn(2)

		if scrambleDouble == 0 {
			if scrambleApostrophe == 1 {
				currentScramble = scrambleMoveLetter + "'"
			} else {
				currentScramble = scrambleMoveLetter
			}
		} else {
			if lastScrambleDouble == 1 {
				currentScramble = scrambleMoveLetter
			} else {
				currentScramble = scrambleMoveLetter + "2"
			}
		}
		if scrambleMoveLetter == lastScrambleLetter {
			i--
			continue
		} else {
			Scramble = append(Scramble, currentScramble)
		}
		lastScrambleLetter = scrambleMoveLetter
		lastScrambleDouble = scrambleDouble
	}
	return Scramble
}

func Timer() time.Duration {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println("Press the spacebar to start the timer")
	timerStart := time.Now()
	timerStopped := true

	termbox.Sync()
	width, height := termbox.Size()
	x := (width - 14) / 2
	y := height / 2

	termbox.SetCursor(x, y)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	elapsedTime := ""
	printTimer(elapsedTime, x, y)
	termbox.HideCursor()
	termbox.Flush()

	go func() {
		for {
			if !timerStopped {
				timerNow := time.Now()
				timerElapsed := timerNow.Sub(timerStart)
				elapsedTime := fmt.Sprintf("%.3f sec", timerElapsed.Seconds())
				printTimer(elapsedTime, x, y)
				termbox.Flush()
			}
			time.Sleep(30 * time.Millisecond)
		}
	}()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeySpace {
			timerStopped = !timerStopped
			if timerStopped {
				break
			}
			timerStart = time.Now()
		}
	}
	timerStop := time.Now()
	return timerStop.Sub(timerStart)
}

func printTimer(str string, x, y int) {
	for i, ch := range str {
		termbox.SetCell(x+i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func displayTimeASCII(duration time.Duration) {
	timeDecimal := fmt.Sprintf("%.3f", duration.Seconds())
	asciiTime := fillAsciiContainer(timeDecimal)
	fmt.Println(asciiTime.line1 + "\n" + asciiTime.line2 + "\n" + asciiTime.line3 + "\n" + asciiTime.line4 + "\n" + asciiTime.line5)
}

func fillAsciiContainer(timeDecimal string) asciiContainer {
	asciiMap := map[rune][]string{
		'0': ([]string{"   ___  ", "  / _ \\ ", " | | | |", " | |_| |", "  \\___/ "}),
		'1': ([]string{"  _ ", " / |", " | |", " | |", " |_|"}),
		'2': ([]string{"  ____  ", " |___ \\ ", "   __) |", "  / __/ ", " |_____|"}),
		'3': ([]string{"  _____ ", " |___ / ", "   |_ \\ ", "  ___) |", " |____/ "}),
		'4': ([]string{"  _  _   ", " | || |  ", " | || |_ ", " |__   _|", "    |_|  "}),
		'5': ([]string{"  ____  ", " | ___| ", " |___ \\ ", "  ___) |", " |____/ "}),
		'6': ([]string{"   __   ", "  / /_  ", " | '_ \\ ", " | (_) |", "  \\___/ "}),
		'7': ([]string{"  _____ ", " |___  |", "    / / ", "   / /  ", "  /_/   "}),
		'8': ([]string{"   ___  ", "  ( _ ) ", "  / _ \\ ", " | (_) |", "  \\___/ "}),
		'9': ([]string{"   ___  ", "  / _ \\ ", " | (_) |", "  \\__, |", "    /_/ "}),
		'.': ([]string{"    ", "    ", "    ", "  _ ", " (_)"}),
	}
	var asciiTime asciiContainer
	for _, v := range timeDecimal {
		asciiTime.line1 += asciiMap[v][0]
		asciiTime.line2 += asciiMap[v][1]
		asciiTime.line3 += asciiMap[v][2]
		asciiTime.line4 += asciiMap[v][3]
		asciiTime.line5 += asciiMap[v][4]

	}
	asciiTime.line1 += "                "
	asciiTime.line2 += "  ___  ___  ___ "
	asciiTime.line3 += " / __|/ _ \\/ __|"
	asciiTime.line4 += " \\__ \\  __/ (__ "
	asciiTime.line5 += " |___/\\___|\\___|"
	return asciiTime
}

func resultToJson(scramble []string, duration time.Duration) error {
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
