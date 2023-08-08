package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	USerID     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}