package models

import (
	"errors"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Define a custom type for the driver
type Driver string

// Function to validate the driver
func (d Driver) IsValid() error {
	for _, valid := range validDrivers {
		if d == valid {
			return nil
		}
	}
	return errors.New("invalid driver: " + string(d))
}


const (
	Postgres Driver = "postgres"
	MySQL    Driver = "mysql"
	SQLite   Driver = "sqlite"
)

// List of valid drivers
var validDrivers = []Driver{Postgres, MySQL, SQLite}

type Program struct {
	Project string
	driver 	Driver
	Exit 	bool
}


// ExitCLI checks if the Project has been exited, and closes
// out of the CLI if it has
func (p *Program) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

// Function to set the driver with validation
func (p *Program) SetDriver(driver string) error {
	d := Driver(strings.ToLower(driver))
	if err := d.IsValid(); err != nil {
		return err
	}
	p.driver = d
	return nil
}

func (p *Program) GetDriver() string {
	return string(p.driver)
}