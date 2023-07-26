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

	highlight := 0
	updateMenu(highlight)

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
			highlight--
			updateMenu(highlight)
		}
		if keySeq.Ch == 's' {

			highlight++
			updateMenu(highlight)
		}
	}

	terminalHeight, terminalWidth := termbox.Size()
	fmt.Println(os.Getuid(), "\n", time.Now(), "\n", terminalHeight, " ", terminalWidth)

}

func updateMenu(highlight int) {
	displayBanner()
	displayOptions(highlight)
}

func displayOptions(highlight int) {
	terminalWidth, terminalHeight := termbox.Size()
	optionPosX := (terminalWidth / 2) - (3)
	optionPosY := (terminalHeight / 8) + 8
	dynamicPosY := optionPosY
	for i, option := range Options {
		dynamicPosX := optionPosX - len(Options[i])/2

		for _, char := range option {
			if i != highlight {
				termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorLightMagenta, termbox.ColorDefault)
				dynamicPosX++
			} else {
				termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorLightMagenta|termbox.AttrBold, termbox.ColorDefault)
				dynamicPosX++
			}
		}

		dynamicPosY += 2
	}
	displaySelector(optionPosX, optionPosY, highlight)
	termbox.Sync()

}

func displaySelector(dynamicPosX, dynamicPosY, highlight int) {
	var lSelectorPosY, lSelectorPosX []int
	lSelectorPosY = append(lSelectorPosY, dynamicPosY, dynamicPosY+2, dynamicPosY+4, dynamicPosY+6, dynamicPosY+8)
	lSelectorPosX = append(lSelectorPosX, dynamicPosX-6, dynamicPosX-8, dynamicPosX-10, dynamicPosX-4, dynamicPosX-4)
	var rSelectorPosX []int
	rSelectorPosY := lSelectorPosY
	rSelectorPosX = append(rSelectorPosX, dynamicPosX+6, dynamicPosX+7, dynamicPosX+9, dynamicPosX+4, dynamicPosX+3)

	// Clean all selectors
	for i := 0; i < 5; i++ {
		termbox.SetCell(lSelectorPosX[i], lSelectorPosY[i], ' ', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(rSelectorPosX[i], rSelectorPosY[i], ' ', termbox.ColorDefault, termbox.ColorDefault)

	}

	// Apply selector for the highlighted option
	termbox.SetCell(lSelectorPosX[highlight], lSelectorPosY[highlight], '>', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
	termbox.SetCell(rSelectorPosX[highlight], rSelectorPosY[highlight], '<', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)

	// termbox.SetCell(lSelectorPosX[highlight], lSelectorPosY[highlight], '>', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
	// termbox.SetCell(dynamicPosX+18, dynamicPosY, '<', termbox.ColorWhite|termbox.AttrBlink, termbox.ColorDefault)
	termbox.Sync()
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

	termbox.Sync()
}

// menu options:
// "Speedcube", "Scramble", "View leaderboard", "Timer", "Quit"
