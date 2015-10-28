package soap

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

const ()

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(10000)
	log.SetLogger("console", "")
}

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!\n")
}

/*
1.客户端和服务端的双向校验
*/
func ServerSertificate2() {
	pool := x509.NewCertPool()
	caCertPath := "cert_server/ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Info("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("cert_server/server.crt", "cert_server/server.key")
	if err != nil {
		log.Info("ListenAndServeTLS err:", err)
	}
}
func Client2() {
	pool := x509.NewCertPool()
	caCertPath := "certs/cert_server/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Info("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("certs/cert_server/client.crt", "certs/cert_server/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
			//InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://server:8081")
	if err != nil {
		log.Info("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
