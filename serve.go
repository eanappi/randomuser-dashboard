package main

func run() {
	err := http.ListenAndServe(localhost:8080, nil)
	if err != nil {
		log.Fatal("Error initializing server:", err)
		return
	}
}

