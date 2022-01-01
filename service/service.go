package service

import (
	"freework/cache"
	"freework/model"
	"net/http"
)

func Get(c *cache.Cache) *model.Response {
	items := c.Items()
	res := model.Response{
		Code:    http.StatusCreated,
		Method:  http.MethodGet,
		Message: "All Data",
		Data:    items}
	return &res
}
func GetByKey(c *cache.Cache, key string) *model.Response {
	item, err := c.Get(key)
	if err == true {
		res := model.Response{
			Code:    http.StatusCreated,
			Method:  http.MethodGet,
			Message: "Value Found",
			Data:    item}
		return &res
	}
	res := model.Response{
		Code:    http.StatusNotFound,
		Method:  http.MethodGet,
		Message: "Value not Found",
		Data:    ""}
	return &res
}
func Set(c *cache.Cache, data model.KeyValue) *model.Response {
	err := c.Set(data.Key, data.Value, cache.NoExpiration)
	if err == nil {
		res := model.Response{
			Code:    http.StatusCreated,
			Method:  http.MethodPost,
			Message: "Added into Store",
			Data:    data}
		return &res
	}
	res := model.Response{
		Code:    http.StatusBadRequest,
		Method:  http.MethodPost,
		Message: "Not Added into Store",
		Data:    nil}
	return &res

}
func Flush(c *cache.Cache) *model.Response {
	c.Flush()
	res := model.Response{
		Code:    http.StatusOK,
		Method:  http.MethodGet,
		Message: "Flushed Store",
		Data:    ""}
	return &res
}
func Delete(c *cache.Cache, data string) *model.Response {
	item := c.Delete(data)
	res := model.Response{
		Code:    http.StatusOK,
		Method:  http.MethodDelete,
		Message: "Deleted into Store",
		Data:    item}
	return &res
}
