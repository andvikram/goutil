package printly

import (
	"fmt"
	"time"
)

// PrintText will print text to console,
// erase and then print the next one
func PrintText(text string, sleepDuration int) {
	fmt.Print(text)
	durationInMillisecond := time.Millisecond * time.Duration(sleepDuration)
	time.Sleep(durationInMillisecond)
	fmt.Print("\r")
	for i := 0; i < len(text); i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
}
