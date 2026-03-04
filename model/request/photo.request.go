package request

import "github.com/google/uuid"

type PhotoCreateRequest struct {
	CategoryID uuid.UUID `json:"category_id" form:"category_id" validate:"required"`
}
