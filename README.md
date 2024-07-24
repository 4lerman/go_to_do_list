# To-Do List

## Overview

This project is a simple To-Do List service implemented in Go. It provides a RESTful API to manage tasks, including creating, retrieving, updating, and deleting tasks. The service uses a sync.Map for in-memory storage and supports Swagger for API documentation.

## Features

- Create Tasks: Add new tasks with a title and activation date.
- Retrieve Tasks: Get a list of tasks with filtering options based on their status.
- Update Tasks: Modify existing tasks.
- Delete Tasks: Remove tasks from the list.
- Mark Tasks as Done: Change the task status to 'done'.
- Swagger Documentation: Access interactive API documentation.

## Prerequisites

- Go 1.22 or later
- Docker
- Docker Compose

## Installation

1. Clone the repository:

   ```sh
   git clone <repository-url>
   cd go_to_do_list
   ```

2. Build and run the Docker container:

   ```sh
   docker-compose up --build
   ```

## Usage

1. Access Swagger Documentation: Open http://localhost:8080/swagger/ to view and interact with the API documentation.

2. API Endpoints:

- GET /api/v1/todo-list/tasks: Retrieve tasks. Optionally filter by status (active or done).
- POST /api/v1/todo-list/tasks: Create a new task.
- PUT /api/v1/todo-list/tasks/{id}: Update an existing task.
- DELETE /api/v1/todo-list/tasks/{id}: Delete a task.
- PUT /api/v1/todo-list/tasks/{id}/done: Mark a task as done.
