package main

import (
	"fmt"
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
	} else if Uinput == "Scramble" {
		fmt.Println("Get a random scramble between 10-16 moves")
	} else if Uinput == "View leaderboard" {
		fmt.Println("Nothing much to see here")
	} else if Uinput == "Timer" {
		clearScreen()
		fmt.Println("Start/stop the timer with spacebar")
		fmt.Println("Press any key to continue to the timer..")
		fmt.Scanln()
		clearScreen()
		duration := Timer()
		fmt.Printf("Timer stopped, Elapsed time: %.3f seconds\n", duration.Seconds())
		// return to menu or quit prompt
	}
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
