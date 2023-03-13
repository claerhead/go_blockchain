package main

import (
	"github.com/claerhead/go_blockchain/explorer"
	"github.com/claerhead/go_blockchain/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
