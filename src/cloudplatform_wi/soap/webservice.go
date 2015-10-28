package soap

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const ()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

/*
1.服务端采用证书，客户端采用普通方式访问
*/
func ServerSertificate() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
	_, err := os.Open("cert_server/server.crt")
	if err != nil {
		log.Info("open ssl key file fail")
		panic(err)

	}
	http.ListenAndServeTLS(":8082", "cert_server/server.crt",
		"cert_server/server.key", nil)
}
func Client() {
	//x509.Certificate.
	pool := x509.NewCertPool()
	//caCertPath := "etcdcerts/ca.crt"
	caCertPath := "certs/cert_server/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Info("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	//pool.AddCert(caCrt)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://server:8081")

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
	log.Info(resp.Status)
}
