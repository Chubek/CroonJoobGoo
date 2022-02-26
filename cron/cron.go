package cron

import (
	"CroonJoobGoo/mssqlrunner"
	"CroonJoobGoo/mysqlrunner"
	"CroonJoobGoo/postgresqlrunner"
	"time"
)

type ExecTime struct {
	Days       []int `json:"days"`
	Hours      []int `json:"hours"`
	Minutes    []int `json:"minutes"`
	AllHours   bool  `json:"all_hours"`
	AllMinutes bool  `json:"all_minutes"`
	AllDays    bool  `json:"all_days"`
}

type SqlType = string

const (
	MSSQL      SqlType = "MSSQL"
	MYSQL      SqlType = "MYSQL"
	POSTGRESQL SqlType = "POSTGRESQL"
)

type RunnerConfig struct {
	Type       SqlType  `json:"sql_type"`
	ExecTime   ExecTime `json:"exec_time"`
	ConfigPath string   `json:"config_path"`
}

func Contains(iterable []int, el int) bool {
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

	dayIs := Contains(baseTime.Days, day) || baseTime.AllDays
	hourIs := Contains(baseTime.Hours, hour) || baseTime.AllHours
	minIs := Contains(baseTime.Minutes, minute) || baseTime.AllMinutes

	if dayIs && hourIs && minIs {
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

		if diffSecs%58 == 0 {
			if CheckTimeIs(&config.ExecTime) {
				switch config.Type {
				case POSTGRESQL:
					postgresqlrunner.RunScript(config.ConfigPath, time.Now().Format("15:04:05"))
				case MSSQL:
					mssqlrunner.RunScript(config.ConfigPath, time.Now().Format("15:04:05"))
				case MYSQL:
					mysqlrunner.RunScript(config.ConfigPath, time.Now().Format("15:04:05"))

				}

				time.Sleep(5 * time.Second)
			}

		}

	}

}
