package main

import "gin_Chat/router"

func main() {
	r := router.Router()
	r.Run()
}
