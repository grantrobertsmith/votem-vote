package voteService

import(
	"fmt"

	"github.com/grantrobertsmith/votem-vote/db/models"
)

type RoundResult struct {
	HasMajority		bool
	RoundTabulation	map[string]int
	BallotCount		int
}

type RCVContestState struct {
	EliminatedContestOptions	map[string]bool
}

var rcvContestsStates map[string]*RCVContestState = make(map[string]*RCVContestState)

func TabulateAndDisplayRCV(ballots models.VotedBallotSlice, contestKey string) {
	// initialize state of contest options that have been eliminated and auto-add the abstention option for the contest
	if (contestKey == "rcv") {
		newEmptyStateStruct := RCVContestState{}
		rcvContestsStates["rcv"] = &newEmptyStateStruct
		rcvContestsStates[contestKey].EliminatedContestOptions = make(map[string]bool)
		rcvContestsStates[contestKey].EliminatedContestOptions["-- ABSTAIN FROM ADDITIONAL SELECTION --"] = true
	}

	var foundWinner bool
	roundNumber := 1

	for !foundWinner {
		roundResult := processRound(ballots, contestKey, roundNumber)
		foundWinner = roundResult.HasMajority
		displayRoundResult(roundResult, roundNumber)
		roundNumber++
	}
}

func processRound(ballots models.VotedBallotSlice, contestKey string, roundNumber int) (roundResult RoundResult) {

	roundResult.RoundTabulation, roundResult.BallotCount = tabulateRoundResults(ballots, contestKey)
	roundResult.HasMajority = checkForMajority(roundResult)

	if !roundResult.HasMajority {
		eliminateLowestPreference(roundResult, contestKey)
	}

	return roundResult
}

func tabulateRoundResults(ballots models.VotedBallotSlice, contestKey string) (results map[string]int, ballotCount int) {
	eliminatedContestOptions := rcvContestsStates[contestKey].EliminatedContestOptions
	results = make(map[string]int)

	for _, ballot := range ballots {
		var voterSelectionsAsStrings []string

		// This can be made extensible using reflect and validating that the data for a given contestKey is a string
		// Better yet, use a contest data model to store what type of contest, call this based on type and have the data already referenced by string
		if (contestKey == "rcv") {
			voterSelections := ballot.RCV
			voterSelections.Unmarshal(&voterSelectionsAsStrings)
		}

		votedOption := voterSelectionsAsStrings[0]
		if !eliminatedContestOptions[votedOption] {
			results[votedOption]++
			ballotCount++
		}
	}

	return results, ballotCount
}

func checkForMajority(result RoundResult) bool {
	for _, optionCount := range result.RoundTabulation {
		if (float64(optionCount) / float64(result.BallotCount) >= .5) {
			return true;
		}
	}

	return false
}

func eliminateLowestPreference(result RoundResult, contestKey string) {
	// var optionsToEliminate []string
	// var lowestTotal int


}

func displayRoundResult(roundResult RoundResult, roundNumber int) {
	var foundMajorityStatement string
	if (roundResult.HasMajority) {
		foundMajorityStatement = "a majority favorite was found. (Or, it was a tie.)"

	} else {
		foundMajorityStatement = "the winner is unclear."
	}

	fmt.Printf("\tAfter round %d, %s.", roundNumber, foundMajorityStatement)
	DisplayResults(roundResult.RoundTabulation)
}
