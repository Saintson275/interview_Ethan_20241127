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
- [Players Management Endpoint](#players-management-endpoint)
- [Rooms Management Endpoint](#rooms-management-endpoint)
- [Endless Challenge Endpoint](#endless-challenge-endpoint)
- [Game Log Endpoint](#game-log-endpoint)
- [Payment Processing Endpoint](#payment-processing-endpoint)

## Installation

To run this project locally, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Saintson275/interview_Ethan_20241127.git
   
2. **Start the development server with docker**:

   ```bash
   docker-compose up -d
   
3. **Start the development server with Terminal**:

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
1. Fiber Web Framework
Built with the high-performance Fiber web framework, enabling the rapid development of a lightweight RESTful API with clean and intuitive code structure.

2. Database Management (GORM)
Integrated with GORM as the Object-Relational Mapper (ORM) for seamless interactions with a PostgreSQL database, supporting model definition, migrations, and complex queries.

3. Payment Functionality
Implements support for multiple payment methods with simulated logic, including:
    - Credit Card Payments
 - Validation for card number, expiry date, and CVV, with simulated transaction failure.
 - Bank Transfers
 - Validation for account and bank details, supporting simulated transaction failure.
 - Third-Party Payment Platforms
 - Extensible to integrate third-party payment APIs with mock transaction logic.
 - Blockchain Payments
 - Verification of wallet addresses with low-probability failure simulation.

4. Modular Design and Routing
A modular architecture utilising Fiber’s routing system, divided into distinct modules (e.g., Players, Levels, Rooms, Payments), ensuring maintainability and ease of expansion.

5. Data Modelling and Migrations
Supports automated database table generation and model migration using GORM, covering key entities such as:
 - Player Management
 - Level Management
 - Room Management
 - Reservation Management
 - Challenge Management
 - Payment Management

6. Unit and Integration Testing
Unit Tests
Comprehensive tests for the payment module, ensuring edge case handling and robust error management.
Integration Tests
End-to-end tests for the payment API, verifying request handling, database interaction, and response accuracy.

7. Dockerised Deployment
Provides Docker and Docker Compose support for streamlined local development and production deployment, including:
 - Configuration of a PostgreSQL database container.
 - Automated service orchestration.

8. Error Handling and Logging
 - Global Error Middleware
 - Handles application-wide errors, returning user-friendly JSON responses.
 - Logging all requests and database operations to facilitate debugging and monitoring.

9. Environment Variable Management
Employs dotenv for environment variable handling, enabling secure and centralised management of configuration details, including database credentials and server settings.

## Technologies Used
- [Go](https://go.dev)
- [Fiber](https://gofiber.io)
- [Postgres](https://www.postgresql.org)
- [Docker](https://www.docker.com)
- [Postman](https://www.postman.com)
- [Vercel](https://vercel.com) (Hosting database)

## Players Management Endpoint

1. **[Get] Fetch all players (requestUrl:http://localhost:8080/api/players)**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "Samuel",
       "level": "等级一"
     },
      ...
   ]

2. **[Post] Create a player (RequestUrl:http://localhost:8080/api/players)**:
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
     },
   ]
3. **[Get] Fetch players by ID (requestUrl:http://localhost:8080/api/players/{id})**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "Samuel",
       "level": "等级一"
     }
   ]

4. **[Put] Update player by ID (RequestUrl:http://localhost:8080/api/players/{id})**:
   ```json
   *requestBody:
   {
     "id": 1,  
   }
   *response:
   [
     {
       "id": 1,
       "name": "Samuel_change",
       "level": "等级二"
     }
   ]
   
5. **[Delete] Delete player by ID (RequestUrl:http://localhost:8080/api/players/{id})**:
   ```json
   *response:
   {message: delete player successful}
   
6. **[Get] Fetch all levels (requestUrl:http://localhost:8080/api/levels)**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "等级一"
     },
      ...
   ]

7. **[Post] Create a new level (RequestUrl:http://localhost:8080/api/levels)**:
   ```json
   *requestBody:
   {
     "name": "string"
   }

   *response:
   [
     {
       "id": 2,
       "name": "等级二",
     },
   ]

## Rooms Management endpoint

1. **[Get] Fetch all rooms (requestUrl:http://localhost:8080/api/rooms)**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "TestingRoom",
       "description": "Aroom",
       "status": "available"
     },
      ...
   ]

2. **[Post] Create a player (RequestUrl:http://localhost:8080/api/rooms)**:
   ```json
   *requestBody:
   {
     "name": "string",
     "description": "string",
     "status": "string"
   }

   *response:
   [
     {
       "id": 2,
       "name": "PokerRoom",
       "description": "Aroom for Poker",
       "status": "available"
     },
   ]
3. **[Get] Fetch rooms by ID (requestUrl:http://localhost:8080/api/rooms/{id})**:
   ```json
   *response:
   [
     {
       "id": 2,
       "name": "PokerRoom",
       "description": "Aroom for Poker",
       "status": "available"
     },
   ]

4. **[Put] Update room by ID (RequestUrl:http://localhost:8080/api/rooms/{id})**:
   ```json
   *requestBody:
   {
    "description": "A room for testing",
   },
   *response:
   {
    "id": 1,
    "name": "TestingRoom",
    "description": "A room for testing",
    "status": "available"
   },
   
5. **[Delete] Delete room by ID (RequestUrl:http://localhost:8080/api/room/{id})**:
   ```json
   *response:
   {"message: delete room successful"}
   
6. **[Get] Fetch all reservations (requestUrl:http://localhost:8080/api/reservations?room_id=&date=&limit=)**:
   ```json
   *response:
   [
     {
        "id": 1,
        "room_id": 1,
        "player_id": 1,
        "date": "2024-12-24T00:00:00Z",
        "time": "15:00",
        "player": {
            "id": 1,
            "name": "Samuel",
            "level": "等级一"
        },
     }
      ...
   ]

7. **[Post] Create a new reservation (RequestUrl:http://localhost:8080/api/reservations)**:
   ```json
   *requestBody:
   {
     "player_id": "int",
     "room_id": "int",
     "date": "string",
     "time": "string"
   }

   *response:
   [
     {
        "id": 1,
        "room_id": 1,
        "player_id": 1,
        "date": "2024-12-24T00:00:00Z",
        "time": "15:00",
        "player": {
            "id": 1,
            "name": "Samuel",
            "level": "等级一"
         },
     }
   ]

## Endless Challenge Endpoint

1. **[Get] Fetch all results (requestUrl:http://localhost:8080/api/challenges)**:
   ```json
   *response:
   [
     {
       "id": 1,
       "name": "TestingRoom",
       "description": "Aroom",
       "status": "available"
     },
      ...
   ]

2. **[Post] Join the challenge with fixed amount of $20.01 (RequestUrl:http://localhost:8080/api/challenges)**:
   ```json
   *requestBody:
   {
     "player_id": "int"
   }

   *response:
   [
     {
       "challenge_id": 3,
       "status": false,
       "win_amount": 0
     }
   ]

## Game Log Endpoint

1. **[Get] Fetch all logs by 'action'/'startTime'/'endTime' (requestUrl:http://localhost:8080/api/logs?action=&start_time=&end_time=)**:
   ```json
   *response:
   [
     {
        "id": 1,
        "player_id": 1,
        "action": "login",
        "timestamp": "2024-11-27T19:51:30.543507Z",
        "details": "Player Samuel login"
     },
      ...
   ]

2. **[Post] Create a log (RequestUrl:http://localhost:8080/api/logs)**:
   ```json
   *requestBody:
   {
     "player_id": "int",
     "action": "string",
     "details": "string"
   }

   *response:
   [
     {
       "id": 2,
       "player_id": 1,
       "action": "logout",
       "timestamp": "2024-11-28T14:45:11.052136Z",
       "details": "Player Samuel logout"
     }
   ]

## Payment Processing Endpoint

1. **[Get] Fetch all payments by id (requestUrl:http://localhost:8080/api/payments/{id})**:
   ```json
   *response:
   [
     {
       "id": 3,
       "player_id": 1,
       "method": "credit_card",
       "amount": 100.5,
       "details": {
           "card_number": "4111111111111111",
           "expiry_date": "12/25",
           "cvv": "123"
        },
       "transaction_id": "CC-6478508272011940677",
       "status": "success",
       "created_at": "2024-11-28T14:53:07.6153091Z"
    }
   ]

2. **[Post] Create a payment (RequestUrl:http://localhost:8080/api/payments)**:
   ```json
   *requestBody(credit card):
   {
    "player_id":int,
    "method": "string",
    "amount": float64,
    "details": {
        "card_number": "string",
        "expiry_date": "string",
        "cvv": "string"
     }
   }

   *requestBody(bank transfer):
   {
    "player_id":int,
    "method": "string",
    "amount": float64,
    "details": {
        "account_number": "string",
        "bank_name": "string"
     }
   }

   *requestBody(third party platform):
   {
    "player_id":int,
    "method": "string",
    "amount": float64,
    "details": {
        "platform": "string",
        "email": "string"
     }
   }

   *requestBody(blockchain):
   {
    "player_id":int,
    "method": "string",
    "amount": float64,
    "details": {
        "wallet_address": "string",
        "transaction_id": "string"
     }
   }

   *response(e.g credit card):
   [
     {
       "id": 3,
       "player_id": 1,
       "method": "credit_card",
       "amount": 100.5,
       "details": {
           "card_number": "4111111111111111",
           "expiry_date": "12/25",
           "cvv": "123"
        },
       "transaction_id": "CC-6478508272011940677",
       "status": "success",
       "created_at": "2024-11-28T14:53:07.6153091Z"
   }
   ]
