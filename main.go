package main

import (
	"qiublog/model"
	"qiublog/routes"
)

func main() {
	//err := hll.Defaults(hll.Settings{
	//	Log2m:             31,
	//	Regwidth:          5,
	//	ExplicitThreshold: 1,
	//	SparseEnabled:     true,
	//})
	//
	//if err != nil {
	//	panic(fmt.Sprintf("hll初始化失败,%s", err))
	//}
	//h := hll.Hll{}
	//for i := 1; i < 30000; i++ {
	//	h.AddRaw(uint64(i))
	//}
	//panic(h.Cardinality())
	model.InitDb()
	routes.InitRouter()
}
