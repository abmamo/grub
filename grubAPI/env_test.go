package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	// env variable
	testEnv := []byte("TEST_ENV = \"test\"")
	// write to test env file
	err := ioutil.WriteFile(".env.test", testEnv, 0644)
	// check writing succeeded
	if err != nil {
		t.Errorf("write env var failed.")
	}
	// get environment variable from test file
	testVar := getEnvironment("TEST_ENV", ".env.test")
	// check correct value
	if testVar != "test" {
		t.Errorf("get environment var failed.")
	}
	// remove test environment
	os.Remove(".env.test")
}
