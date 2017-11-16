package main

import (
	"fmt"
	"os"
	// "reflect"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/grantrobertsmith/votem-vote/db/models"
)

func main() {
	pgUser := os.Getenv("VOTEM_PG_USER")
	db, err := sql.Open("postgres", `dbname=votem_vote_dev host=localhost sslmode=disable user=` + pgUser)

	ballots, err := models.VotedBallots(db).All()
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n\nRank Choice Vote Results: ")
	tabulateAndDisplayRCV(ballots)

	fmt.Println("\n\n\nUnexpired Term Results: ")
	utResults := tabulateResultsChooseOne(ballots, "unexpired_term")
	displayResults(utResults)

	fmt.Println("\n\n\nVote for 2 Results: ")
	vf2Results := tabulateResultsChooseMultiple(ballots, "vote_for_2")
	displayResults(vf2Results)

	fmt.Println("\n\n\nBallot Initiative Results: ")
	biResults := tabulateResultsChooseOne(ballots, "ballot_issue")
	displayResults(biResults)

	fmt.Println("\n\n\n")
}

func tabulateAndDisplayRCV(ballots models.VotedBallotSlice) {


}

func tabulateResultsChooseMultiple(ballots models.VotedBallotSlice, contestKey string) (results map[string]int) {
	results = make(map[string]int)

	// This can be made extensible using reflect and validating that the data is a slice of type.JSON
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

func tabulateResultsChooseOne(ballots models.VotedBallotSlice, contestKey string) (results map[string]int) {
	results = make(map[string]int)

	// This can be made extensible using reflect and validating that the data is a string
	for _, ballot := range ballots {
		if contestKey == "unexpired_term" {
			results[ballot.UnexpiredTerm]++

		} else if contestKey == "ballot_issue" {
			results[ballot.BallotIssue]++
		}
	}

	return results
}

func displayResults(results map[string]int) {
	for contestOption, voteCount := range results {
		fmt.Printf("\n%s: %d", contestOption, voteCount)
	}
}
