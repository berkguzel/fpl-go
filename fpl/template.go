package fpl

import (
	"time"
)

type General struct {
	Events []struct {
		ID                     int    `json:"id"`
		Name                   string `json:"name"`
		DeadlineTime           string `json:"deadline_time"`
		AverageEntryScore      int    `json:"average_entry_score"`
		Finished               bool   `json:"finished"`
		DataChecked            bool   `json:"data_checked"`
		HighestScoringEntry    int    `json:"highest_scoring_entry"`
		DeadlineTimeEpoch      int    `json:"deadline_time_epoch"`
		DeadlineTimeGameOffset int    `json:"deadline_time_game_offset"`
		HighestScore           int    `json:"highest_score"`
		IsPrevious             bool   `json:"is_previous"`
		IsCurrent              bool   `json:"is_current"`
		IsNext                 bool   `json:"is_next"`
		ChipPlays              []struct {
			ChipName  string `json:"chip_name"`
			NumPlayed int    `json:"num_played"`
		} `json:"chip_plays"`
		MostSelected      int `json:"most_selected"`
		MostTransferredIn int `json:"most_transferred_in"`
		TopElement        int `json:"top_element"`
		TopElementInfo    struct {
			ID     int `json:"id"`
			Points int `json:"points"`
		} `json:"top_element_info"`
		TransfersMade     int `json:"transfers_made"`
		MostCaptained     int `json:"most_captained"`
		MostViceCaptained int `json:"most_vice_captained"`
	} `json:"events"`
	GameSettings struct {
		LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
		LeagueJoinPublicMax          int           `json:"league_join_public_max"`
		LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
		LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
		LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
		LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
		LeaguePrefixPublic           string        `json:"league_prefix_public"`
		LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
		LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
		LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
		LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
		CupStartEventID              int           `json:"cup_start_event_id"`
		CupStopEventID               int           `json:"cup_stop_event_id"`
		CupQualifyingMethod          string        `json:"cup_qualifying_method"`
		CupType                      string        `json:"cup_type"`
		SquadSquadplay               int           `json:"squad_squadplay"`
		SquadSquadsize               int           `json:"squad_squadsize"`
		SquadTeamLimit               int           `json:"squad_team_limit"`
		SquadTotalSpend              int           `json:"squad_total_spend"`
		UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
		UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
		UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
		StatsFormDays                int           `json:"stats_form_days"`
		SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
		TransfersCap                 int           `json:"transfers_cap"`
		TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
		LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
		Timezone                     string        `json:"timezone"`
	} `json:"game_settings"`
	Phases []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		StartEvent int    `json:"start_event"`
		StopEvent  int    `json:"stop_event"`
	} `json:"phases"`
	Teams []struct {
		Code                int         `json:"code"`
		Draw                int         `json:"draw"`
		Form                interface{} `json:"form"`
		ID                  int         `json:"id"`
		Loss                int         `json:"loss"`
		Name                string      `json:"name"`
		Played              int         `json:"played"`
		Points              int         `json:"points"`
		Position            int         `json:"position"`
		ShortName           string      `json:"short_name"`
		Strength            int         `json:"strength"`
		TeamDivision        interface{} `json:"team_division"`
		Unavailable         bool        `json:"unavailable"`
		Win                 int         `json:"win"`
		StrengthOverallHome int         `json:"strength_overall_home"`
		StrengthOverallAway int         `json:"strength_overall_away"`
		StrengthAttackHome  int         `json:"strength_attack_home"`
		StrengthAttackAway  int         `json:"strength_attack_away"`
		StrengthDefenceHome int         `json:"strength_defence_home"`
		StrengthDefenceAway int         `json:"strength_defence_away"`
		PulseID             int         `json:"pulse_id"`
	} `json:"teams"`
	TotalPlayers int `json:"total_players"`
	Elements     []struct {
		ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
		ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
		Code                             int         `json:"code"`
		CostChangeEvent                  int         `json:"cost_change_event"`
		CostChangeEventFall              int         `json:"cost_change_event_fall"`
		CostChangeStart                  int         `json:"cost_change_start"`
		CostChangeStartFall              int         `json:"cost_change_start_fall"`
		DreamteamCount                   int         `json:"dreamteam_count"`
		ElementType                      int         `json:"element_type"`
		EpNext                           string      `json:"ep_next"`
		EpThis                           string      `json:"ep_this"`
		EventPoints                      int         `json:"event_points"`
		FirstName                        string      `json:"first_name"`
		Form                             string      `json:"form"`
		ID                               int         `json:"id"`
		InDreamteam                      bool        `json:"in_dreamteam"`
		News                             string      `json:"news"`
		NewsAdded                        time.Time   `json:"news_added"`
		NowCost                          int         `json:"now_cost"`
		Photo                            string      `json:"photo"`
		PointsPerGame                    string      `json:"points_per_game"`
		SecondName                       string      `json:"second_name"`
		SelectedByPercent                string      `json:"selected_by_percent"`
		Special                          bool        `json:"special"`
		SquadNumber                      interface{} `json:"squad_number"`
		Status                           string      `json:"status"`
		Team                             int         `json:"team"`
		TeamCode                         int         `json:"team_code"`
		TotalPoints                      int         `json:"total_points"`
		TransfersIn                      int         `json:"transfers_in"`
		TransfersInEvent                 int         `json:"transfers_in_event"`
		TransfersOut                     int         `json:"transfers_out"`
		TransfersOutEvent                int         `json:"transfers_out_event"`
		ValueForm                        string      `json:"value_form"`
		ValueSeason                      string      `json:"value_season"`
		WebName                          string      `json:"web_name"`
		Minutes                          int         `json:"minutes"`
		GoalsScored                      int         `json:"goals_scored"`
		Assists                          int         `json:"assists"`
		CleanSheets                      int         `json:"clean_sheets"`
		GoalsConceded                    int         `json:"goals_conceded"`
		OwnGoals                         int         `json:"own_goals"`
		PenaltiesSaved                   int         `json:"penalties_saved"`
		PenaltiesMissed                  int         `json:"penalties_missed"`
		YellowCards                      int         `json:"yellow_cards"`
		RedCards                         int         `json:"red_cards"`
		Saves                            int         `json:"saves"`
		Bonus                            int         `json:"bonus"`
		Bps                              int         `json:"bps"`
		Influence                        string      `json:"influence"`
		Creativity                       string      `json:"creativity"`
		Threat                           string      `json:"threat"`
		IctIndex                         string      `json:"ict_index"`
		InfluenceRank                    int         `json:"influence_rank"`
		InfluenceRankType                int         `json:"influence_rank_type"`
		CreativityRank                   int         `json:"creativity_rank"`
		CreativityRankType               int         `json:"creativity_rank_type"`
		ThreatRank                       int         `json:"threat_rank"`
		ThreatRankType                   int         `json:"threat_rank_type"`
		IctIndexRank                     int         `json:"ict_index_rank"`
		IctIndexRankType                 int         `json:"ict_index_rank_type"`
		CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
		CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
		DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
		DirectFreekicksText              string      `json:"direct_freekicks_text"`
		PenaltiesOrder                   interface{} `json:"penalties_order"`
		PenaltiesText                    string      `json:"penalties_text"`
	} `json:"elements"`
	ElementStats []struct {
		Label string `json:"label"`
		Name  string `json:"name"`
	} `json:"element_stats"`
	ElementTypes []struct {
		ID                 int    `json:"id"`
		PluralName         string `json:"plural_name"`
		PluralNameShort    string `json:"plural_name_short"`
		SingularName       string `json:"singular_name"`
		SingularNameShort  string `json:"singular_name_short"`
		SquadSelect        int    `json:"squad_select"`
		SquadMinPlay       int    `json:"squad_min_play"`
		SquadMaxPlay       int    `json:"squad_max_play"`
		UIShirtSpecific    bool   `json:"ui_shirt_specific"`
		SubPositionsLocked []int  `json:"sub_positions_locked"`
		ElementCount       int    `json:"element_count"`
	} `json:"element_types"`
}

