package rest

type Api struct{}

func (self *Api) With(ctx *Context)                     {}
func (self *Api) SetupItems(ctx *Context, ids []string) {}
func (self *Api) View(ctx *Context) *Errs               { return nil }
func (self *Api) List(ctx *Context) *Errs               { return nil }
func (self *Api) Create(ctx *Context) *Errs             { return nil }
func (self *Api) Update(ctx *Context) *Errs             { return nil }
func (self *Api) Delete(ctx *Context) *Errs             { return nil }
