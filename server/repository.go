package server

import "github.com/manfromdownunder/eKonyv/core"

type Repository struct {
	servers core.IrcServers
}

func NewRepository() *Repository {
	return &Repository{servers: core.IrcServers{}}
}
