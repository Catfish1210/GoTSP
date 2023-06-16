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

	timerStart := time.Now()
	timerStopped := true
	go func() {
		for {
			if !timerStopped {
				timerNow := time.Now()
				timerElapsed := timerNow.Sub(timerStart)
				fmt.Printf("\rTIME: %.3f seconds", timerElapsed.Seconds())
			}
			time.Sleep(1000)
		}
	}()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeySpace {
			timerStopped = !timerStopped

			if timerStopped {
				break
			}

			fmt.Println("\nTimer Started")
			timerStart = time.Now()
		}
	}
	timerStop := time.Now()
	return timerStop.Sub(timerStart)

}
