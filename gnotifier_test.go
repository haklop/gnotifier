package gnotifier

import (
	"testing"
)

func Test_Notification(t *testing.T) {
	n := Notification("Hey", "Hello")

	//assert defaults
	if n.GetConfig().Title != "Hey" {
		t.Error("NewNotification doesn't have a Title specified")
	}

	if n.GetConfig().Message != "Hello" {
		t.Error("NewNotification doesn't have a Message specified")
	}
}

func Test_Notification_Title_Validity(t *testing.T) {
	n := Notification("", "Hello")

	err := n.Push()
	if err == nil {
		t.Error("Notification should trigger an error, title is mandatory")
	}
}

func Test_Notification_Message_Validity(t *testing.T) {
	n := Notification("Title", "")

	err := n.Push()
	if err == nil {
		t.Error("Notification should trigger an error, message is mandatory")
	}
}

func Test_Push(t *testing.T) {
	n := Notification("Title!", "Testing Push")
	err := n.Push()

	//assert defaults
	if err != nil {
		t.Error("Test_Push failed with error: ", err)
	}
}
