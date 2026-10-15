package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunVersionPrintsTheVersion(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run([]string{"-version"}, &out, &errOut); code != 0 {
		t.Fatalf("exit code %d, stderr %q", code, errOut.String())
	}
	if strings.TrimSpace(out.String()) == "" {
		t.Error("-version printed nothing")
	}
}

func TestRunUnknownFlagFails(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run([]string{"-nope"}, &out, &errOut); code != 2 {
		t.Errorf("exit code %d, want 2", code)
	}
}
