package entity

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBookEntity(t *testing.T) {
	id := uuid.New()
	book := Book{
		ID:     id,
		Title:  "Test Book",
		Author: "Test Author",
		Cover:  "test_cover.jpg",
	}

	assert.Equal(t, id, book.ID)
	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, "Test Author", book.Author)

	data, err := json.Marshal(book)
	assert.NoError(t, err)

	var bookMap map[string]interface{}
	err = json.Unmarshal(data, &bookMap)
	assert.NoError(t, err)

	assert.Equal(t, "Test Book", bookMap["title"])
	assert.Equal(t, "Test Author", bookMap["author"])
}
