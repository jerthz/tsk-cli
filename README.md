
# tsk-cli

`tsk-cli` is a command-line tool written in Go for efficiently managing your tasks directly from the terminal. It allows you to create, track, and organize tasks while logging start and end times for better time management.

## Features

- Add, edit, and delete tasks
- Automatically record start and end times for tasks
- Filter and sort tasks by date, priority, or status
- Use kanban, daily, today etc commands 

## Installation

Make sure you have Go installed on your system (version 1.18 or later).

```bash
git clone https://github.com/jerthz/tsk-cli.git
cd tsk-cli
go build -o tsk
```

Then, add the binary to your PATH:

```bash
export PATH=$PATH:$(pwd)
```

Or move the bin tsk in any directory you already have in your $PATH

## Usage

### Add a Task using NVIM editor (not configurable right now)

```bash
tsk add
```

### List Tasks

```bash
tsk list
```

### Start a Task

```bash
tsk start <task_id>
```

## Contribution

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a branch (`git checkout -b feature-new-feature`)
3. Commit your changes (`git commit -m 'Add a new feature'`)
4. Push your branch (`git push origin feature-new-feature`)
5. Open a Pull Request

## Authors

- **Jerthz** - Lead Developer
