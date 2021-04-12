package book

import "io"

type NewBookRequest struct {
	File io.ReadCloser
}
