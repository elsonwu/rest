package rest

func newStore() *store {
	s := new(store)
	s.datas = map[string][]interface{}{}
	s.ids = map[string][]string{}
	return s
}

type store struct {
	datas map[string][]interface{}
	ids   map[string][]string
}

func (self *store) Add(name string, id string) {
	if ids, ok := self.ids[name]; ok {
		for _, _id := range ids {
			if _id == id {
				return
			}
		}
	}

	self.ids[name] = append(self.ids[name], id)
}

func (self *store) Empty(name string) {
	delete(self.ids, name)
}

func (self *store) Append(name string, data interface{}) {
	self.datas[name] = append(self.datas[name], data)
}

func (self *store) All(name string) []interface{} {
	if ds, ok := self.datas[name]; ok {
		return ds
	}

	return nil
}

func (self *store) Datas() map[string][]interface{} {
	return self.datas
}

func (self *store) SetupIds(ctx *Context) {
	if 0 == len(self.ids) {
		return
	}

	for apiName, ids := range self.ids {
		if 0 == len(ids) {
			self.Empty(apiName)
			continue
		}

		api := ctx.Handler().Get(apiName)
		if api != nil {
			for _, id := range ids {
				api.Fill(ctx, id)
			}

			self.Empty(apiName)
			self.SetupIds(ctx)
		}
	}
}
