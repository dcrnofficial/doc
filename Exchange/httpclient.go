package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

//marshalledStr is jsonRPC requestbody
//requestflag 1 mean request dcrnwallet jsonRPC server
//			  2 mean request dcrn jsonRPC server
func sendPostRequest(marshalledStr string, requestFlag int) (string, error) {
	marshalledJSON := []byte(marshalledStr)
	//Your dcrnwallet rpc port. Default 19110
	var url string
	if requestFlag == 1 {
		url = "https://127.0.0.1:19110"
	} else {
		url = "https://127.0.0.1:19109"
	}

	fmt.Println(string(marshalledJSON))

	bodyReader := bytes.NewReader(marshalledJSON)
	httpRequest, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return "", err
	}
	httpRequest.Close = true
	httpRequest.Header.Set("Content-Type", "application/json")

	//the config in
	//linux: ~/.dcrd/dcrd.cfg rpcuser and rpcpass
	//windows: c:/Users/<your user>/AppData/Local/Dcrd/dcrd.cfg rpcuser and rpcpass
	httpRequest.SetBasicAuth("root", "root")

	httpClient, err := newHTTPClient(requestFlag)
	if err != nil {
		return "", err
	}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %w", err)
		return "", err
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		if len(respBytes) == 0 {
			return "", fmt.Errorf("%d %s", httpResponse.StatusCode,
				http.StatusText(httpResponse.StatusCode))
		}
		return "", fmt.Errorf("%s", respBytes)
	}

	fmt.Println(string(respBytes))

	return string(respBytes), nil
}

func newHTTPClient(requestFlag int) (*http.Client, error) {
	// Configure proxy if needed.
	var dial func(network, addr string) (net.Conn, error)

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	tlsConfig = &tls.Config{
		InsecureSkipVerify: false,
	}

	//linux: ~/.dcrwallet/rpc.cert
	//windows: C:\Users\<your user>\AppData\Local\Dcrwallet\rpc.cert
	var filename string
	if requestFlag == 1 {
		filename = "C:\\Users\\Apologise\\AppData\\Local\\Dcrwallet\\rpc.cert"
	} else {
		filename = "C:\\Users\\Apologise\\AppData\\Local\\Dcrd\\rpc.cert"
	}
	pem, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return nil, fmt.Errorf("invalid certificate file")
	}
	tlsConfig.RootCAs = pool

	// Create and return the new HTTP client potentially configured with a
	// proxy and TLS.
	client := http.Client{
		Transport: &http.Transport{
			Dial:            dial,
			TLSClientConfig: tlsConfig,
		},
	}
	return &client, nil
}
