package gnotifier

import (
	"github.com/guelfey/go.dbus"
)

func (n *notifier) Push() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", n.Config.Title, n.Config.Message, []string{},
		map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		return call.Err
	}

	return nil
}
