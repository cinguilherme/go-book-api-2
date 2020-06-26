package health

import "github.com/cinguilherme/go-api-fiber/database"

//Health struct type
type Health struct {
	Status       string       `json:"status"`
	Dependencies []Dependency `json:"dependencies"`
}

//Dependency struct type
type Dependency struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Options string `json:"options"`
}

func NewHealth() Health {

	var health Health
	health.Status = "UP"

	db := dbDepCheck()
	redis := redisCheck()

	health.Dependencies = []Dependency{db, redis}

	return health
}

func dbDepCheck() Dependency {
	dbx := database.DBConn
	status := "DOWN"
	options := "Unnable to connect"

	if dbx != nil {
		status = "UP"
		options = "connection is up"
	}

	var db Dependency
	db.Name = "postgres"
	db.Status = status
	db.Options = options
	return db
}

func redisCheck() Dependency {
	var redis Dependency
	redis.Name = "redis"
	redis.Status = "not implemented"
	redis.Options = "not implemented"
	return redis
}
