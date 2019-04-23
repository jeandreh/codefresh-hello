package cmd

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	reader *os.File
	writer *os.File
	err    error
)

func TestMain(m *testing.M) {
	reader, writer, err = os.Pipe()
	if err != nil {
		os.Exit(1)
	}
	rootCmd.SetOutput(writer)
	os.Exit(m.Run())
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"should print help by default",
			"help for codefresh-hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
			writer.Close()
			defer reader.Close()
			output, err := ioutil.ReadAll(reader)
			if err != nil {
				reader.Close()
				t.Fatalf("Unexpected empty output")
			}
			if !strings.Contains(string(output), tt.want) {
				t.Fatalf("Help output text not found")
			}
		})
	}
}
