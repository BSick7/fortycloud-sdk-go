package api

import "fmt"

func (ap *Api) FindSubnet(cidr string, gatewayID string) (*Subnet, error) {
	subnets, err := ap.Subnets.All()
	if err != nil {
		return nil, fmt.Errorf("error retrieving subnets: %s", err)
	}

	for _, subnet := range subnets {
		if subnet.Cidr == cidr && subnet.GatewayID() == gatewayID {
			return &subnet, nil
		}
	}

	return nil, nil
}
