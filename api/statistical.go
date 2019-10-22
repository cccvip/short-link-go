package api

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

const (
	insertIp = "insert into ip_statistical(ip,short) values(?,?) "
	updateIp = "update ip_statistical set total = total+1 where id=?"
)

func (a *App) storeIp(ip, short string) {
	row := a.Db.QueryRow("select id from ip_statistical where ip=? and short=? ")
	var id int
	row.Scan(&id)
	if id == 0 {
		stmt, err := a.Db.Prepare(insertIp)
		if err != nil {
			log.Print(err.Error())
		}
		res, err := stmt.Exec(ip, short)
		if err != nil {
			log.Print(err.Error())
		}
		_, _ = res.RowsAffected()
	} else {
		stmt, err := a.Db.Prepare(updateIp)
		if err != nil {
			log.Print(err.Error())
		}
		res, err := stmt.Exec(id)
		if err != nil {
			log.Print(err.Error())
		}
		_, _ = res.RowsAffected()
	}
}
