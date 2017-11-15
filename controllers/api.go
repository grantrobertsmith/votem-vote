package controllers

import (
	"time"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	. "github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/grantrobertsmith/votem-vote/password-service/password"
	"github.com/grantrobertsmith/votem-vote/db/models"
	"github.com/grantrobertsmith/votem-vote/db"
)

type Response struct {
	Success	bool	`json:"success"`
	Message string	`json:"message,omitempty"`
}

// Authentication Endpoint
func (m Main) Authenticate(w http.ResponseWriter, r *http.Request) error {
	type LoginInfo struct {
		Email		string	`json:"email"`
		Password 	string	`json:"password"`
	}

	type Response struct {
		Success		bool	`json:"success"`
		Message		string	`json:"message,omitempty"`
		AuthToken	string	`json:"authToken"`
	}

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return err
    }

	var loginInfo LoginInfo
	err = json.Unmarshal(body, &loginInfo)
    if err != nil {
        return err
    }

	humans, err := models.Humen(db.DB, Where("primary_email = ?", loginInfo.Email)).All()
	if err != nil {
		return err
	}

	human := humans[0]
	if human == nil || !password.CheckPasswordHash(loginInfo.Password, human.PasswordHash) {
		response, _ := json.Marshal(Response{
			Success: false,
			Message: "Incorrect Email or Password.",
		})
		w.Write(response)
		return nil
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"timestamp": time.Now(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("votem_secret")))
	if err != nil {
		return err
	}

	response, _ := json.Marshal(Response{
		Success: true,
		AuthToken: tokenString,
	})
	w.Write(response)
	return nil
}

// Registration Endpoint
func (m Main) Register(w http.ResponseWriter, r *http.Request) error {
	type User struct {
		Human	models.Human	`json:"human"`
		NameRef	models.NameRef	`json:"name_ref"`
		Password string         `json:"password"`
	}

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return err
    }

	var user User
	err = json.Unmarshal(body, &user)
    if err != nil {
        return err
    }

	user.Human.PasswordHash, err = password.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// if err := validateRegistration(user); err != nil {
	// 	return error
	// }

	if err := user.NameRef.Insert(db.DB); err != nil {
		return err
	}

	user.Human.NameRefID = user.NameRef.ID
	if err := user.Human.Insert(db.DB); err != nil {
		return err
	}

	response, _ := json.Marshal(Response{
		Success: true,
	})
	w.Write(response)
	return nil
}

// Ballot Submission Endpoint
func (m Main) Vote(w http.ResponseWriter, r *http.Request) error {
	type VotedBallot struct {
		VoterEmail	string		`json:"voterEmail"`
		RCV			[]string	`json:"rvc"`
		UT			string		`json:"ut"`
		VF2			[]string	`json:"vf2"`
		BI			string		`json:"bi"`
	}

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return err
    }

	var votedBallot VotedBallot
	err = json.Unmarshal(body, &votedBallot)
    if err != nil {
        return err
    }

	response, _ := json.Marshal(Response{
		Success: true,
	})
	w.Write(response)
	return nil
}
