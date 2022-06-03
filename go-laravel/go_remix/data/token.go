package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"net/http"
	"strings"
	"time"

	up "github.com/upper/db/v4"
)

type TokenModel struct {

}

type Token struct {
	Id int `db:"id" json:"id"`
	UserId int `db:"user_id" json:"userId"`
	FirstName string `db:"first_name" json:"firstName"`
	Email string `db:"email" json:"email"`
	PlainText string `db:"token" json:"token"`
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
		up.Cond{"user_id": userID, "expiry > ": time.Now()},
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

func (t *TokenModel) GetByTokenValue(tokenValue string) (*Token, error) {
	var token Token

	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{ "token": tokenValue })
	err := res.One(&token)

	if err != nil { return nil, err }
	return &token, nil
}

func (t *TokenModel) GetTokensForUser(userID int) ([]*Token, error) {
	var tokens []*Token
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{ "user_id": userID })
	err := res.All(tokens)

	if err != nil { return nil, err }
	return tokens, nil
}

func (t *TokenModel) Get(id int) (*Token, error) {
	var token Token

	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{ "id": id })
	err := res.One(&token)
	if err != nil { return nil, err }
	return &token, nil
}

func (t *TokenModel) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil { return err }
	return nil
}

func (t *TokenModel) DeleteByTokenValue(tokenValue int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{ "token": tokenValue})
	err := res.Delete()
	if err != nil { return err }
	return nil
}

func (t *TokenModel) Insert(token Token, user User) error {
	collection := upper.Collection(t.Table())
	// first delete existing tokens
	err := t.deleteTokensForUser(user.ID)
	if err != nil { return err }
	token.CreatedAt = time.Now()
	token.UserId = user.ID
	token.FirstName = user.FirstName
	token.Email = user.Email

	// If you wanted you could attach the id to the token and return it instead.
	_, err = collection.Insert(token)

	if err != nil { return err }

	return nil
}

func (t *TokenModel) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserId: userID,
		Expiry: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)

	if err != nil { return nil, err }

	// Ge the plaintext in exactly the correct number of chars
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]
	return token, nil
}

func (t *TokenModel) deleteTokensForUser(userID int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"user_id": userID})
	err := res.Delete()
	if err != nil { return err }
	return nil
}

func (t *TokenModel) AuthenticateToken(r *http.Request) (*Token, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no authorization heder received")
	}

	tokenValue := headerParts[1]

	token, err := t.GetByTokenValue(tokenValue)

	if err != nil || !t.ValidateToken(token) { return nil, errors.New("no valid token found") }

	return token, nil
}

func (t *TokenModel) ValidateToken(token *Token) bool {
	// Choosing to not check the user exists specifically, because assuming that the db constraints will work for this
	if len(token.PlainText) != 26 || token.Expiry.Before(time.Now()) || token.UserId == 0 { return false}

	return true
}
