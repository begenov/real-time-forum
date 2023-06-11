package domain

type WSEvent struct {
	Type        string
	Body        any
	RecipientID int
}
