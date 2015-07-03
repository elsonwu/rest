package rest

// in order to store a sorted map,
// we store all ids in map[string]*ids
type ids struct {
	arr []string
	m   map[string]bool
}

func newIds() *ids {
	return &ids{
		arr: make([]string, 0, 20),
		m:   make(map[string]bool),
	}
}

func (i *ids) add(id string) {
	if _, ok := i.m[id]; ok {
		return
	}

	i.m[id] = true
	i.arr = append(i.arr, id)
}

func (i *ids) isEmpty() bool {
	return len(i.m) == 0
}

func newStore() *store {
	s := new(store)
	s.dataMap = map[string][]interface{}{}
	s.idsMap = map[string]*ids{}
	s.meta = map[string]interface{}{}
	return s
}

type store struct {
	dataMap map[string][]interface{}
	idsMap  map[string]*ids
	meta    map[string]interface{}
}

func (self *store) Meta() map[string]interface{} {
	return self.meta
}

func (self *store) AddMeta(name string, val interface{}) {
	self.meta[name] = val
}

func (self *store) AddId(name string, id string) {
	if nil == self.idsMap[name] {
		self.idsMap[name] = newIds()
	}

	self.idsMap[name].add(id)
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

func (self *store) fillByIds(ctx IContext) {
	count := len(self.idsMap)
	if 0 == count {
		return
	}

	for apiName, idsMap := range self.idsMap {
		if 0 == count || idsMap.isEmpty() {
			self.empty(apiName)
			return
		}

		api := ctx.Handler().Get(apiName)
		if api == nil {
			continue
		}

		for _, id := range idsMap.arr {
			api.Fill(ctx, id)
		}

		self.empty(apiName)
		self.fillByIds(ctx)
	}
}
