package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/okeefem2/simple_bank/db/mock"
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

// Custom matcher code

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := password.CheckPassword(arg.HashedPassword, e.password)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v)", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{
		arg,
		password,
	}
}

func createRandomNewUser() db.User {
	return db.User{
		Username:       util.RandomOwner(),
		HashedPassword: util.RandomString(12),
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
}

func TestCreateUserAPI(t *testing.T) {
	user := createRandomNewUser()
	// hashedPassword, err := password.HashPassword(user.Password)
	// require.NoError(t, err)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Created",
			body: gin.H{
				"username": user.Username,
				"password": user.HashedPassword,
				"fullName": user.FullName,
				"email":    user.Email,
			},
			buildStubs: func(mockStore *mockdb.MockStore) {

				//  build stubs
				// This function is expect to get called 1 time with any context, and the user id from our new random user
				// And the function should return the user passed and nil for the error
				userArg := db.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				// Remember the hashed passowrd here is acting as the plain password
				mockStore.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(userArg, user.HashedPassword)).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// I feel like some of this can be lifted out?
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockStore := mockdb.NewMockStore(ctrl)

			tc.buildStubs(mockStore)

			server := NewServer(mockStore)

			// A recorder is the "official" way of testing http requests without having to manually make the call
			// I wonder what is happening under the hood, if a real request is being made? we aren't actually calling start ourselves
			recorder := httptest.NewRecorder()
			jsonBody, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, userParams db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var parsedUser db.User
	err = json.Unmarshal(data, &parsedUser)
	require.NoError(t, err)
	require.NotEmpty(t, parsedUser)
	require.Equal(t, userParams.Username, parsedUser.Username)
	require.Equal(t, userParams.FullName, parsedUser.FullName)
	require.Equal(t, userParams.Email, parsedUser.Email)
}
