package Gotsp

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func Speedcube() {
	ClearScreen()
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
	timerStatus := false
	terminalWidth, terminalHeight := termbox.Size()
	dynamicPosX := terminalWidth / 2
	dynamicPosY := (terminalHeight / 2) - 1
	displayText(dynamicPosX, dynamicPosY, "Apply the scramble with white on the top and green on the front:")
	scramble := GenerateScramble()
	var scrambleString string
	for i := 0; i < len(scramble); i++ {
		scrambleString += scramble[i]
		if i != len(scramble)-1 {
			scrambleString += " "
		}
	}
	dynamicPosY++
	displayText(dynamicPosX, dynamicPosY, scrambleString)
	dynamicPosY += 2
	displayText(dynamicPosX, dynamicPosY, "Press spacebar if the scramble is applied & you're ready to start the timer..")
	userChoice := -1
	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				ClearScreen()
				termbox.Close()
				os.Exit(0)
				break
			}
			if keySeq.Key == termbox.KeySpace && timerStatus != true {
				duration := Timer2()
				ResultToJson(scramble, duration)
				timerStatus = true
				timeDecimal := fmt.Sprintf("%.3f", duration.Seconds())
				asciiTime := FillAsciiContainer(timeDecimal)
				cleanTermbox()
				terminalWidth, terminalHeight := termbox.Size()
				asciiPosX := (terminalWidth / 2)
				asciiPosY := (terminalHeight / 2) - (terminalHeight / 5)
				displayText(asciiPosX, asciiPosY+1, asciiTime.line1)
				displayText(asciiPosX, asciiPosY+2, asciiTime.line2)
				displayText(asciiPosX, asciiPosY+3, asciiTime.line3)
				displayText(asciiPosX, asciiPosY+4, asciiTime.line4)
				displayText(asciiPosX, asciiPosY+5, asciiTime.line5)
				updateUserOption(asciiPosX, asciiPosY+7, -1)
				termbox.Sync()
			}
			if keySeq.Key == termbox.KeySpace && timerStatus == true {
				if userChoice == 0 {
					ClearScreen()
					termbox.Close()
					Menu2()
					os.Exit(0)
					break
				}
				if userChoice == 1 {
					ClearScreen()
					termbox.Close()
					os.Exit(0)
					break
				}
			}
		}
		if keySeq.Ch == 'a' {
			userChoice = 0
			updateUserOption((terminalWidth / 2), ((terminalHeight/2)-(terminalHeight/5))+7, userChoice)
			termbox.Sync()
		}
		if keySeq.Ch == 'd' {
			userChoice = 1
			updateUserOption((terminalWidth / 2), ((terminalHeight/2)-(terminalHeight/5))+7, userChoice)
			termbox.Sync()
		}
	}
}

func updateUserOption(x, y, highlighted int) {
	exitOptions := []string{
		"Return to Menu",
		"Quit",
	}
	// Clean previous selection
	xj, _ := termbox.Size()
	for xi := 0; xi < xj; xi++ {
		termbox.SetCell(xi, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	// Display selection w/ highlight
	for i, option := range exitOptions {
		if i == 0 {
			x -= 15
		}
		if i == 1 {
			x += 10
		}

		for j, ch := range option {
			if highlighted == 0 && i == 0 {
				if j == 0 {
					termbox.SetCell((x - 2), y, '>', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
				}
				if j == len(option)-1 {
					termbox.SetCell((x + 2), y, '<', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
				}
				termbox.SetCell(x, y, ch, termbox.ColorLightMagenta|termbox.AttrBold, termbox.ColorDefault)
				x++
			} else if highlighted == 1 && i == 1 {
				if j == 0 {
					termbox.SetCell((x - 2), y, '>', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
				}
				if j == len(option)-1 {
					termbox.SetCell((x + 2), y, '<', termbox.ColorWhite|termbox.AttrBlink|termbox.AttrBold, termbox.ColorDefault)
				}
				termbox.SetCell(x, y, ch, termbox.ColorLightMagenta|termbox.AttrBold, termbox.ColorDefault)
				x++
			} else {
				termbox.SetCell(x, y, ch, termbox.ColorLightMagenta, termbox.ColorDefault)
				x++
			}
		}
	}
}

// Center to x y pos
func displayText(x, y int, text string) {
	dynamicPosX, dynamicPosY := x, y
	dynamicPosX -= (len(text) / 2)
	for _, char := range text {
		termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorDefault, termbox.ColorDefault)
		dynamicPosX++
	}
	termbox.Sync()
}
