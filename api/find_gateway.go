package api

import (
	"fmt"
	"log"
	"time"
)

type MissingGatewayError struct {
	PublicIP string
}

func (err MissingGatewayError) Error() string {
	return fmt.Sprintf("could not find gateway with Public IP=%s.", err.PublicIP)
}

func (ap *Api) FindGatewayByPublicIP(public_ip string, wait bool) (*Gateway, error) {
	findFunc := func() (*Gateway, error) {
		gws, err := ap.Gateways.All()
		if err != nil {
			return nil, fmt.Errorf("error retrieving gateways: %s", err)
		}

		for _, cur := range gws {
			if cur.PublicIP == public_ip {
				return &cur, nil
			}
		}
		return nil, nil
	}

	start := time.Now()
	timeout, err := time.ParseDuration(ap.findTimeout)
	if err != nil {
		return nil, fmt.Errorf("invalid find gateway timeout: %s", err)
	}
	delay, _ := time.ParseDuration("10s")

	for {
		gw, err := findFunc()
		if err != nil {
			return nil, err
		}
		if gw != nil {
			return gw, nil
		}
		if !wait {
			return nil, nil
		}
		if time.Since(start) > timeout {
			return nil, MissingGatewayError{PublicIP: public_ip}
		}

		log.Printf("waiting %d to locate gateway (%s)...\n", delay.String(), public_ip)
		time.Sleep(delay)
	}
}
