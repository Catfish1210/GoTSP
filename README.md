# GoTSP
Terminal Speedcubing Program written in Go


# Build/Notes

### Imports

* https://github.com/nsf/termbox-go -- to check for keystrokes in terminal
* https://github.com/AlecAivazis/survey/v2 -- for UI/Selector

### GenerateScramble
Scramblers are imperative in cubing. Without them, you cannot scramble a cube according to regulations, ensure that people have the same scrambles in competition, or practice effectively.

Normal scrambles are 20 moves long, as there is no point going over that because God’s Number says that all states of the cube can be solved in less than 20 moves. Moves such as R2 or U2 are counted as one move.

A letter by itself means a clockwise rotation of a face, while the letter followed by an apostrophe (') means a counterclockwise turn.

WCA lays out its requirements for cube scrambling in a competition. Nearly all cubers will also follow the most basic of these regulations at home, **such as scrambling with white on the top and green on the front.**

In this function, current time `time.Now()` is used as a seed for the Scramble generator

##### Flow
**(1)** Scrambles that count as a double(M2, B2, ..etc) wont have an apostrophe.
**(2)** Scramble doesn't contain two of the same letter moves in a row.
**(3)** Scramble doesn't contain two double-moves in a row.


# Add:
- [x]Timer
- []Format Timer return value in 'min:sec:ns' if return value > 60sec
- [x]Func GenerateScramble
- [x]General UI
- []Timer using ascii
- []Save attempts in json
- []Display scramble result with #'s and color
- []Display PBtimeDelta on Timer
- []Add timer documentation
- []Improve GenerateScramble flow with using prefix aproach ("2"||"'"||"")


