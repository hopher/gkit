package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// 健康检查
	r.HandleFunc("/healthz", Healthz)

	// 404 不存在
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// 405 请求方法(GET POST ...)不允许
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	return r
}

// Healthz 服务健康检察 (匆删)
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(Response{
		Code: 0,
		Msg:  "OK",
	})
	w.Write(data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(Response{
		Code: 404,
		Msg:  "Not Found",
	})
	w.Write(data)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(Response{
		Code: 405,
		Msg:  "Method Not Allowed",
	})
	w.Write(data)
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
