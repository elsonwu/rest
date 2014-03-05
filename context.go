package rest

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func NewContext() *Context {
	c := new(Context)
	c.params = url.Values{}
	return c
}

type Context struct {
	req                   *http.Request
	params                url.Values
	user                  User
	runParseMultipartForm bool
	AutoSetUser           func()
}

func (self *Context) User() User {
	if self.user.Id == "" {
		self.AutoSetUser()
	}

	return self.user
}

func (self *Context) SetUser(user User) {
	self.user = user
}

func (self *Context) SetReq(req *http.Request) {
	self.req = req
}

func (self *Context) Req() *http.Request {
	return self.req
}

func (self *Context) JsonDecode(out interface{}) error {
	return json.NewDecoder(self.req.Body).Decode(out)
}

func (self *Context) initParams() {
	if self.runParseMultipartForm {
		return
	}

	// 32m
	self.req.ParseMultipartForm(32 << 20)
	self.params = self.req.Form
	self.runParseMultipartForm = true
}

func (self *Context) Params() url.Values {
	self.initParams()
	return self.params
}
