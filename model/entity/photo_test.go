package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhotoEntity(t *testing.T) {
	photo := Photo{
		ID:         1,
		Image:      "test_photo.jpg",
		CategoryID: 10,
	}

	assert.Equal(t, uint(1), photo.ID)
	assert.Equal(t, "test_photo.jpg", photo.Image)
	assert.Equal(t, uint(10), photo.CategoryID)

	data, err := json.Marshal(photo)
	assert.NoError(t, err)

	var photoMap map[string]interface{}
	err = json.Unmarshal(data, &photoMap)
	assert.NoError(t, err)

	assert.Equal(t, "test_photo.jpg", photoMap["image"])
	assert.Equal(t, float64(10), photoMap["category_id"])
}
