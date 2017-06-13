package qqwry

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// NewResponse 创建一个新的response对象
func NewResponse(w http.ResponseWriter, r *http.Request) Response {
	r.ParseForm()
	return Response{
		w: w,
		r: r,
	}
}

// ReturnSuccess 返回正确的信息
func (r *Response) ReturnSuccess(data interface{}) {
	r.Return(data, 200)
}

// ReturnError 返回错误信息
func (r *Response) ReturnError(statuscode, code int, errMsg string) {
	r.Return(map[string]interface{}{"errcode": code, "errmsg": errMsg}, statuscode)
}

// Return 向客户返回回数据
func (r *Response) Return(data interface{}, code int) {
	rs, err := json.Marshal(data)
	if err != nil {
		code = 500
		rs = []byte(fmt.Sprintf(`{"errcode":500, "errmsg":"%s"}`, err.Error()))
	}

	r.w.WriteHeader(code)
	r.w.Header().Add("Content-Type", "application/json")
	r.w.Write(rs)
}