type TeamResponse struct {
	Code                int         `json:"code"`
	Draw                int         `json:"draw"`
	Form                interface{} `json:"form"`
	ID                  int         `json:"id"`
	Loss                int         `json:"loss"`
	Name                string      `json:"name"`
	Played              int         `json:"played"`
	Points              int         `json:"points"`
	Position            int         `json:"position"`
	ShortName           string      `json:"short_name"`
	Strength            int         `json:"strength"`
	TeamDivision        interface{} `json:"team_division"`
	Unavailable         bool        `json:"unavailable"`
	Win                 int         `json:"win"`
	StrengthOverallHome int         `json:"strength_overall_home"`
	StrengthOverallAway int         `json:"strength_overall_away"`
	StrengthAttackHome  int         `json:"strength_attack_home"`
	StrengthAttackAway  int         `json:"strength_attack_away"`
	StrengthDefenceHome int         `json:"strength_defence_home"`
	StrengthDefenceAway int         `json:"strength_defence_away"`
	PulseID             int         `json:"pulse_id"`
}

type EventsResponse struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	DeadlineTime           string `json:"deadline_time"`
	AverageEntryScore      int    `json:"average_entry_score"`
	Finished               bool   `json:"finished"`
	DataChecked            bool   `json:"data_checked"`
	HighestScoringEntry    int    `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int    `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int    `json:"deadline_time_game_offset"`
	HighestScore           int    `json:"highest_score"`
	IsPrevious             bool   `json:"is_previous"`
	IsCurrent              bool   `json:"is_current"`
	IsNext                 bool   `json:"is_next"`
	ChipPlays              []struct {
		ChipName  string `json:"chip_name"`
		NumPlayed int    `json:"num_played"`
	} `json:"chip_plays"`
	MostSelected      int `json:"most_selected"`
	MostTransferredIn int `json:"most_transferred_in"`
	TopElement        int `json:"top_element"`
	TopElementInfo    struct {
		ID     int `json:"id"`
		Points int `json:"points"`
	} `json:"top_element_info"`
	TransfersMade     int `json:"transfers_made"`
	MostCaptained     int `json:"most_captained"`
	MostViceCaptained int `json:"most_vice_captained"`
}

type PhasesResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}

type PlayerDetailedInfo struct {
	ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
	Code                             int         `json:"code"`
	CostChangeEvent                  int         `json:"cost_change_event"`
	CostChangeEventFall              int         `json:"cost_change_event_fall"`
	CostChangeStart                  int         `json:"cost_change_start"`
	CostChangeStartFall              int         `json:"cost_change_start_fall"`
	DreamteamCount                   int         `json:"dreamteam_count"`
	ElementType                      int         `json:"element_type"`
	EpNext                           string      `json:"ep_next"`
	EpThis                           string      `json:"ep_this"`
	EventPoints                      int         `json:"event_points"`
	FirstName                        string      `json:"first_name"`
	Form                             string      `json:"form"`
	ID                               int         `json:"id"`
	InDreamteam                      bool        `json:"in_dreamteam"`
	News                             string      `json:"news"`
	NewsAdded                        time.Time   `json:"news_added"`
	NowCost                          int         `json:"now_cost"`
	Photo                            string      `json:"photo"`
	PointsPerGame                    string      `json:"points_per_game"`
	SecondName                       string      `json:"second_name"`
	SelectedByPercent                string      `json:"selected_by_percent"`
	Special                          bool        `json:"special"`
	SquadNumber                      interface{} `json:"squad_number"`
	Status                           string      `json:"status"`
	Team                             int         `json:"team"`
	TeamCode                         int         `json:"team_code"`
	TotalPoints                      int         `json:"total_points"`
	TransfersIn                      int         `json:"transfers_in"`
	TransfersInEvent                 int         `json:"transfers_in_event"`
	TransfersOut                     int         `json:"transfers_out"`
	TransfersOutEvent                int         `json:"transfers_out_event"`
	ValueForm                        string      `json:"value_form"`
	ValueSeason                      string      `json:"value_season"`
	WebName                          string      `json:"web_name"`
	Minutes                          int         `json:"minutes"`
	GoalsScored                      int         `json:"goals_scored"`
	Assists                          int         `json:"assists"`
	CleanSheets                      int         `json:"clean_sheets"`
	GoalsConceded                    int         `json:"goals_conceded"`
	OwnGoals                         int         `json:"own_goals"`
	PenaltiesSaved                   int         `json:"penalties_saved"`
	PenaltiesMissed                  int         `json:"penalties_missed"`
	YellowCards                      int         `json:"yellow_cards"`
	RedCards                         int         `json:"red_cards"`
	Saves                            int         `json:"saves"`
	Bonus                            int         `json:"bonus"`
	Bps                              int         `json:"bps"`
	Influence                        string      `json:"influence"`
	Creativity                       string      `json:"creativity"`
	Threat                           string      `json:"threat"`
	IctIndex                         string      `json:"ict_index"`
	InfluenceRank                    int         `json:"influence_rank"`
	InfluenceRankType                int         `json:"influence_rank_type"`
	CreativityRank                   int         `json:"creativity_rank"`
	CreativityRankType               int         `json:"creativity_rank_type"`
	ThreatRank                       int         `json:"threat_rank"`
	ThreatRankType                   int         `json:"threat_rank_type"`
	IctIndexRank                     int         `json:"ict_index_rank"`
	IctIndexRankType                 int         `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
	DirectFreekicksText              string      `json:"direct_freekicks_text"`
	PenaltiesOrder                   interface{} `json:"penalties_order"`
	PenaltiesText                    string      `json:"penalties_text"`
}

type ElementStatsResponse struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type GameSettingsResponse struct {
	LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
	LeagueJoinPublicMax          int           `json:"league_join_public_max"`
	LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
	LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
	LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
	LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
	LeaguePrefixPublic           string        `json:"league_prefix_public"`
	LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
	LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
	LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
	LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
	CupStartEventID              int           `json:"cup_start_event_id"`
	CupStopEventID               int           `json:"cup_stop_event_id"`
	CupQualifyingMethod          string        `json:"cup_qualifying_method"`
	CupType                      string        `json:"cup_type"`
	SquadSquadplay               int           `json:"squad_squadplay"`
	SquadSquadsize               int           `json:"squad_squadsize"`
	SquadTeamLimit               int           `json:"squad_team_limit"`
	SquadTotalSpend              int           `json:"squad_total_spend"`
	UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
	UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
	UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
	StatsFormDays                int           `json:"stats_form_days"`
	SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
	TransfersCap                 int           `json:"transfers_cap"`
	TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
	LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
	Timezone                     string        `json:"timezone"`
}
type ElementTypesResponse struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SingularNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_play"`
	SquadMaxPlay       int    `json:"squad_max_play"`
	UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
	ElementCount       int    `json:"element_count"`
}
type Fixture []struct {
	Code                 int           `json:"code"`
	Event                interface{}   `json:"event"`
	Finished             bool          `json:"finished"`
	FinishedProvisional  bool          `json:"finished_provisional"`
	ID                   int           `json:"id"`
	KickoffTime          interface{}   `json:"kickoff_time"`
	Minutes              int           `json:"minutes"`
	ProvisionalStartTime bool          `json:"provisional_start_time"`
	Started              interface{}   `json:"started"`
	TeamA                int           `json:"team_a"`
	TeamAScore           interface{}   `json:"team_a_score"`
	TeamH                int           `json:"team_h"`
	TeamHScore           interface{}   `json:"team_h_score"`
	Stats                []interface{} `json:"stats"`
	TeamHDifficulty      int           `json:"team_h_difficulty"`
	TeamADifficulty      int           `json:"team_a_difficulty"`
	PulseID              int           `json:"pulse_id"`
}

