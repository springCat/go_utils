package syn

import (
	"sync"
)

var RWMutexMapPool = sync.Pool{}

func init() {
	RWMutexMapPool.New = func() interface{} {
		return &RWMutexMap{
			items: make(map[interface{}]interface{}, 20),
		}
	}
}

type RWMutexMap struct {
	items map[interface{}]interface{}
	lock  sync.RWMutex
}

func NewRWMutexMap() *RWMutexMap {
	m := RWMutexMapPool.Get().(*RWMutexMap)
	return m
}

// Dispose will add this set back into the pool.
func (m *RWMutexMap) Recycle() {
	m.lock.Lock()
	for k := range m.items {
		delete(m.items, k)
	}
	RWMutexMapPool.Put(m)
	m.lock.Unlock()
}

func (m *RWMutexMap) Put(key interface{}, item interface{}) bool {
	m.lock.RLock()
	if source := m.items[key]; source == nil {
		m.lock.RUnlock()
		m.lock.Lock()
		if source := m.items[key]; source == nil {
			m.items[key] = item
			m.lock.Unlock()
			return true
		}
		m.lock.Unlock()
		return false
	}
	m.lock.RUnlock()
	return false
}

func (m *RWMutexMap) Keys() []interface{} {
	m.lock.RLock()
	s := make([]interface{},len(m.items))
	for k := range m.items {
		s = append(s,k)
	}
	m.lock.RUnlock()
	return s
}

func (m *RWMutexMap) Values() []interface{} {
	m.lock.RLock()
	s := make([]interface{},len(m.items))
	for _,v := range m.items {
		s = append(s,v)
	}
	m.lock.RUnlock()
	return s
}

func (m *RWMutexMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	if source := m.items[key]; source != nil {
		m.lock.RUnlock()
		return m.items[key],true
	}
	m.lock.RUnlock()
	return nil,false

}

func (m *RWMutexMap) Contains(key interface{}) bool {
	m.lock.RLock()
	source := m.items[key]
	m.lock.RUnlock()
	return source == nil
}

func (m *RWMutexMap) Clear() {
	m.lock.Lock()
	m.items = map[interface{}]interface{}{}
	m.lock.Unlock()
}

func (m *RWMutexMap) Len() int64 {
	m.lock.RLock()
	size := int64(len(m.items))
	m.lock.RUnlock()
	return size
}
