package rest

type Api struct {
	name string
}

func (self *Api) With(ctx *Context, dataItem interface{}) {}
func (self *Api) Fill(ctx *Context, id string)            {}
func (self *Api) View(ctx *Context) *Errs                 { return nil }
func (self *Api) List(ctx *Context) *Errs                 { return nil }
func (self *Api) Create(ctx *Context) *Errs               { return nil }
func (self *Api) Update(ctx *Context) *Errs               { return nil }
func (self *Api) Delete(ctx *Context) *Errs               { return nil }
