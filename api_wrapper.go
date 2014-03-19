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

func (self *ApiWrapper) LoopWith(ctx *Context) {
	for _, dataItem := range ctx.Store().All(self.api.DataName()) {
		self.api.With(ctx, dataItem)
	}
}

func (self *ApiWrapper) after(ctx *Context) {
	ctx.Handler().After(ctx, self)
}

func (self *ApiWrapper) Fill(ctx *Context, id string) {
	self.api.Fill(ctx, id)
	self.LoopWith(ctx)
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
