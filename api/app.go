package api

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/carl-xiao/short-link-go/middleware"
	"github.com/carl-xiao/short-link-go/modules"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Cli    *modules.RedisClient
}

type shortReq struct {
	Url    string
	Expire int64
}
type shortRes struct {
	Shortlink string
}

//初始化结构
func (a *App) Initliaze() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.initliazeRoutes()
	a.Cli = modules.RedisInit()
}

//初始化路由
func (a *App) initliazeRoutes() {
	a.Router.HandleFunc("/ping", Pong).Methods("GET", "POST")
	a.Router.HandleFunc("/api/short", a.createShortLink).Methods("POST")
	a.Router.HandleFunc("/api/info", a.getShortLinkInfo).Methods("GET")
	a.Router.HandleFunc("/{link:[a-zA-Z0-9]{1,12}}", a.redirect).Methods("GET")
	a.Router.HandleFunc("/api/panic", testPanic).Methods("GET")

	m := middleware.Middware{}
	a.Router.Use(m.RecoverHandler)
	a.Router.Use(m.LoggingHandler)
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("resp"))
}

/**
创建短链
*/
func (a *App) createShortLink(w http.ResponseWriter, r *http.Request) {
	link := r.Form.Get("link")
	valid := validation.Validation{}
	valid.Required(link, "name")
	if valid.HasErrors() {
		log.Println("link不存在")
		return
	}

	result, err := a.Cli.ShortenUrl(link)
	if err != nil {
		responseErrorMsg(w, modules.StatusError{Code: 500, Err: errors.New("redis异常")})
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))
}

/**
获取短链详情信息
*/
func (a *App) getShortLinkInfo(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	if link == "" {
		responseErrorMsg(w, modules.StatusError{Code: modules.INVALID_PARAMS, Err: errors.New(modules.MsgFlags[modules.INVALID_PARAMS])})
	}
}

/**
重定向信息
*/
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	//dict := mux.Vars(r)
	//link := dict["link"]
	//
	//result, err := a.Cli.UnShortUrl(link)
	//url := result.(string)
	//if err != nil {
	//
	//}
	url := "http://www.baidu.com"
	http.Redirect(w, r, url, 302)
}

/**
测试panic错误
*/
func testPanic(w http.ResponseWriter, r *http.Request) {
	panic("测试一下错误信息")
}

func responseErrorMsg(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case modules.StatusError:
		w.WriteHeader(e.Status())
		fmt.Fprintf(w, "%s", e.Error())
		log.Printf("Http %d--%s", e.Status(), e.Error())
	default:
		log.Printf("default %s", e.Error())
	}
}
