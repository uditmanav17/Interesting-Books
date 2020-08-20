package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	Title string
	Date
}

// unexported event title
type Event2 struct {
	title string
	Date
}

func (e *Event2) Title() string {
	return e.title
}

func (e *Event2) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid Title! (count > 30)")
	}
	e.title = title
	return nil
}
