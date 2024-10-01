# Task Tracker CLI

** Task Tracker** is a CLI project used to trarck and manage your tasks.
Challenge from task-tracker: https://roadmap.sh/projects/task-tracker

## Features

- **Add tasks** with a description.
- **List tasks** by status: `todo`, `in-progress`, `done`, or all tasks.
- **Update the status** of a task.
- **Delete tasks** by ID.
- Runs inside **Docker**, so you don’t need to install Go locally.
- Task data is persisted in a JSON file.

## Prerequisites

- [Docker](https://www.docker.com/) installed on your machine.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/yask-tracker.git
   cd yask-tracker

   ```

2. Build the Docker image:

   ```
   docker build -t yask-tracker .

   ```

3. Make sure there is a tasks.json file in the data folder. If it doesn't exist, it will be created automatically when you add your first task.

4. To get help on how to use the commands:

   ```
   docker run -it --rm yask-tracker help
   ```
