/*

没有比这个更简单的例子了。

*/

package main

import (
	"github.com/zhangxinngang/wukong/engine"
	"github.com/zhangxinngang/wukong/types"
	"log"
)

var (
	// searcher是线程安全的
	searcher = engine.Engine{}
)

func main() {
	// 初始化
	searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: "../data/dictionary.txt"})
	defer searcher.Close()

	// 将文档加入索引
	searcher.IndexDocument(0, types.DocumentIndexData{Content: "lambgini"})
	searcher.IndexDocument(1, types.DocumentIndexData{Content: "lam pra"})
	searcher.IndexDocument(2, types.DocumentIndexData{Content: "plam"})

	// 强制索引刷新
	searcher.FlushIndex()

	// 搜索输出格式见types.SearchResponse结构体
	log.Print(searcher.Search(types.SearchRequest{Text: "lam"}))
}
