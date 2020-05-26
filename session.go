package main

import (
	"fmt"
	"time"
)

type session struct {
	id string
	timestamp time.Time
	userLogin string
}

func (ctx *context)validateSession(id string) (res string, err error) {
	s, ok := ctx.sessions[id];

	if !ok {
		res = fmt.Sprintf("No session with id: %s", id)
		return
	}

	const SessionDuration = time.Duration(30 * 24 * time.Hour);

	if time.Now().UTC().Sub(s.timestamp) > SessionDuration {
		res = fmt.Sprintf("Session expired at %s", s.timestamp.Add(SessionDuration).UTC())
		return
	}

	return
}

// TODO
func (ctx *context)validateLoginCredentials(req loginRequest) (res string, err error) {
	res = ""
	return
}

// TODO
func (ctx *context)validateRegisterCredentials(req registerRequest) (res string, err error) {
	res = ""
	return
}

// TODO
func (ctx *context)establishSession(req loginRequest) (res loginResponse, err error) {
	res.Login = "abc"
	res.SessionId = "def"
	return
}

// TODO
func (ctx *context)registerUser(req registerRequest) (res registerResponse, err error) {
	res.Login = "abc"
	res.SessionId = "def"
	return
}
