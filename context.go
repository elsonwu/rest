package rest

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func NewContext() IContext {
	c := new(Context)
	c.params = &urlValues{&url.Values{}}
	c.tempData = newTempData()
	return c
}

type urlValues struct {
	*url.Values
}

func (self *urlValues) Int(key string) int {
	i, _ := strconv.Atoi(self.Get(key))
	return i
}

func (self *urlValues) Bool(key string) bool {
	v := self.Get(key)
	return v == "true" || v == "1"
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
	reqBody               []byte
	params                *urlValues
	user                  IUser
	runParseMultipartForm bool
	store                 *store
	handler               *Handler
	AutoSetUser           func()
	tempData              *tempData
}

func (self *Context) User() IUser {
	if self.user == nil {
		self.AutoSetUser()
		if self.user == nil {
			self.user = new(User)
		}
	}

	return self.user
}

func (self *Context) SetAutoSetUserFunc(fn func()) {
	self.AutoSetUser = fn
}

func (self *Context) Handler() *Handler {
	return self.handler
}

func (self *Context) SetHandler(handler *Handler) {
	self.handler = handler
}

func (self *Context) SetUser(u IUser) {
	self.user = u
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

func (self *Context) TempData() *tempData {
	return self.tempData
}
