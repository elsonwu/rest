package rest

type IApi interface {
	DataName() string
	With(ctx *Context, dataItem interface{})
	Fill(ctx *Context, id string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
}

type IApiWrapper interface {
	LoopWith(ctx *Context)
	Fill(ctx *Context, id string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
}
