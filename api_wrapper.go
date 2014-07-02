package rest

func NewApiWrapper(api IApi) *ApiWrapper {
	aw := new(ApiWrapper)
	aw.api = api
	return aw
}

func defaultAfter(ctx IContext) {}

type ApiWrapper struct {
	api IApi
}

func (self *ApiWrapper) Init() {
	self.api.Init()
}

func (self *ApiWrapper) Api() IApi {
	return self.api
}

func (self *ApiWrapper) LoopWith(ctx IContext) {
	for _, dataItem := range ctx.Store().all(self.api.DataName()) {
		self.With(ctx, dataItem)
	}
}

func (self *ApiWrapper) after(ctx IContext) {
	self.LoopWith(ctx)
	ctx.Store().fillByIds(ctx)
}

func (self *ApiWrapper) Fill(ctx IContext, id string) {
	self.api.Fill(ctx, id)
	self.LoopWith(ctx)
}

func (self *ApiWrapper) With(ctx IContext, dataItem interface{}) {
	self.api.With(ctx, dataItem)
}

func (self *ApiWrapper) View(ctx IContext) []error {
	es := self.api.View(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) List(ctx IContext) []error {
	es := self.api.List(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Create(ctx IContext) []error {
	es := self.api.Create(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Update(ctx IContext) []error {
	es := self.api.Update(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) Delete(ctx IContext) []error {
	es := self.api.Delete(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) UpdateAll(ctx IContext) []error {
	es := self.api.UpdateAll(ctx)
	self.after(ctx)
	return es
}

func (self *ApiWrapper) DeleteAll(ctx IContext) []error {
	es := self.api.DeleteAll(ctx)
	self.after(ctx)
	return es
}
