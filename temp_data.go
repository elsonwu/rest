package rest

import "sync"

func newTempData() *tempData {
	return &tempData{
		data: make(map[string]interface{}),
	}
}

type tempData struct {
	data map[string]interface{}
	lock sync.RWMutex
}

func (td *tempData) Get(key string) interface{} {
	td.lock.RLock()
	defer td.lock.RUnlock()
	if d, ok := td.data[key]; ok {
		return d
	}

	return nil
}

func (td *tempData) Set(key string, val interface{}) {
	td.lock.Lock()
	td.data[key] = val
	td.lock.Unlock()
}
