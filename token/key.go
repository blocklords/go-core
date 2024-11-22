package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// 加密 key 的方法, 实现 IKEY 就行
const (
	private = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAp85KUowPoHHyZlg4/qPnIvbB9Gr//sMjSTHGxkNyFiKfwgFE\nBrNAQR1vzC1C5rMC9KzHLTOexJHrq/AkUO8yhroVsibNBYEFLStLj2C3uXkQL1bX\nGEKj/LXjAU+sav10AW/UKFzIOfN4QTmPKUh1A7s+q2EEoKpmeFkekyYM7nBMzlh0\n8nlEZmKaNSppL8fq3JNpKiKP1Z6oZxXxpnO7QMFOMVNg5VDVhv0pxN0G09F3z8Ko\nWPxF55qJu/PSden1LsQlKvEh33ev4BgSFuuPtJGLBh9Fk17CvuWihNGkwGFcc2A5\nioC9UYjhJZS7Qxp8EVtubFEzli5fMG/muyzIDQIDAQABAoIBAQCR5VXRN11OzkNG\noGXNX4vSZmBztaQlSFwhg1mjf3htrmTgNGGEwcyX0JQnHSMRmYp0WNRDhKIBni0d\nLIkmpRF0+c1rOzj+FBMAFqh3XEvgwlVEE2in+yjAyxM3TKJH011M8oGvJhwf5oMj\nknvaFNlICUCPmKaBWiYFdNaUcXzEwQ5A2ApNyc8FRs6XvEV5tXnysLGInoDhxfBU\nNyB/a+BZZeFJntG4CXtCQ0ophBWVJ984p/E7YN+O4Ne9GIt2zaohN2ImqG+ZlRAs\nFjoS9iljnxR6h9sZbbpDG5epsEyrmQiGAJFMDsWv/552g6m0SThccqZExxT4847V\naT3VQMUBAoGBANR+oitxIqMby1cofwSlmqTtEDKfcgskOipcoea9L55JtYsJdLsH\ntyJawyySR/jUWyFghFhuAsGfbjBLEd856YApQpYPl0przMjIVU+HJgaNvFDWWkhx\nrJoztWW+CZOoowUwLMycHBJUjZUPXXHI6jrClUSLbl5Gp5/ecx5UQWRxAoGBAMop\nZ7PsLUG04eWfqURFw9+f7JPJ27JI4JL8T/iffhbd91EmCUK5OvDZ45BFOe07FY+Z\nPvNegI6nBgXiUM/jTwqC2Tm3ZsC1jA5kBRUCGnvvUVYI5W//dQDzCSfQ+KsWrPG2\nI8/Jv+Zje8xfrDmJfKE/oPgL/jIPWUf4E7xHOXtdAoGBAImIAIwfaGyrQ5uAwV0P\nlhyitrYdDqH5a5AZbkw6LETFrjN0BlI69yPMHMCPWPfK8cSThHT7ltscxiOJouKY\nx/FEQy1+n8vyI5PcXaLgdRMOz1B+u+ZhdHZFe2WDbw1bu09TU9uGOoD+qrhMPo2z\nnS403ImFuQRZtIo7XsTFgaFxAoGBAMlLgR7+Q/HxEh16ZSi97tN0gjSGAmP7fOHe\nqiJ9bSeHzQLYRNBTcATycEzvIUa+VjGt/aiGqKtiU/T37E+Tnthwgauemom4O8T4\ngrbwaT6OhQaNxSdHzlErriofQfvZkEr9eZsk4BefZ12QxgRkidxlZvqVtn5SGiw3\nMC+BHBNhAoGAFrTFCNYIzdNp0xxMqsNHHfZ/QFHk3ol7ywDUN09vMxXrVFtPO1MT\n4Krta0CpkcdH6h+eOWlghY9Bs8imR2JVqIBANtAwI3kdb4nZBnM993KIs0YFqh4/\nNtvG/x40NBSbPBICea8FqADjfhHGNDXkO45h1LZ/8bEIOU+PJ5W7FlM=\n-----END RSA PRIVATE KEY-----"
	public  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp85KUowPoHHyZlg4/qPn\nIvbB9Gr//sMjSTHGxkNyFiKfwgFEBrNAQR1vzC1C5rMC9KzHLTOexJHrq/AkUO8y\nhroVsibNBYEFLStLj2C3uXkQL1bXGEKj/LXjAU+sav10AW/UKFzIOfN4QTmPKUh1\nA7s+q2EEoKpmeFkekyYM7nBMzlh08nlEZmKaNSppL8fq3JNpKiKP1Z6oZxXxpnO7\nQMFOMVNg5VDVhv0pxN0G09F3z8KoWPxF55qJu/PSden1LsQlKvEh33ev4BgSFuuP\ntJGLBh9Fk17CvuWihNGkwGFcc2A5ioC9UYjhJZS7Qxp8EVtubFEzli5fMG/muyzI\nDQIDAQAB\n-----END PUBLIC KEY-----"
)

type (
	Key struct {
		private *rsa.PrivateKey
		public  *rsa.PublicKey
	}

	KFn func(k *Key)
)

func DefaultPrivate() *rsa.PrivateKey {
	privateBlock, _ := pem.Decode([]byte(private))
	if privateBlock == nil {
		panic(fmt.Errorf("private key: malformed or missing PEM format (RSA)"))
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		if keyP, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes); err == nil {
			pKey, ok := keyP.(*rsa.PrivateKey)
			if !ok {
				panic(fmt.Errorf("private key: expected a type of *rsa.PrivateKey"))
			}

			privateKey = pKey
		} else {
			panic(err)
		}
	}

	return privateKey
}

func DefaultPublic() *rsa.PublicKey {
	publicBlock, _ := pem.Decode([]byte(public))
	if publicBlock == nil {
		panic("jwt: Could not decode public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(publicBlock.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			panic(err)
		}
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("public key: expected a type of *rsa.PublicKey"))
	}
	return publicKey
}

func WithPrivate(private *rsa.PrivateKey) KFn {
	return func(k *Key) {
		k.private = private
	}
}
func WithPublic(public *rsa.PublicKey) KFn {
	return func(k *Key) {
		k.public = public
	}
}

func (k *Key) Private() *rsa.PrivateKey {
	return k.private
}
func (k *Key) Public() *rsa.PublicKey {
	return k.public
}

func NewKey(fns ...KFn) *Key {
	k := &Key{
		private: DefaultPrivate(),
		public:  DefaultPublic(),
	}

	for _, fn := range fns {
		fn(k)
	}
	return k
}
