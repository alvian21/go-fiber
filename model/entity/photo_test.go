package entity

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPhotoEntity(t *testing.T) {
	id := uuid.New()
	categoryId := uuid.New()
	photo := Photo{
		ID:         id,
		Image:      "test_photo.jpg",
		CategoryID: categoryId,
	}

	assert.Equal(t, id, photo.ID)
	assert.Equal(t, "test_photo.jpg", photo.Image)
	assert.Equal(t, categoryId, photo.CategoryID)

	data, err := json.Marshal(photo)
	assert.NoError(t, err)

	var photoMap map[string]interface{}
	err = json.Unmarshal(data, &photoMap)
	assert.NoError(t, err)

	assert.Equal(t, "test_photo.jpg", photoMap["image"])
	assert.Equal(t, categoryId.String(), photoMap["category_id"])
}
