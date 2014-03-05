package rest

type ApiWrapper struct {
	Api IApi
}

func (self *ApiWrapper) View(ctx *Context) (Data, *Errs) {
	return self.Api.View(ctx)
}

func (self *ApiWrapper) List(ctx *Context) (Data, *Errs) {
	return self.Api.List(ctx)
}

func (self *ApiWrapper) Create(ctx *Context) (Data, *Errs) {
	return self.Api.Create(ctx)
}

func (self *ApiWrapper) Update(ctx *Context) (Data, *Errs) {
	return self.Api.Update(ctx)
}

func (self *ApiWrapper) Delete(ctx *Context) (Data, *Errs) {
	return self.Api.Delete(ctx)
}
