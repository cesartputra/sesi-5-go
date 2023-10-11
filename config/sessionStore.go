package config

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("sesi-5-go"))
