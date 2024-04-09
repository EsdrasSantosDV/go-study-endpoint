package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type country struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Region  string `json:"region"`
}

var countries = []country{
	{ID: "1", Name: "India", Capital: "New Delhi", Region: "Asia"},
	{ID: "2", Name: "USA", Capital: "Washington D.C.", Region: "North America"},
	{ID: "3", Name: "UK", Capital: "London", Region: "Europe"},
	{ID: "4", Name: "Australia", Capital: "Canberra", Region: "Australia"},
	{ID: "5", Name: "Brazil", Capital: "Brasília", Region: "South America"},
	{ID: "6", Name: "South Africa", Capital: "Pretoria", Region: "Africa"},
	{ID: "7", Name: "Russia", Capital: "Moscow", Region: "Europe"},
	{ID: "8", Name: "China", Capital: "Beijing", Region: "Asia"},
}

func main() {
	router := gin.Default()
	router.GET("/countries", getCountries)
	router.Run("localhost:8080")
}

func getCountries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, countries)
}
