package Room

import (
	"context"
	"fmt"
)

type Equipment struct {
	name     string
	color    string
	material string
}

func GetRoomEquipmentFromDB() ([]Equipment, error) {
	getDatabase()
	conn := getDatabase()

	rows, err := conn.Query(context.Background(), "select name, color, material from \"RoomEquipment\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipment []Equipment
	for rows.Next() {
		var item Equipment
		if err := rows.Scan(&item.name, &item.color, &item.material); err != nil {
			return nil, err
		}
		equipment = append(equipment, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return equipment, nil
}

func PrintRoomEquipment() {
	equipment, err := GetRoomEquipmentFromDB()
	FindError(err)

	fmt.Printf("\n\nВ комнате есть такая техника:")
	for i := 0; i < len(equipment); i++ {
		fmt.Print("\n\nНазвание: ", equipment[i].name, "\nЦвет: ", equipment[i].color, "\nМатериал: ", equipment[i].material)
	}
}
