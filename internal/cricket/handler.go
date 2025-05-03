package cricket

import (
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

	return c.JSON(DoCricketBettingSettlement(crtResult, sampleGuess))

}
