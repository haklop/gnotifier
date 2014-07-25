package gnotifier

import (
	"github.com/deckarep/gosx-notifier"
)

func (n *notifier) Push() error {
	err := n.IsValid()
	if err != nil {
		return err
	}

	notification := gosxnotifier.NewNotification(n.Config.Message)
	notification.Title = n.Config.Title

	err := note.Push()
	return err
}
