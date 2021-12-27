package id

import "github.com/rs/xid"

func Generate() string {
	return xid.New().String()
}
