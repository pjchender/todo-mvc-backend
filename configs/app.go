package configs

type App struct {
	Name string `default:"todo-mvc"`
	Mode string `default:"dev"`
}
