package main

import (
	"strconv"
	"strings"
)

var real = `forward 2
down 9
up 6
forward 1
down 5
down 7
down 9
forward 9
down 8
up 7
forward 2
up 6
forward 4
down 5
down 9
up 1
down 9
forward 8
forward 6
forward 6
forward 5
forward 9
up 3
up 5
forward 1
down 4
down 7
forward 2
up 3
down 8
forward 1
down 2
forward 3
up 1
up 1
up 7
forward 5
up 8
forward 8
forward 8
down 6
forward 1
forward 5
forward 4
forward 6
forward 5
down 6
down 9
forward 9
down 8
forward 6
down 5
forward 9
up 3
up 1
down 8
down 7
down 9
forward 7
down 8
down 9
down 5
down 3
forward 1
forward 6
down 1
forward 9
down 5
forward 7
up 2
down 8
forward 1
down 4
down 9
down 4
up 5
forward 4
forward 6
forward 1
down 3
forward 1
down 6
up 5
up 4
forward 6
forward 1
forward 1
down 2
up 4
up 3
up 2
up 6
down 6
forward 1
down 8
forward 1
up 6
forward 7
down 5
forward 4
forward 6
down 4
forward 4
down 4
down 4
forward 2
forward 8
down 5
down 1
down 8
up 5
up 8
down 5
forward 4
down 6
up 7
forward 2
down 3
forward 2
forward 2
down 9
down 3
up 6
forward 8
up 2
up 9
forward 4
down 1
down 5
forward 4
down 2
down 3
forward 5
down 4
forward 7
up 4
forward 6
up 8
forward 1
up 9
down 4
forward 2
down 1
forward 7
down 3
down 2
forward 5
down 3
down 9
down 9
up 5
forward 2
down 8
up 9
forward 4
down 3
forward 3
forward 6
up 2
forward 3
down 1
down 1
down 1
forward 7
forward 4
forward 7
down 5
down 6
down 2
forward 6
down 3
up 6
forward 4
down 8
up 1
forward 8
down 2
down 5
forward 4
down 9
forward 2
forward 2
down 3
forward 3
down 1
forward 2
down 7
forward 3
forward 9
up 9
forward 6
forward 2
down 1
down 5
forward 6
forward 6
down 3
up 3
forward 9
down 7
down 2
down 4
down 7
forward 5
up 4
forward 8
down 5
forward 7
down 7
up 7
down 8
forward 9
up 5
forward 1
down 2
forward 5
down 9
forward 3
down 5
forward 8
forward 3
up 5
down 2
up 3
forward 2
up 1
up 5
down 8
forward 2
down 5
up 4
up 5
up 2
forward 9
forward 6
down 9
up 9
forward 6
forward 4
forward 3
forward 7
up 1
down 2
down 6
down 1
forward 8
down 1
forward 6
down 8
forward 8
down 7
down 6
down 5
forward 2
up 8
up 6
up 5
down 1
forward 1
down 1
down 5
forward 7
forward 3
down 1
forward 5
forward 5
forward 8
down 1
up 2
down 6
up 6
forward 6
forward 6
down 3
forward 9
up 4
forward 4
down 6
up 1
forward 6
down 2
down 5
down 2
down 6
up 5
down 1
down 1
forward 3
forward 7
forward 3
up 2
down 8
down 4
down 1
down 5
down 1
down 9
forward 6
down 6
down 4
down 6
down 8
forward 4
down 6
down 7
forward 8
down 4
up 4
down 1
forward 1
forward 4
forward 1
up 9
down 7
forward 7
down 4
forward 1
up 4
forward 4
down 5
down 7
forward 5
forward 7
forward 1
forward 1
forward 9
forward 9
up 3
forward 4
down 2
forward 9
up 8
forward 3
up 5
down 3
down 8
forward 8
down 6
forward 1
down 6
down 6
up 9
down 2
forward 8
up 9
down 7
up 9
up 8
up 1
forward 6
forward 9
down 2
forward 8
down 1
up 4
forward 4
forward 7
up 2
forward 4
down 5
forward 3
down 2
down 7
down 4
down 2
up 5
down 5
down 5
down 4
up 1
forward 7
down 6
forward 5
forward 1
down 4
up 9
down 5
forward 7
forward 5
down 6
down 3
down 9
down 1
forward 6
up 2
down 7
down 3
down 6
up 3
down 4
down 4
forward 9
down 3
forward 2
down 9
down 8
up 4
down 2
forward 2
down 5
down 4
down 4
down 2
forward 6
down 3
forward 1
down 4
forward 7
down 5
up 4
down 6
forward 8
down 6
forward 2
forward 4
forward 5
forward 7
forward 4
forward 5
down 8
down 7
forward 3
forward 5
up 7
forward 1
down 4
forward 5
forward 4
forward 4
down 5
down 8
forward 8
down 1
down 1
down 5
up 5
forward 6
down 6
forward 3
forward 4
forward 7
forward 4
down 8
forward 2
down 4
forward 4
down 1
up 2
forward 6
up 1
down 7
down 9
forward 7
forward 2
up 3
down 2
down 9
down 5
up 7
forward 1
forward 8
down 8
up 3
down 3
forward 9
up 4
down 5
up 5
down 1
up 8
forward 9
down 3
up 6
forward 6
forward 1
down 1
forward 9
down 8
forward 8
down 6
up 9
down 4
up 3
up 9
forward 2
down 2
down 2
forward 3
down 2
forward 5
forward 4
up 8
forward 9
up 7
forward 2
down 5
down 6
forward 8
up 7
forward 4
forward 3
up 5
down 8
forward 3
up 2
down 3
forward 6
down 9
down 2
down 6
down 2
forward 7
forward 5
forward 7
down 8
forward 2
down 2
forward 8
up 8
forward 4
forward 3
up 5
down 3
forward 3
up 8
up 7
down 4
down 1
forward 2
down 1
up 6
up 4
down 3
up 1
forward 7
forward 7
forward 7
forward 8
down 1
forward 5
down 6
forward 9
forward 7
forward 7
down 4
up 4
down 6
down 9
up 4
up 2
up 6
forward 4
up 4
up 6
down 2
forward 4
down 9
forward 9
forward 9
down 1
forward 7
down 2
down 7
down 8
down 8
down 9
up 9
down 5
forward 5
forward 7
forward 4
down 7
forward 8
forward 1
down 8
up 9
down 7
forward 9
forward 4
forward 8
down 9
forward 4
down 3
forward 3
down 1
down 1
down 2
up 5
down 2
down 1
down 8
forward 3
up 2
forward 7
down 3
down 8
down 1
forward 4
forward 7
down 5
forward 6
down 6
down 2
forward 6
down 3
up 4
down 7
forward 7
up 1
up 9
down 1
down 2
down 8
down 7
up 1
forward 7
down 2
forward 4
forward 6
forward 9
down 6
forward 2
up 8
down 2
up 2
up 5
down 8
up 6
down 9
forward 6
down 8
down 6
down 1
up 7
up 6
down 8
forward 2
up 7
forward 5
forward 7
forward 7
up 5
forward 2
down 9
up 2
up 8
up 2
down 3
down 7
forward 9
down 3
up 9
forward 8
up 8
forward 4
forward 8
forward 6
up 1
down 3
up 1
down 1
forward 2
forward 1
forward 4
forward 7
up 8
down 9
up 2
down 7
forward 4
down 3
forward 4
forward 2
down 9
forward 8
forward 5
forward 3
down 6
forward 4
forward 4
forward 9
forward 4
up 5
down 7
up 6
forward 5
down 5
forward 4
down 5
forward 7
forward 3
forward 5
down 5
forward 4
down 5
up 4
down 8
up 3
down 3
up 5
forward 4
forward 5
down 6
forward 6
forward 1
forward 8
down 6
down 9
up 5
forward 2
forward 8
up 6
down 6
forward 2
down 8
forward 7
forward 7
down 5
forward 5
forward 8
forward 1
down 4
down 2
down 5
up 4
forward 3
forward 5
down 4
down 7
down 4
up 9
up 6
forward 1
down 8
up 8
up 9
forward 2
forward 1
down 6
forward 6
down 4
forward 7
up 2
up 1
forward 4
down 1
forward 8
forward 3
up 7
up 5
down 1
forward 8
forward 6
up 6
forward 9
down 5
down 9
forward 2
down 3
up 1
up 7
down 1
forward 8
up 9
down 1
down 5
down 7
down 5
down 5
down 5
up 9
forward 9
forward 7
forward 4
forward 6
down 5
down 3
forward 9
forward 1
down 1
down 8
up 4
down 9
forward 9
up 1
down 5
forward 8
up 6
forward 3
down 6
up 8
down 7
forward 3
forward 6
down 7
forward 6
forward 4
forward 4
down 4
forward 6
forward 5
down 6
forward 6
down 7
forward 6
forward 3
up 4
up 2
up 6
down 2
down 8
forward 5
forward 1
up 4
forward 7
forward 9
up 6
down 7
down 3
up 5
forward 5
down 8
up 1
down 1
down 3
down 2
down 1
forward 5
down 3
down 5
forward 7
forward 9
down 3
forward 7
forward 5
forward 4
forward 2
forward 7
forward 8
forward 6
down 8
forward 5
forward 6
forward 6
down 8
down 2
forward 4
down 7
forward 6
down 7
down 4
forward 6
up 6
forward 4
forward 9
forward 2
forward 3
forward 1
down 8
down 3
forward 4
up 3
forward 7
forward 1
down 7
down 8
forward 1
up 8
forward 8
up 8
down 5
forward 6
down 8
down 4
down 9
up 1
down 3
forward 6
down 6
forward 7
forward 3
down 6
down 6
forward 4
down 4
down 1
down 8
forward 2
forward 8
forward 8
down 6
forward 9
down 9
down 5
down 5
forward 7
down 1
forward 1
down 1
down 6
down 1
forward 1
up 6
up 9
forward 5
down 6
forward 8
forward 6
down 7
forward 1
forward 4
forward 9
forward 2
forward 4
down 2
forward 1
forward 8
down 1
down 1
forward 4
down 5
down 3
down 9
down 2
up 8
down 7
down 1
down 9
forward 2
forward 2
up 3
forward 3
down 3
forward 5
forward 9
down 7
up 7
down 9
forward 3
forward 7
down 1
forward 8
down 8
forward 1
down 8
down 6
forward 2
down 3
down 1
down 8
forward 3
up 5
down 7
up 2
up 8
forward 5
up 7
down 6
up 7
down 9
forward 5
up 4
forward 9
down 5
up 7
down 2
up 2
up 7
forward 5
down 6
forward 4
down 4
down 3
forward 2
up 2
down 5
forward 8
down 3
up 7
down 1
down 7
forward 7
forward 4
forward 7
down 2
down 9
down 6
down 9
down 2
down 9
down 7
down 5
forward 4
up 5
up 7
forward 2
forward 7
down 3
down 3
forward 4`


