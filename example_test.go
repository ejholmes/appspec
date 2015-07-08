package appspec

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
)

var app = &App{
	Name: "acme-inc",
	Processes: []Process{
		{
			Type:      "web",
			Instances: 1,
			Containers: []Container{
				{
					Image:   "remind101/acme-inc:latest",
					Command: []string{"./bin/web"},
					Env: map[string]string{
						"RAILS_ENV": "production",
					},
					Memory: 1073741824,
					CPU:    1024,
					Links:  []string{"statsd"},
					Ports: []Port{
						{
							Exposure: "public",
						},
					},
				},
				{
					Image:  "statsd:latest",
					Memory: 1073741824,
					CPU:    1024,
				},
			},
		},
		{
			Type:      "worker",
			Instances: 5,
			Containers: []Container{
				{
					Image:   "remind101/acme-inc:latest",
					Command: []string{"./bin/worker"},
					Env: map[string]string{
						"RAILS_ENV": "production",
					},
					Memory: 1073741824,
					CPU:    1024,
				},
			},
		},
	},
}

func TestExample(t *testing.T) {
	f, err := os.Create("appspec.json")
	if err != nil {
		t.Fatal(err)
	}

	w := io.MultiWriter(f, os.Stdout)
	raw, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	io.Copy(w, bytes.NewBuffer(raw))
}
