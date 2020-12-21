package tls_utils

import (
	"net/http"
	"testing"
)

func TestTLS_Generate(t *testing.T) {
	tls := new(TLS)

	err := tls.Generate(
		2048,
		true,
		`CN`,
		`woshinibb`,
		`woshinibb`,
		`Zhejiang`,
		`Hangzhou`,
		`cuisite.cn`,
		[]string{`cuisite@x.c0n`},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = tls.SetIssuer(tls)
	if err != nil {
		t.Fatal(err)
	}

	err = tls.ExportToPemFile(`ca.key`, `ca.crt`)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_web(t *testing.T) {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tls"))
	})

	err := http.ListenAndServeTLS(":443", `ca.crt`, `ca.key`, nil)
	if err != nil {
		t.Fatal(err)
	}

}
