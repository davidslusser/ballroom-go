package main

import (
	"math/rand"
	"time"
)

func SimulateEvent(req CalculationRequest) float64 {
	totalDances := req.DanceDurationMinutes / 5
	rand.Seed(time.Now().UnixNano())

	type pair struct {
		leader   string
		follower string
	}

	// Track unique partners per participant
	leaderPartners := make(map[string]map[string]bool)
	followerPartners := make(map[string]map[string]bool)

	for i := 0; i < totalDances; i++ {
		// Build potential pairs based on common styles
		var compatiblePairs []pair
		for lid, lStyles := range req.LeaderKnowledge {
			for fid, fStyles := range req.FollowerKnowledge {
				if hasCommonStyle(lStyles, fStyles) {
					compatiblePairs = append(compatiblePairs, pair{leader: lid, follower: fid})
				}
			}
		}

		if len(compatiblePairs) == 0 {
			break
		}

		selected := compatiblePairs[rand.Intn(len(compatiblePairs))]

		if _, exists := leaderPartners[selected.leader]; !exists {
			leaderPartners[selected.leader] = make(map[string]bool)
		}
		if _, exists := followerPartners[selected.follower]; !exists {
			followerPartners[selected.follower] = make(map[string]bool)
		}

		leaderPartners[selected.leader][selected.follower] = true
		followerPartners[selected.follower][selected.leader] = true
	}

	totalUniquePartners := 0
	participantCount := len(req.LeaderKnowledge) + len(req.FollowerKnowledge)

	for _, partners := range leaderPartners {
		totalUniquePartners += len(partners)
	}
	for _, partners := range followerPartners {
		totalUniquePartners += len(partners)
	}

	if participantCount == 0 {
		return 0
	}
	return float64(totalUniquePartners) / float64(participantCount)
}

func hasCommonStyle(a, b []string) bool {
	set := make(map[string]bool)
	for _, style := range a {
		set[style] = true
	}
	for _, style := range b {
		if set[style] {
			return true
		}
	}
	return false
}
