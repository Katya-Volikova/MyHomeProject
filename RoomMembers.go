package main

import "fmt"

type Neighbours struct {
	name    string
	age     string
	country string
	faculty string
}

func createNeighbour(name string, age string, country string, faculty string) Neighbours {
	return Neighbours{
		name:    name,
		age:     age,
		country: country,
		faculty: faculty,
	}
}

func initializeNeighbours() []Neighbours {
	var neighbours []Neighbours

	// TODO брать данные из бд и использовать здесь
	//for i := 0; i < 3; i++ {
	//	neighbours = append(neighbours, createNeighbour(neighboursMock[0][i], neighboursMock[1][i], neighboursMock[2][i], neighboursMock[3][i]))
	//}

	return neighbours
}

func infoAboutNeighbours(neighbours []Neighbours) {
	fmt.Print("В комнате живет 3 девушки")
	for i := 0; i < 3; i++ {
		fmt.Print("\nИмя: ", neighbours[i].name, "\nЕй ", neighbours[i].age, "\nОна учится на факультете - ", neighbours[i].faculty, "\n")
	}
}
