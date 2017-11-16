package main

import (
	"fmt"
	"os"
	// "reflect"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/grantrobertsmith/votem-vote/db/models"
	"github.com/grantrobertsmith/votem-vote/services/voteService"
)

func main() {
	pgUser := os.Getenv("VOTEM_PG_USER")
	db, err := sql.Open("postgres", `dbname=votem_vote_dev host=localhost sslmode=disable user=` + pgUser)

	ballots, err := models.VotedBallots(db).All()
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n\nRank Choice Vote Results: ")
	voteService.TabulateAndDisplayRCV(ballots, "rcv")

	fmt.Println("\n\n\nUnexpired Term Results: ")
	utResults := voteService.TabulateResultsChooseOne(ballots, "unexpired_term")
	voteService.DisplayResults(utResults)

	fmt.Println("\n\n\nVote for 2 Results: ")
	vf2Results := voteService.TabulateResultsChooseMultiple(ballots, "vote_for_2")
	voteService.DisplayResults(vf2Results)

	fmt.Println("\n\n\nBallot Initiative Results: ")
	biResults := voteService.TabulateResultsChooseOne(ballots, "ballot_issue")
	voteService.DisplayResults(biResults)

	fmt.Println("\n\n\n")
}
