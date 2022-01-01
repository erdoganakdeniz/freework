package service

import (
	"freework/cache"
	"freework/models"
	"net/http"
)

func Get(c *cache.Cache) *models.Response {
	items := c.Items()
	res := models.Response{
		Code:    http.StatusCreated,
		Method:  http.MethodGet,
		Message: "All Data",
		Data:    items}
	return &res
}
func GetByKey(c *cache.Cache, key string) *models.Response {
	item, err := c.Get(key)
	if err == true {
		res := models.Response{
			Code:    http.StatusCreated,
			Method:  http.MethodGet,
			Message: "Value Found",
			Data:    item}
		return &res
	}
	res := models.Response{
		Code:    http.StatusNotFound,
		Method:  http.MethodGet,
		Message: "Value not Found",
		Data:    ""}
	return &res
}
func Set(c *cache.Cache, data models.KeyValue) *models.Response {
	err := c.Set(data.Key, data.Value, cache.DefaultExpiration)
	if err == nil {
		res := models.Response{
			Code:    http.StatusCreated,
			Method:  http.MethodPost,
			Message: "Added into Store",
			Data:    data}
		return &res
	}
	res := models.Response{
		Code:    http.StatusBadRequest,
		Method:  http.MethodPost,
		Message: "Not Added into Store",
		Data:    nil}
	return &res

}
func Flush(c *cache.Cache) *models.Response {
	c.Flush()
	res := models.Response{
		Code:    http.StatusOK,
		Method:  http.MethodGet,
		Message: "Flushed Store",
		Data:    ""}
	return &res
}
func Delete(c *cache.Cache, data string) *models.Response {
	item := c.Delete(data)
	res := models.Response{
		Code:    http.StatusOK,
		Method:  http.MethodDelete,
		Message: "Deleted into Store",
		Data:    item}
	return &res
}
