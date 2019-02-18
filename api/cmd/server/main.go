package main

import (
	"talk-docker/pkg/rest"
	"talk-docker/pkg/storage"
)

func main() {
	repo, err := storage.NewRepository()

	if err != nil {
		panic(err)
	}

	server := rest.NewServer(repo)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
