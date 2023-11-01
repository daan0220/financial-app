package main

func main() {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
