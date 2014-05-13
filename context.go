package rest

import (
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func NewContext() *Context {
	c := new(Context)
	c.params = &urlValues{&url.Values{}}

	// Default "Decode" method
	c.Decode = func(out interface{}) error {
		defer c.req.Body.Close()
		return json.NewDecoder(c.req.Body).Decode(out)
	}

	return c
}

type urlValues struct {
	*url.Values
}

func (self *urlValues) Int(key string) int {
	i, _ := strconv.Atoi(self.Get(key))
	return i
}

type Request struct {
	*http.Request
}

func (self *Request) Ip() string {
	ip, _, _ := net.SplitHostPort(self.RemoteAddr)
	return ip
}

type Context struct {
	req                   *Request
	params                *urlValues
	user                  User
	runParseMultipartForm bool
	store                 *store
	handler               *Handler
	AutoSetUser           func()
	Decode                func(out interface{}) error
}

func (self *Context) User() User {
	if self.user.Id == "" {
		self.AutoSetUser()
	}

	return self.user
}

func (self *Context) Handler() *Handler {
	return self.handler
}

func (self *Context) SetHandler(handler *Handler) {
	self.handler = handler
}

func (self *Context) SetUser(user User) {
	self.user = user
}

func (self *Context) Store() *store {
	if self.store == nil {
		self.store = newStore()
	}

	return self.store
}

func (self *Context) SetReq(req *http.Request) {
	self.req = &Request{
		Request: req,
	}
}

func (self *Context) Req() *Request {
	return self.req
}

func (self *Context) initParams() {
	if self.runParseMultipartForm {
		return
	}

	// 32m
	self.req.ParseMultipartForm(32 << 20)
	self.params = &urlValues{&self.req.Form}
	self.runParseMultipartForm = true
}

func (self *Context) Params() *urlValues {
	self.initParams()
	return self.params
}
