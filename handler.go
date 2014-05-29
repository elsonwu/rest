package rest

import "reflect"

func NewHandler() *Handler {
	handler := new(Handler)
	handler.apis = apis{}
	return handler
}

type apis map[string]*apiItem

type apiItem struct {
	apiType reflect.Type
	api     IApiWrapper
}

func (self *apiItem) Type() reflect.Type {
	if self.apiType == nil {
		self.apiType = reflect.TypeOf(self.api.Api())
	}

	return self.apiType
}

type Handler struct {
	apis apis
}

func (self *Handler) Api(val interface{}) IApiWrapper {
	typ := reflect.TypeOf(val)
	for _, a := range self.apis {
		if a.Type() == typ {
			return a.api
		}
	}

	return nil
}

func (self *Handler) Add(name string, api *ApiWrapper) {
	if self.Has(name) {
		panic("Api " + name + " already exists")
	}

	api.Init()
	self.apis[name] = &apiItem{
		api: api,
	}
}

func (self *Handler) Has(apiName string) bool {
	for name, _ := range self.apis {
		if name == apiName {
			return true
		}
	}

	return false
}

func (self *Handler) Get(apiName string) IApiWrapper {
	for name, a := range self.apis {
		if name == apiName {
			return a.api
		}
	}

	return nil
}
