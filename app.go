package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// getPointer()

	stepClick()

}
func stepClick() {
	duration := time.Duration(2) * time.Second

	// เลือก applet
	robotgo.MovesClick(869, 565, 1, "left")
	robotgo.MoveClick(869, 565, 1, "left")

	time.Sleep(duration)

	// click setting
	robotgo.MovesClick(934, 340, 1, "left")
	robotgo.MoveClick(934, 340, 1, "left")

	time.Sleep(duration)

	// scroll down
	robotgo.ScrollMouse(100, "down")

	// click delete button
	robotgo.MovesClick(768, 617, 1, "left")
	robotgo.MoveClick(768, 617, 1, "left")

	time.Sleep(duration)

	// click confirm
	robotgo.MovesClick(955, 563, 1, "left")
	robotgo.MoveClick(955, 563, 1, "left")

	time.Sleep(duration)

	// click confirm
	robotgo.MovesClick(955, 563, 1, "left")
	robotgo.MoveClick(955, 563, 1, "left")

	time.Sleep(duration)

	robotgo.ScrollMouse(200, "up")

	// click main menu
	robotgo.MovesClick(129, 262, 1, "left")
	robotgo.MoveClick(129, 262, 1, "left")
	time.Sleep(duration)

}
func getPointer() {
	duration := time.Duration(2) * time.Second
	time.Sleep(duration)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}
