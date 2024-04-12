package structs

import (
	"time"

	uuid "github.com/google/uuid"
)

type Sub struct {
	Plan  string `json:"plan"`
	Price int    `json:"price"`
}

type Subscriptions struct {
	Sub      []Sub `json:"subscriptions"`
	Activate bool  `json:"activate"`
}

type SubscriptionList struct {
	Subscriptions Subscriptions `json:"subscription"`
}

type SubscriptionWithoutId struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" json:"id"`
	Price     float32   `json:"price" validate:"required"`
	Frequency int       `json:"frequency"`
	Plan      string    `json:"plan" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