type WeeklyFixture []struct {
	Code                 int    `json:"code"`
	Event                int    `json:"event"`
	Finished             bool   `json:"finished"`
	FinishedProvisional  bool   `json:"finished_provisional"`
	ID                   int    `json:"id"`
	KickoffTime          string `json:"kickoff_time"`
	Minutes              int    `json:"minutes"`
	ProvisionalStartTime bool   `json:"provisional_start_time"`
	Started              bool   `json:"started"`
	TeamA                int    `json:"team_a"`
	TeamAScore           int    `json:"team_a_score"`
	TeamH                int    `json:"team_h"`
	TeamHScore           int    `json:"team_h_score"`
	Stats                []struct {
		Identifier string        `json:"identifier"`
		A          []interface{} `json:"a"`
		H          []struct {
			Value   int `json:"value"`
			Element int `json:"element"`
		} `json:"h"`
	} `json:"stats"`
	TeamHDifficulty int `json:"team_h_difficulty"`
	TeamADifficulty int `json:"team_a_difficulty"`
	PulseID         int `json:"pulse_id"`
}

type Player struct {
	Fixtures []struct {
		ID                   int         `json:"id"`
		Code                 int         `json:"code"`
		TeamH                int         `json:"team_h"`
		TeamHScore           interface{} `json:"team_h_score"`
		TeamA                int         `json:"team_a"`
		TeamAScore           interface{} `json:"team_a_score"`
		Event                int         `json:"event"`
		Finished             bool        `json:"finished"`
		Minutes              int         `json:"minutes"`
		ProvisionalStartTime bool        `json:"provisional_start_time"`
		KickoffTime          string      `json:"kickoff_time"`
		EventName            string      `json:"event_name"`
		IsHome               bool        `json:"is_home"`
		Difficulty           int         `json:"difficulty"`
	} `json:"fixtures"`
	History []struct {
		Element          int       `json:"element"`
		Fixture          int       `json:"fixture"`
		OpponentTeam     int       `json:"opponent_team"`
		TotalPoints      int       `json:"total_points"`
		WasHome          bool      `json:"was_home"`
		KickoffTime      time.Time `json:"kickoff_time"`
		TeamHScore       int       `json:"team_h_score"`
		TeamAScore       int       `json:"team_a_score"`
		Round            int       `json:"round"`
		Minutes          int       `json:"minutes"`
		GoalsScored      int       `json:"goals_scored"`
		Assists          int       `json:"assists"`
		CleanSheets      int       `json:"clean_sheets"`
		GoalsConceded    int       `json:"goals_conceded"`
		OwnGoals         int       `json:"own_goals"`
		PenaltiesSaved   int       `json:"penalties_saved"`
		PenaltiesMissed  int       `json:"penalties_missed"`
		YellowCards      int       `json:"yellow_cards"`
		RedCards         int       `json:"red_cards"`
		Saves            int       `json:"saves"`
		Bonus            int       `json:"bonus"`
		Bps              int       `json:"bps"`
		Influence        string    `json:"influence"`
		Creativity       string    `json:"creativity"`
		Threat           string    `json:"threat"`
		IctIndex         string    `json:"ict_index"`
		Value            int       `json:"value"`
		TransfersBalance int       `json:"transfers_balance"`
		Selected         int       `json:"selected"`
		TransfersIn      int       `json:"transfers_in"`
		TransfersOut     int       `json:"transfers_out"`
	} `json:"history"`
	HistoryPast []struct {
		SeasonName      string `json:"season_name"`
		ElementCode     int    `json:"element_code"`
		StartCost       int    `json:"start_cost"`
		EndCost         int    `json:"end_cost"`
		TotalPoints     int    `json:"total_points"`
		Minutes         int    `json:"minutes"`
		GoalsScored     int    `json:"goals_scored"`
		Assists         int    `json:"assists"`
		CleanSheets     int    `json:"clean_sheets"`
		GoalsConceded   int    `json:"goals_conceded"`
		OwnGoals        int    `json:"own_goals"`
		PenaltiesSaved  int    `json:"penalties_saved"`
		PenaltiesMissed int    `json:"penalties_missed"`
		YellowCards     int    `json:"yellow_cards"`
		RedCards        int    `json:"red_cards"`
		Saves           int    `json:"saves"`
		Bonus           int    `json:"bonus"`
		Bps             int    `json:"bps"`
		Influence       string `json:"influence"`
		Creativity      string `json:"creativity"`
		Threat          string `json:"threat"`
		IctIndex        string `json:"ict_index"`
	} `json:"history_past"`
}
type PlayerFixture struct {
	ID                   int         `json:"id"`
	Code                 int         `json:"code"`
	TeamH                int         `json:"team_h"`
	TeamHScore           interface{} `json:"team_h_score"`
	TeamA                int         `json:"team_a"`
	TeamAScore           interface{} `json:"team_a_score"`
	Event                int         `json:"event"`
	Finished             bool        `json:"finished"`
	Minutes              int         `json:"minutes"`
	ProvisionalStartTime bool        `json:"provisional_start_time"`
	KickoffTime          string      `json:"kickoff_time"`
	EventName            string      `json:"event_name"`
	IsHome               bool        `json:"is_home"`
	Difficulty           int         `json:"difficulty"`
}

