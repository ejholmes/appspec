package appspec_test

import "github.com/ejholmes/appspec"

func Example() {
	app := &appspec.App{
		Name: "acme-inc",
		Processes: []appspec.Process{
			{
				Type:      "web",
				Instances: 1,
				Containers: []appspec.Container{
					{
						Image:   "remind101/acme-inc:latest",
						Command: []string{"./bin/web"},
						Env: map[string]string{
							"RAILS_ENV": "production",
						},
						Memory: 1073741824,
						CPU:    1024,
						Links:  []string{"statsd"},
						Ports: []appspec.Port{
							Exposure:  "public",
							Host:      9000,
							Container: 8080,
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
				Containers: []appspec.Container{
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
}
