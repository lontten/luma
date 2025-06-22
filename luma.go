package luma

import (
	"fmt"
	"sync"
)

type LumaSearch struct {
	invertedIndex map[string]map[int64]int64
	writeMu       sync.Mutex
	docID         int64
	documents     map[int64]string //文档的 索引 对应的  文档内容
}

func NewLumaSearch() *LumaSearch {
	return &LumaSearch{
		invertedIndex: make(map[string]map[int64]int64),
		writeMu:       sync.Mutex{},
		docID:         0,
		documents:     make(map[int64]string),
	}
}

func (ls *LumaSearch) Insert(doc string) {
	ls.writeMu.Lock()
	defer ls.writeMu.Unlock()

	ls.docID++
	ls.documents[ls.docID] = doc

	var v = []rune(doc)
	tokens := cut(doc, 1, len(v))
	ls.pushToken(tokens)
}

type QueryOption struct {
	value   string // 搜索词
	minGram int    // 最小切词长度
	maxGram int    // 最大切词长度
	limit   int    // 返回的文档数量
}

func (ls *LumaSearch) Search(q QueryOption) ([]string, error) {
	var docWeight = make(map[int64]int64)
	var minGram = 1
	if q.minGram > 0 {
		minGram = q.minGram
	}
	var maxGram = StrLen(q.value)
	if q.maxGram > 0 {
		maxGram = q.maxGram
	}
	if minGram > maxGram {
		return []string{}, fmt.Errorf("minGram must be less than maxGram")
	}

	keys := cut(q.value, minGram, maxGram)
	for _, key := range keys {
		m := ls.invertedIndex[key]
		for idx, weight := range m {
			if _, ok := docWeight[idx]; !ok {
				docWeight[idx] = weight
			} else {
				docWeight[idx] += weight
			}
		}
	}

	om := NewTable[int64, int64]()
	for k, v := range docWeight {
		om.Set(k, v)
	}
	om.SortByValue(false)

	var docs []string
	var limit = 10
	if q.limit > 0 {
		limit = q.limit
	}

	for _, e := range om.TopN(limit) {
		fmt.Printf("docID: %d, weight: %d\n", e.Key, e.Value)
		docs = append(docs, ls.documents[e.Key])
	}

	return docs, nil
}

func (ls *LumaSearch) pushToken(tokens []string) {
	var index = ls.docID
	for _, token := range tokens {
		var num int64 = 1 << StrLen(token)
		idMap, ok := ls.invertedIndex[token]
		if !ok {
			idMap = map[int64]int64{}
		}
		weight, ok := idMap[index]
		if !ok {
			idMap[index] = num
		} else {
			idMap[index] = weight + num
		}
		ls.invertedIndex[token] = idMap
	}
}

// cut 返回 s 的所有长度为 minGram 到 maxGram 的连续子串
func cut(s string, minGram, maxGram int) []string {
	runes := []rune(s)
	n := len(runes)
	var result []string
	if minGram < 1 {
		minGram = 1
	}
	if maxGram > n {
		maxGram = n
	}
	for length := minGram; length <= maxGram; length++ {
		for i := 0; i <= n-length; i++ {
			result = append(result, string(runes[i:i+length]))
		}
	}
	return result
}
func StrLen(s string) int {
	return len([]rune(s))
}
