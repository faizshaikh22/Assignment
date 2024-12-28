# Pathfinding Grid Application

This project is a simple pathfinding application that allows users to select start and end points on a grid and visualize the path between them.

## Prerequisites

- Node.js and npm (for the frontend)
- Go (for the backend)

## Setup Instructions

### Backend

1. **Navigate to the backend directory:**

   ```bash
   cd backend
   ```

2. **Run the Go server:**

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

### Frontend

1. **Navigate to the frontend directory:**

   ```bash
   cd frontend
   ```

2. **Install dependencies:**

   ```bash
   npm install
   ```

3. **Start the React application:**

   ```bash
   npm start
   ```

   The application will open in your default web browser at `http://localhost:3000`.

## Usage

- Click on a cell in the grid to select the start point.
- Click on another cell to select the end point.
- The path will be calculated and displayed in blue on the grid.
- Click again to reset the selection.

## Troubleshooting

- Ensure that both the backend and frontend servers are running.
- Check the console for any error messages if the path is not displayed.