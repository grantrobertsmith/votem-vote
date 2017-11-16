package voteService

import(
	"fmt"

	"github.com/grantrobertsmith/votem-vote/db/models"
)

func TabulateResultsChooseMultiple(ballots models.VotedBallotSlice, contestKey string) (results map[string]int) {
	results = make(map[string]int)

	// This can be made extensible using reflect and validating that the data is a slice of type.JSON
	// Better yet, use a contest data model to store what type of contest, call this based on type and have the data already referenced by string
	for _, ballot := range ballots {
		if contestKey == "vote_for_2" {
			voterSelections := ballot.VoteFor2
			// fmt.Printf("%T %s", voterSelections, voterSelections)
			// fmt.Println(voterSelections, &voterSelections)
			var voterSelectionsAsStrings []string
			voterSelections.Unmarshal(&voterSelectionsAsStrings)

			for _, selection := range voterSelectionsAsStrings {
				results[selection]++
			}
		}
	}

	return results
}

func TabulateResultsChooseOne(ballots models.VotedBallotSlice, contestKey string) (results map[string]int) {
	results = make(map[string]int)

	// This can be made extensible using reflect and validating that the data for a given contextKey is a string
	// Better yet, use a contest data model to store what type of contest, call this based on type and have the data already referenced by string
	for _, ballot := range ballots {
		if contestKey == "unexpired_term" {
			results[ballot.UnexpiredTerm]++

		} else if contestKey == "ballot_issue" {
			results[ballot.BallotIssue]++
		}
	}

	return results
}

func DisplayResults(results map[string]int) {
	for contestOption, voteCount := range results {
		fmt.Printf("\n%s: %d", contestOption, voteCount)
	}
}
