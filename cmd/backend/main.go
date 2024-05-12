package main

import (
	connection "github.com/Sreeram0045/backend/internal/db"
	ServerOpen "github.com/Sreeram0045/backend/internal/server"
)

func main() {
	connection.ConnectToDB()
	ServerOpen.ServerStart()
}
