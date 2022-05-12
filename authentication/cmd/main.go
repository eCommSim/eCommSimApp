package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/eCommSim/authentication/database"
	"github.com/eCommSim/authentication/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	opts := os.Args
	if len(opts) < 2 || opts[1] == "internal" {
		// app is running in cluster
		database.DatabasePort_g = 5432
	} else {
		// if we are running this outside the kubernetes environment
		out, err := exec.Command("kubectl", "get", "svc", "pgbouncer", "-o", "jsonpath=\"{.spec.ports[0].nodePort}\"").Output()
		if err != nil {
			log.Fatal("failed to find database connection service port:", err)
		}
		svcPort, err := strconv.ParseInt(strings.Split(string(out), "\"")[1], 10, 32)
		if err != nil {
			log.Fatal("failed to parse database connection service port:", err)
		}
		database.DatabasePort_g = svcPort
	}
	database.CreateUserTable()
	route := gin.Default()
	route.GET("/signup", routes.Signup)
	route.GET("/getuser", routes.Getuser)
	route.GET("/login", routes.Signin)

	route.Run(":8080")
}
