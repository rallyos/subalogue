package session

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore("", []byte("subalogue-auth-session"))
	gob.Register(map[string]interface{}{})
	return nil
}
