package mysqlrunner_test

import (
	"CroonJoobGoo/mysqlrunner"
	"testing"
)

func TestMysqlrunner(t *testing.T) {
	mysqlrunner.RunScript("../example/runner_test.json", "")
}
