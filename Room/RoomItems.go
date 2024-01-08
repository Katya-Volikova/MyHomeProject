package Room

import (
	"context"
	"fmt"
)

type Items struct {
	name     string
	size     string
	color    string
	material string
}

func GetRoomItemsFromDB() ([]Items, error) {
	getDatabase()
	conn := getDatabase()

	rows, err := conn.Query(context.Background(), "select name, size, color, material from \"RoomItems\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Items
	for rows.Next() {
		var item Items
		if err := rows.Scan(&item.name, &item.size, &item.color, &item.material); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func PrintRoomItems() {
	items, err := GetRoomItemsFromDB()
	FindError(err)

	fmt.Printf("\n\nВ комнате есть такие предметы:")
	for i := 0; i < len(items); i++ {
		fmt.Print("\n\nНазвание: ", items[i].name, "\nРазмер: ", items[i].size, " м", "\nЦвет: ", items[i].color, "\nМатериал: ", items[i].material)
	}
}
