package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Translate(c *gin.Context) {
	reqUrl := "https://api-free.deepl.com/v2/translate"
	// contentType := "application/x-www-form-urlencoded"
	data := url.Values{
		"auth_key":    {""},
		"text":        {"Hello"},
		"source_lang": {"EN"},
		"target_lang": {"JA"},
	}
	res, err := http.PostForm(reqUrl, data)
	if err != nil {
		log.Fatal("DeepL API request failed: ", err)
	}
	var resp map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resp)
	log.Print("-----", resp)
	c.JSON(http.StatusOK, gin.H{"hoge": "Hello"})
}