type PlayerHistory struct {
	Element          int       `json:"element"`
	Fixture          int       `json:"fixture"`
	OpponentTeam     int       `json:"opponent_team"`
	TotalPoints      int       `json:"total_points"`
	WasHome          bool      `json:"was_home"`
	KickoffTime      time.Time `json:"kickoff_time"`
	TeamHScore       int       `json:"team_h_score"`
	TeamAScore       int       `json:"team_a_score"`
	Round            int       `json:"round"`
	Minutes          int       `json:"minutes"`
	GoalsScored      int       `json:"goals_scored"`
	Assists          int       `json:"assists"`
	CleanSheets      int       `json:"clean_sheets"`
	GoalsConceded    int       `json:"goals_conceded"`
	OwnGoals         int       `json:"own_goals"`
	PenaltiesSaved   int       `json:"penalties_saved"`
	PenaltiesMissed  int       `json:"penalties_missed"`
	YellowCards      int       `json:"yellow_cards"`
	RedCards         int       `json:"red_cards"`
	Saves            int       `json:"saves"`
	Bonus            int       `json:"bonus"`
	Bps              int       `json:"bps"`
	Influence        string    `json:"influence"`
	Creativity       string    `json:"creativity"`
	Threat           string    `json:"threat"`
	IctIndex         string    `json:"ict_index"`
	Value            int       `json:"value"`
	TransfersBalance int       `json:"transfers_balance"`
	Selected         int       `json:"selected"`
	TransfersIn      int       `json:"transfers_in"`
	TransfersOut     int       `json:"transfers_out"`
}

