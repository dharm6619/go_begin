package loops

import (
	"fmt"
	"time"
)

func SecondsTimer(timeInSec int) string {

	for timer := timeInSec; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Time Up")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
	return "Time Up"
}
