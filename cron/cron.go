package cron

import (
	"time"
	"CroonJoobGoo/mssqlrunner"
	"CroonJoobGoo/mysqlrunner"
	"CroonJoobGoo/postgresqlrunner"

)

type ExecTime struct {
	Days    []int
	Hours   []int
	Minutes []int
}


type SqlType	= string

const (
	MSSQL 		SqlType 	= "MSSQL"
	MYSQL 		SqlType		= "MYSQL"
	POSTGRESQL	SqlType		= "POSTGRESQL"
)

type RunnerConfig struct {
	Type  		SqlType 	`json:"sql_type"`
	ExecTime	ExecTime	`json:"exec_time"`
	ConfigPath	string		`json:"config_path"`
}

func Contains(iterable []int, el int) bool
{
	for _, v := range iterable {
		if v == el {
			return true
		}
	}

	return false
}

func CheckTimeIs(baseTime *ExecTime) bool {
	now := time.Now()

	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()

	dayIs := Contains(baseTime.Days, day)
	hourIs := Contains(baseTime.Hours, hour)
	minIs := Contains(baseTime.Minutes, minute)

	if  dayIs && hourIs && minIs {
		return true
	}

	return false
}



func RoutineExec(config RunnerConfig) {
	tZero := time.Now()

	var tOne time.Time
	var diff time.Duration
	var diffSecs uint64

	for {
		tOne = time.Now()
		diff = tOne.Sub(tZero)

		diffSecs = uint64(diff.Seconds())

		if diffSecs % 5 == 0 {
			if CheckTimeIs(&config.ExecTime) {
				switch config.Type {
				case POSTGRESQL:
					postgresqlrunner.RunScript(config.ConfigPath, )time.Now().Format("15:04:05")
				case MSSQL:
					 mssqlrunner.RunScript(config.ConfigPath, time.Now().Format("15:04:05"))
				case MYSQL:
					mysqlrunner.RunScript(config.ConfigPath, time.Now().Format("15:04:05"))

				}

			}
		}



	}

}


