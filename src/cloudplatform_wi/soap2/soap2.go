package soap2

import (
	"encoding/xml"
	"fmt"
	"github.com/radoslav/soap"
)

const (
	zhangsan = ""
)

func Soap() {
	env := &soap.Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    "http://www.example.pl/ws/test/universal",
		Header: &soap.Header{
			WsseSecurity: &soap.WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &soap.UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &soap.Username{},
					Password: &soap.Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Username.Value = "test"
	env.Header.WsseSecurity.UsernameToken.Password.Value = "pass"
	env.Body = &soap.Body{} // interface

	output, err := xml.MarshalIndent(env, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))
}
