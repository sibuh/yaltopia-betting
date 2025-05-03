package volleyball

type VolleyballPrematch struct {
	Success int `json:"success"`
	Results []struct {
		FI      string `json:"FI"`
		EventID string `json:"event_id"`
		Main    struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				GameLines struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name,omitempty"`
						Header   string `json:"header"`
						Handicap string `json:"handicap,omitempty"`
					} `json:"odds"`
				} `json:"game_lines"`
				CorrectSetScore struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"correct_set_score"`
				MatchTotalOddEven struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"match_total_odd_even"`
				Set1Lines struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"set_1_lines"`
				Set1ToGoToExtraPoints struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"set_1_to_go_to_extra_points"`
				Set1TotalOddEven struct {
					ID   string        `json:"id"`
					Name string        `json:"name"`
					Odds []interface{} `json:"odds"`
					Open int           `json:"open"`
				} `json:"set_1_total_odd_even"`
			} `json:"sp"`
		} `json:"main"`
		Others []struct {
			UpdatedAt string `json:"updated_at"`
			Sp        struct {
				Set1Lines struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"set_1_lines"`
			} `json:"sp"`
		} `json:"others"`
		Schedule struct {
			UpdatedAt string `json:"updated_at"`
			Key       string `json:"key"`
			Sp        struct {
				Main []struct {
					ID       string `json:"id"`
					Odds     string `json:"odds"`
					Name     string `json:"name"`
					Handicap string `json:"handicap,omitempty"`
				} `json:"main"`
			} `json:"sp"`
		} `json:"schedule"`
	} `json:"results"`
}

type VolleyballResult struct {
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
		Ss     string `json:"ss"`
		Scores struct {
			Num1 struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"1"`
			Num2 struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"2"`
			Num3 struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"3"`
			Num4 struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"4"`
			Num5 struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"5"`
		} `json:"scores"`
		Stats struct {
			PointsWonOnServe []string `json:"points_won_on_serve"`
			LongestStreak    []string `json:"longest_streak"`
		} `json:"stats"`
		Events []struct {
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"events"`
		Extra struct {
			HomePos    string `json:"home_pos"`
			AwayPos    string `json:"away_pos"`
			Bestofsets string `json:"bestofsets"`
			Round      string `json:"round"`
		} `json:"extra"`
		InplayCreatedAt string `json:"inplay_created_at"`
		InplayUpdatedAt string `json:"inplay_updated_at"`
		ConfirmedAt     string `json:"confirmed_at"`
		Bet365ID        string `json:"bet365_id"`
	} `json:"results"`
}
