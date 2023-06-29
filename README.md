# GoTSP
Terminal Speedcubing Program written in Go


# Build/Notes

### Imports

* https://github.com/nsf/termbox-go -- to check for keystrokes in terminal
* https://github.com/AlecAivazis/survey/v2 -- for UI/Selector

### Timer
Timer function that measures elapsed time between 2 spacebar presses.

It can be used in various applications where measuring time is essential,
such as performance testing &| time-based challenges (like Speedcubing).

Timer is centered and refreshed with using termbox-go.

Tldr: the Timer function allows users to start and stop the timer and returns the elapsed time between.

##### Flow
1. Wait for user to start timer with spacebar
2. Once timer is started, a goroutine is launched to update the elapsed time continously
3. goroutine calculates the elapsed time and refreshes the display accordingly
4. If user presses spacebar, the timer stops and breaks out of the event-loop
5. Then the function records the stop time and returns the duration between the start and stop times 


### GenerateScramble
Scramblers are imperative in cubing. Without them, you cannot scramble a cube according to regulations, ensure that people have the same scrambles in competition, or practice effectively.

Normal scrambles are 20 moves long, as there is no point going over that because God’s Number says that all states of the cube can be solved in less than 20 moves. Moves such as R2 or U2 are counted as one move.

A letter by itself means a clockwise rotation of a face, while the letter followed by an apostrophe (') means a counterclockwise turn.

WCA lays out its requirements for cube scrambling in a competition. Nearly all cubers will also follow the most basic of these regulations at home, **such as scrambling with white on the top and green on the front.**

In this function, current time `time.Now()` is used as a seed for the Scramble generator

##### Flow
1. Scrambles that count as a double(M2, B2, ..etc) wont have an apostrophe.
2. Scramble doesn't contain two of the same letter moves in a row.
3. Scramble doesn't contain two double-moves in a row.


### DisplayTimeASCII
##### Flow

# Add:
- [x]Timer
- [x]Format Timer return value in 'min:sec' if return value > 60sec
- [x]Func GenerateScramble
- [x]General UI
- [x]Timer result using ascii
- [x]Timer result using ascii
- []Add a function wrapper to survey.question
- []Save attempts in json {Working on it}
- []Display PBtimeDelta on Timer
- [x]Add timer documentation
- []Add UI flow to all options.
- []Center UI with termbox
- []Scramble from []string to string for json/leaderboard
##### 
***Completed 7/14***

# Fix:
panic: runtime error: index out of range [2] with length 2
goroutine 1 [running]:
main.fillAsciiContainer(0x338c37160c, {0xc000193c30, 0x2, 0x7?})
main.go:223 +0xede
main.displayTimeASCII(0x338c37160c)
main.go:189 +0x62d
main.ui()
main.go:48 +0x469
main.main()
main.go:17 +0x57

