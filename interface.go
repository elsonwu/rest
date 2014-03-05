package rest

type IApi interface {
	View(ctx *Context) (Data, *Errs)
	List(ctx *Context) (Data, *Errs)
	Create(ctx *Context) (Data, *Errs)
	Update(ctx *Context) (Data, *Errs)
	Delete(ctx *Context) (Data, *Errs)
}

type IApiWrapper interface {
	View(ctx *Context) (Data, *Errs)
	List(ctx *Context) (Data, *Errs)
	Create(ctx *Context) (Data, *Errs)
	Update(ctx *Context) (Data, *Errs)
	Delete(ctx *Context) (Data, *Errs)
}
