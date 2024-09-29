# Go REST CLI Generator

## Overview

The Go REST Application is a CLI tool that helps you scaffold a REST API project in Go. It allows users to create a new project, choose a database, and generate the necessary files and folders to kickstart their development.

## Features

- Interactive command-line interface to guide users through the setup process.
- Supports PostgreSQL, MySQL, SQLite, and MongoDB as database options.
- Automatically generates project structure and templates based on user input.
- Validates templates and ensures that all necessary files are created.

## Getting Started

### Prerequisites

- Go (version 1.18 or higher) installed on your machine.
- Git installed for version control (optional, but recommended).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-rest.git
   cd go-rest

2. Build the application

    ```bash
    go build -o go-rest


### Running the Application

To start the application, run the following command:

    ./go-rest

### Usage

When you run the application, it will prompt you for the following:

1. Project Name: Enter the name of your project.
2. Database Selection: Choose the database you want to use (PostgreSQL, MySQL or SQLite).

The application will then create the project structure and generate the necessary files based on your selections.

## Project Structure

After running the application, your project structure will look something like this:


```
my-awesome-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── database/
│   │   └── database.go
│   └── server/
│       ├── routes.go
│       └── server.go
└── .air
└── .env
└── Dockerfile
└── go.mod
└── go.sum
└── README.md
```
## Contributing

If you'd like to contribute to the project, feel free to open an issue or submit a pull request. Contributions are always welcome!




