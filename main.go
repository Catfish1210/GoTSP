package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	fmt.Println("Press spacebar to start/stop the timer")
	duration := Timer()
	fmt.Printf("Timer stopped, Elapsed time: %.3f seconds\n", duration.Seconds())
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
				elapsedTime := fmt.Sprintf("%.3f seconds", timerElapsed.Seconds())

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
