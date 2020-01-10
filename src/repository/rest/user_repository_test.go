package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...	")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"me@test.com","password":"testpass"}`,
		RespHTTPCode: -1,
	})
	repository := userRepository{}
	user, err := repository.LoginUser("test@mail.com", "testpassword")
	assert.Nil(t, user, "User should be nil")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when request to login api", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"me@test.com","password":"testpass"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "no record matching","status": "404","error": "not_found"}`,
	})
	repository := userRepository{}
	user, err := repository.LoginUser("test@mail.com", "testpassword")
	assert.Nil(t, user, "User should be nil")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when login", err.Message)
}

func TestLoginUserInvalidLoginCredential(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"me@test.com","password":"testpass"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "no record matching","status": 404,"error": "not_found"}`,
	})
	repository := userRepository{}
	user, err := repository.LoginUser("test@mail.com", "testpassword")
	assert.Nil(t, user, "User should be nil")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "no record matching", err.Message)
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"me@test.com","password":"testpass"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
			"id": "5",
			"first_name": "deni",
			"last_name": "dharmatin",
			"email": "dharmatin@gmail.com.sg"
		  }`,
	})
	repository := userRepository{}
	user, err := repository.LoginUser("test@mail.com", "testpassword")
	assert.Nil(t, user, "User should be nil")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when unmarshaling", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"me@test.com","password":"testpass"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
			"id": 5,
			"first_name": "deni",
			"last_name": "dharmatin",
			"email": "dharmatin@gmail.com.sg"
		  }`,
	})
	repository := userRepository{}
	user, err := repository.LoginUser("test@mail.com", "testpassword")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 5)
	assert.EqualValues(t, user.FirstName, "deni")
	assert.EqualValues(t, user.LastName, "dharmatin")
	assert.EqualValues(t, user.Email, "dharmatin@gmail.com.sg")
}
