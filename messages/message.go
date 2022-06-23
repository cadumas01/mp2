package messages

import (
	"fmt"
	"strings"
)

type Message struct {
	To      string
	From    string
	Content string
}

func (m Message) String() string {
	return fmt.Sprintf("To: %s\nFrom:  %s\nContent: %s\n",
		m.To, m.From, m.Content)
}

func NewMessage(to string, from string, content string) *Message {
	return &Message{to, from, content}
}

// potential issue: no error checking
func ParseMessage(s string) *Message {
	fields := strings.Split(s, "\n")

	// Trim text labels in string
	to := strings.TrimPrefix(fields[0], "To: ")
	from := strings.TrimPrefix(fields[1], "From: ")
	content := strings.TrimPrefix(fields[2], "Content: ")

	return NewMessage(to, from, content)
}
