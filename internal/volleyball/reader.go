package volleyball

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"yaltopia-betting/internal/utils"
)

func ReadVolleyBallPrematch(path string) (VolleyballPrematch, error) {
	vPrematch := VolleyballPrematch{}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read prematch data:", err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("failed to read opened file:")
	}

	if err := json.Unmarshal(buf, &vPrematch); err != nil {
		log.Fatal("failed to unmarshal json data: ", err)
	}
	return vPrematch, nil

}
func ReadVolleyballResult(path string) (VolleyballResult, error) {
	var vbResult VolleyballResult

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read prematch data:", err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("failed to read opened file:")
	}

	if err := json.Unmarshal(buf, &vbResult); err != nil {
		log.Fatal("failed to unmarshal json data: ", err)
	}
	if err != nil {
		return VolleyballResult{}, err
	}
	return vbResult, nil
}

type SampleVolleyballGuess struct {
	Winner       string
	OverUnder    string
	CorrectScore string
}

func MakeVolleyballGuess(input VolleyballPrematch) SampleVolleyballGuess {

	return SampleVolleyballGuess{
		Winner:       WinnerGuess(input),
		OverUnder:    "Not Applicable",
		CorrectScore: CorrectScoreGuess(input),
	}
}

func WinnerGuess(input VolleyballPrematch) string {

	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].Main.Sp.GameLines.Odds))
		randomOdd := input.Results[0].Main.Sp.GameLines.Odds[randomIndex]
		if randomOdd.Header != "" {
			return randomOdd.Header
		}
	}

}

func OverUnderGuess(input VolleyballPrematch) string {
	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].Main.Sp.MatchTotalOddEven.Odds))
		randomOdd := input.Results[0].Main.Sp.GameLines.Odds[randomIndex]
		if randomOdd.Header != "" {
			return randomOdd.Handicap
		}
	}
}

func CorrectScoreGuess(input VolleyballPrematch) string {
	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].Main.Sp.CorrectSetScore.Odds))
		randomOdd := input.Results[0].Main.Sp.CorrectSetScore.Odds[randomIndex]
		if randomOdd.Header != "" {
			return randomOdd.Name
		}
	}
}
