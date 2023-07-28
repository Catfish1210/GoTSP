package Gotsp

import (
	"fmt"
	"os"

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
			} else if keySeq.Key == termbox.KeyEnter || keySeq.Key == termbox.KeySpace {
				selectMenuItem(highlight)
				termbox.Close()
			}
		}
		if keySeq.Ch == 'w' && highlight > 0 {
			highlight--
			updateMenu(highlight)
		}
		if keySeq.Ch == 's' && highlight < 4 {
			highlight++
			updateMenu(highlight)
		}
	}

}

func selectMenuItem(highlight int) {
	termbox.Close()
	if highlight == 0 {
		Speedcube()
	}
	if highlight == 2 {
		ViewLeaderboard()
	}
}

func ViewLeaderboard() {
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

	a := GetTop10()
	updateTop10Selection(a, 0)
	//

	//

	termbox.Sync()

	for {
		keySeq := <-keyPress
		if keySeq.Type == termbox.EventKey {
			if keySeq.Key == termbox.KeyEsc || keySeq.Key == termbox.KeyCtrlC {
				ClearScreen()
				termbox.Close()
				os.Exit(0)
				break
			}
			if keySeq.Key == termbox.KeySpace {

			}

			if keySeq.Ch == 'w' {

			}
			if keySeq.Ch == 's' {

			}
		}
	}
}

func updateTop10Selection(top10 [][]string, selected int) {
	terminalWidth, _ := termbox.Size()
	// dynamicPosX := terminalWidth / 2 - (len(Top10Banner)/2)
	dynamicPosY := 2

	for _, line := range Top10Banner {
		dynamicPosX := (terminalWidth / 2) - (len(Top10Banner[0]) / 2)
		for _, ch := range line {
			termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorMagenta|termbox.AttrBold, termbox.ColorDefault)
			dynamicPosX++
		}
		dynamicPosY++
	}
	dynamicPosY++
	dynamicPosX := (terminalWidth / 2) - (len(top10[0][0]+top10[0][1]+top10[0][2]+top10[0][3]) / 2)
	termbox.SetCell(dynamicPosX, dynamicPosY, 'A', termbox.ColorRed, termbox.ColorDefault)

	for _, scoreEntry := range top10 {

		// Display position:
		for i, pos := range scoreEntry[1] {
			termbox.SetCell(dynamicPosX, dynamicPosY, pos, termbox.ColorRed, termbox.ColorDefault)
			if i > 1 {
				dynamicPosX++
			}

		}

		// Display time:
		// Display Date

		dynamicPosY++
		// if selected and enter --> display the scramble

	}

}

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

func cleanTermbox() {
	w, h := termbox.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}

func displayText(x, y int, text string) {
	dynamicPosX, dynamicPosY := x, y
	dynamicPosX -= (len(text) / 2)
	for _, char := range text {
		termbox.SetCell(dynamicPosX, dynamicPosY, char, termbox.ColorDefault, termbox.ColorDefault)
		dynamicPosX++
	}
	termbox.Sync()
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
