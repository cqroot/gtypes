package orderedmap

import "github.com/cqroot/gtypes/linkedlist"

type OrderedMap[K comparable, V any] struct {
	ll *linkedlist.LinkedList[K]
	kv map[K]V
}

func New[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		ll: linkedlist.New[K](),
		kv: make(map[K]V),
	}
}

func (m *OrderedMap[K, V]) Len() int {
	return m.ll.Size()
}

func (m *OrderedMap[K, V]) Has(key K) bool {
	_, ok := m.kv[key]
	return ok
}

func (m *OrderedMap[K, V]) Put(key K, val V) {
	if !m.Has(key) {
		m.ll.Add(key)
	}
	m.kv[key] = val
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := m.kv[key]
	return val, exists
}

func (m *OrderedMap[K, V]) Remove(key K) {
	if !m.Has(key) {
		return
	}

	delete(m.kv, key)
	m.ll.Remove(m.ll.IndexOf(key))
}

func (m *OrderedMap[K, V]) ForEach(f func(key K, val V) error) error {
	return m.ll.ForEach(func(i int, k K) error {
		err := f(k, m.kv[k])
		if err != nil {
			return err
		}
		return nil
	})
}
