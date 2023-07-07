//go:build tools
// +build tools

//go:generate ent generate ./ent/schema
//go:generate ent describe ./ent/schema

package main
