package main

import (
    "fmt"
    "github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
    api := api.NewApi("https://api.fortycloud.net/restapi/v0.4", "")
    api.Auth.Set("", "", "")
    
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
}