package electionday

import (
	"strconv"
)

func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
}

func DisplayResult(result *ElectionResult) string {
	return result.Name + " (" + strconv.Itoa(result.Votes) + ")"
}

func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, ok := results[candidate]
	if ok {
		results[candidate]--
	}
}
