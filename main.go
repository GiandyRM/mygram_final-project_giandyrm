package main

import (
	"log"
	commentDelivery "mygram_final-project_giandyrm/comment/delivery"
	commentRepository "mygram_final-project_giandyrm/comment/repository"
	commentUseCase "mygram_final-project_giandyrm/comment/usecase"
	"mygram_final-project_giandyrm/database"
	photoDelivery "mygram_final-project_giandyrm/photo/delivery"
	photoRepository "mygram_final-project_giandyrm/photo/repository"
	photoUseCase "mygram_final-project_giandyrm/photo/usecase"
	socialMediaDelivery "mygram_final-project_giandyrm/socialmedia/delivery"
	socialMediaRepository "mygram_final-project_giandyrm/socialmedia/repository"
	socialMediaUseCase "mygram_final-project_giandyrm/socialmedia/usecase"
	userDelivery "mygram_final-project_giandyrm/user/delivery"
	userRepository "mygram_final-project_giandyrm/user/repository"
	userUseCase "mygram_final-project_giandyrm/user/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := database.StartDB()
	routers := gin.Default()
	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})
	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	userDelivery.NewUserHandler(routers, userUseCase)

	photoRepository := photoRepository.NewPhotoRepository(db)
	photoUseCase := photoUseCase.NewPhotoUsecase(photoRepository)
	photoDelivery.NewPhotoHandler(routers, photoUseCase)

	commentRepository := commentRepository.NewCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentUseCase(commentRepository)
	commentDelivery.NewCommentHandler(routers, commentUseCase, photoUseCase)

	socialMediaRepository := socialMediaRepository.NewSocialMediaRepository(db)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(socialMediaRepository)
	socialMediaDelivery.NewSocialMediaHandler(routers, socialMediaUseCase)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	port := os.Getenv("PORT")
	if len(os.Args) > 1 {
		requestPort := os.Args[1]
		if requestPort != "" {
			port = requestPort
		}
	}
	if port == "" {
		port = "6969"
	}
	routers.Run(":" + port)
}
