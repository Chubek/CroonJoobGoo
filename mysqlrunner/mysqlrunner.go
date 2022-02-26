package mysqlrunner

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type Command struct {
	Statement string      `json:"statement"`
	Args      interface{} `json:"args"`
}

type Profile struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Host     string    `json:"Host"`
	Database string    `json:"database"`
	Port     string    `json:"port"`
	Commands []Command `json:"commands"`
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func RunScript(profileLocation string, time string) {
	log.Println("SQL exec started at ", time, " for config file ", profileLocation)

	read, err := ioutil.ReadFile(profileLocation)

	handleErr(err)

	var profile Profile

	err = json.Unmarshal(read, &profile)

	handleErr(err)

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%sd@tcp(%s:%s)/%s", profile.Username, profile.Password, profile.Host, profile.Port, profile.Database))
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		handleErr(err)
	}(db)

	log.Println("Logging Exec for file ", profileLocation)

	for _, comm := range profile.Commands {
		_, err = db.Exec(comm.Statement, comm.Args)

		log.Println("Exec statement ", comm.Statement, " done!")

		handleErr(err)
	}

	log.Println("Done for ", profileLocation)
}
