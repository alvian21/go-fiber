package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryEntity(t *testing.T) {
	category := Category{
		ID:   1,
		Name: "Test Category",
		Photos: []Photo{
			{ID: 1, Image: "photo1.jpg"},
		},
	}

	assert.Equal(t, uint(1), category.ID)
	assert.Equal(t, "Test Category", category.Name)
	assert.Len(t, category.Photos, 1)

	data, err := json.Marshal(category)
	assert.NoError(t, err)

	var categoryMap map[string]interface{}
	err = json.Unmarshal(data, &categoryMap)
	assert.NoError(t, err)

	assert.Equal(t, "Test Category", categoryMap["name"])
	assert.NotNil(t, categoryMap["photos"])
}
