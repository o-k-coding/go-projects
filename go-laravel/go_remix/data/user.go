package data

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	up "github.com/upper/db/v4"
)

type UserModel struct {
	tokens TokenModel
}

type User struct {
	ID int `db:"id,omitempty"` // map this property to "id" in the DB and remove if empty
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Email string `db:"email"`
	Active int `db:"user_active"`
	Password string `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Token Token `db:"-"` // - means ignore because it doesn't exist in the db yet
}

// This will always give the name of the table that corresponds to type User
// This may be useful as an interface for any of the data model types
func (u *UserModel) Table() string {
	return "users"
}


// Consider the API surface of passing pointers to users or not.
// To avoid side effects, should pass by value, then return any changed data from the function.
// that's more functional. Though these functions both change and save the data, so ideally
// We would break those things apart.
// Example insert, pass user, user is inserted, id is returned. Caller then sets the id on the user? Not sure about that. Fine line.


func (u *UserModel) GetAll(condition up.Cond) ([]*User, error) {
	collection := upper.Collection(u.Table())
	var users []*User
	res := collection.Find(condition).OrderBy("last_name")
	err := res.All(&users)

	if err != nil { return nil, err }

	return users, nil
}

func (u *UserModel) GetByEmail(email string) (*User, error) {
	var user User
	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{ "email =": email})
	err := res.One(&user)
	if err != nil { return nil, err }
	u.LoadUserToken(&user)

	return &user, nil
}

func (u *UserModel) Get(id int) (*User, error) {

	var user User

	collection := upper.Collection(u.Table())

	res := collection.Find(up.Cond{"id =": id})

	err := res.One(&user)

	if err != nil { return nil, err }
	u.LoadUserToken(&user)

	return &user, nil
}

func (u *UserModel)  LoadUserToken(user *User) error {
	// Could use an include syntax for joining potentially
	token, err := u.tokens.LoadUserToken(user.ID)
	if err != nil { return err }
	user.Token = *token
	return nil
}

func (u *UserModel) Update(user *User) error {
	user.UpdatedAt = time.Now()
	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{"id =": user.ID})
	err := res.Update(user)

	if err != nil { return err }

	return nil
}

func (u * UserModel) Delete(id int) error {
	collection := upper.Collection(u.Table())
	res := collection.Find(id) // this is the same as collection.Find(up.Cond{"id =": u.Id})
	err := res.Delete()
	if err != nil { return err }
	return nil
}

func (u *UserModel) Insert(user *User) (int, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil { return 0, err }

	// TODO need to salt!
	user.Password = string(passHash)
	user.CreatedAt = time.Now()
	// I always struggle with setting updated at when created... I think it should be null at first to indicate a lack of updates
	table := u.Table()

	fmt.Println("Table")
	fmt.Println(table)
	collection := upper.Collection(u.Table())
	res, err := collection.Insert(user)
	if err != nil { return 0, err }
	id := getInsertID(res.ID())
	user.ID = id
	return id, nil
}

// Might pass the id here and use that to load the user?
func (u *UserModel) ResetPassword(user User, password string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil { return err }
	user.Password = string(newHash)
	err = u.Update(&user)
	if err != nil { return err }
	return nil
}

func (u *UserModel) PasswordMatches(id int, plainText string) (bool, error) {
	user, err := u.Get(id)
	if err != nil { return false, err }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))

	if err != nil {
		switch {
			case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
				return false, nil
			default:
				return false, err
		}
	}

	return true, err
}

// IDEA

// Create an interface for the CRUD operations...
// And create optional HOOKS as well? like pre insert, post insert etc. for things like business logic?
// I like the pattern better where the DB saving and logic are completely decoupled. Like ideally the password hashing would
// be done completely separately before the caller inserts... however I can also see that as being a little risky.
// That relies on the dev knowing that they called the code in the correct order

// remember interfaces are duck typed in go, so this is really only useful if you want some generic functions for
// calling DB model functionality.


// So my DDD/FP brain likes having the separate DB model struct for the functions.
// I can see why it's not "needed" though, since adding it to the User struct is really not causing issues since
// The functions aren't "part" of the data structure at runtime (I believe) like they would be in JS.
// But I like the separation of Data and functionality (FP brain)
