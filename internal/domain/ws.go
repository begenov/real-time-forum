package domain

type WSEvent struct {
	Type        string `json:"type"`
	Body        any    `json:"body"`
	RecipientID int    `json:"to_user_id"`
}
