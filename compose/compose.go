// Pacakge compose provides methods for decoding a docker-compose.yml file into
// an appspec.
package compose

import (
	"io"

	"github.com/ejholmes/appspec"
)

// Decode decodes the yml formatted docker-compose.yml file into an appspec.App.
func Decode(r io.Reader) (appspec.App, error) {
	return appspec.App{}, nil
}