type PlayerHistoryPast struct {
	SeasonName      string `json:"season_name"`
	ElementCode     int    `json:"element_code"`
	StartCost       int    `json:"start_cost"`
	EndCost         int    `json:"end_cost"`
	TotalPoints     int    `json:"total_points"`
	Minutes         int    `json:"minutes"`
	GoalsScored     int    `json:"goals_scored"`
	Assists         int    `json:"assists"`
	CleanSheets     int    `json:"clean_sheets"`
	GoalsConceded   int    `json:"goals_conceded"`
	OwnGoals        int    `json:"own_goals"`
	PenaltiesSaved  int    `json:"penalties_saved"`
	PenaltiesMissed int    `json:"penalties_missed"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards        int    `json:"red_cards"`
	Saves           int    `json:"saves"`
	Bonus           int    `json:"bonus"`
	Bps             int    `json:"bps"`
	Influence       string `json:"influence"`
	Creativity      string `json:"creativity"`
	Threat          string `json:"threat"`
	IctIndex        string `json:"ict_index"`
}
type GameWeek struct {
	Elements []struct {
		ID    int `json:"id"`
		Stats struct {
			Minutes         int    `json:"minutes"`
			GoalsScored     int    `json:"goals_scored"`
			Assists         int    `json:"assists"`
			CleanSheets     int    `json:"clean_sheets"`
			GoalsConceded   int    `json:"goals_conceded"`
			OwnGoals        int    `json:"own_goals"`
			PenaltiesSaved  int    `json:"penalties_saved"`
			PenaltiesMissed int    `json:"penalties_missed"`
			YellowCards     int    `json:"yellow_cards"`
			RedCards        int    `json:"red_cards"`
			Saves           int    `json:"saves"`
			Bonus           int    `json:"bonus"`
			Bps             int    `json:"bps"`
			Influence       string `json:"influence"`
			Creativity      string `json:"creativity"`
			Threat          string `json:"threat"`
			IctIndex        string `json:"ict_index"`
			TotalPoints     int    `json:"total_points"`
			InDreamteam     bool   `json:"in_dreamteam"`
		} `json:"stats"`
		Explain []struct {
			Fixture int `json:"fixture"`
			Stats   []struct {
				Identifier string `json:"identifier"`
				Points     int    `json:"points"`
				Value      int    `json:"value"`
			} `json:"stats"`
		} `json:"explain"`
	} `json:"elements"`
}

type Manager struct {
	ID                       int       `json:"id"`
	JoinedTime               time.Time `json:"joined_time"`
	StartedEvent             int       `json:"started_event"`
	FavouriteTeam            int       `json:"favourite_team"`
	PlayerFirstName          string    `json:"player_first_name"`
	PlayerLastName           string    `json:"player_last_name"`
	PlayerRegionID           int       `json:"player_region_id"`
	PlayerRegionName         string    `json:"player_region_name"`
	PlayerRegionIsoCodeShort string    `json:"player_region_iso_code_short"`
	PlayerRegionIsoCodeLong  string    `json:"player_region_iso_code_long"`
	SummaryOverallPoints     int       `json:"summary_overall_points"`
	SummaryOverallRank       int       `json:"summary_overall_rank"`
	SummaryEventPoints       int       `json:"summary_event_points"`
	SummaryEventRank         int       `json:"summary_event_rank"`
	CurrentEvent             int       `json:"current_event"`
	Leagues                  struct {
		Classic []struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			ShortName      string      `json:"short_name"`
			Created        time.Time   `json:"created"`
			Closed         bool        `json:"closed"`
			Rank           interface{} `json:"rank"`
			MaxEntries     interface{} `json:"max_entries"`
			LeagueType     string      `json:"league_type"`
			Scoring        string      `json:"scoring"`
			AdminEntry     interface{} `json:"admin_entry"`
			StartEvent     int         `json:"start_event"`
			EntryCanLeave  bool        `json:"entry_can_leave"`
			EntryCanAdmin  bool        `json:"entry_can_admin"`
			EntryCanInvite bool        `json:"entry_can_invite"`
			HasCup         bool        `json:"has_cup"`
			CupLeague      interface{} `json:"cup_league"`
			CupQualified   interface{} `json:"cup_qualified"`
			EntryRank      int         `json:"entry_rank"`
			EntryLastRank  int         `json:"entry_last_rank"`
		} `json:"classic"`
		H2H []interface{} `json:"h2h"`
		Cup struct {
			Matches []interface{} `json:"matches"`
			Status  struct {
				QualificationEvent   int    `json:"qualification_event"`
				QualificationNumbers int    `json:"qualification_numbers"`
				QualificationRank    int    `json:"qualification_rank"`
				QualificationState   string `json:"qualification_state"`
			} `json:"status"`
			CupLeague int `json:"cup_league"`
		} `json:"cup"`
	} `json:"leagues"`
	Name                       string      `json:"name"`
	Kit                        interface{} `json:"kit"`
	LastDeadlineBank           int         `json:"last_deadline_bank"`
	LastDeadlineValue          int         `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int         `json:"last_deadline_total_transfers"`
}
type ManagerResponse struct {
}

