package cricket

type CricketPrematch struct {
	Success int `json:"success"`
	Results []struct {
		FI        string `json:"FI"`
		EventID   string `json:"event_id"`
		OneStOver struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				OneStOverTotalRuns struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"1st_over_total_runs"`
				OneStOverTotalRunsOddEven struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID   string `json:"id"`
						Odds string `json:"odds"`
						Name string `json:"name"`
					} `json:"odds"`
				} `json:"1st_over_total_runs_odd_even"`
			} `json:"sp"`
		} `json:"1st_over"`
		Innings1 struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				OneStInningsScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"1st_innings_score"`
				OneStInningsOfMatchBowledOut struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"1st_innings_of_match_bowled_out?"`
			} `json:"sp"`
		} `json:"innings_1"`
		Main struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				ToWinTheMatch struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"to_win_the_match"`
				TeamTopBatter struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name,omitempty"`
						Header string `json:"header"`
						Team   string `json:"team,omitempty"`
					} `json:"odds"`
				} `json:"team_top_batter"`
				TeamTopBowler struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"team_top_bowler"`
				PlayerOfTheMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_of_the_match"`
				OneStWicketMethod struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"1st_wicket_method"`
				PlayerPerformance struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"player_performance"`
				RajasthanRoyalsVsMumbaiIndians struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
				} `json:"rajasthan_royals_vs_mumbai_indians-"`
			} `json:"sp"`
		} `json:"main"`
		Match struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				TopMatchBatter struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"top_match_batter"`
				TopMatchBowler struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"top_match_bowler"`
				HighestOpeningPartnership struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"highest_opening_partnership"`
				TeamToMakeHighest1St6OversScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"team_to_make_highest_1st_6_overs_score"`
				ToGoToSuperOver struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"to_go_to_super_over?"`
				RunsAtFallOf1StWicket struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"runs_at_fall_of_1st_wicket"`
				MatchRunsAtFallOfFirstWicket3Way struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"match_runs_at_fall_of_first_wicket_(3_way)"`
				MatchRunsAtFallOfFirstWicketTeam3Way struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"match_runs_at_fall_of_first_wicket_team_(3_way)"`
				OneStWicketMethod struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"1st_wicket_method"`
				Match1StWicketMethod2Way struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"match_1st_wicket_method_(2_way)"`
				Match1StWicketMethodTeam2Way struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"match_1st_wicket_method_team_(2_way)"`
				AFiftyToBeScored struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"a_fifty_to_be_scored"`
				AHundredToBeScoredInTheMatch struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"a_hundred_to_be_scored_in_the_match"`
				MostMatchSixes struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID   string `json:"id"`
						Odds string `json:"odds"`
						Name string `json:"name"`
					} `json:"odds"`
				} `json:"most_match_sixes"`
				MostMatchFours struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID   string `json:"id"`
						Odds string `json:"odds"`
						Name string `json:"name"`
					} `json:"odds"`
				} `json:"most_match_fours"`
				MostRunOutsFielding struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"most_run_outs_(fielding)"`
				OneStScoringShotOfTheMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"1st_scoring_shot_of_the_match"`
				HighestIndividualScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID   string `json:"id"`
						Odds string `json:"odds"`
						Name string `json:"name"`
					} `json:"odds"`
				} `json:"highest_individual_score"`
				HighestMatchIndividualScoreTeam struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"highest_match_individual_score_team"`
				SixBoundariesInAnOverMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"six_boundaries_in_an_over_match"`
				RajasthanRoyalsVsMumbaiIndians struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
				} `json:"rajasthan_royals_vs_mumbai_indians-"`
			} `json:"sp"`
		} `json:"match"`
		Others []struct {
			UpdatedAt string `json:"updated_at"`
			Sp        struct {
				TeamTopBowler struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name,omitempty"`
						Header string `json:"header"`
						Team   string `json:"team,omitempty"`
					} `json:"odds"`
					Open int `json:"open"`
				} `json:"team_top_bowler"`
			} `json:"sp"`
		} `json:"others"`
		Player struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				PlayerOfTheMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_of_the_match"`
				PlayerToScoreMost6STeam struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_to_score_most_6s_team"`
				PlayerToScoreMostMatch6S struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_to_score_most_match_6s"`
				PlayerToScoreMostMatch4STeam struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_to_score_most_match_4s_team"`
				PlayerToScoreMostMatch4S struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"player_to_score_most_match_4s"`
				BatterMatchRuns struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"batter_match_runs"`
				BatterMilestones struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Name2  string `json:"name2"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"batter_milestones"`
				BowlerTotalMatchWickets struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"bowler_total_match_wickets"`
				BowlerMilestones struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Name2  string `json:"name2"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"bowler_milestones"`
				PlayerPerformance struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"player_performance"`
				BatterMatchesMostRuns struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"batter_matches_(most_runs)"`
				RaceTo10Runs struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"race_to_10_runs"`
				BatterToScoreAFiftyInTheMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"batter_to_score_a_fifty_in_the_match"`
				BatterToScoreAHundredInTheMatch struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"batter_to_score_a_hundred_in_the_match"`
				BatterTotalMatchFours struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"batter_total_match_fours"`
				BatterTotalMatchSixes struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Name2    string `json:"name2"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"batter_total_match_sixes"`
			} `json:"sp"`
		} `json:"player"`
		Schedule struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				Main []struct {
					ID   string `json:"id"`
					Odds string `json:"odds"`
					Name string `json:"name"`
				} `json:"main"`
			} `json:"sp"`
		} `json:"schedule"`
		Team struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				TeamTopBatter struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name,omitempty"`
						Header string `json:"header"`
						Team   string `json:"team,omitempty"`
					} `json:"odds"`
				} `json:"team_top_batter"`
				TeamTopBowler struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"team_top_bowler"`
			} `json:"sp"`
		} `json:"team"`
	} `json:"results"`
}

//
type CricketResult struct {
	Success int `json:"success"`
	Results []struct {
		ID         string `json:"id"`
		SportID    string `json:"sport_id"`
		Time       string `json:"time"`
		TimeStatus string `json:"time_status"`
		League     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Cc   string `json:"cc"`
		} `json:"league"`
		Home struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			Cc      string `json:"cc"`
		} `json:"home"`
		Away struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			Cc      string `json:"cc"`
		} `json:"away"`
		Ss    string `json:"ss"`
		Extra struct {
			StadiumData struct {
				ID           string `json:"id"`
				Name         string `json:"name"`
				City         string `json:"city"`
				Country      string `json:"country"`
				Capacity     string `json:"capacity"`
				Googlecoords string `json:"googlecoords"`
			} `json:"stadium_data"`
		} `json:"extra"`
		HasLineup       int    `json:"has_lineup"`
		InplayCreatedAt string `json:"inplay_created_at"`
		InplayUpdatedAt string `json:"inplay_updated_at"`
		ConfirmedAt     string `json:"confirmed_at"`
		Bet365ID        string `json:"bet365_id"`
	} `json:"results"`
}