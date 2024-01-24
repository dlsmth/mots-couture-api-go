package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mc.com/rest-api/models"
)

func getWords(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{"message": "Hello Go World!"})
	words, err := models.GetAllWords()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch words."})
		return
	}
	context.JSON(http.StatusOK, words)
}

func getWord(context *gin.Context) {
	wordId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse word id."})
		return
	}
	event, err := models.GetWordById(wordId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch word."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createWord(context *gin.Context) {
	var word models.Word
	err := context.ShouldBindJSON(&word)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	word.ID = 1
	err = word.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save word.", "word": word})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Word created", "word": word})
}

func updateWord(context *gin.Context) {
	wordId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse word id."})
		return
	}
	// userId := context.GetInt64("userId")
	word, err := models.GetWordById(wordId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse fetch word."})
		return
	}
	if word.Mot == "arrete" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update word."})
		return
	}
	var updatedWord models.Word
	err = context.ShouldBindJSON(&updatedWord)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse the word data."})
		return
	}

	updatedWord.ID = wordId
	err = updatedWord.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the word."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Word updated successfully"})
}

func deleteWord(context *gin.Context) {
	wordId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse word id."})
		return
	}
	// userId := context.GetInt64("userId")
	word, err := models.GetWordById(wordId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch word."})
		return
	}
	// if word.UserID != userId {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete word."})
	// 	return
	// }
	err = word.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete word."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Word deleted successfully."})
}
