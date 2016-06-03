package api

import (
	"os"
	"testing"
)

func TestAccPreCheck(t *testing.T) {
	if v := os.Getenv("FORTYCLOUD_ACCESS_KEY"); v == "" {
		t.Fatal("FORTYCLOUD_ACCESS_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("FORTYCLOUD_SECRET_KEY"); v == "" {
		t.Fatal("FORTYCLOUD_SECRET_KEY must be set for acceptance tests")
	}
}
