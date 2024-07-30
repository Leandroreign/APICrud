package model

type Community struct {
	Name string
}

// Communitites slice de comunidades
type Communitites []Community

// Person estructura de una persona
type Person struct {
	// Name nombre de la persona
	Name string
	// Age edad de la persona
	Age uint8
	// Communities comunidades a la que pertenece la persona
	Communities Communitites
}

type Persons []Person
