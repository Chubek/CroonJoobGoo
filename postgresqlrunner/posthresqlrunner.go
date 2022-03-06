package postgresqlrunner

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
	"strings"
	"context"

)

type Profile struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Host     string   `json:"Host"`
	Database string   `json:"database"`
	Port     string   `json:"port"`
	Commands []string `json:"commands"`
	Queries	 []string  `json:"queries"`
	SaveFile string    `json:"save_file"`
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

	db, err := pgx.Connect(context.Background(), loginString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	log.Println("Logging Exec for file ", profileLocation)

	for _, comm := range profile.Commands {
		commMod := strings.Trim(comm, "")
		commMod = strings.Trim(commMod, " ")
		commMod = strings.Trim(commMod, "\n")
		commMod = strings.Trim(commMod, "\r")

		_, err = db.Exec(context.Background(), commMod)

		log.Println("Exec statement ", comm, " done! But maybe there are errors...")

		handleErr(err)
	}
	log.Println("Logging Query for file ", profileLocation)

	var allRows []interface{}

	for _, comm := range profile.Queries {
		commMod := strings.Trim(comm, "")
		commMod = strings.Trim(commMod, " ")
		commMod = strings.Trim(commMod, "\n")
		commMod = strings.Trim(commMod, "\r")

		res, err := db.Query(context.Background(), commMod)

		log.Println("Query ", comm, " done! But maybe there are errors, if there are, they will show up...")

		handleErr(err)

		for res.Next() {
			var rowInterface interface{}

			res.Scan(&rowInterface)

			log.Println("Row: ", rowInterface)

			allRows = append(allRows, rowInterface)

		}

		
	}
	log.Println("Saving to text file...")

	file, _ := json.MarshalIndent(allRows, "", " ")

	_ = ioutil.WriteFile(profile.SaveFile, file, 0644)

	log.Println("Done for ", profileLocation, " if no errors appeared, it should be A-Ok.")
}
