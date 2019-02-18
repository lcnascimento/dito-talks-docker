package main

import (
	"talk-docker/pkg"
	"talk-docker/pkg/storage"
)

func main() {
	repo, err := storage.NewRepository()

	if err != nil {
		panic(err)
	}

	for _, talk := range talks() {
		if err := repo.CreateTalk(talk); err != nil {
			panic(err)
		}
	}
}

func talks() []pkg.Talk {
	return []pkg.Talk{
		pkg.Talk{
			Name:    "Dockerizando uma aplicação Web",
			Place:   "Dito",
			Speaker: "Luís Nascimento",
		},
	}
}
