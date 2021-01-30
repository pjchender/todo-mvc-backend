package configs

type Database struct {
	DSN     string `default:"sslmode=disable host=localhost port=5432 dbname=todo_mvc"`
	TestDSN string `default:"sslmode=disable host=localhost port=5432 dbname=todo_mvc_test"`
}
