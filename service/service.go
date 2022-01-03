package service

import (
	"freework/models"
	"freework/store"
	"net/http"
)

func Get(s *store.Store) *models.Response {
	items := s.Items()
	res := models.Response{
		Code:    http.StatusOK,
		Method:  http.MethodGet,
		Message: "All Data",
		Data:    items}
	return &res
}
func GetByKey(s *store.Store, key string) *models.Response {
	item, err := s.Get(key)
	if err == true {
		res := models.Response{
			Code:    http.StatusOK,
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
func Set(s *store.Store, data models.KeyValue) *models.Response {
	err := s.Set(data.Key, data.Value, store.DefaultExpiration)
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
func Flush(s *store.Store) *models.Response {
	s.Flush()
	res := models.Response{
		Code:    http.StatusOK,
		Method:  http.MethodGet,
		Message: "Flushed Store",
		Data:    ""}

	return &res
}
func Delete(s *store.Store, data string) *models.Response {
	item := s.Delete(data)
	res := models.Response{
		Code:    http.StatusOK,
		Method:  http.MethodDelete,
		Message: "Deleted into Store",
		Data:    item}
	return &res
}
