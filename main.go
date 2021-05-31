package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const configFilePath = "./config.yaml"

const (
	adminUsername = "admin"
	adminPassword = "adminpw"
)

type CAService interface {
}

type CAClient struct {
	*msp.Client
}

var _ CAService = (*CAClient)(nil)

func main() {
	cli, err := New(configFilePath)
	if err != nil {
		panic(err)
	}

	// aa, err := cli.GetAllAffiliations()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", aa)

	// fabric-ca-server start 后已注册了 admin（user 表里有 admin），再次注册 admin 会报错
	// es, err := cli.Register(&msp.RegistrationRequest{
	// 	Name: "admin",
	// 	Type: "user",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(es)

	// // Enroll 可多次调用
	// if err := cli.Enroll(adminUsername, msp.WithSecret(adminPassword)); err != nil {
	// 	panic(err)
	// }
	// id, err := cli.GetSigningIdentity(adminUsername)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Printf("%s enroll successfully\n", adminUsername)
	// fmt.Printf("%#v\n", id)

	// Type 任意输，没有报错
	// es, err := cli.Register(&msp.RegistrationRequest{
	// 	Name: "asb",
	// 	Type: "user",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(es)

	_, err = cli.Register(&msp.RegistrationRequest{
		Name:        "asbba",
		Type:        "user",
		Secret:      "abcdefg",
		Affiliation: "org0.sales.department1",
	})
	if err != nil {
		panic(err)
	}

	if err := cli.Enroll("asbba", msp.WithSecret("abcdefg"), msp.WithProfile("tls")); err != nil {
		panic(err)
	}
}

func New(configFilePath string) (*CAClient, error) {
	sdk, err := fabsdk.New(config.FromFile(configFilePath))
	if err != nil {
		return nil, err
	}
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		return nil, err
	}
	return &CAClient{mspClient}, nil
}
