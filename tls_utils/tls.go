package tls_utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	rand2 "math/rand"
	"time"
)

type TLS struct {
	pri     *rsa.PrivateKey
	crt     *x509.Certificate
	crt_der []byte
}

//						位数					国家，	组织，		组织单位，			省，		市，     通用名
func (t *TLS) Generate(bits int, IsCA bool, Country, Organization, OrganizationalUnit, Province, Locality, CommonName string, EmailAddresses []string) error {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	certificate := &x509.Certificate{
		SerialNumber: big.NewInt(rand2.Int63()),
		Subject: pkix.Name{
			Country:            []string{Country},
			Organization:       []string{Organization},
			OrganizationalUnit: []string{OrganizationalUnit},
			Province:           []string{Province},
			CommonName:         CommonName,
			Locality:           []string{Locality},
			//ExtraNames:         info.Names,//多余的名字
		},
		NotBefore:             time.Now(),                                                                 //证书的开始时间
		NotAfter:              time.Now().AddDate(1, 0, 0),                                                //证书的结束时间
		BasicConstraintsValid: true,                                                                       //基本的有效性约束
		IsCA:                  IsCA,                                                                       //是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, //证书用途
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		EmailAddresses:        EmailAddresses,
	}

	t.pri = private
	t.crt = certificate

	return nil
}

// 如果ca是nil 或者 是 t 本身 那么这是自签名证书
func (t *TLS) SetIssuer(ca *TLS) error {
	if ca == nil {
		ca = t
	}
	buf, err := x509.CreateCertificate(rand.Reader, t.crt, ca.crt, &t.pri.PublicKey, ca.pri)
	if err != nil {
		return err
	}
	crt, err := x509.ParseCertificate(buf)
	if err != nil {
		return err
	}
	t.crt = crt
	t.crt_der = buf
	return nil
}

// private is pkcs1 pattern
func (t *TLS) LoadFromPem(pri_pem, crt_pem []byte) error {
	p, _ := pem.Decode(pri_pem)
	c, _ := pem.Decode(crt_pem)

	pri, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return err
	}

	crt, err := x509.ParseCertificate(c.Bytes)
	if err != nil {
		return err
	}
	t.pri = pri
	t.crt = crt
	t.crt_der = c.Bytes
	return nil
}

func (t *TLS) LoadFromPemFile(pri_pem_file, crt_pem_file string) error {
	p, err := ioutil.ReadFile(pri_pem_file)
	if err != nil {
		return err
	}

	c, err := ioutil.ReadFile(crt_pem_file)
	if err != nil {
		return err
	}

	return t.LoadFromPem(p, c)
}

func (t *TLS) ExportToPem() (pri_pem, crt_pem []byte, err error) {
	pri_der := x509.MarshalPKCS1PrivateKey(t.pri)
	pri_pem = pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		//Headers: nil,
		Bytes: pri_der,
	})
	crt_pem = pem.EncodeToMemory(&pem.Block{
		Type: "CERTIFICATE",
		//Headers: nil,
		Bytes: t.crt_der,
	})
	return
}

func (t *TLS) ExportToPemFile(pri_pem_file, crt_pem_file string) error {
	p, c, err := t.ExportToPem()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pri_pem_file, p, 0666)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(crt_pem_file, c, 0666)
	if err != nil {
		return err
	}

	return nil

}
