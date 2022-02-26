package CroonJoobGoo

import (
	"CroonJoobGoo/cron"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	configFile := "./config.json"

	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	var config []cron.RunnerConfig

	bytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(len(config))
	for _, conf := range config {
		conf := conf
		go func() {
			cron.RoutineExec(conf)
		}()
	}
	wg.Wait()

}
