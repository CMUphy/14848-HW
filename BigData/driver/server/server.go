package main
import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

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

func establishProxy(){
	JupyterNotebook := os.Getenv("JUPYTER_NOTEBOOK")
	//JupyterNotebook:="http://localhost"
	JupyterUrl := JupyterNotebook + ":8888"
	//jupyterUrl := "https://www.baidu.com"

	SonarQube := os.Getenv("SONARQUBE")
	//Spark:="http://localhost"
	SonarQubeUrl := SonarQube + ":9000"

	Spark := os.Getenv("SONARQUBE")
	//Spark:="http://localhost"
	SparkUrl := Spark + ":8080"

	Hadoop := os.Getenv("SONARQUBE")
	//Spark:="http://localhost"
	HadoopUrl := Hadoop + ":32007"

	go proxyServer(HadoopUrl,":2333")
	go proxyServer(SparkUrl,":2334")
	go proxyServer(JupyterUrl,":2335")
	go proxyServer(SonarQubeUrl,":2336")

}

func sendRequest(conn net.Conn, text string) {
	conn.Write([]byte(text))
	fmt.Printf("send over %s\n",text)
}

func main() {

	//建立socket，监听端口
	addr,_:=net.ResolveTCPAddr("tcp4", ":1024")

	netListen, err := net.ListenTCP("tcp", addr)
	CheckError(err)
	defer netListen.Close()
	establishProxy()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}
//处理连接
func handleConnection(conn net.Conn) {
	currentIP := "localhost"
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		receiveData := string(buffer[:n])
		switch receiveData {
		case "1":
			sendRequest(conn,"http://"+currentIP+":2333")
		case "2":
			sendRequest(conn,"http://"+currentIP+":2334")
		case "3":
			sendRequest(conn,"http://"+currentIP+":2335")
		case "4":
			sendRequest(conn,"http://"+currentIP+":2336")
		}
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func getClientIp() (string ,error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "",err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(),nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")

}
