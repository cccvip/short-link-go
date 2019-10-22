package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/carl-xiao/short-link-go/middleware"
	"github.com/carl-xiao/short-link-go/modules"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Cli    *modules.RedisClient
	Db     *sql.DB
}

//初始化结构
func (a *App) Initliaze() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.initliazeRoutes()
	a.Cli = modules.RedisInit()
	a.Db = modules.Dbinit()
}

//初始化路由
func (a *App) initliazeRoutes() {
	a.Router.HandleFunc("/ping", Pong).Methods("GET", "POST")
	a.Router.HandleFunc("/api/short", a.createShortLink).Methods("POST")
	a.Router.HandleFunc("/api/info", a.getShortLinkInfo).Methods("GET")
	a.Router.HandleFunc("/{link:[a-zA-Z0-9]{1,12}}", a.redirect).Methods("GET")
	a.Router.HandleFunc("/api/panic", testPanic).Methods("GET")
	a.Router.HandleFunc("/index.html", a.index).Methods("GET")
	//处理静态文件
	a.Router.Handle("/static/{type}/{file}", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	a.Router.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	m := middleware.Middware{}
	a.Router.Use(m.RecoverHandler)
	a.Router.Use(m.LoggingHandler)
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Pong"))
}

/**
创建短链
*/
func (a *App) createShortLink(w http.ResponseWriter, r *http.Request) {
	link := r.FormValue("link")
	valid := validation.Validation{}
	valid.Required(link, "link")
	if valid.HasErrors() {
		responseErrorMsg(w, modules.StatusError{Code: 500, Err: errors.New("link不存在")})
	} else {
		result, err := a.Cli.ShortenUrl(link)
		if err != nil || result == "" {
			responseErrorMsg(w, modules.StatusError{Code: 500, Err: err})
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(result))
		}
	}
}

/**
获取短链详情信息
*/
func (a *App) getShortLinkInfo(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	if link == "" {
		responseErrorMsg(w, modules.StatusError{Code: modules.INVALID_PARAMS, Err: errors.New(modules.MsgFlags[modules.INVALID_PARAMS])})
	}
	result, err := a.Cli.ShortLinkInfo(link)
	if err != nil {
		responseErrorMsg(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(result.(string)))
	}
}

/**
重定向信息
*/
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	dict := mux.Vars(r)
	link := dict["link"]
	result, err := a.Cli.UnShortUrl(link)
	if err != nil {
		responseErrorMsg(w, modules.StatusError{Code: modules.INVALID_PARAMS, Err: errors.New(modules.MsgFlags[modules.INVALID_PARAMS])})
	} else {
		//1 统计短链uv、pv
		ip := ClientIP(r)
		//2: Db存放
		a.storeIp(ip, link)
		//3: 跳转
		http.Redirect(w, r, result, 302)
	}
}
func (a *App) index(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("index.html")
	var err error
	tpl, err = tpl.ParseFiles("template/index.html")
	if err != nil {
		log.Printf("parse template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("execute template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
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
