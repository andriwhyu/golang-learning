package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
	"time"
)

// triangular wave algorithm
//func ballPosition(currentVal, limit int) int {
//	// cycleLen will make sure it will calculate increasing until the limit and decreasing to the starting point
//	cycleLen := 2 * limit
//	position := currentVal % cycleLen
//
//	if position < limit {
//		return position
//	}
//
//	return cycleLen - position
//}

func main() {

	//1st approach without slices
	//x, y := 46, 10
	//var (
	//	i, iX, iY int
	//)
	//screen.Clear()
	//
	//for {
	//	xPos := ballPosition(iX, x)
	//	yPos := ballPosition(iY, y)
	//
	//	screen.MoveTopLeft()
	//	for row := 0; row <= y; row++ {
	//		for col := 0; col <= x; col++ {
	//			printObj := " "
	//			if row == yPos && col == xPos {
	//				printObj = "⚾️"
	//			}
	//			fmt.Print(printObj)
	//		}
	//		fmt.Println()
	//	}
	//	time.Sleep(50 * time.Millisecond)
	//
	//	i++
	//	iX += 5
	//	iY += 2
	//}
	//fmt.Println(2)

	// 2nd approach with slices
	const (
		maxFrame = 1200
		speed    = time.Second / 20

		emptyCell = ' '
		ballCell  = '⚾'
	)

	var (
		// rune is use for store a single character.
		printObj      rune
		width, height = screen.Size()

		px, py int
		//py, px = screen.Size()
		vx, vy = 1, 1
	)

	ballWidth := runewidth.RuneWidth(ballCell)

	height -= ballWidth
	width /= ballWidth // divide the width with the rune width

	board := make([][]bool, width)
	// buffer is used to store the printed value in memory and print it at the end of the program.
	// print in a loop will occupy a big process due print will directly access the I/O, and it will consume resources.
	// (width*2+1) * height purpose is to facilitate extra whitespace and new line.
	buf := make([]rune, 0, (width*2+1)*height)

	for row := range board {
		board[row] = make([]bool, height)
	}

	board[px][py] = true

	screen.Clear()
	for i := 0; i < maxFrame; i++ {
		buf = buf[:0]

		screen.MoveTopLeft()
		for y := range board[0] {
			for x := range board {
				printObj = emptyCell
				if board[x][y] {
					printObj = ballCell
				}
				buf = append(buf, printObj, ' ')
			}
			buf = append(buf, '\n')
		}

		fmt.Println(string(buf))
		time.Sleep(speed)

		board[px][py] = false

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		board[px][py] = true
	}
}
