package luma

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// 定义键值对结构
type KeyValue[K, V constraints.Ordered] struct {
	Key   K
	Value V
}

// 可排序Map结构
type Table[K, V constraints.Ordered] struct {
	data   map[K]V
	arr    []KeyValue[K, V]
	sorted bool
}

func NewTable[K, V constraints.Ordered]() *Table[K, V] {
	return &Table[K, V]{
		data: make(map[K]V),
	}
}

// 添加或更新元素
func (m *Table[K, V]) Set(key K, value V) {
	m.data[key] = value
	m.sorted = false // 标记需要重新排序
}

// 按键排序 (升序/降序)
func (m *Table[K, V]) SortByKey(asc bool) {
	m.prepareSort()
	sort.Slice(m.arr, func(i, j int) bool {
		if asc {
			return m.arr[i].Key < m.arr[j].Key
		}
		return m.arr[i].Key > m.arr[j].Key
	})
	m.sorted = true
}

// 按值排序 (升序/降序)
func (m *Table[K, V]) SortByValue(asc bool) {
	m.prepareSort()
	sort.Slice(m.arr, func(i, j int) bool {
		if asc {
			return m.arr[i].Value < m.arr[j].Value
		}
		return m.arr[i].Value > m.arr[j].Value
	})
	m.sorted = true
}

// 准备排序数据
func (m *Table[K, V]) prepareSort() {
	m.arr = make([]KeyValue[K, V], 0, len(m.data))
	for k, v := range m.data {
		m.arr = append(m.arr, KeyValue[K, V]{k, v})
	}
	m.sorted = false
}

// 获取前N个元素
func (m *Table[K, V]) TopN(n int) []KeyValue[K, V] {
	if !m.sorted {
		m.SortByKey(true) // 默认按键排序
	}
	return m.slice(0, n)
}

// 获取后N个元素
func (m *Table[K, V]) BottomN(n int) []KeyValue[K, V] {
	if !m.sorted {
		m.SortByKey(true) // 默认按键排序
	}
	return m.slice(len(m.arr)-n, len(m.arr))
}

// 切片工具函数
func (m *Table[K, V]) slice(start, end int) []KeyValue[K, V] {
	if start < 0 {
		start = 0
	}
	if end > len(m.arr) {
		end = len(m.arr)
	}
	if start >= end {
		return []KeyValue[K, V]{}
	}
	return append([]KeyValue[K, V]{}, m.arr[start:end]...)
}

// 获取所有排序后的元素
func (m *Table[K, V]) All() []KeyValue[K, V] {
	if !m.sorted {
		m.SortByKey(true)
	}
	return append([]KeyValue[K, V]{}, m.arr...)
}
