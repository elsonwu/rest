package rest

type IApi interface {
	Init()
	DataName() string
	With(ctx *Context, dataItem interface{})
	Fill(ctx *Context, id string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
	UpdateAll(ctx *Context) *Errs
	DeleteAll(ctx *Context) *Errs
}

type IApiWrapper interface {
	Init()
	Api() IApi
	LoopWith(ctx *Context)
	With(ctx *Context, dataItem interface{})
	Fill(ctx *Context, id string)
	View(ctx *Context) *Errs
	List(ctx *Context) *Errs
	Create(ctx *Context) *Errs
	Update(ctx *Context) *Errs
	Delete(ctx *Context) *Errs
	UpdateAll(ctx *Context) *Errs
	DeleteAll(ctx *Context) *Errs
}
