package main

import (
    "fmt"
    "github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
    api := fortycloud.NewApi("https://api.fortycloud.net/restapi/v0.4", "https://www1.fortycloud.net")
	api.SetApiCredentials("", "", "")
	api.SetFormsCredentials("", "")
	
    subnets, err := api.PrivateSubnets.All(nil)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
	for _,subnet := range subnets {
    	fmt.Printf("%+v\n", subnet)
	}
	
	conns, err2 := api.Connections.All(0, 0, nil)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}
	for _,conn := range conns {
		fmt.Printf("%+v\n", conn)
	}
    
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