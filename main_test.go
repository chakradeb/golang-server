package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
)

//func TestSlashRoute(t *testing.T) {
//	recorder := httptest.NewRecorder()
//	request := httptest.NewRequest("GET", "/", nil)
//	http.
//
//	assert.Equal(t, "Hello", recorder.Body)
//}

func TestLandingPageHandler(t *testing.T) {
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	req,_ := http.NewRequest("GET","/",nil)
	res := httptest.NewRecorder()
	arg := args{res, req}
	tests := []struct {
		name string
		args args
	}{
		{"/",arg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LandingPageHandler(tt.args.res, tt.args.req)
		})
	}
}
