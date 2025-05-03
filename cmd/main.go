package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"yaltopia-betting/internal/volleyball"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{

		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
	}))
	// ... (define routes and other settings) ...
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/vb", VolleyballBetting)

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("User ID: " + id)
	})
	app.Listen(":3000") // Start the server on port 3000
}

const (
	Won  string = "WON"
	Lost string = "LOST"
)

func VolleyballBetting(c *fiber.Ctx) error {
	prematch, _ := volleyball.ReadVolleyBallPrematch("./data/volleyball_prematch.json")
	guess := volleyball.MakeVolleyballGuess(prematch)
	fmt.Println("This is bettor guess:==>", guess)
	result, _ := volleyball.ReadVolleyballResult("./data/volleyball_result.json")

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
	if scoresInt[0] > scoresInt[1] && winner == "1" {
		return Won
	}
	return Lost
}
func SettleCorrectScoreGuess(score, guess string) string {
	if score == guess {
		return Won
	} else {
		return Lost
	}

}

func DoSettlement(score string, bettorGuess volleyball.SampleVolleyballGuess) map[string]string {
	return map[string]string{
		"CorrectScore": SettleCorrectScoreGuess(score, bettorGuess.CorrectScore),
		"WinLoss":      SettleWinnerGuess(score, bettorGuess.Winner),
		"OverUnder":    "odds for OverUnder do not exist",
	}

}
