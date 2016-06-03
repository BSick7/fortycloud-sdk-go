package main

import (
	"fmt"
	"github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
	conf := api.DefaultApiConfig()
	conf.AccessKey = ""
	conf.SecretKey = ""
	conn := api.NewApi(conf)

	/*
		gateways, err := conn.Gateways.All()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		for _, gw := range gateways {
			fmt.Printf("%+v\n", gw)
		}
	*/

	ns, err2 := conn.Subnets.Create(&api.Subnet{
		Name: "Test Subnet 1",
		Cidr: "10.5.0.0/16",
	})
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}

	fmt.Printf("%+v\n", *ns)

	/*
		conn, err := api.Connections.Get(468)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		conn.Name = "ip-10-2-21-176<--->ip-10-1-11-23"

		_, err2 := api.Connections.Update(conn)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			return
		}
	*/

	/*
		servers, err := api.Servers.All()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("%+v", servers)
		fmt.Println("")


		script, err := api.Scripts.Get("Default Global Settings", true)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("%+v\n", script)
		fmt.Println("")
	*/
}