type ManagerLeaguesClassic struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	ShortName      string      `json:"short_name"`
	Created        time.Time   `json:"created"`
	Closed         bool        `json:"closed"`
	Rank           interface{} `json:"rank"`
	MaxEntries     interface{} `json:"max_entries"`
	LeagueType     string      `json:"league_type"`
	Scoring        string      `json:"scoring"`
	AdminEntry     interface{} `json:"admin_entry"`
	StartEvent     int         `json:"start_event"`
	EntryCanLeave  bool        `json:"entry_can_leave"`
	EntryCanAdmin  bool        `json:"entry_can_admin"`
	EntryCanInvite bool        `json:"entry_can_invite"`
	HasCup         bool        `json:"has_cup"`
	CupLeague      interface{} `json:"cup_league"`
	CupQualified   interface{} `json:"cup_qualified"`
	EntryRank      int         `json:"entry_rank"`
	EntryLastRank  int         `json:"entry_last_rank"`
}

type ManagerLeaguesCup struct {
	Matches []interface{} `json:"matches"`
	Status  struct {
		QualificationEvent   int    `json:"qualification_event"`
		QualificationNumbers int    `json:"qualification_numbers"`
		QualificationRank    int    `json:"qualification_rank"`
		QualificationState   string `json:"qualification_state"`
	} `json:"status"`
	CupLeague int `json:"cup_league"`
}

type ManagerCup struct{}

type Weekly struct {
	Current []struct {
		Event              int `json:"event"`
		Points             int `json:"points"`
		TotalPoints        int `json:"total_points"`
		Rank               int `json:"rank"`
		RankSort           int `json:"rank_sort"`
		OverallRank        int `json:"overall_rank"`
		Bank               int `json:"bank"`
		Value              int `json:"value"`
		EventTransfers     int `json:"event_transfers"`
		EventTransfersCost int `json:"event_transfers_cost"`
		PointsOnBench      int `json:"points_on_bench"`
	} `json:"current"`
	Past  []interface{} `json:"past"`
	Chips []interface{} `json:"chips"`
}

type WeeklyInfo struct {
	Event              int `json:"event"`
	Points             int `json:"points"`
	TotalPoints        int `json:"total_points"`
	Rank               int `json:"rank"`
	RankSort           int `json:"rank_sort"`
	OverallRank        int `json:"overall_rank"`
	Bank               int `json:"bank"`
	Value              int `json:"value"`
	EventTransfers     int `json:"event_transfers"`
	EventTransfersCost int `json:"event_transfers_cost"`
	PointsOnBench      int `json:"points_on_bench"`
}

