package entities

type SenderMessage struct {
	To      string
	Text    string
	Buttons []interface{}
	Data    []interface{}
}
