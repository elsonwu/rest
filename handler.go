package rest

func DefaultAfter(ctx *Context, aw *ApiWrapper) {
	aw.LoopWith(ctx)
	ctx.Store().SetupIds(ctx)
}

func NewHandler() *Handler {
	handler := new(Handler)
	handler.apis = apis{}
	handler.After = DefaultAfter
	return handler
}

type apis map[string]IApiWrapper

type Handler struct {
	apis  apis
	After func(*Context, *ApiWrapper)
}

func (self *Handler) Add(name string, api *ApiWrapper) {
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
	for name, ApiWrapper := range self.apis {
		if name == apiName {
			return ApiWrapper
		}
	}

	return nil
}
