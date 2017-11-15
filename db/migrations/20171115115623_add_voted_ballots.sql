-- +mig Up
CREATE TABLE voted_ballot (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    rcv jsonb,
    unexpired_term text,
    vote_for_2 jsonb,
    ballot_issue text
);

-- +mig Down
DROP TABLE voted_ballot;