type Direction string

const (
	forward Direction = "forward"
	down    Direction = "down"
	up      Direction = "up"
)

func MatchDirection(maybeDir string) Direction {
	switch maybeDir {
	case string(forward):
		return forward
	case string(down):
		return down
	case string(up):
		return up
	}

	return ""
}

type Movement struct {
	Direction Direction
	Distance  int64
}

var test = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func Parse(ms string) []*Movement {
	mss := strings.Split(ms, "\n")
	moves := make([]*Movement, 0)
	for _, m := range mss {
		toParse := strings.Split(m, " ")
		num, _ := strconv.ParseInt(toParse[1], 10, 64)
		moves = append(moves, &Movement{
			Direction: MatchDirection(toParse[0]),
			Distance:  num,
		})
	}

	return moves
}

func Calc(moves []*Movement) (depth, length int64) {
	for _, m := range moves {
		switch m.Direction {
		case down:
			depth += -m.Distance
		case forward:
			length += m.Distance
		case up:
			depth += m.Distance
		}
	}

	if depth < 0 {
		depth *= -1
	}

	return depth, length
}

func Calc2(moves []*Movement) (depth, length int64) {
	var aim int64

	for _, m := range moves {
		switch m.Direction {
		case down:
			aim += m.Distance
		case forward:
			length += m.Distance
			depth += aim * m.Distance
		case up:
			aim += -m.Distance
		}
	}

	if depth < 0 {
		depth *= -1
	}

	return depth, length
}


func main() {
	a, b := Calc(Parse(real))
	print(a*b)
	
	a, b = Calc2(Parse(real))
	print(a*b)
}
