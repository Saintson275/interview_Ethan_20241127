# Go postgresql Restful APIs for OXO
This project is a backend management system for the OXO series of games built with the Go+fiber+Postgresql framework. It includes 5 modules:
- Players Managment Module
- Game Room Management Module
- Endless Challenge Module
- Game Log Module
- Payment Processing Module

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Folder Structure](#folder-structure)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Approach Taken](#approach-taken)
- [Results](#results)

## Installation

To run this project locally, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Saintson275/interview_Ethan_20241127.git
   
2. **Start the development server with docker::**:

   ```bash
   docker-compose up -d
   
3. **Start the development server with Terminal:**:

   ```bash
   go run main.go
   

   
## Usage
- Start the Docker service and use [Postman](https://www.postman.com) to test the API of each module.
- Create Post: Navigate to /create-post and fill out the form to add a new post with a title and content.

## Folder Structure
The project structure is organized as follows:
       .


    ├── handler/                # Contains business logic and route handlers
    │   ├── players.go          # Handlers for player-related operations
    │   ├── levels.go           # Handlers for level-related operations
    │   ├── rooms.go            # Handlers for room-related operations
    │   ├── reservations.go     # Handlers for reservation-related operations
    │   ├── challenge.go        # Handlers for game challenge operations
    │   ├── gameLog.go          # Handlers for game log operations
    │   ├── payments.go         # Handlers for payment-related operations
    │   └── routes.go           # Centralized route registration
    ├── models/                 # Defines the data structures and database models
    │   ├── players.go          # Player model definition
    │   ├── levels.go           # level model definition
    │   ├── rooms.go            # Room model definition
    │   ├── reservations.go     # Reservation model definition
    │   ├── challenge.go        # Challenge model definition
    │   ├── gameLog.go          # Game log model definition
    │   └── payments.go         # Payment model definition
    ├── storage/                # Contains database connection and storage logic
    │   └── postgre.go          # PostgreSQL database connection configuration
    ├── tests/                  # Contains unit and integration tests
    │   ├── payment_integration_test.go                  # Integration tests for payment-related APIs
    │   └── payment_test.go     # Unit tests for payment-related logic
    ├── Dockerfile              # Docker configuration for building the application container
    ├── docker-compose.yml      # Docker Compose configuration for multi-container setups
    ├── main.go                 # Main application entry point
    └── README.md               # Project README file

## Features
- Redux Toolkit: State management with Redux Toolkit, including creating slices and defining API queries.
- React Router: Navigation and routing between different components (HomePage, AddPostForm).
- API Integration: Integration with a mock API (jsonplaceholder.typicode.com) to fetch and add posts.
- Responsive Design: Uses Tailwind CSS for responsive and mobile-first design principles.
- Error Handling: Basic error handling for API requests and form submissions.
- Dialogs and Modals: Utilizes Shadcn Alert Dialog for displaying modal dialogs for success/error messages.

## Technologies Used
- [Go](https://go.dev)
- [Fiber](https://gofiber.io)
- [Postgres](https://www.postgresql.org)
- [Docker](https://www.docker.com)
- [Postman](https://www.postman.com)
- [Vercel](https://vercel.com) (Hosting database)

## Players management endpoint

1.**Get all players RequestUrl:http://localhost:8080/api/players**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "Samuel",
       "level": "等级一"
     }
   ]

2.**Post: Register a new player and receive a request in JSON format, including the player's name and level. Returns the new player's ID**:
   RequestUrl:http://localhost:8080/api/players
   ```json
   *requestBody:
   {
     "name": "string",  
     "level": "string"
   }

   *response:
   [
     {
       "id": 1,
       "name": "Samuel",
       "level": "等级一"
     }
   ]
   
