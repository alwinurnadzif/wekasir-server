package config

import "fmt"

var pages = struct {
	users          string
	branches       string
	branchAccesses string
}{
	users:          "users",
	branches:       "branches",
	branchAccesses: "branchAccesses",
}

var Pages = map[string]string{
	"users-create": fmt.Sprintf("%v.create", pages.users),
	"users-view":   fmt.Sprintf("%v.view", pages.users),
	"users-delete": fmt.Sprintf("%v.delete", pages.users),
	"users-update": fmt.Sprintf("%v.update", pages.users),

	"branches-create": fmt.Sprintf("%v.create", pages.branches),
	"branches-view":   fmt.Sprintf("%v.view", pages.branches),
	"branches-delete": fmt.Sprintf("%v.delete", pages.branches),
	"branches-update": fmt.Sprintf("%v.update", pages.branches),

	"branchAccesses-create": fmt.Sprintf("%v.create", pages.branchAccesses),
	"branchAccesses-view":   fmt.Sprintf("%v.view", pages.branchAccesses),
	"branchAccesses-delete": fmt.Sprintf("%v.delete", pages.branchAccesses),
	"branchAccesses-update": fmt.Sprintf("%v.update", pages.branchAccesses),
}
