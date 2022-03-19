package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	CommandType uint32
	Name        string
}

func NewPlayer(commandType uint32) *player {
	return &player{
		CommandType: commandType,
	}
}

func (p *player) RequestName() {
	fmt.Printf("input name for Player %d: ", p.CommandType)
	fmt.Scanf("%s\n", &p.Name)
}

func (p *player) RequestMove() (int, int, error) {
	fmt.Printf("Player %d:\n", p.CommandType)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.ToLower(strings.Replace(str, "\n", "", -1))
	str = strings.Trim(str, " ")
	ijStr := strings.Split(str, " ")
	if len(ijStr) != 2 {
		return 0, 0, fmt.Errorf("incorrect move")
	}
	ijInt := make([]int, 2)
	for i, s := range ijStr {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			return 0, 0, fmt.Errorf("incorrect move")
		}
		ijInt[i] = intVal
		if intVal != 0 && intVal != 1 && intVal != 2 {
			return 0, 0, fmt.Errorf("incorrect move")
		}
	}
	return ijInt[0], ijInt[1], nil
}
