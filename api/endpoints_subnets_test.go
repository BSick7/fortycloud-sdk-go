package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEndpointsSubnets_Create(t *testing.T) {
	ma := NewMockApi()

	subnet := &Subnet{
		Name: "Test Subnet 1",
		Cidr: "10.5.0.0/16",
	}
	wantBody := fmt.Sprintf(`{"subnet":{"name":"%s","description":"","cidr":"10.5.0.0/16","disableAutoNAT":false}}`, subnet.Name)

	ma.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/restapi/v0.4/subnets" {
			if b, err := ioutil.ReadAll(r.Body); err != nil {
				t.Errorf("could not read body: %s", r.URL.Path)
			} else {
				got := string(b)
				if got != wantBody {
					t.Errorf("mismatched body: got\n%s\nwant\n%s", got, wantBody)
				}
			}
			type result struct {
				Subnet Subnet `json:"subnet"`
			}
			data, err := json.Marshal(result{
				Subnet: Subnet{
					Id:             "123",
					Name:           subnet.Name,
					Cidr:           subnet.Cidr,
					Description:    "",
					DisableAutoNAT: false,
				},
			})
			if err != nil {
				t.Errorf("could not marshal response json")
			}
			w.Write(data)

		} else {
			http.NotFound(w, r)
		}
	}))

	newSubnet, err := ma.Api.Subnets.Create(subnet)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if newSubnet.Id != "123" {
		t.Errorf("expected subnet id: 123, got %s", newSubnet.Id)
	}
}

func TestAccSubnetCRUD(t *testing.T) {
	api := NewApi(DefaultApiConfig())
	newSubnet, err := api.Subnets.Create(&Subnet{
		Name: "Test Subnet 1",
		Cidr: "10.5.0.0/16",
	})
	if err != nil {
		t.Fatalf("Error creating subnet. %s", err)
	}
	if newSubnet.Cidr != "10.5.0.0/16" {
		t.Errorf("Expected: '10.5.0.0/16'; Actual: '%s'", newSubnet.Cidr)
	}
	if newSubnet.Name != "Test Subnet 1" {
		t.Errorf("Expected: 'Test Subnet 1'; Actual: '%s'", newSubnet.Name)
	}

	newSubnet.Cidr = "10.5.0.0/24"
	_, err = api.Subnets.Update(newSubnet.Id, newSubnet)
	if err != nil {
		t.Fatalf("Error updating subnet. %s", err)
	}

	subnet, err2 := api.Subnets.Get(newSubnet.Id)
	if err2 != nil {
		t.Fatalf("Error getting subnet. %s", err2)
	}
	if subnet.Cidr != "10.5.0.0/24" {
		t.Errorf("Expected: '10.5.0.0/24'; Actual: '%s'", subnet.Cidr)
	}

	subnets, err3 := api.Subnets.All()
	if err3 != nil {
		t.Fatalf("Error getting all subnets. %s", err3)
	}
	if len(subnets) < 1 {
		t.Errorf("Expected at least 1 subnet.")
	}

	err = api.Subnets.Delete(subnet.Id)
	if err != nil {
		t.Fatalf("Error deleting subnet. %s", err)
	}
}
