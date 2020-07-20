package models

import "strconv"

//Model 商品模型
type Model struct {
	ID   int
	Name string
}

//NewProduct 单个商品
func NewProduct(id int, name string) *Model {
	return &Model{ID: id, Name: name}
}

//NewProductList 商品列表
func NewProductList(n int) []*Model {
	ret := make([]*Model, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProduct(100+i, "proname"+strconv.Itoa(100+i)))
	}
	return ret
}
