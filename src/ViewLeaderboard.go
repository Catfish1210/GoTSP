package Gotsp

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

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
	selected := 0
	updateTop10Selection(GetTop10(), selected)
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
			if keySeq.Key == termbox.KeySpace || keySeq.Key == termbox.KeyEnter {
				updateTop10Selection(GetTop10(), selected)
				go updateChosenScramble(GetTop10(), selected)
			}
			if keySeq.Ch == 'w' && selected > 0 {
				selected--
				updateTop10Selection(GetTop10(), selected)
			}
			if keySeq.Ch == 's' && selected < 9 {
				selected++
				updateTop10Selection(GetTop10(), selected)
			}
		}
	}
}

func updateTop10Selection(top10 [][]string, selected int) {
	terminalWidth, _ := termbox.Size()
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
	for i, scoreEntry := range top10 {
		dynamicPosX := (terminalWidth / 2) - (len(top10[0][0]+top10[0][1]+top10[0][2]) / 2) - 1
		pos := i + 1
		posString := fmt.Sprint(pos) + ". "
		if pos == 10 {
			posString = fmt.Sprint(pos) + "."
		}
		for _, ch := range posString {
			termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorDefault, termbox.ColorDefault)
			if i == selected {
				termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorRed|termbox.AttrBold, termbox.ColorDefault)
			}
			dynamicPosX++
		}
		dynamicPosX++
		for _, ch := range scoreEntry[1] {
			termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorDefault, termbox.ColorDefault)
			if i == selected {
				termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorRed|termbox.AttrBold, termbox.ColorDefault)
			}
			dynamicPosX++
		}
		dynamicPosX++
		for _, ch := range scoreEntry[2] {
			termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorDefault, termbox.ColorDefault)
			if i == selected {
				termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorRed|termbox.AttrBold, termbox.ColorDefault)
			}
			dynamicPosX++
		}
		dynamicPosY++
	}
	termbox.Sync()
}

func showScramble(scramble string, dynamicPosX, dynamicPosY int) {
	cleanX := dynamicPosX
	for _, ch := range scramble {
		termbox.SetCell(dynamicPosX, dynamicPosY, ch, termbox.ColorDefault, termbox.ColorDefault)

		dynamicPosX++
	}
	termbox.Sync()
	time.Sleep(5 * time.Second)
	dynamicPosX = cleanX
	for i := 0; i < len(scramble); i++ {
		termbox.SetCell(dynamicPosX, dynamicPosY, ' ', termbox.ColorDefault, termbox.ColorDefault)
		dynamicPosX++
	}
	termbox.Sync()
}

func updateChosenScramble(top10 [][]string, selected int) {
	dynamicY := 9 + selected
	terminalWidth, _ := termbox.Size()
	dynamicX := (terminalWidth / 2) - (len(top10[0][0]+top10[0][1]+top10[0][2]) / 2) + 22
	delScramble := dynamicX
	scrambleString := "> " + top10[selected][3]
	for i, ch := range scrambleString {
		if i != 1 && ch != ' ' {
			termbox.SetCell(dynamicX, dynamicY, ch, termbox.ColorMagenta|termbox.AttrBlink, termbox.ColorDefault)
		}
		dynamicX++
	}
	termbox.Sync()
	time.Sleep(6 * time.Second)
	for _, ch := range scrambleString {
		_ = ch
		termbox.SetCell(delScramble, dynamicY, ' ', termbox.ColorDefault, termbox.ColorDefault)
		delScramble++
	}
	termbox.Sync()
}
