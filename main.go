package main

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/api"
	"os"
)

func main() {
	conf := api.DefaultApiConfig()
	conf.AccessKey = ""
	conf.SecretKey = ""
	conn := api.NewApi(conf)

	gw, err := conn.FindGatewayByPublicIP("52.203.17.110")
	if err != nil {
		fmt.Printf("error finding gateway: %s", err)
		os.Exit(1)
	}

	gw, err = conn.Gateways.Get(gw.Id)
	if err != nil {
		fmt.Printf("error getting gateway: %s", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", gw)

	/*
		ns, err2 := conn.Subnets.Create(&api.Subnet{
			Name: "Test Subnet 1",
			Cidr: "10.5.0.0/16",
		})
		if err2 != nil {
			fmt.Println("Error: ", err2)
			return
		}

		fmt.Printf("%+v\n", *ns)
	*/
}
