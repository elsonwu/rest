package rest

type Api struct {
	name string
}

func (self *Api) Init()                                   {}
func (self *Api) With(ctx *Context, dataItem interface{}) {}
func (self *Api) Fill(ctx *Context, id string)            {}
func (self *Api) View(ctx *Context) []error               { return nil }
func (self *Api) List(ctx *Context) []error               { return nil }
func (self *Api) Create(ctx *Context) []error             { return nil }
func (self *Api) Update(ctx *Context) []error             { return nil }
func (self *Api) Delete(ctx *Context) []error             { return nil }
func (self *Api) UpdateAll(ctx *Context) []error          { return nil }
func (self *Api) DeleteAll(ctx *Context) []error          { return nil }
