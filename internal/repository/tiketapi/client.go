package tiketapi

import "gotick/internal/transport/httpclient"

type Repository struct {
	Client *httpclient.Client
}

func New(client *httpclient.Client) *Repository {
	return &Repository{
		Client: client,
	}
}