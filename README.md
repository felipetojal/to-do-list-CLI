#Go Task CLI

A lightweight and efficient command-line interface (CLI) task manager, developed in Go. This project uses local JSON persistence and demonstrates the use of flags, file manipulation, and data structures in Go.

# Features

Task Creation: Adds tasks with auto-incrementing IDs (Sequence).

Optional Description: Support for titles and detailed descriptions.

Listing: Formatted view of all pending and completed tasks.

Completion: Marks tasks as done.

Removal: Deletes tasks while maintaining list integrity.

Persistence: Data is automatically saved in the tasks.json file.

# Installation and Execution

Prerequisites: Go installed (version 1.22 or higher).

Clone the repository:

git clone [https://github.com/felipetojal/to-do-list-CLI.git](https://github.com/felipetojal/to-do-list-CLI.git)
cd to-do-list-CLI


Run directly via Go:

go run . -help


(Optional) Build the binary executable:

go build -o task-cli .
# Now you can just run: ./task-cli -list


# Usage Guide

1. Add a Task

Use the -add flag for the title. Optionally use -desc for a description.

Simple
go run . -add "Study Go"

With description (Use quotes if it contains spaces)
go run . -add "Finish Project" -desc "Write the README and push to Git"


2. List Tasks

Displays ID, Title, Description, Status, and Creation Date.

go run . -list


3. Complete a Task

Use the -complete flag passing the task ID.

go run . -complete 1


4. Delete a Task

Use the -delete flag passing the task ID.

go run . -delete 1


# Project Structure

The project follows a simple architecture for CLIs:

main.go: Entry point. Manages Flags (Inputs) and control flow.

tasks.json: Local database (automatically generated on first run).

Load/Save Logic: Persistence functions using encoding/json and os.File.

Developed by Felipe Tojal for Software Engineering study purposes.
