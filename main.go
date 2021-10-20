package main



//  original 
import (
//	"log"
	"os"
	"time"
//         "io/fs"
        "strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"strconv"
	"fmt"
//	"strings"
         "net/http"   //consume a rest api
         "io/ioutil"   //read the response body from the API
	
// modules for the github repository functionality and the in memory repository 
        billy "github.com/go-git/go-billy/v5"
        memfs "github.com/go-git/go-billy/v5/memfs"
        git "github.com/go-git/go-git/v5"
        httpgit "github.com/go-git/go-git/v5/plumbing/transport/http"
        memory "github.com/go-git/go-git/v5/storage/memory"	
        "github.com/go-git/go-git/v5/plumbing/object"
)


// variables for the the in memory repository  and the filesystem to handle internally the local git repo
var storer *memory.Storage
var fsmemory billy.Filesystem



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
	        c.IndentedJSON(500, resultados)
		return
	}
       if s2, err := strconv.ParseFloat(elemento2, 64); err == nil {
         fmt.Println(s2) // 3.14159265
	       s2final =s2
	}else{
         	resultados[0].Status ="NOK";
	        resultados[0].Resultado = "second parameter is expected numeric";
	        c.IndentedJSON(500, resultados)
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



//func addInGit(filenombre string )  bolean{
func addInGit() {
	
	fmt.Println("addInGit 8th  delcare in memory")
	
        storer = memory.NewStorage()
        fsmemory = memfs.New()
        fmt.Println("addInGit   set auth")
        // Authentication
        auth := &httpgit.BasicAuth{
//                Username: "youtochibots",
                Username: "izendejass600@gmail.com",
//                Password: "Impo",



        }

	fmt.Println("oct 5 addInGit   define github repository and login ")
        repository := "https://github.com/youtochibots/bot.git"
        r, err := git.Clone(storer, fsmemory, &git.CloneOptions{
                URL:  repository,
                Auth: auth,
        })
	
	fmt.Println("addInGit   login in ok")
	
        if err != nil {
                fmt.Printf("%v", err)
                return 
        }
        fmt.Println("Repository cloned")
        fmt.Println("Now display the contnets of the repository ")



	// Getting the latest commit on the current branch
        fmt.Println("Oct 19 -git log -1")

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()

	// ... retrieving the commit object
	commit01, err := r.CommitObject(ref.Hash())

	fmt.Println(commit01)

	// List the tree from HEAD
        fmt.Println(" Oct 19 - git ls-tree -r HEAD")

	// ... retrieve the tree from the commit
	tree, err := commit01.Tree()

	// ... get the files iterator and print the file
        var file001content string

	tree.Files().ForEach(func(f *object.File) error {
//		fmt.Printf("100644 blob %s    %s\n", f.Hash, f.Name)
                if(f.Name =="packages/leon/data/answers/en.json" ){
                    fmt.Println(" Oct 19 - file001 is set ")
                    file001content, err= f.Contents()
	          fmt.Printf(f.Contents())

                }
		return nil
	})

 

        // Create new file  packages/leon/data/answers/en.json
        filePath00 := "packages/leon/data/answers/en333.json"
        newFile00, err := fsmemory.Create(filePath00)
        if err != nil {
                return 
        }

       resultsFile001 := strings.Replace(file001content, "Tochi.", "SOY TOCHI Y QUE", -1)

        newFile00.Write( []byte(resultsFile001 )  ) 


	if err := newFile00.Close(); err != nil {
		return
	}


////////////////////////////////////////////////////////////////////////////////////////
/////GET room  content from API and use it for files answer, expressions , config and py
//////////////////////////////////////////////////////////////////////////////////////////

        fmt.Println("Geet the content of the card from the learning Room using a API")

         var apicontent = getNewContentFromApi("luna100-cards" )   //get the content for room files using an API

        fmt.Println("Repository cloned and now work with new files")

        w, err := r.Worktree()
        if err != nil {
                fmt.Printf("%v", err)
                return 
        }

	fmt.Println("addInGit   create new file")
        // Create new file
        filePath := "mipath/my-new-carlos005_2.txt"
        newFile, err := fsmemory.Create(filePath)
        if err != nil {
                return 
        }
        newFile.Write([]byte("My new file carlos005"))
        newFile.Close()

        // Create new file  packages/leon/data/answers/en.json
        filePath01 := "packages/leon/data/answers/en_json_2.txt"
        newFile01, err := fsmemory.Create(filePath01)
        if err != nil {
                return 
        }
        var answersContent = getNewAnswers(apicontent,"miembrosfamiliaveracruz","miembros_familia_veracruz"  ) //get the content for the added answers for the room

        newFile01.Write([]byte(     	answersContent  )) 
  
        newFile01.Close()

        // Create new file  packages/leon/data/answers/en.json
        filePath02 := "packages/leon/data/expressions/en_json_2.txt"
        newFile02, err := fsmemory.Create(filePath02)
        if err != nil {
                return 
        }

        var expressionsContent = getNewExpressions(apicontent ,"miembrosfamiliaveracruz" ) //get the content for the added expressions for the room

        newFile02.Write([]byte(		expressionsContent          )) 
  
        newFile02.Close()

        // Create new file  packages/leon/config/config.sample.json
        filePath03 := "packages/leon/config/config_sample_json_2.txt"
        newFile03, err := fsmemory.Create(filePath03)
        if err != nil {
                return 
        }

        var configSamplesContent = getNewConfigSamples(apicontent,"miembrosfamiliaveracruz"  ) //get the content for the config samples for the room
        newFile03.Write([]byte(configSamplesContent          )) 
  
        newFile03.Close()


        // Create new file  packages/leon/meaningoflife.py
        filePath04 := "packages/leon/meaningoflife_py_2.txt"
        newFile04, err := fsmemory.Create(filePath04)
        if err != nil {
                return 
        }

        var newPyContent = getNewPyFile(apicontent,"miembros_familia_veracruz"  )  //get the content for the new py file for the room
        newFile04.Write([]byte( newPyContent     )) 
  
        newFile04.Close()



        // Run git status before adding the file to the worktree
        fmt.Println(w.Status())

	fmt.Println("addInGit   git add the file")
        // git add $filePath
        w.Add(filePath)

	fmt.Println("addInGit   git add the file001")
        // git add $filePath
        w.Add(filePath01)

	fmt.Println("addInGit   git add the file002")
        // git add $filePath
        w.Add(filePath02)

	fmt.Println("addInGit   git add the file003")
        // git add $filePath
        w.Add(filePath03)

	fmt.Println("addInGit   git add the file004")
        // git add $filePath
        w.Add(filePath04)

	fmt.Println("addInGit   git add the file000 overwrite/edit")
        // git add $filePath
        w.Add(filePath00)

        // Run git status after the file has been added adding to the worktree
        fmt.Println(w.Status()) //displays A accoridng with doc this is add  https://pkg.go.dev/github.com/go-git/go-git/v5@v5.1.0#StatusCode

	fmt.Println("addInGit   git commit")
        // git commit -m $message
//        w.Commit("Added my new file", &git.CommitOptions{})

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Carlos Z",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)

	CheckIfError(err)

        fmt.Println(obj)
	
	fmt.Println("addInGit   git push")
        //Push the code to the remote
        err = r.Push(&git.PushOptions{
                RemoteName: "origin",
                Auth:       auth,
        })
        if err != nil {
	        fmt.Println("addInGit   git push: depsues- fallo")
                fmt.Printf("%v", err)		
                return 
        }
	fmt.Println("addInGit   git push:despues -ok")
        fmt.Println("Remote updated.", filePath)
        return
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

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func getNewContentFromApi(filename string )  string{
      var resultado string

    fmt.Println("1. Performing Http Get...")
//    resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    resp, err := http.Get("https://youtochibotas.herokuapp.com/v1/api/redisroomallcards/miembrosfamiliaveracruz-cards")



    if err != nil {
//        fmt.Printf(err)
         return "error http get"
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)

    // Convert response body to Todo struct
//    var todoStruct Todo
//    json.Unmarshal(bodyBytes, &todoStruct)
//    fmt.Printf("API Response as struct %+v\n", todoStruct)

     resultado = bodyString      
     return resultado

}

func getNewAnswers(apicontent ,roomname , roomnamespaces string )  string{
      var resultado string

      //get array of elements based on split the apicontent response


       splitcards := strings.Split(apicontent , "colour")  //the response contains list of cards with that starts with this string
       
      cuantascards:= len(splitcards)
      fmt.Printf("las cards son  #",cuantascards)
       answercards :=" "
      for _,card := range splitcards {
                fmt.Println(card)
                if len(card) > 0 {
                         indexsticky := strings.Index(card, "sticky")
                        if indexsticky >0 { //sticky exists in the string
                          //inicio \"text\":\"
                          var indexinicio= strings.Index(card, "\\\"text\\\":\\\"")
                          //fin    \",\"type\":
                          var indexfin = strings.Index(card, "\\\",\\\"type\\\":")
                          if indexinicio >0 && indexfin >0 && indexfin > indexinicio {
    			       inputFmt := card[indexinicio+11:indexfin]     //+11 as there are 11 charcaters in the  string \"text\":\"
                               fmt.Println(inputFmt )
                               answercards = answercards + "\\\""+inputFmt+".\\\",  "  
                               fmt.Println(answercards )
                          }

                        } //end if sticky
	         } //end if len card   
      }
       
      // set the resultado string
      resultado = 
//                "	\"miembrosfamiliaveracruz\": {  "+
                "	\""+ roomname+"\": {  "+
//		" \"miembros_familia_veracruz\": [  "+
		" \""+roomnamespaces+"\": [  "+
		"	\"LEo.\",  "+
		" \"Tommy.\",   "+
	        " \"Bodoque.\",    "+
         	" \"Tochi.\",   "+
         	" \"Weritin.\",   "+
         	" \"Flaco.\"   "+
	 	"	]   "+
	 	"  },     "+
         	"    \"partnerassistant\": {"
               ;
          
      return resultado
}

func getNewExpressions(apicontent string ,roomname  string )  string{
      var resultado string
      resultado =
//                " \"miembrosfamiliaveracruz\": {  "+
                " \""+roomname+"\": {  "+
		" \"run\": {   "+
		" \"expressions\": [  "+
		" \"Member of the familiy in Veracruz?\",   "+
		" \"Psrt of the Family in veracruz\",   "+
		" \"Is part of the Family in veracruz\",  "+
		" \"vive en depa Veracruz\",  "+
		" \"Vivessss en Veracruz\"  "+
		"  ]   "+
		" }  "+
		"},    "+
		" \"partnerassistant\": {  "
                ;
          
      return resultado
}

func getNewConfigSamples(apicontent ,roomname  string )  string{
      var resultado string
      resultado =
//		" \"miembrosfamiliaveracruz\": {  "+
		" \""+roomname +"\": {  "+
		 " \"options\": {}    "+
		 "  },   "+
		 " \"partnerassistant\": {  "  
                 ;
          
      return resultado
}

func getNewPyFile(apicontent , roomnamespaces string )  string{
      var resultado string

      resultado =" #!/usr/bin/env python  \n "+
		" # -*- coding:utf-8 -*-  \n "+
		" import utils            \n"+
		" def run(string, entities):    \n "+
		" 	\"\"\"Leon saysthe family "+roomnamespaces+"  \"\"\"       \n"+
//		" 	return utils.output('end', 'miembros_familia_veracruz', utils.translate('miembros_familia_veracruz'))    " 
		" 	return utils.output('end', '"+roomnamespaces+"', utils.translate('"+roomnamespaces+"'))    " 
                ;
          
      return resultado
}
