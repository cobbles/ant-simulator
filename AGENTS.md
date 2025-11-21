# AGENTS.md

This document provides guidance for AI agents working in the Ant Simulator repository.

## Project Overview

This is a Go-based web application that runs a simulation of ants foraging for food. The backend runs the simulation and exposes the world state via an API, which is then visualized by a frontend (`index.html`).

## Commands

Standard Go commands are used for this project.

-   **Run the application**: `air`
    This will start the web server on `http://localhost:8080`.
-   **Build the application**: `go build .`
-   **Run tests**: `go test ./...`

## Code Organization

The project is structured into several packages:

-   `main.go`: The application entry point. It initializes and starts the web server.
-   `server/`: Contains the HTTP server setup and API route definitions. It handles incoming requests and serves the simulation state.
-   `app/`: The core of the simulation. It manages the world state (ants, food, nest), contains the main simulation loop, and defines the rules for ant movement.
-   `ant/`: Defines the `Ant` data structure and its basic movement logic.
-   `food/`: Defines the `TFood` data structure.
-   `index.html`: The frontend HTML file that visualizes the simulation.

## Naming Conventions and Style

-   The code follows standard Go naming conventions (e.g., `PascalCase` for exported identifiers).
-   Structs intended for JSON serialization are prefixed with `T` in some cases (e.g., `TWorld`, `TNest`, `TFood`).

## Testing Approach

-   Tests are located in `_test.go` files within the same package as the code they are testing (e.g., `ant/ant_test.go`).
-   The standard Go `testing` package is used.
-   Tests should exercise the behaviours and not implementation details

## Development Approach

-   Write golang code using "effective go" conventions
-   Use the canonical TDD approach by Kent Beck to make code changes (Red, Green Refactor)

## Important Gotchas & Patterns

-   **Global World State**: The simulation state is stored in a single global variable `app.World`. This is a critical component and any modifications to it should be done with care, considering potential race conditions.
-   **Broken Tests**: The tests in `ant/ant_test.go` are currently broken. They attempt to call a `MoveAnt` function that is defined in the `app` package, which creates a dependency issue. This indicates that the tests have not been maintained after recent refactoring. Any work on the ant movement logic should include fixing these tests.
-   **Simulation Loop**: The main simulation logic runs in a separate goroutine started in `app/Start()`.
-   **Dependencies**: The project uses `go.mod` for dependency management.
