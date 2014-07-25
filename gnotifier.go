package gnotifier

import (
	"errors"
)

// GNotifier interface
type GNotifier interface {
	Push() error
	GetConfig() *Config
	IsValid() error
}

// Config define the notification options
type Config struct {
	Title      string
	Message    string
	Expiration int32
}

type notifier struct {
	Config *Config
}

func (n *notifier) GetConfig() *Config {
	return n.Config
}

func (n *notifier) IsValid() error {
	if n.GetConfig().Title == "" {
		return errors.New("A Title is mandatory")
	}

	if n.GetConfig().Message == "" {
		return errors.New("A Message is mandatory")
	}
	return nil
}

// Notification is the builder
func Notification(title, message string) GNotifier {
	config := &Config{title, message, 5000}
	n := &notifier{Config: config}
	return n
}
