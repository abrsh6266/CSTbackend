package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gilab.com/pragmaticreviews/golang-gin-poc/initializer"
	"gilab.com/pragmaticreviews/golang-gin-poc/model"
	"gilab.com/pragmaticreviews/golang-gin-poc/utils"
	"github.com/gin-gonic/gin"
)

	func Login(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		respBody, err2 := initializer.HasuraRequest(http.MethodPost, string(body))
		if err2 != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server Error"})
			return
		} else if string(respBody) == "{\"data\":{\"User\": []}}" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email"})
			return
		}
		var result struct {
			Data struct {
				Users []model.User `json:"User"`
			} `json:"data"`
		}
		if err := json.Unmarshal(respBody, &result); err != nil {
			fmt.Println("failed to Parse the data")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
			return
		}
		if result.Data.Users[0].Email==""{
			fmt.Println("Error!!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error"})
			return
		}
		token, err := utils.GenerateToken(result.Data.Users[0].Email,result.Data.Users[0].Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		expiration := time.Now().Add(24 * time.Hour)
        cookie := &http.Cookie{
            Name:     "jwt",
            Value:    token,
            Expires:  expiration,
            HttpOnly: false,
        }
        http.SetCookie(c.Writer, cookie)
		result.Data.Users[0].Token = token
		c.JSON(http.StatusOK,result)
	}
func GetProfile(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		respBody, err2 := initializer.HasuraRequest(http.MethodPost, string(body))
		if err2 != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server Error"})
			return
		} else if string(respBody) == "{\"data\":{\"User\": []}}" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email"})
			return
		}
		var result struct {
			Data struct {
				Users []model.User `json:"User"`
			} `json:"data"`
		}
		if err := json.Unmarshal(respBody, &result); err != nil {
			fmt.Println("failed to Parse the data")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
			return
		}
		c.JSON(http.StatusOK,result)
	}
    func Signup(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	val, err := initializer.HasuraMutationRequest(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else if (string(val))[0:5] == "{\"err" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": string(val)})
		return
	}
	var result struct {
		Data struct {
			InsertUser struct {
				Returning []struct {
					Email string  `json:"email"`
				}`json:"returning"`
			}`json:"insert_User"`
		} `json:"data"`
	}
	if err := json.Unmarshal(val, &result); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}
	c.JSON(http.StatusOK, result)
}
func Validate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "I am logged in",
	})
}
