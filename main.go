package main

import (
    "fmt"
    "github.com/mdl/fortycloud-sdk-go/api"
)

func main() {
    api := fortycloud.NewApi()
	api.SetApiCredentials("", "", "")
	api.SetFormsCredentials("", "")
	
	nodes, err := api.Nodes.All(nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _,node := range nodes {
		fmt.Printf("%+v\n", node)
	}
	
	/*
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
	*/
	
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