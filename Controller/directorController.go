package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"github.com/gin-gonic/gin"
)
func DirectorControl(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}else if string(respBody)[:8]=="{\"errors"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong!"})
		return
	}
	var result struct {
		Data struct {
			OneDirector model.Director `json:"insert_Director_one"`
			Directors []model.Director `json:"Director"`
			Delete_Director_by_pk model.Director `json:"delete_Director_by_pk"`
			Update_Director_by_pk model.Director `json:"update_Director_by_pk"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
		}
		fmt.Println(result)
	c.JSON(http.StatusOK,result)
}