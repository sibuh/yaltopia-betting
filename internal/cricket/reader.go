package cricket

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"yaltopia-betting/internal/common"
	"yaltopia-betting/internal/utils"
)

func ReadCricketPrematch(path string) (CricketPrematch, error) {
	vPrematch := CricketPrematch{}

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

func ReadCricketResult(path string) (CricketResult, error) {
	var vbResult CricketResult

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
		return CricketResult{}, err
	}
	return vbResult, nil
}

type SampleCricketGuess struct {
	Winner       string
	OverUnder    string
	CorrectScore string
}

func MakeCricketGuess(input CricketPrematch) SampleCricketGuess {

	return SampleCricketGuess{
		Winner:       CricketWinnerGuess(input),
		OverUnder:    "No data in result for over/under market to settle bettor's guess", //CricketOverUnderGuess(input),
		CorrectScore: "Not Applicable",
	}
}

func CricketWinnerGuess(input CricketPrematch) string {

	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].Main.Sp.ToWinTheMatch.Odds))
		randomOdd := input.Results[0].Main.Sp.ToWinTheMatch.Odds[randomIndex]
		if randomOdd.Name != "" {
			return randomOdd.Name
		}
	}

}

func CricketOverUnderGuess(input CricketPrematch) map[string]string {
	overUnderGuesses := make(map[string]string)
	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].Innings1.Sp.OneStInningsScore.Odds))
		randomOdd := input.Results[0].Innings1.Sp.OneStInningsScore.Odds[randomIndex]
		if randomOdd.Name != "" {
			overUnderGuesses["Innigs1"] = randomOdd.Name
			break
		}
	}
	for {
		randomIndex := utils.RandomIndex(len(input.Results[0].OneStOver.Sp.OneStOverTotalRuns.Odds))
		randomOdd := input.Results[0].Innings1.Sp.OneStInningsScore.Odds[randomIndex]
		if randomOdd.Name != "" {
			overUnderGuesses["OneStOver"] = randomOdd.Name
			break
		}
	}
	return overUnderGuesses
}

func CricketCorrectScoreGuess(input CricketPrematch) string {
	// for {
	// 	randomIndex := utils.RandomIndex(len(input.Results[0].Main.Sp.CorrectSetScore.Odds))
	// 	randomOdd := input.Results[0].Main.Sp.CorrectSetScore.Odds[randomIndex]
	// 	if randomOdd.Header != "" {
	// 		return randomOdd.Name
	// 	}
	// }
	return ""
}

func DoCricketBettingSettlement(result CricketResult, bettorGuess SampleCricketGuess) map[string]string {
	return map[string]string{
		"WinLoss":      common.SettleWinnerGuess(result.Results[0].Ss, bettorGuess.Winner),
		"OverUnder":    "odds for OverUnder do not exist",
		"CorrectScore": "This market is not available currently",
	}
}
