package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type RProxy struct {
	remote *url.URL
}

func GoReverseProxy(this *RProxy) *httputil.ReverseProxy {
	remote := this.remote

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(request *http.Request) {
		targetQuery := remote.RawQuery
		request.URL.Scheme = remote.Scheme
		request.URL.Host = remote.Host
		request.Host = remote.Host // todo 这个是关键
		request.URL.Path, request.URL.RawPath = joinURLPath(remote, request.URL)

		if targetQuery == "" || request.URL.RawQuery == "" {
			request.URL.RawQuery = targetQuery + request.URL.RawQuery
		} else {
			request.URL.RawQuery = targetQuery + "&" + request.URL.RawQuery
		}
		if _, ok := request.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36")
		}
		fmt.Println("request.URL.Path：", request.URL.Path, "request.URL.RawQuery：", request.URL.RawQuery)
	}

	return proxy
}

// go sdk 源码
func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

// go sdk 源码
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func NewProxy1(selectedUrl string) *httputil.ReverseProxy {

	remote, err := url.Parse(selectedUrl)
	if err != nil {
		panic(err)
	}

	proxy := GoReverseProxy(&RProxy{
		remote: remote,
	})
	return proxy

}

func NewProxy(targetHost string) *httputil.ReverseProxy {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil
	}

	return httputil.NewSingleHostReverseProxy(url)
}


func proxyServer(proxyUrl string, addr string) {
	//被代理的服务器host和port
	fmt.Println(proxyUrl)
	serveErr := http.ListenAndServe(addr, NewProxy(proxyUrl))
	if serveErr != nil {
		panic(serveErr)
	}
}


func main() {
	//JupyterNotebook := os.Getenv("JUPYTER_NOTEBOOK")
	JupyterNotebook:="http://localhost"
	JupyterUrl := JupyterNotebook + ":8888"
	//jupyterUrl := "https://www.baidu.com"

	//SonarQube := os.Getenv("SONARQUBE")
	//Spark:="http://localhost"
	//SonarQubeUrl := SonarQube + ":9000"
	SonarQubeUrl := "https://www.baidu.com"


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	for{
		fmt.Println("Welcome to Big Data Processing Application")
		fmt.Println("Please type the number that corresponds to which application you would like to run:")
		fmt.Println("1. Apache Hadoop")
		fmt.Println("2. Apache Spark")
		fmt.Println("3. Jupyter Notebook")
		fmt.Println("4. SonarQube and SonarScanner")
		fmt.Println("Type the number here >")

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "1":
		case "2":
			go proxyServer(SonarQubeUrl,":2333")
		case "3":
			go proxyServer(JupyterUrl,":2334")
		case "4":
		}
	}
}