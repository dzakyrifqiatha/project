package server

import (
	"context"
)

type donaturBk struct {
	writer ReadWriter
}

func NewDonaturBk(writer ReadWriter) DonaturBukuService {
	return &donaturBk{writer: writer}
}

func (dbk *donaturBk) AddDBk