package main

import (
	"context"
	"log"

	"github.com/Gustavo-Feijo/leago"

	"github.com/Gustavo-Feijo/leago/api/ddragon/realms"
	"github.com/Gustavo-Feijo/leago/examples"
)

func main() {
	// go run examples/ddragon/main.go
	client := leago.NewDDragonClient(
		realms.RealmNA,

		// Options to override the default realm values.
		// Can be used to force older version or different language.
		[]leago.DDragonOption{
			leago.WithVersion("16.1.1"),
			leago.WithLocale("pt_BR"),
		},
	)

	champion, err := client.DDragon.Champion.GetChampionByID(context.Background(), "Aatrox")
	if err != nil {
		log.Fatal(err)
	}

	examples.PrettyPrint(champion)
}
