//go:build integration

// Package integration holds the tests that run against a real Nomad agent
// (docs/development.md#integration).
package integration

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

// nomadAddr is the agent the tests run against; without it they skip.
func nomadAddr(t *testing.T) string {
	t.Helper()
	addr := os.Getenv("NAP_TEST_NOMAD_ADDR")
	if addr == "" {
		t.Skip("NAP_TEST_NOMAD_ADDR is not set")
	}
	return addr
}

// TestNomadAgentIsReachable proves the CI agent is up and says which version
// the integration tests run against.
func TestNomadAgentIsReachable(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, nomadAddr(t)+"/v1/agent/self", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Nomad-Token", os.Getenv("NAP_TEST_NOMAD_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET /v1/agent/self: %s", resp.Status)
	}
	var self struct {
		Config struct {
			Version struct{ Version string }
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&self); err != nil {
		t.Fatal(err)
	}
	if self.Config.Version.Version == "" {
		t.Fatal("the agent reported no version")
	}
	t.Logf("nomad %s", self.Config.Version.Version)
}
