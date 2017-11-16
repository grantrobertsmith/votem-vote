-- +mig Up
CREATE TABLE voted_ballot (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    voter_email text NOT NULL,
    rcv jsonb NOT NULL,
    unexpired_term text NOT NULL,
    vote_for_2 jsonb NOT NULL,
    ballot_issue text NOT NULL
);

-- +mig Down
DROP TABLE voted_ballot;
