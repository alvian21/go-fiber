package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "hashedpassword",
		Address:  "Test Address",
		Phone:    "081234567890",
	}

	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.Password)

	// Test JSON Marshalling (Password should be hidden)
	data, err := json.Marshal(user)
	assert.NoError(t, err)

	var userMap map[string]interface{}
	err = json.Unmarshal(data, &userMap)
	assert.NoError(t, err)

	assert.Equal(t, "Test User", userMap["name"])
	assert.Equal(t, "test@example.com", userMap["email"])
	assert.NotContains(t, userMap, "password")
}
