package common

import (
	"log"
	"strconv"
	"strings"
)

const (
	Won  string = "WIN"
	Lost string = "LOSS"
)

func SettleWinnerGuess(score, winner string) string {

	scores := strings.Split(score, "-")
	scoresInt := []int64{}
	for _, v := range scores {
		intValue, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal("failed to parse int from string: ", err)
		}
		scoresInt = append(scoresInt, intValue)
	}
	if scoresInt[0] > scoresInt[1] && winner == "1" || scoresInt[0] < scoresInt[1] && winner == "2" {
		return Won
	}
	return Lost
}
