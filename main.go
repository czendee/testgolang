package main



//  original 
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"strconv"
	"fmt"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/inicio", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	
	router.GET("/v1/albums", getAlbums)
	
	router.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)
	
	router.Run(":" + port)


}
//termin ortiginall

 

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}


// getAlbums responds with the list of all albums as JSON.
func getMultiplicaByID(c *gin.Context) {
	 elemento1 := c.Param("numero1")
	 elemento2 := c.Param("numero2")
	s1 := 0;
	s2 := 0;
	if s1, err := strconv.ParseFloat(elemento1, 32); err == nil {
             fmt.Println(s1) // 3.1415927410125732
       }
       if s2, err := strconv.ParseFloat(elemento2, 32); err == nil {
         fmt.Println(s2) // 3.14159265
      }
	
	resultado := s1* s2;
	 fmt.Println(resultado) 
        c.IndentedJSON(http.StatusOK, albums)
}
