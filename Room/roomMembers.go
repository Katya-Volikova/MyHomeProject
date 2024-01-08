package Room

import (
	"context"
	"fmt"
)

type Person struct {
	name    string
	age     string
	country string
	faculty string
}

func GetRoomMembersFromDB() ([]Person, error) {
	getDatabase()
	conn := getDatabase()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select name, age, faculty, country from \"RoomMembers\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomMembers []Person
	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.name, &person.age, &person.faculty, &person.country); err != nil {
			return nil, err
		}
		roomMembers = append(roomMembers, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roomMembers, nil
}

func PrintRoomMembers() {
	roomMembers, err := GetRoomMembersFromDB()
	FindError(err)

	PrintTitle("ЖИТЕЛИ")
	for i := 0; i < len(roomMembers); i++ {
		fmt.Print("\n\nИмя: ", roomMembers[i].name, "\nВозраст: ", roomMembers[i].age, " лет", "\nФакультет: ", roomMembers[i].faculty, "\nСтрана: ", roomMembers[i].country)
	}
}
