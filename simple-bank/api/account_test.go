package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/okeefem2/simple_bank/db/mock"
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/token"
	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomNewAccount(owner string) db.Account {
	return db.Account{
		ID:       uuid.New(),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func TestGetAccountAPI(t *testing.T) {
	user := createRandomNewUser()
	account := createRandomNewAccount(user.Username)

	testCases := []struct {
		name          string
		accountID     string
		buildStubs    func(mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	}{
		{
			name:      "OK",
			accountID: account.ID.String(),
			buildStubs: func(mockStore *mockdb.MockStore) {
				//  build stubs
				// This function is expect to get called 1 time with any context, and the account id from our new random account
				// And the function should return the account passed and nil for the error
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
		},

		{
			name:      "NotFound",
			accountID: account.ID.String(),
			buildStubs: func(mockStore *mockdb.MockStore) {
				//  build stubs
				// This function is expect to get called 1 time with any context, and the account id from our new random account
				// And the function should return the account passed and nil for the error
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
		},

		{
			name:      "BadRequest",
			accountID: "non uuid string",
			buildStubs: func(mockStore *mockdb.MockStore) {
				//  build stubs
				// This function is expect to get called 1 time with any context, and the account id from our new random account
				// And the function should return the account passed and nil for the error
				// mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
		},

		{
			name:      "InternalError",
			accountID: account.ID.String(),
			buildStubs: func(mockStore *mockdb.MockStore) {
				//  build stubs
				// This function is expect to get called 1 time with any context, and the account id from our new random account
				// And the function should return the account passed and nil for the error
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
		},

		{
			name:      "Unauthorized",
			accountID: account.ID.String(),
			buildStubs: func(mockStore *mockdb.MockStore) {
				//  build stubs
				// This function is expect to get called 1 time with any context, and the account id from our new random account
				// And the function should return the account passed and nil for the error
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
			},
		},

		{
			name:      "NoAuthorization",
			accountID: account.ID.String(),
			buildStubs: func(mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
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

			server := newTestServer(t, mockStore)

			// A recorder is the "official" way of testing http requests without having to manually make the call
			// I wonder what is happening under the hood, if a real request is being made? we aren't actually calling start ourselves
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%s", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var parsedAccount db.Account
	err = json.Unmarshal(data, &parsedAccount)
	require.NoError(t, err)
	require.Equal(t, parsedAccount, account)
}
