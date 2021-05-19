package web_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryan-rozario/sudoku-solver/solver"
	"github.com/ryan-rozario/sudoku-solver/sudoku"
)

func solveSudokuApi(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var grid [][]int
	json.Unmarshal(reqBody, &grid)
	board_obj := sudoku.NewBoard(grid)
	solver.FillGrid(board_obj)

	json.NewEncoder(w).Encode(board_obj)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func HandleRequests() {
	// creates a new instance of a mux router
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/api/solve-sudoku", homePage).Methods("GET")
	apiRouter.HandleFunc("/api/solve-sudoku", solveSudokuApi).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", apiRouter))
}
