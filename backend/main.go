package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Request struct {
	Start Point `json:"start"`
	End Point `json:"end"`
}

type Response struct {
	Path []Point `json:"path"`
}

func isValid(x, y int) bool {
	return x >= 0 && x < 20 && y >= 0 && y < 20
}

func findPath(start Point, end Point, visited [][]bool) []Point {
	if !isValid(start.X, start.Y) || !isValid(end.X, end.Y) {
		return nil
	}

	type Node struct {
		Point Point
		Path  []Point
	}

	queue := []Node{{Point: start, Path: []Point{start}}}
	visited[start.X][start.Y] = true

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Point.X == end.X && current.Point.Y == end.Y {
			return current.Path
		}

		for _, dir := range directions {
			nextX := current.Point.X + dir[0]
			nextY := current.Point.Y + dir[1]

			if isValid(nextX, nextY) && !visited[nextX][nextY] {
				visited[nextX][nextY] = true
				newPath := append([]Point{}, current.Path...)
				newPath = append(newPath, Point{nextX, nextY})
				queue = append(queue, Node{Point: Point{nextX, nextY}, Path: newPath})
			}
		}
	}

	return nil
}

func handlePath(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	visited := make([][]bool, 20)
	for i := range visited {
		visited[i] = make([]bool, 20)
	}

	path := findPath(req.Start, req.End, visited)
	json.NewEncoder(w).Encode(Response{Path: path})
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/find-path", enableCORS(http.HandlerFunc(handlePath)))
	
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}