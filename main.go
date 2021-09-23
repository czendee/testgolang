package main



//  original 
import (
//	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"strconv"
	"fmt"
// modules for the github repository functionality and the in memory repository 
        billy "github.com/go-git/go-billy/v5"
        memfs "github.com/go-git/go-billy/v5/memfs"
        git "github.com/go-git/go-git/v5"
        httpgit "github.com/go-git/go-git/v5/plumbing/transport/http"
        memory "github.com/go-git/go-git/v5/storage/memory"

)

// variables for the the in memory repository  and the filesystem to handle internally the local git repo
var storer *memory.Storage
var fs billy.Filesystem

func main() {
	port := os.Getenv("PORT")


	if port == "" {
//		log.Fatal("$PORT must be set")
                port ="8090"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/inicio", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	
	
	router.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)

	router.GET("v2/addFileGit/:nombrearchivo/:numero2", getAddFileGit)
	
	router.Run(":" + port)


}
//termin ortiginall

 



// multiplica represents data about multiplicacion.
type multiplica struct {
    Status     string  `json:"status"`
    Resultado  string  `json:"resultado"`

}

var resultados = []multiplica{
    {Status: "ok", Resultado: "Blue Train"},

}


// getMultiplicaByID responds with the stauts and the result as JSON.
func getMultiplicaByID(c *gin.Context) {
	 elemento1 := c.Param("numero1")
	 elemento2 := c.Param("numero2")
	var s1final float64 = 0
	var s2final float64 = 0
	
	 
	if s1, err := strconv.ParseFloat(elemento1, 64); err == nil {
             fmt.Println(s1) // 3.1415927410125732
		s1final =s1;
	}else{
		resultados[0].Status ="NOK";
	        resultados[0].Resultado = "first parameter is expected numeric";
	        c.IndentedJSON(http.StatusOK, resultados)
		return
	}
       if s2, err := strconv.ParseFloat(elemento2, 64); err == nil {
         fmt.Println(s2) // 3.14159265
	       s2final =s2
	}else{
         	resultados[0].Status ="NOK";
	        resultados[0].Resultado = "second parameter is expected numeric";
	        c.IndentedJSON(http.StatusOK, resultados)
		return
	}
	
	resultado := s1final* s2final;
	 fmt.Println(resultado) 
	sresultado := fmt.Sprintf("%f", resultado)
	
	resultados[0].Status ="OK";
	resultados[0].Resultado = sresultado
	d:=resultados[0].Resultado 
	fmt.Println(d+"si")

//	response:= json.NewEncoder(w).Encode(map[string]string{"status": "OK"})	
	
        c.IndentedJSON(http.StatusOK, resultados)
}


// getAddFileGit  agrega un archivo al github repository ,responds with the stauts and the result as JSON.
func getAddFileGit(c *gin.Context) {
	 elemento1 := c.Param("nombrearchivo")
	 elemento2 := c.Param("numero2")
	var s2final float64 = 0
	
	 
       if s2, err := strconv.ParseFloat(elemento2, 64); err == nil {
         fmt.Println(s2) // 3.14159265
	       s2final =s2
	}else{
         	resultados[0].Status ="NOK";
	        resultados[0].Resultado = "second parameter is expected numeric";
	        c.IndentedJSON(http.StatusOK, resultados)
		return
	}
//logic

        addInGit()


//prapre 	
	resultado := s2final;
	 fmt.Println(resultado) 
	sresultado := fmt.Sprintf("%f", resultado)
	
	resultados[0].Status ="OK"+elemento1;
	resultados[0].Resultado = sresultado
	d:=resultados[0].Resultado 
	fmt.Println(d+"si")

	
        c.IndentedJSON(http.StatusOK, resultados)
}

//func addInGit(filenombre string )  bolean{
func addInGit() {
	
	fmt.Println("addInGit   delcare in memory")
	
        storer = memory.NewStorage()
        fs = memfs.New()
        fmt.Println("addInGit   set auth")
        // Authentication
        auth := &httpgit.BasicAuth{
                Username: "youtochibots",
                Password: "Impos",
        }

	fmt.Println("addInGit   define github repository and login ")
        repository := "https://github.com/youtochibots/bot.git"
        r, err := git.Clone(storer, fs, &git.CloneOptions{
                URL:  repository,
//                Auth: auth,
        })
	
	fmt.Println("addInGit   login in ok")
	
        if err != nil {
                fmt.Printf("%v", err)
                return 
        }
        fmt.Println("Repository cloned")

        w, err := r.Worktree()
        if err != nil {
                fmt.Printf("%v", err)
                return 
        }

	fmt.Println("addInGit   create new file")
        // Create new file
        filePath := "my-new-ififif.txt"
        newFile, err := fs.Create(filePath)
        if err != nil {
                return 
        }
        newFile.Write([]byte("My new file carlos"))
        newFile.Close()

        // Run git status before adding the file to the worktree
        fmt.Println(w.Status())

	fmt.Println("addInGit   git add the file")
        // git add $filePath
        w.Add(filePath)

        // Run git status after the file has been added adding to the worktree
        fmt.Println(w.Status())

	fmt.Println("addInGit   git commit")
        // git commit -m $message
        w.Commit("Added my new file", &git.CommitOptions{})

	
	fmt.Println("addInGit   git push")
        //Push the code to the remote
        err = r.Push(&git.PushOptions{
                RemoteName: "origin",
                Auth:       auth,
        })
        if err != nil {
                return 
        }
        fmt.Println("Remote updated.", filePath)
        return
}
