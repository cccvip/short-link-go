package modules

type Storage interface {
	//转换为短链接
	ShortenUrl(url string) (string, error)
	//获取短链接的详细信息
	ShortUrlInfo(url string) (interface{}, error)
	//逆向查看原链接
	UnshortUrl(url string) (string, error)
}
