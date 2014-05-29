package rest

func newStore() *store {
	s := new(store)
	s.dataMap = map[string][]interface{}{}
	s.idsMap = map[string]map[string]bool{}
	return s
}

type store struct {
	dataMap map[string][]interface{}
	idsMap  map[string]map[string]bool
}

func (self *store) AddId(name string, id string) {
	if nil == self.idsMap[name] {
		self.idsMap[name] = map[string]bool{}
	}

	self.idsMap[name][id] = true
}

func (self *store) empty(name string) {
	delete(self.idsMap, name)
}

func (self *store) AddRecord(name string, data interface{}) {
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
			continue
		}

		for id, _ := range idsMap {
			api.Fill(ctx, id)
		}

		self.empty(apiName)
		self.fillByIds(ctx)
	}
}
