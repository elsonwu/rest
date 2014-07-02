package rest

type Api struct {
	name string
}

func (self *Api) Init()                                   {}
func (self *Api) With(ctx IContext, dataItem interface{}) {}
func (self *Api) Fill(ctx IContext, id string)            {}
func (self *Api) View(ctx IContext) []error               { return nil }
func (self *Api) List(ctx IContext) []error               { return nil }
func (self *Api) Create(ctx IContext) []error             { return nil }
func (self *Api) Update(ctx IContext) []error             { return nil }
func (self *Api) Delete(ctx IContext) []error             { return nil }
func (self *Api) UpdateAll(ctx IContext) []error          { return nil }
func (self *Api) DeleteAll(ctx IContext) []error          { return nil }
