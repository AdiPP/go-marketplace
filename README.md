# Go Marketplace

Go Marketplace is a sample project demonstrating a marketplace application built with Go using Clean Architecture principles. It features event-driven processing using a memory-based queue adapter.

## Table of Contents

- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Project Structure

The project is organized into several packages, each responsible for different aspects of the application:

- `pkg/domain/event`: Defines the domain events used in the application.
- `pkg/domain/queue`: Defines the interfaces for the queue system.
- `pkg/infrastructure/queue`: Provides a memory-based queue adapter implementation.
- `pkg/interface/controller`: Contains the HTTP controllers for handling API requests.
- `pkg/interface/listener`: Contains the event listeners for handling events.
- `pkg/usecase`: Contains the use cases that implement the business logic.

## Prerequisites

Before running this project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.22.2 or higher)

## Installation

1. Clone the repository:

```sh
git clone https://github.com/AdiPP/go-marketplace.git
cd go-marketplace
```

2. Install the dependencies:

```sh
go mod tidy
```

## Usage

To run the application, use the following command:

```sh
go run main.go
```

The application will start an HTTP server listening on port 8080. You can create orders by sending a POST request to the `/create-order` endpoint.

## License

This project is licensed under the MIT License.