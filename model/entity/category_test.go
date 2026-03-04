package entity

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoryEntity(t *testing.T) {
	id := uuid.New()
	photoId := uuid.New()
	category := Category{
		ID:   id,
		Name: "Test Category",
		Photos: []Photo{
			{ID: photoId, Image: "photo1.jpg"},
		},
	}

	assert.Equal(t, id, category.ID)
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
