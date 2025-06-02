//package ent
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.Options{})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
