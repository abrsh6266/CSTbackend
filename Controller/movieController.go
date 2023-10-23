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

func GetMovie(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	respBody, err := initializer.HasuraRequest(http.MethodPost, string(body))
	if err != nil {
		fmt.Println(err)	
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "some sort of error"})
		return
	}
	var result struct {
		Data struct {
			Movies []model.Movie `json:"Movie"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	fmt.Println(result)
	ctx.JSON(http.StatusOK,result)
}
func AddMovie(c *gin.Context) {
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
			OneMovie model.Movie `json:"insert_Movie_one"`
			OneSchedule model.Schedule `json:"insert_Schedule_one"`
			OneCast model.Cast `json:"insert_Cast_one"`
			OneMessage model.Message `json:"insert_Message_one"`
			OneDirectorMovieRel model.DirectorMovieRel `json:"insert_DirectorMovieRel_one"`
			Movies []model.Movie `json:"Movie"`
			Users []model.User `json:"User"`
			Schedules []model.Schedule `json:"Schedule"`
			Stars []model.Star `json:"Star"`
			Messages []model.Message `json:"Message"`
			Directors []model.Director `json:"Director"`
			OneRating model.Rating `json:"insert_Rating_one"`
			Bookmarks []model.Bookmark `json:"Bookmark"`
			Delete_Bookmark_by_pk model.Movie `json:"delete_Bookmark_by_pk"`
			Delete_Movie_by_pk model.Movie `json:"delete_Movie_by_pk"`
			Delete_Schedule_by_pk model.Movie `json:"delete_Schedule_by_pk"`
			Update_Movie_by_pk model.Movie `json:"update_Movie_by_pk"`
			Update_Star_by_pk model.Movie `json:"update_Star_by_pk"`
			Update_Director_by_pk model.Movie `json:"update_Director_by_pk"`
			Update_User_by_pk model.User `json:"update_User_by_pk"`
			Update_Message_by_pk model.Message `json:"update_Message_by_pk"`
			OneBookmark model.Bookmark `json:"insert_Bookmark_one"`
			User_aggregate struct{
				Aggregate struct{
					Count int `json:"count"`
				} `json:"aggregate"`
			} `json:"User_aggregate"`
			Rate_aggregate struct{
				Aggregate struct{
					Count int `json:"count"`
				} `json:"aggregate"`
			} `json:"Rating_aggregate"`
			Bookmark_aggregate struct{
				Aggregate struct{
					Count int `json:"count"`
				} `json:"aggregate"`
			} `json:"Bookmark_aggregate"`
			Movie_aggregate struct{
				Aggregate struct{
					Count int `json:"count"`
				} `json:"aggregate"`
			} `json:"Movie_aggregate"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse category data"})
		return
		}
	c.JSON(http.StatusOK,result)
}
func RemoveMovie(c *gin.Context) {
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
			DeletedMaterial struct{
				AffectedRows int  `json:"affected_rows"`
			} `json:"delete_Department"`
		} `json:"data"`
	}
	if json.Unmarshal(respBody,&result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
		return
		}
	c.JSON(http.StatusOK,result)
}