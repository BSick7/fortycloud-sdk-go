package main

import (
    "fmt"
    "github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
    svc := fortycloud.NewService(nil, &fortycloud.Authentication {
        Credentials: fortycloud.Credentials {
            Username: "",
            Password: "",
        },
        Tenant: "",
    })
    
    err := svc.Authenticate()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    fmt.Println(svc.Auth.Token)
    fmt.Println(svc.Auth.Expires)
    
    var servers []fortycloud.Server
    servers, err = svc.Servers().All()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    fmt.Printf("%+v", servers)
    fmt.Println("")
    
    script, err := svc.Scripts().Get("Default Global Settings", true)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    fmt.Printf("%+v\n", script)
    fmt.Println("")
}