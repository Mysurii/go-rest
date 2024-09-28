package steps

type Option struct {
	Title       string
	Description string
}

var DriverOptions = []Option{
	{
		Title:       "Postgres",
		Description: "A pure Go driver and toolkit for PostgreSQL.",
	},
	{
		Title:       "Mysql",
		Description: "MySQL-Driver for Go's database/sql package.",
	},
	{
		Title:       "Sqlite",
		Description: "Sqlite3 driver that conforms to the built-in database/sql interface.",
	},
}
