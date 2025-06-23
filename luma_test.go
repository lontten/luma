package luma

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ls := NewLumaSearch()
	//ls.Insert("搜索")
	//ls.Insert("搜索倒排开发")
	//ls.Insert("搜索倒排开发2")
	//ls.Insert("搜索倒排开发快速")
	//ls.Insert("搜索倒排开发快速快速")
	//ls.Insert("搜索倒排开发精确快速快速")
	ls.Insert("确定")
	ls.Insert("精明")
	ls.Insert("精")
	//ls.Insert("精确")

	strings, err := ls.Search(QueryOption{MinGram: 2, Value: "精确快速"})
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range strings {
		fmt.Println(s)
	}

}
