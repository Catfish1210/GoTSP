package Gotsp

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func Menu2() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	keyPress := make(chan termbox.Event)
	go func() {
		for {
			keySeq := termbox.PollEvent()
			keyPress <- keySeq
		}
	}()

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				termbox.Close()
				break
			}
		}
		if keySeq.Ch == 'w' {
			displayBanner()
		}
	}

	terminalHeight, terminalWidth := termbox.Size()
	fmt.Println(os.Getuid(), "\n", time.Now(), "\n", terminalHeight, " ", terminalWidth)

}

func displayBanner() {
	terminalWidth, terminalHeight := termbox.Size()
	bannerPosX := (terminalWidth / 2) - (23)
	bannerPosY := terminalHeight / 8

	dynamicPosY := bannerPosY
	for _, line := range Banner {
		dynamicPosX := bannerPosX
		for _, char := range line {
			termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorLightBlue, termbox.ColorDefault)
			dynamicPosX++
		}
		dynamicPosY++
	}
	termbox.SetCell(bannerPosX, dynamicPosY, '0', termbox.ColorRed, termbox.ColorDefault)

	termbox.Sync()
}

// menu options:
// "Speedcube", "Scramble", "View leaderboard", "Timer", "Quit"