type LeagueInfo struct {
	League struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		Created     time.Time   `json:"created"`
		Closed      bool        `json:"closed"`
		MaxEntries  interface{} `json:"max_entries"`
		LeagueType  string      `json:"league_type"`
		Scoring     string      `json:"scoring"`
		AdminEntry  int         `json:"admin_entry"`
		StartEvent  int         `json:"start_event"`
		CodePrivacy string      `json:"code_privacy"`
		HasCup      bool        `json:"has_cup"`
		CupLeague   interface{} `json:"cup_league"`
		Rank        interface{} `json:"rank"`
	} `json:"league"`
	NewEntries struct {
		HasNext bool `json:"has_next"`
		Page    int  `json:"page"`
		Results []struct {
			Entry           int       `json:"entry"`
			EntryName       string    `json:"entry_name"`
			JoinedTime      time.Time `json:"joined_time"`
			PlayerFirstName string    `json:"player_first_name"`
			PlayerLastName  string    `json:"player_last_name"`
		} `json:"results"`
	} `json:"new_entries"`
	Standings struct {
		HasNext bool          `json:"has_next"`
		Page    int           `json:"page"`
		Results []interface{} `json:"results"`
	} `json:"standings"`
}

type StandingsResponse struct {
	Entry      int    `json:"entry"`
	EntryName  string `json:"entry_name"`
	EventTotal int    `json:"event_total"`
	ID         int    `json:"id"`
	LastRank   int    `json:"last_rank"`
	PlayerName string `json:"player_name"`
	Rank       int    `json:"rank"`
	RankSort   int    `json:"rank_sort"`
	Total      int    `json:"total"`
}
type NewEntriesResponse struct {
	HasNext bool `json:"has_next"`
	Page    int  `json:"page"`
	Results []struct {
		Entry           int       `json:"entry"`
		EntryName       string    `json:"entry_name"`
		JoinedTime      time.Time `json:"joined_time"`
		PlayerFirstName string    `json:"player_first_name"`
		PlayerLastName  string    `json:"player_last_name"`
	} `json:"results"`
}
type MyTeam struct {
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		SellingPrice  int  `json:"selling_price"`
		Multiplier    int  `json:"multiplier"`
		PurchasePrice int  `json:"purchase_price"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
	} `json:"picks"`
	Chips []struct {
		StatusForEntry string        `json:"status_for_entry"`
		PlayedByEntry  []interface{} `json:"played_by_entry"`
		Name           string        `json:"name"`
		Number         int           `json:"number"`
		StartEvent     int           `json:"start_event"`
		StopEvent      int           `json:"stop_event"`
		ChipType       string        `json:"chip_type"`
	} `json:"chips"`
	Transfers struct {
		Cost   int    `json:"cost"`
		Status string `json:"status"`
		Limit  int    `json:"limit"`
		Made   int    `json:"made"`
		Bank   int    `json:"bank"`
		Value  int    `json:"value"`
	} `json:"transfers"`
}

type TransferHistory struct {
	ElementIn      int       `json:"element_in"`
	ElementInCost  int       `json:"element_in_cost"`
	ElementOut     int       `json:"element_out"`
	ElementOutCost int       `json:"element_out_cost"`
	Entry          int       `json:"entry"`
	Evnt           int       `json:"evnt,omitempty"`
	Time           time.Time `json:"time"`
	Event          int       `json:"event,omitempty"`
}

type TeamWeekly struct {
	ActiveChip    string        `json:"active_chip"`
	AutomaticSubs []interface{} `json:"automatic_subs"`
	EntryHistory  struct {
		Event              int `json:"event"`
		Points             int `json:"points"`
		TotalPoints        int `json:"total_points"`
		Rank               int `json:"rank"`
		RankSort           int `json:"rank_sort"`
		OverallRank        int `json:"overall_rank"`
		Bank               int `json:"bank"`
		Value              int `json:"value"`
		EventTransfers     int `json:"event_transfers"`
		EventTransfersCost int `json:"event_transfers_cost"`
		PointsOnBench      int `json:"points_on_bench"`
	} `json:"entry_history"`
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		Multiplier    int  `json:"multiplier"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
	} `json:"picks"`
}
