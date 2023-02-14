module github.com/carl-xiao/short-link-go

go 1.12

replace github.com/carl-xiao/short-link-go/api => ./short-link-go/api

require (
	github.com/OwnLocal/goes v1.0.0 // indirect
	github.com/astaxie/beego v1.12.2
	github.com/go-ini/ini v1.48.0
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.7.3
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726 // indirect
	github.com/siddontang/ledisdb v0.0.0-20181029004158-becf5f38d373 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/zserge/lorca v0.1.8
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/ini.v1 v1.48.0 // indirect
)
