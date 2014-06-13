package rest

type IApi interface {
	Init()
	DataName() string
	With(ctx *Context, dataItem interface{})
	Fill(ctx *Context, id string)
	View(ctx *Context) []error
	List(ctx *Context) []error
	Create(ctx *Context) []error
	Update(ctx *Context) []error
	Delete(ctx *Context) []error
	UpdateAll(ctx *Context) []error
	DeleteAll(ctx *Context) []error
}

type IApiWrapper interface {
	Init()
	Api() IApi
	LoopWith(ctx *Context)
	With(ctx *Context, dataItem interface{})
	Fill(ctx *Context, id string)
	View(ctx *Context) []error
	List(ctx *Context) []error
	Create(ctx *Context) []error
	Update(ctx *Context) []error
	Delete(ctx *Context) []error
	UpdateAll(ctx *Context) []error
	DeleteAll(ctx *Context) []error
}
