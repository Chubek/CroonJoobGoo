package postgresqlrunner

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"io/ioutil"
	"log"
	"strings"
)

type Profile struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Host     string   `json:"Host"`
	Database string   `json:"database"`
	Port     string   `json:"port"`
	Commands []string `json:"commands"`
	Queries	 []string  `json:"queries"`
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

	loginString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", profile.Username,
		profile.Password, profile.Host, profile.Port, profile.Database)

	fmt.Println("Connecting with string ", loginString)

	db, err := sql.Open("pgx", loginString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		handleErr(err)
	}(db)

	log.Println("Logging Exec for file ", profileLocation)

	for _, comm := range profile.Commands {
		commMod := strings.Trim(comm, "")
		commMod = strings.Trim(commMod, " ")
		commMod = strings.Trim(commMod, "\n")
		commMod = strings.Trim(commMod, "\r")

		_, err = db.Exec(commMod)

		log.Println("Exec statement ", comm, " done! But maybe there are errors...")

		handleErr(err)
	}
	log.Println("Logging Query for file ", profileLocation)

	for _, comm := range profile.Queries {
		commMod := strings.Trim(comm, "")
		commMod = strings.Trim(commMod, " ")
		commMod = strings.Trim(commMod, "\n")
		commMod = strings.Trim(commMod, "\r")

		_, err = db.Query(commMod)

		log.Println("Query ", comm, " done! But maybe there are errors...")

		handleErr(err)
	}


	log.Println("Done for ", profileLocation)
}
