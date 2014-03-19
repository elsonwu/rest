package rest

type IApi interface {
	With(ctx *Context)
	SetupItems(ctx *Context, ids []string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
}

type IApiWrapper interface {
	With(ctx *Context)
	SetupItems(ctx *Context, ids []string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
}
