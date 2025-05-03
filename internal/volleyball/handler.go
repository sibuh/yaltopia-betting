package volleyball

import (
	"log"
	"yaltopia-betting/internal/common"

	"github.com/gofiber/fiber/v2"
)

func VolleyballBetting(c *fiber.Ctx) error {
	prematch, _ := ReadVolleyBallPrematch("./data/volleyball_prematch.json")
	guess := MakeVolleyballGuess(prematch)

	result, _ := ReadVolleyballResult("./data/volleyball_result.json")

	if err := c.JSON(DoSettlement(result.Results[0].Ss, guess)); err != nil {
		log.Fatal("error occurred when writing body", err)
	}
	return nil
}

type Settlment struct {
	Winner       string
	OverUnder    string
	CorrectScore string
}

func SettleCorrectScoreGuess(score, guess string) string {
	if score == guess {
		return common.Won
	} else {
		return common.Lost
	}

}

func DoSettlement(score string, bettorGuess SampleVolleyballGuess) map[string]string {
	return map[string]string{
		"WinLoss":      common.SettleWinnerGuess(score, bettorGuess.Winner),
		"OverUnder":    "odds for OverUnder do not exist",
		"CorrectScore": SettleCorrectScoreGuess(score, bettorGuess.CorrectScore),
	}

}
