package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetExcercises",
		"GET",
		"/exercises",
		GetExercises,
	},
	Route{
		"AddExcercise",
		"POST",
		"/exercises",
		AddExercise,
	},
	Route{
		"GetExcercise",
		"GET",
		"/exercises/{id}",
		GetExercise,
	},
	Route{
		"StartExercise",
		"POST",
		"/exercises/start/{id}",
		StartExercise,
	},
}
