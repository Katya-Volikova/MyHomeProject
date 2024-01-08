package Room

import (
	"context"
	"fmt"
)

type Neighbours struct {
	name    string
	age     string
	country string
	faculty string
}

func GetRoomMembersFromDB() ([]Neighbours, error) {
	getDatabase()
	conn := getDatabase()

	rows, err := conn.Query(context.Background(), "select name, age, faculty, country from \"RoomMembers\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var neighbours []Neighbours
	for rows.Next() {
		var person Neighbours
		if err := rows.Scan(&person.name, &person.age, &person.faculty, &person.country); err != nil {
			return nil, err
		}
		neighbours = append(neighbours, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return neighbours, nil
}

func PrintRoomMembers() {
	neighbours, err := GetRoomMembersFromDB()
	FindError(err)

	fmt.Printf("\n\nВ комнате живут:")
	for i := 0; i < len(neighbours); i++ {
		fmt.Print("\n\nИмя: ", neighbours[i].name, "\nВозраст: ", neighbours[i].age, " лет", "\nФакультет: ", neighbours[i].faculty, "\nСтрана: ", neighbours[i].country)
	}
}
