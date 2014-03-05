package rest

func NewHandler() *Handler {
	handler := new(Handler)
	handler.apis = apis{}
	return handler
}

type apis map[string]IApiWrapper

type Handler struct {
	apis apis
}

func (self *Handler) Add(name string, api IApiWrapper) {
	if self.Has(name) {
		panic("Api " + name + " already exists")
	}

	self.apis[name] = api
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
	for name, apiWrapper := range self.apis {
		if name == apiName {
			return apiWrapper
		}
	}

	return nil
}
