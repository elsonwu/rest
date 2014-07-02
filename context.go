package rest

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func NewContext() IContext {
	c := new(Context)
	c.params = &urlValues{&url.Values{}}

	// Default "decode"
	c.SetDecodeFunc(func(out interface{}) error {
		if c.reqBody == nil {
			defer c.req.Body.Close()
			v, err := ioutil.ReadAll(c.req.Body)
			c.reqBody = v
			if err != nil {
				return err
			}
		}

		return json.Unmarshal(c.reqBody, out)
	})

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
	decode                func(out interface{}) error
}

func (self *Context) User() IUser {
	if self.user == nil {
		self.AutoSetUser()
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

func (self *Context) SetDecodeFunc(fn func(out interface{}) error) {
	self.decode = fn
}

func (self *Context) Decode(out interface{}) error {
	return self.decode(out)
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
