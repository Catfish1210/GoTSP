package Gotsp

import (
	"math/rand"
	"time"
)

func GenerateScramble() []string {
	var Scramble []string
	Moves := []string{"U", "D", "R", "L", "F", "B", "M"}
	seed := time.Now().UnixNano()
	RandomGenerator := rand.New(rand.NewSource(seed))

	scrambleLength := RandomGenerator.Intn(8) + 14
	var lastScrambleLetter string
	var lastScrambleDouble int
	for i := 0; i < scrambleLength; i++ {
		var currentScramble string
		scrambleMove := RandomGenerator.Intn(7)
		scrambleMoveLetter := Moves[scrambleMove]
		scrambleDouble := RandomGenerator.Intn(2)
		scrambleApostrophe := RandomGenerator.Intn(2)

		if scrambleDouble == 0 {
			if scrambleApostrophe == 1 {
				currentScramble = scrambleMoveLetter + "'"
			} else {
				currentScramble = scrambleMoveLetter
			}
		} else {
			if lastScrambleDouble == 1 {
				currentScramble = scrambleMoveLetter
			} else {
				currentScramble = scrambleMoveLetter + "2"
			}
		}
		if scrambleMoveLetter == lastScrambleLetter {
			i--
			continue
		} else {
			Scramble = append(Scramble, currentScramble)
		}
		lastScrambleLetter = scrambleMoveLetter
		lastScrambleDouble = scrambleDouble
	}
	return Scramble
}
