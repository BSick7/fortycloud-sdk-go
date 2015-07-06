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
    }
    fmt.Println("Done")
}