package Gotsp

import (
	"fmt"
	"time"
)

type asciiContainer struct {
	line1 string
	line2 string
	line3 string
	line4 string
	line5 string
}

func DisplayTimeASCII(duration time.Duration) {
	timeDecimal := fmt.Sprintf("%.3f", duration.Seconds())
	asciiTime := FillAsciiContainer(timeDecimal)
	fmt.Println(asciiTime.line1 + "\n" + asciiTime.line2 + "\n" + asciiTime.line3 + "\n" + asciiTime.line4 + "\n" + asciiTime.line5)
}

func FillAsciiContainer(timeDecimal string) asciiContainer {
	asciiMap := map[rune][]string{
		'0': ([]string{"   ___  ", "  / _ \\ ", " | | | |", " | |_| |", "  \\___/ "}),
		'1': ([]string{"  _ ", " / |", " | |", " | |", " |_|"}),
		'2': ([]string{"  ____  ", " |___ \\ ", "   __) |", "  / __/ ", " |_____|"}),
		'3': ([]string{"  _____ ", " |___ / ", "   |_ \\ ", "  ___) |", " |____/ "}),
		'4': ([]string{"  _  _   ", " | || |  ", " | || |_ ", " |__   _|", "    |_|  "}),
		'5': ([]string{"  ____  ", " | ___| ", " |___ \\ ", "  ___) |", " |____/ "}),
		'6': ([]string{"   __   ", "  / /_  ", " | '_ \\ ", " | (_) |", "  \\___/ "}),
		'7': ([]string{"  _____ ", " |___  |", "    / / ", "   / /  ", "  /_/   "}),
		'8': ([]string{"   ___  ", "  ( _ ) ", "  / _ \\ ", " | (_) |", "  \\___/ "}),
		'9': ([]string{"   ___  ", "  / _ \\ ", " | (_) |", "  \\__, |", "    /_/ "}),
		'.': ([]string{"    ", "    ", "    ", "  _ ", " (_)"}),
	}
	var asciiTime asciiContainer
	for _, v := range timeDecimal {
		asciiTime.line1 += asciiMap[v][0]
		asciiTime.line2 += asciiMap[v][1]
		asciiTime.line3 += asciiMap[v][2]
		asciiTime.line4 += asciiMap[v][3]
		asciiTime.line5 += asciiMap[v][4]

	}
	asciiTime.line1 += "                "
	asciiTime.line2 += "  ___  ___  ___ "
	asciiTime.line3 += " / __|/ _ \\/ __|"
	asciiTime.line4 += " \\__ \\  __/ (__ "
	asciiTime.line5 += " |___/\\___|\\___|"
	return asciiTime
}
