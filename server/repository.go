package server

import "github.com/eKonyv/eKonyv/core"

type Repository struct {
	servers core.IrcServers
}

func NewRepository() *Repository {
	return &Repository{servers: core.IrcServers{}}
}
