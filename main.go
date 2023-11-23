package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dreamkanathip/cruiseship/entity"
)


func main() {

entity.SetupDatabase()

r := gin.Default()

r.Run()

}