package rest

func NewApiWrapper(api IApi) *ApiWrapper {
	aw := new(ApiWrapper)
	aw.api = api
	return aw
}

func defaultAfter(ctx *Context) {}

type ApiWrapper struct {
	api IApi
}

func (self *ApiWrapper) With(ctx *Context) {
	self.api.With(ctx)
}

func (self *ApiWrapper) after(ctx *Context) {
	ctx.Handler().After(ctx, self)
}

func (self *ApiWrapper) SetupItems(ctx *Context, ids []string) {
	self.api.SetupItems(ctx, ids)
	self.With(ctx)
}

func (self *ApiWrapper) View(ctx *Context) *Errs {
	es := self.api.View(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) List(ctx *Context) *Errs {
	es := self.api.List(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Create(ctx *Context) *Errs {
	es := self.api.Create(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Update(ctx *Context) *Errs {
	es := self.api.Update(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Delete(ctx *Context) *Errs {
	es := self.api.Delete(ctx)
	self.after(ctx)
	return es
}
