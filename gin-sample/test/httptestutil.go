package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func performRequest(r *gin.Engine, method, uri string, param interface{}) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performGet(r *gin.Engine, uri string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	return w
}
