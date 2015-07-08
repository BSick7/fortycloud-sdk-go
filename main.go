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