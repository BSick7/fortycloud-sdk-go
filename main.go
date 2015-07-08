package main

import (
    "fmt"
    "github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
    api := fortycloud.NewApi("https://api.fortycloud.net/restapi/v0.4", "https://www1.fortycloud.net")
	api.SetCredentials("", "", "")
    
    subnets, err := api.PrivateSubnets.All()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    fmt.Printf("%+v", subnets)
    fmt.Println("")
    
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