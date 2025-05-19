package main

import "testing"

func TestSimulateEvent_Basic(t *testing.T) {
	req := CalculationRequest{
		TotalLeaders:         2,
		TotalFollowers:       2,
		DanceStyles:          []string{"Waltz"},
		LeaderKnowledge:      map[string][]string{"1": {"Waltz"}, "2": {"Waltz"}},
		FollowerKnowledge:    map[string][]string{"A": {"Waltz"}, "B": {"Waltz"}},
		DanceDurationMinutes: 10, // 2 dances
	}

	result := SimulateEvent(req)
	if result <= 0 {
		t.Errorf("Expected positive average, got %f", result)
	}
}

func TestSimulateEvent_NoCommonStyles(t *testing.T) {
	req := CalculationRequest{
		TotalLeaders:         2,
		TotalFollowers:       2,
		DanceStyles:          []string{"Waltz", "Tango"},
		LeaderKnowledge:      map[string][]string{"1": {"Waltz"}, "2": {"Waltz"}},
		FollowerKnowledge:    map[string][]string{"A": {"Tango"}, "B": {"Tango"}},
		DanceDurationMinutes: 30,
	}

	result := SimulateEvent(req)
	if result != 0 {
		t.Errorf("Expected 0 average due to no common styles, got %f", result)
	}
}

func TestSimulateEvent_ImbalancedParticipants(t *testing.T) {
	req := CalculationRequest{
		TotalLeaders:         1,
		TotalFollowers:       5,
		DanceStyles:          []string{"Waltz"},
		LeaderKnowledge:      map[string][]string{"1": {"Waltz"}},
		FollowerKnowledge:    map[string][]string{"A": {"Waltz"}, "B": {"Waltz"}, "C": {"Waltz"}, "D": {"Waltz"}, "E": {"Waltz"}},
		DanceDurationMinutes: 25, // 5 dances
	}

	result := SimulateEvent(req)
	if result <= 0 {
		t.Errorf("Expected positive average, got %f", result)
	}
}

func TestSimulateEvent_ZeroDuration(t *testing.T) {
	req := CalculationRequest{
		TotalLeaders:         2,
		TotalFollowers:       2,
		DanceStyles:          []string{"Waltz"},
		LeaderKnowledge:      map[string][]string{"1": {"Waltz"}, "2": {"Waltz"}},
		FollowerKnowledge:    map[string][]string{"A": {"Waltz"}, "B": {"Waltz"}},
		DanceDurationMinutes: 0,
	}

	result := SimulateEvent(req)
	if result != 0 {
		t.Errorf("Expected 0 average for 0 duration, got %f", result)
	}
}

func TestSimulateEvent_EmptyKnowledge(t *testing.T) {
	req := CalculationRequest{
		TotalLeaders:         0,
		TotalFollowers:       0,
		DanceStyles:          []string{},
		LeaderKnowledge:      map[string][]string{},
		FollowerKnowledge:    map[string][]string{},
		DanceDurationMinutes: 60,
	}

	result := SimulateEvent(req)
	if result != 0 {
		t.Errorf("Expected 0 average for no participants, got %f", result)
	}
}
