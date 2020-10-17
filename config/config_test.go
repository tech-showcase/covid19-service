package config

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedOutput := Config{
		ServiceName: "dummy-service-name",
		Address:     "http://dummy.address/",
		Tracer: Tracer{
			AgentAddress: "dummy-tracer-agent-address",
		},
		Log: Log{
			Filepath: "dummy-filepath",
		},
	}

	os.Setenv("CONFIG_FILEPATH", "..")
	os.Setenv("CONFIG_FILENAME", ".env.example")

	config, err := Parse()

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}
