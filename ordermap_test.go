package luma

import (
	"fmt"
	"testing"
)

func TestNewSortableMap(t *testing.T) {
	sm := NewTable[string, int]()
	sm.Set("apple", 5)
	sm.Set("banana", 2)
	sm.Set("cherry", 8)
	sm.Set("date", 3)

	fmt.Println("--- Sort by key asc---")
	sm.SortByKey(true)
	for _, e := range sm.All() {
		fmt.Printf("%s: %d\n", e.Key, e.Value)
	}

	fmt.Println("\n--- Sort by value desc---")
	sm.SortByValue(false)
	for _, e := range sm.All() {
		fmt.Printf("%s: %d\n", e.Key, e.Value)
	}

	fmt.Println("\n--- Top 2 ---")
	for _, e := range sm.TopN(2) {
		fmt.Printf("%s: %d\n", e.Key, e.Value)
	}

	fmt.Println("\n--- Bottom 2 (reverse order) ---")
	// 先按键排序获取原始顺序
	sm.SortByKey(true)
	for _, e := range sm.BottomN(2) {
		fmt.Printf("%s: %d\n", e.Key, e.Value)
	}
}
