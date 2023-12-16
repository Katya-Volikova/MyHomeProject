package main

type Equipment struct {
	name     string
	color    string
	material string
}

func createEquipment(name string,
	color string,
	material string,
) Equipment {
	return Equipment{
		name:     name,
		color:    color,
		material: material,
	}
}

func initializeEquipment() []Equipment {
	var equipment []Equipment

	// TODO брать данные из бд и использовать здесь
	//for i := 0; i < 5; i++ {
	//	equipment = append(equipment, createEquipment(itemsMock[0][i], itemsMock[1][i], itemsMock[2][i]))
	//}

	return equipment
}
