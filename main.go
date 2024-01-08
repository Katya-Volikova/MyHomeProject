package main

import (
	_ "github.com/lib/pq"
	"main/Room"
)

type dbInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("load config err | ", err)
	//}
	//
	//pgCfg := dbInfo{}
	//
	//err = envconfig.Process("db", &pgCfg)
	//if err != nil {
	//	log.Fatal("parse config err | ", err)
	//}
	//
	//DSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pgCfg.Host, pgCfg.Port, pgCfg.User, pgCfg.Password, pgCfg.Name)
	//fmt.Println(DSN)
	//
	//db, err := sql.Open("postgres", DSN)
	//if err != nil {
	//	log.Fatal("SQL OPEN ERR | ", err)
	//}
	//
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal("PING ERR | ", err)
	//}
	//
	//fmt.Println("+vibe")

	Room.BuildRoom()

}
