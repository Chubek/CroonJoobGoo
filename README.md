# Go SQL Parallax Cron!

This project allows you to run parallex CRON jobs in Go. The CRON system is written
in pure Go. 

Basically, you have a master config file, let's call it `config.json` which:

A) Either sits in the same folder as the executable or the file, called, called `config.json`;
Or B) You can pass to the executable.

This JSON file is an array with objects that look like the ones you see in `config.json` of this project,
for example:

```
[
  {
    "sql_type": "MYSQL",
    "exec_time": {
        "days": [26, 27, 28],
        "hours": [],
        "minutes": [13, 15, 16],
        "all_minutes": false,
        "all_hours": true,
        "all_days": false
    },
    "config_path": "./example/runner_test.json"
  }
]

```

The array can contain as many objects as you desire. If you don't wanna pass a day and want to check "all_days", pass an empty array.
Same is true about hour and day. Otherwise pass days, hours, and minutes; and it will run the job at that time.

Supported SQL types are: MYSQL, MSSQL, POSTGRESSQL (keep in mind, it should be in capital!)

The sub-config file looks like this:

```
{
  "username": "superchubak",
  "password": "supsup",
  "host": "localhost",
  "database": "test_db",
  "port": "3306",
  "commands": [
    "CREATE TABLE Test (firstname varchar(255))",
    "INSERT INTO Suppose (firstname) VALUES  ('zuzie')"
  ]
}
```

There's one in `example` folder. 

You can have as many configs as you like, it will run in parallel ;)

If you have a question direct them to Chubak#7400 on Discord.