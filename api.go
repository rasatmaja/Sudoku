package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//untuk mengirimkan data menggunakan path '/' dan menggunakan method POST
	router.HandleFunc("/", sudoku).Methods("POST")
	return router
}

func handleRequest() {
	//membuart dan menjalankan service API pada port 8232
	log.Fatal(http.ListenAndServe(":8232", routes()))
}

func sudoku(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	//menentukan Header request
	httpResponse.Header().Set("Content-Type", "application/json")

	//digunakan untuk mengambil semua data yang ada di dalam request
	data, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		panic(err)
	}

	//digunakan untuk mem-parsing data dari string ke array 2D.
	var inputBoard [][]int
	if err := json.Unmarshal([]byte(data), &inputBoard); err != nil {
		panic(err)
	}

	//mencetak inputan sudoku kedalam console/terminal
	println("input sudoku:")
	printToConsole(&inputBoard)

	//memindahkan inputan sudoku kedalam variable baru
	board := inputBoard

	/*
	 * memanggil method untuk memastikan inputan sudoku sudah sesuaiS
	 *
	 */
	if isPosibleToSolve(&board) {
		/*
		 * memanggil method untuk menyelesaikan sudoku
		 * dan di passing-kan argumen berupan variable yang menyimpan inputan sudoku
		 * dalam kasus ini akan di passing-kan berupa alamat memory-nya (passing by reference)
		 *
		 */
		solve(&board)

		//mencetak hasil board sukodu yang sudah terselesaikan (solved)
		println(" ")
		println("ouput sudoku:")
		printToConsole(&board)

		/*
		 * membuat berupa response data sudoku yang
		 * sudah selesai ke dalam bentuk JSON
		 *
		 */
		response := board
		json.NewEncoder(httpResponse).Encode(response)
	} else {
		/*
		 * membuat berupa response input invalid
		 * ke dalam bentuk JSON
		 *
		 */
		response := "Input invalid. Please check your input"
		json.NewEncoder(httpResponse).Encode(response)

		println(" ")
		println(response)
	}

}
