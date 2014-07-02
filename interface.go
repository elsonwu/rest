package rest

type IApi interface {
	Init()
	DataName() string
	With(ctx IContext, dataItem interface{})
	Fill(ctx IContext, id string)
	View(ctx IContext) []error
	List(ctx IContext) []error
	Create(ctx IContext) []error
	Update(ctx IContext) []error
	Delete(ctx IContext) []error
	UpdateAll(ctx IContext) []error
	DeleteAll(ctx IContext) []error
}

type IApiWrapper interface {
	Init()
	Api() IApi
	LoopWith(ctx IContext)
	With(ctx IContext, dataItem interface{})
	Fill(ctx IContext, id string)
	View(ctx IContext) []error
	List(ctx IContext) []error
	Create(ctx IContext) []error
	Update(ctx IContext) []error
	Delete(ctx IContext) []error
	UpdateAll(ctx IContext) []error
	DeleteAll(ctx IContext) []error
}
