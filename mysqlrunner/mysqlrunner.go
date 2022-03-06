package mysqlrunner

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Profile struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Host     string   `json:"Host"`
	Database string   `json:"database"`
	Port     string   `json:"port"`
	Commands []string `json:"commands"`
	Queries  []string `json:"queries"`
	SaveFile string   `json:"save_file"`
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

	loginString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		profile.Username, profile.Password, profile.Host, profile.Port, profile.Database)

	db, err := sql.Open("mysql", loginString)
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

		log.Println("Exec statement ", comm, " done!")

		handleErr(err)
	}

	log.Println("Logging Query for file ", profileLocation)

	var allRows [][]string

	for _, comm := range profile.Queries {
		commMod := strings.Trim(comm, "")
		commMod = strings.Trim(commMod, " ")
		commMod = strings.Trim(commMod, "\n")
		commMod = strings.Trim(commMod, "\r")

		res, err := db.Query(commMod)

		log.Println("Query ", comm, " done! But maybe there are errors...")

		handleErr(err)

		log.Println("Saving query results (if no error, and if any)...")

		var currRow []string

		for res.Next() {
			var rowInterface string

			res.Scan(&rowInterface)

			fmt.Println("Got ", rowInterface)

			currRow = append(currRow, rowInterface)

		}
		allRows = append(allRows, currRow)
	}

	log.Println("Saving to text file...")

	file, _ := json.MarshalIndent(allRows, "", " ")

	_ = ioutil.WriteFile(profile.SaveFile, file, 0644)

	log.Println("Done for ", profileLocation)
}
