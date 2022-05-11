package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type TokenModel struct {

}

type Token struct {
	Id int `db:"id" json:"id"`
	UserId int `db:"user_id" json:"userId"`
	FirstName int `db:"first_name" json:"firstName"`
	Email int `db:"email" json:"email"`
	PlainText int `db:"-" json:"token"`
	Hash []byte `db:"token_hash" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Expiry time.Time `db:"expiry" json:"expiry"`
}

// This will always give the name of the table that corresponds to type User
// This may be useful as an interface for any of the data model types
func (t *TokenModel) Table() string {
	return "tokens"
}

func (t *TokenModel) LoadUserToken(userID int) (*Token, error) {
	// Could use an include syntax for joining potentially
	var token Token
	collection := upper.Collection(t.Table())
	res := collection.Find(
		up.Cond{"user_id =": userID, "expiry < ": time.Now()},
	).OrderBy("created_at desc")
	err := res.One(&token)
	if err != nil {
		// Any error other than these two is considered fatal for this function
		// Otherwise we should be fine returning a user without a token
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}
	return &token, nil
}
