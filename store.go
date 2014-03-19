package rest

func newStore() *store {
	s := new(store)
	s.dataMap = map[string][]interface{}{}
	s.idsMap = map[string][]string{}
	return s
}

type store struct {
	dataMap map[string][]interface{}
	idsMap  map[string][]string
}

func (self *store) Add(name string, id string) {
	if idsMap, ok := self.idsMap[name]; ok {
		for _, _id := range idsMap {
			if _id == id {
				return
			}
		}
	}

	self.idsMap[name] = append(self.idsMap[name], id)
}

func (self *store) empty(name string) {
	delete(self.idsMap, name)
}

func (self *store) Append(name string, data interface{}) {
	self.dataMap[name] = append(self.dataMap[name], data)
}

func (self *store) all(name string) []interface{} {
	if ds, ok := self.dataMap[name]; ok {
		return ds
	}

	return nil
}

func (self *store) DataMap() map[string][]interface{} {
	return self.dataMap
}

func (self *store) fillByIds(ctx *Context) {
	count := len(self.idsMap)
	if 0 == count {
		return
	}

	for apiName, idsMap := range self.idsMap {
		if 0 == count {
			self.empty(apiName)
			return
		}

		api := ctx.Handler().Get(apiName)
		if api == nil {
			return
		}

		for _, id := range idsMap {
			api.Fill(ctx, id)
		}

		self.empty(apiName)
		self.fillByIds(ctx)
	}
}
