package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"unicode/utf8"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type GenerateCasesRequestBody struct {
	Factors string `json:"factors"`
}

func SetupRouter() *gin.Engine {
	// set gin as release mode
	if os.Getenv("NODE_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS settings
	// https://github.com/gin-contrib/cors#using-defaultconfig-as-start-point
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"*",
	}
	if os.Getenv("NODE_ENV") != "development" {
		config.AllowOrigins = []string{
			"https://pairwise.yuuniworks.com",
		}
	}
	router.Use(cors.New(config))

	// enable gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	router.POST("/generate_cases", func(c *gin.Context) {
		// extract request body
		var generateCasesRequestBody GenerateCasesRequestBody
		c.BindJSON(&generateCasesRequestBody)

		// log body
		bodyBytes, err := json.Marshal(generateCasesRequestBody)
		if err != nil {
			log.Print(err)
			c.Status(500)
			return
		}

		// limit multi-byte chars
		if len(bodyBytes) != utf8.RuneCount(bodyBytes) {
			c.JSON(400, "Multi-byte characters cannot be used for now.")
			return
		}

		// limit size of factors
		if len(bodyBytes) > 5000 {
			c.JSON(400, "Test factors are too large. Maximum size is limited to 5KB.")
			return
		}

		// log factors
		bodyString := string(bodyBytes)
		log.Print(bodyString)

		// create temporary test-factors file
		factors := []byte(generateCasesRequestBody.Factors)
		tmpfile, err := ioutil.TempFile("", "temp-test-factors-")
		if err != nil {
			log.Print(err)
			c.Status(500)
			return
		}
		defer os.Remove(tmpfile.Name()) // clean up
		if _, err := tmpfile.Write(factors); err != nil {
			log.Print(err)
			c.Status(500)
			return
		}
		if err := tmpfile.Close(); err != nil {
			log.Print(err)
			c.Status(500)
			return
		}

		// exec `pict` command
		pictCmd := exec.Command("pict", tmpfile.Name())
		pictOut, err := pictCmd.CombinedOutput() // get both stdin & stdout
		if err != nil {
			c.JSON(400, string(pictOut))
			return
		}

		c.JSON(200, string(pictOut))
	})

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
