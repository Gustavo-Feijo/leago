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

	// DDragon client initialization can use IO if version and locale are not provided and need to be fetched.
	client, err := leago.NewDDragonClient(
		realms.RealmNA,

		// Options to override the default realm values.
		// Can be used to force older version or different language.
		[]leago.DDragonOption{
			leago.WithVersion("16.1.1"),
			leago.WithLocale("pt_BR"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	champion, err := client.DDragon.Champion.GetChampionByID(context.Background(), "Aatrox")
	if err != nil {
		log.Fatal(err)
	}

	examples.PrettyPrint(champion)
}
