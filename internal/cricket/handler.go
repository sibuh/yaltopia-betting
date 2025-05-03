package cricket

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CricketBetting(c *fiber.Ctx) error {
	crtPrematch, err := ReadCricketPrematch("./data/cricket_prematch.json")
	if err != nil {
		return err
	}
	crtResult, err := ReadCricketResult("./data/cricket_result.json")
	if err != nil {
		return err
	}
	sampleGuess := MakeCricketGuess(crtPrematch)

	fmt.Println(fmt.Sprintf("score: %s guess:%s", crtResult.Results[0].Ss, sampleGuess.Winner))

	return c.JSON(DoCricketBettingSettlement(crtResult, sampleGuess))

}
