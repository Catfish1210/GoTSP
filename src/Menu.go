package Gotsp

import (
	"fmt"
	"os"
)

func Menu() {
	Uinput := GetUserOption("Welcome to GoTSP", []string{"Speedcube", "Scramble", "View leaderboard", "Timer", "Quit"})

	if Uinput == "Speedcube" {
		fmt.Println("Speedcube time")
		ClearScreen()
		fmt.Println("Apply the scramble with white on the top and green on the front:")
		scramble := GenerateScramble()
		fmt.Println(scramble)
		fmt.Println("\nPress enter if the scramble is applied..")
		fmt.Scanln()

		duration := Timer()
		ResultToJson(scramble, duration)
		ClearScreen()
		DisplayTimeASCII(duration)

	} else if Uinput == "Scramble" {
		ClearScreen()
		fmt.Println("Apply the scramble with white on the top and green on the front:")
		fmt.Println(GenerateScramble())
	} else if Uinput == "View leaderboard" {
		fmt.Println("Nothing much to see here")
	} else if Uinput == "Timer" {
		ClearScreen()
		fmt.Println("Start/stop the timer with spacebar")
		fmt.Println("Press any key to continue to the timer..")
		fmt.Scanln()
		ClearScreen()
		duration := Timer()
		if duration.Seconds() > 60 {
			minutes := int(duration.Minutes())
			seconds := int(duration.Seconds()) - (minutes * 60)
			fmt.Printf("Timer stopped, Elapsed time: %dmin %dsec\n", minutes, seconds)
		} else {
			fmt.Printf("Timer stopped, Elapsed time: %.3f seconds\n", duration.Seconds())
		}
	} else if Uinput == "Quit" {
		ClearScreen()
		os.Exit(0)
	}
	fmt.Println()
	Uinput = GetUserOption("What would you like to do?", []string{"Return to Main menu", "Quit"})
	if Uinput == "Quit" {
		ClearScreen()
		os.Exit(0)
	} else {
		ClearScreen()
		Menu()
	}
}
