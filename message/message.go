package message

import (
	"fmt"
	"strings"
	"time"
)

// Template represents a message template
type Template struct {
	Content string
}

// MessageRecord tracks sent messages
type MessageRecord struct {
	ProfileURL string
	SentAt     time.Time
}

// MessageService handles follow-up messaging
type MessageService struct {
	Sent map[string]MessageRecord
}

// NewMessageService creates a new service
func NewMessageService() *MessageService {
	return &MessageService{
		Sent: make(map[string]MessageRecord),
	}
}

// RenderTemplate fills dynamic variables in a template
func RenderTemplate(t Template, variables map[string]string) string {
	msg := t.Content
	for k, v := range variables {
		msg = strings.ReplaceAll(msg, "{{"+k+"}}", v)
	}
	return msg
}

// SendFollowUp simulates sending a follow-up message
func (m *MessageService) SendFollowUp(profileURL string, message string) {
	// prevent duplicate messages
	if _, exists := m.Sent[profileURL]; exists {
		fmt.Println("Message already sent to:", profileURL)
		return
	}

	fmt.Println("Sending message to:", profileURL)
	fmt.Println("Message content:", message)

	m.Sent[profileURL] = MessageRecord{
		ProfileURL: profileURL,
		SentAt:     time.Now(),
	}
}
