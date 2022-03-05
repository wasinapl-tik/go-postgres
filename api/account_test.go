package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/wasinaphatlilawatthananan/go-postgres/db/mock"
	db "github.com/wasinaphatlilawatthananan/go-postgres/db/sqlc"
	"github.com/wasinaphatlilawatthananan/go-postgres/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl :=gomock.NewController(t)
	defer ctrl.Finish()

	store  := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(),gomock.Eq(account.ID)).
	Times(1).
	Return(account,nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d",account.ID)
	request ,err :=http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)

	server.router.ServeHTTP(recorder,request)

	require.Equal(t,http.StatusOK,recorder.Code)
}

func randomAccount() db.Accounts {
	return db.Accounts{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}