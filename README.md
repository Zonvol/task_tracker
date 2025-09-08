# Task Tracker CLI

A simple command-line interface (CLI) application for tracking tasks. This application allows users to create, update, and delete tasks, providing a straightforward way to manage personal or team tasks.

## Project Structure

```
task-tracker-cli
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── task
│   │   └── task.go      # Task struct and methods
│   └── storage
│       └── storage.go   # Task storage interface and implementation
├── go.mod                # Module dependencies
├── go.sum                # Module dependency checksums
└── README.md             # Project documentation
```

## Installation

To install the project, clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd task-tracker-cli
```

Then, run the following command to download the dependencies:

```bash
go mod tidy
```

## Usage

To run the application, use the following command:

```bash
go run cmd/main.go
```

## Features

- Create new tasks
- Update existing tasks
- Delete tasks
- List all tasks

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.