package main

import (
	"fmt"
	"math/rand"
	//	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nsf/termbox-go"
)

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
		duration := Timer()
		displayTimeASCII(duration)
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

type asciiContainer struct {
	line1 string
	line2 string
	line3 string
	line4 string
	line5 string
}

func fillAsciiContainer(duration time.Duration, timeSlice [][]int) asciiContainer {
	ascii0 := []string{"   ___  ", "  / _ \\ ", " | | | |", " | |_| |", "  \\___/ "}
	ascii1 := []string{"  _ ", " / |", " | |", " | |", " |_|"}
	ascii2 := []string{"  ____  ", " |___ \\ ", "   __) |", "  / __/ ", " |_____|"}
	ascii3 := []string{"  _____ ", " |___ / ", "   |_ \\ ", "  ___) |", " |____/ "}
	ascii4 := []string{"  _  _   ", " | || |  ", " | || |_ ", " |__   _|", "    |_|  "}
	ascii5 := []string{"  ____  ", " | ___| ", " |___ \\ ", "  ___) |", " |____/ "}
	ascii6 := []string{"   __   ", "  / /_  ", " | '_ \\ ", " | (_) |", "  \\___/ "}
	ascii7 := []string{"  _____ ", " |___  |", "    / / ", "   / /  ", "  /_/   "}
	ascii8 := []string{"   ___  ", "  ( _ ) ", "  / _ \\ ", " | (_) |", "  \\___/ "}
	ascii9 := []string{"   ___  ", "  / _ \\ ", " | (_) |", "  \\__, |", "    /_/ "}
	// asciispace := []string{" ", " ", " ", " ", " "}
	asciidot := []string{"    ", "    ", "    ", "  _ ", " (_)"}

	asciiMap := map[int][]string{
		0: ascii0,
		1: ascii1,
		2: ascii2,
		3: ascii3,
		4: ascii4,
		5: ascii5,
		6: ascii6,
		7: ascii7,
		8: ascii8,
		9: ascii9,
	}

	var isMinute bool
	var isDoubleSeconds bool

	if duration.Seconds() > 60 {
		isMinute = true
		//
	} else {
		isMinute = false
		if duration.Seconds() >= 10 {
			isDoubleSeconds = true
		} else {
			isDoubleSeconds = false
		}
	}
	fmt.Println(isMinute)

	var asciiTime asciiContainer
	for i := 0; i <= 5; i++ {
		//hotfix
		if (i == 4 && isDoubleSeconds == false) || (i == 5 && isDoubleSeconds == true) {
			break
		}
		//
		numvalue := timeSlice[0][i]
		asciiTime.line1 += asciiMap[numvalue][0]
		asciiTime.line2 += asciiMap[numvalue][1]
		asciiTime.line3 += asciiMap[numvalue][2]
		asciiTime.line4 += asciiMap[numvalue][3]
		asciiTime.line5 += asciiMap[numvalue][4]
		if (i == 0 && isDoubleSeconds == false) || (i == 1 && isDoubleSeconds == true) {
			asciiTime.line1 += asciidot[0]
			asciiTime.line2 += asciidot[1]
			asciiTime.line3 += asciidot[2]
			asciiTime.line4 += asciidot[3]
			asciiTime.line5 += asciidot[4]
		}
	}
	return asciiTime
}

func slicefyInt(num int) []int {
	var numslice []int
	for num > 0 {
		numslice = append(numslice, num%10)
		num /= 10
	}
	for i, j := 0, len(numslice)-1; i < j; i, j = i+1, j-1 {
		numslice[i], numslice[j] = numslice[j], numslice[i]
	}
	return numslice
}

func displayTimeASCII(duration time.Duration) string {
	// fmt.Println(duration)
	var asciiTime asciiContainer
	if duration.Seconds() > 60 {
		minutes := int(duration.Minutes())
		seconds := int(duration.Seconds()) - (minutes * 60)
		asciiTime = fillAsciiContainer(duration, [][]int{slicefyInt(seconds), slicefyInt(minutes)})
	} else {
		seconds := int(duration)
		asciiTime = fillAsciiContainer(duration, [][]int{slicefyInt(seconds)})
	}

	// fmt.Println("========")
	fmt.Println(asciiTime.line1 + "\n" + asciiTime.line2 + "\n" + asciiTime.line3 + "\n" + asciiTime.line4 + "\n" + asciiTime.line5)
	return (asciiTime.line1 + "\n" + asciiTime.line2 + "\n" + asciiTime.line3 + "\n" + asciiTime.line4 + "\n" + asciiTime.line5)
}
