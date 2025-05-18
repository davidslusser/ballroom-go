package main

// CalculationRequest defines the input payload structure
type CalculationRequest struct {
	TotalLeaders         int                 `json:"total_leaders"`
	TotalFollowers       int                 `json:"total_followers"`
	DanceStyles          []string            `json:"dance_styles"`
	LeaderKnowledge      map[string][]string `json:"leader_knowledge"`
	FollowerKnowledge    map[string][]string `json:"follower_knowledge"`
	DanceDurationMinutes int                 `json:"dance_duration_minutes"`
}

// CalculationResponse defines the output payload structure
type CalculationResponse struct {
	AverageDancePartners float64 `json:"average_dance_partners"`
}
