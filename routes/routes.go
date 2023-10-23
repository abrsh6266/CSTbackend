package routes

import (
	"fmt"
	"log"
	controller "gilab.com/pragmaticreviews/golang-gin-poc/Controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)
type RegistrationPayload struct {
    Event struct {
        Op    string `json:"op"`
        Data  User   `json:"data"`
    } `json:"event"`
}

type User struct {
    Old map[string]interface{} `json:"old"`
    New map[string]interface{} `json:"new"`
}
// SetupRouter configures the application's routes
func SetupRouter() *gin.Engine {
    r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Set AllowAllOrigins to true
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
    
	// Public routes
    r.POST("/register", controller.Signup)
    r.POST("/login", controller.Login)
	r.POST("/movie/get",controller.GetMovie)
	r.POST("/profile/get",controller.GetProfile)
	r.POST("/star/get",controller.CastControl)
	r.POST("/director/get",controller.DirectorControl)
	r.POST("/movie/add",controller.AddMovie)
	r.POST("/sendemail", func(c *gin.Context) {
		var payload RegistrationPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		username := payload.Event.Data.New["username"].(string)
		userEmail := payload.Event.Data.New["email"].(string)

		m := gomail.NewMessage()
		m.SetHeader("From", "abrsh6266@gmail.com")
		m.SetHeader("To", userEmail)
		m.SetHeader("Subject", "Habesh Cinema")
		m.SetBody("text/html", fmt.Sprintf("<strong>Hello %s,</strong><br>You have successfully registered. log into our website and explore movies and buy tickets! <a href='http://localhost:3000/login'>Login</a>", username))

		d := gomail.NewDialer("smtp.gmail.com", 587, "abrsh6266@gmail.com", "xgwx esod euwh njuw")

		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(200, gin.H{"message": "Email sent successfully"})
	})
	r.POST("/sendemailmessage", func(c *gin.Context) {
		var payload RegistrationPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		userEmail := payload.Event.Data.New["userEmail"].(string)

		m := gomail.NewMessage()
		m.SetHeader("From", "abrsh6266@gmail.com")
		m.SetHeader("To", userEmail)
		m.SetHeader("Subject", "Habesh Cinema")
		m.SetBody("text/html", fmt.Sprintf("<strong>Hello Dear,</strong><br>You have Recieved new Message. log into our website and take a look the message! <a href='http://localhost:3000/login'>Login</a>"))

		d := gomail.NewDialer("smtp.gmail.com", 587, "abrsh6266@gmail.com", "xgwx esod euwh njuw")
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(200, gin.H{"message": "Email sent successfully"})
	})
    return r
}