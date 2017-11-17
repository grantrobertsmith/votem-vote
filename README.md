## Getting Started 

### Before you begin, please make sure you have Golang and Postgres installed and running locally:

1. Postgres can be found at https://www.postgresql.org/download/
2. Golang setup can be found at https://golang.org/doc/install
3. Useful instructions for getting Node, NPM and Gulp setup can be found at https://github.com/volatiletech/abcweb#how-do-i-install-nodejs-npm-and-gulp


### Install the following command line tools from volatile tech through go:

`go get -u github.com/volatiletech/abcweb`

`go get -u github.com/volatiletech/sqlbuilder`

`go get -u github.com/volatiletech/mig`


### Set environment variables:

Edit `db/sqlboiler.toml` to reference the appropriate postgres user.

`export votem_secret=something_secret_$5432jfiowe23`

`export VOTEM_PG_USER=<your postgres user>`


### Setup the database:

`createdb votem_vote_dev`

`cd db/migrations && mig up postgres "user=$VOTEM_PG_USER dbname=votem_vote_dev sslmode=disable""`

`cd db && sqlboiler --wipe postgres`


### Starting the application:

`abcweb dev`


### Viewing in browser:

http://localhost:4000


I expect these instructions are general and your machine may require additional system 
adminstration to deploy this project locally. For further assistance, please contact me 
at mailto:grantrobertsmith@gmail.com.


## For the audiance:

This project is largely incomplete, and it may be expanded upon in many ways. The application 
provided here begins - scratches the surface - of an architecture for a voting system. This module 
is designed to facillitate voters wanting to cast a ballot securely over the internet. Other services 
such as voter registration and ballot construction are required to fully support the digital voting 
ecosystem.


### Template project work

Please see ATTRIBUTION and LICENSE files for work that is _not_ the contributor's own.


### Project work

The ballot UI construction, controller and services are the author's own. The server side services, scripts 
and database schema are the author's own.


### Future work

A few next steps:

1. Finish auth token implementation
2. Write unit tests for new code
3. Implement a blockchain integration for keeping an audit trail of voter actions outside of the database


## Auto-Generated README Follows...

This project was generated by the Go ABCWeb web app framework: 
https://github.com/volatiletech/abcweb -- See here for readme.

### Example run 

`VOTEM_VOTE_ENV` is the environment variable that specifies which
config.toml configuration mode/section to use.

`./votem-vote` will use the default mode of `prod`.

`VOTEM_VOTE_ENV=dev ./votem-vote`
`VOTEM_VOTE_ENV=prod ./votem-vote`

### Configuration

This project loads configuration in the order of:

1. command line argument default values
2. config.toml
3. environment variables
4. supplied command line arguments

This means that values passed into the command line will
override values passed in through the config.toml and env vars, and so on.
