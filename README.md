# Golang API Application

This repository houses a Golang-based API application designed for managing organizations. The application includes features such as token management, CRUD operations for organizations, user invitations, and integration with MongoDB using Docker.

## Project Structure

The project structure is designed to assist you in getting started quickly. You can modify it as needed for your specific requirements.


- **pkg/**: Core logic of the application divided into different packages.
  - **api/**: API handling components.
    - **handlers/**: API route handlers.
        **userHandler.go** : handle user requests
        **orgHandler.go** : handle organizaation requests 
    - **middleware/**: Middleware functions.
      -**middleware.go** : **contains AuthMiddleware for bearer Token "5425861"** 
    - **routes/**: Route definitions.
        -**routes.go** 
  - **controllers/**: Business logic for each route.
      **userController.go** : user functions
      **orgController.go** : organization functions
  - **database/**: Database-related code.
    - **mongodb/**
      - **models/**: Data models.
          **user.go** user model
          **organization.go** organization model
      - **repository/**: Database operations.
          **connect.go** :connect database and create collections
          **userQueries.go** :user queries
          **orgQueries.go** :organization queries
  - **utils/**: Utility functions.
  - **app.go**: Application initialization and setup. **Must update mongodb localhost if want to run locally**
- **main.go**: The entry point of the application.
- **Dockerfile**: Instructions for building the application image.
- **docker-compose.yaml**: Configuration for Docker Compose.

- **config/**: Configuration files for the application.
  - **app-config.yaml**: General application settings.
  - **database-config.yaml**: Database connection details.

- **tests/**: Directory for tests.
  - **e2e/**: End-to-End tests.
  - **unit/**: Unit tests.

- **.gitignore**: Specifies files and directories to be ignored by Git.

## Getting Started

To begin working with the application, follow the instructions in the project documentation. Feel free to adjust the project structure as needed based on your preferences and evolving project requirements.
