package mockery

import "io"

type RequesterIface interface {
	Get() io.Reader
}
