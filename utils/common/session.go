package common

import "github.com/kataras/iris/v12/sessions"

var (
	Session = sessions.New(sessions.Config{
		Cookie: sessions.DefaultCookieName,
	})
)
