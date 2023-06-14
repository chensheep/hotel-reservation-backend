package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/chensheep/hotel-reservation-backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) types.User {
	params := types.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "abc@test.com",
		Password:  "abc123456",
	}
	body, err := json.Marshal(params)
	require.NoError(t, err)
	req := httptest.NewRequest(fiber.MethodPost, testUserEndpoint, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testEnv.app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	var user types.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	require.NoError(t, err)
	require.Equal(t, params.FirstName, user.FirstName)
	require.Equal(t, params.LastName, user.LastName)
	require.Equal(t, params.Email, user.Email)
	require.NotEmpty(t, user.ID)

	return user
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	// time.Sleep(10 * time.Second)

	testUrl := testUserEndpoint + "/" + user.ID.Hex()
	// fmt.Println(testUrl)
	resp, err := testEnv.app.Test(httptest.NewRequest(fiber.MethodGet, testUrl, nil))
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	gotUser := types.User{}
	err = json.NewDecoder(resp.Body).Decode(&gotUser)
	require.NoError(t, err)

	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.FirstName, gotUser.FirstName)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.Email, gotUser.Email)
}

func TestGetUsers(t *testing.T) {
	createRandomUser(t)
	createRandomUser(t)

	resp, err := testEnv.app.Test(httptest.NewRequest(fiber.MethodGet, testUserEndpoint, nil))
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	var users []types.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users), 2)
}

func TestPostUser(t *testing.T) {
	createRandomUser(t)
}

func TestPutUser(t *testing.T) {
	user := createRandomUser(t)

	params := types.UpdateUserParams{
		FirstName: "Jane123",
		LastName:  "Doe123",
	}
	body, err := json.Marshal(params)
	require.NoError(t, err)
	req := httptest.NewRequest(fiber.MethodPut, testUserEndpoint+"/"+user.ID.Hex(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rsp, err := testEnv.app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, rsp.StatusCode)

	// var updatedUser types.User
	// err = json.NewDecoder(rsp.Body).Decode(&updatedUser)
	// require.NoError(t, err)
	// require.Equal(t, params.FirstName, updatedUser.FirstName)
	// require.Equal(t, params.LastName, updatedUser.LastName)
	// require.Equal(t, user.Email, updatedUser.Email)
	// require.Equal(t, user.ID, updatedUser.ID)
}
