package controllers

import (
	entities "api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type tweetController struct {
	tweets []entities.Tweet
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := validate.Struct(tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.tweets = append(t.tweets, *tweet)
	ctx.JSON(http.StatusOK, gin.H{"message": "Tweet create successfully", "tweet": tweet})
}

func (t *tweetController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, tweet := range t.tweets {
		if tweet.ID == id {
			newTweet := entities.NewTweet()
			if err := ctx.BindJSON(&newTweet); err != nil {
				return
			}
			t.tweets[i] = *newTweet

			ctx.JSON(http.StatusOK, gin.H{"message": "Tweet updated successfully", "tweet": newTweet})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[:i], t.tweets[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
}
