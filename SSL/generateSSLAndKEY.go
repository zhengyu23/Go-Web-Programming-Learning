package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// SSL证书实际上就是一个将扩展密钥用法(extend key usage)设置成了服务器身份验证操作的X.509证书,所以程序在生成
// 证书时使用了crypto/x509标准库.

// X.509证书: 简称SSL证书
func main() {
	generateSSLAndKey()
}

func generateSSLAndKey() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max) //	记录由CA分发的唯一号码,一个非常长的随机整数
	subject := pkix.Name{                         //	创建一个专有名称(distinguished name),并设置成证书标题
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 使用 Certificate 结构来对证书进行配置
	template := x509.Certificate{
		SerialNumber: serialNumber, // 记录由CA分发的唯一号码
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 25 * time.Hour),           //	有效期一年
		KeyUsage:     x509.KeyUsageKeyEncipherment,                   //	表明该X.509证书用于服务器身份验证
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, //	表明该X.509证书用于服务器身份验证
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},             //	只能在IP地址"127.0.0.1"上运行
	}

	// crypto/rsa 标准库 GenerateKey函数生成RSA私钥, RSA私钥的结构中包含一个能公开访问的公钥(public key)
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// CreateCertificate函数接收 Certificate结构/公钥和私钥,创建出一个经过DER编码格式化的字节切片
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// 使用 encoding/pem 标准库将证书编码到 cert.pem 文件里
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 使用 encoding/pem 标准库将密钥编码保存到 key.pem 文件里
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
