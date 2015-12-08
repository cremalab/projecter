package main

type provider interface {
	AddRoutes(map[string][]route, config) map[string][]route
}
