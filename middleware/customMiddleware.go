package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middware struct {
}

//记录请求参数以及消耗时间
func (m *Middware) LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Print(t2.Sub(t1).Seconds())
	})
}

//Recover panic
func (m *Middware) RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
