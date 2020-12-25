package paysharestoml

import "log"

// ExampleGetTOML gets the payshares.toml file for coins.asia
func ExampleGetTOML() {
	_, err := DefaultClient.GetPaysharesToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
