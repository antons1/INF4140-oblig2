package main

import ()

var (
	// A list of potential names for passengers
	personNames = []string{
		"Fredrik",
		"Per",
		"Hans",
		"Håkon",
		"Sigrid",
		"Ulla",
		"Ragnhild",
		"Pernille",
		"Karoline",
		"Kristoffer",
		"Anders",
		"Kari",
		"Frida",
		"Kjell",
		"Inge",
		"Kristiane",
		"Arne",
		"Knut",
		"Tor",
		"Geir",
		"Terje",
		"Øivind",
		"Marius",
		"Øystein",
		"Peder",
		"Hilde",
		"Bjørg",
		"Anna",
		"Solveig",
		"Randi",
		"Ida",
		"Tone",
		"Tove",
	}

	// A list of potential names for the coaster
	coasterNames = []string{
		"Adventure Express",
		"Alpine Bobsled",
		"Dahlonega Mine Train",
		"Drachen Fire",
		"Flight of the Hippogriff",
		"Roaring Waterfalls",
		"Raging Bull",
		"The Tornado",
		"Flying Saucers",
		"Terrible Pterodactyl",
		"Mister Bones Wild Ride",
	}
)

// GetPersonName returns a random name for a person
func GetPersonName() string {
	ind := GetRandomNumber(0, len(personNames))

	return personNames[ind]
}

// GetCoasterName returns a random name for a coaster
func GetCoasterName() string {
	ind := GetRandomNumber(0, len(coasterNames))

	return coasterNames[ind]
}
