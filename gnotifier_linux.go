package gnotifier

import (
	"github.com/guelfey/go.dbus"
)

func (n *notifier) Push() error {
	err := n.IsValid()
	if err != nil {
		return err
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", n.Config.Title, n.Config.Message, []string{},
		map[string]dbus.Variant{}, n.Config.Expiration)
	if call.Err != nil {
		return call.Err
	}

	return nil
}
