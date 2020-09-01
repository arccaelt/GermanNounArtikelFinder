package routes

import (
	"dao"
	"encoding/json"
	"io/ioutil"
	"model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const QUERY_PATTERN string = "noun = ?"
const BAD_REQUEST_ERROR_MESSAGE string = "The request must contain the word for which the article is searched and only one word is allowed"
const NOUN_NOT_FOUND_MESSAGE string = "Unknown noun"

func handleWordArticle(context *gin.Context) {
	requetBodyJson := getRequestBodyAsJSON(context)

	if !isValidGetWordArticleRequest(requetBodyJson) {
		context.JSON(http.StatusBadRequest, BAD_REQUEST_ERROR_MESSAGE)
		return
	}

	queryWord := cleanQueryWord(requetBodyJson.RawText)
	var foundWord dao.Word
	dao.DBORM.Where(QUERY_PATTERN, queryWord).First(&foundWord)

	if foundWord.Noun == "" {
		context.JSON(http.StatusNotFound, NOUN_NOT_FOUND_MESSAGE)
		return
	}

	mappedDAOWord := mapDAOToResponseStruct(&foundWord)
	context.JSON(http.StatusOK, mappedDAOWord)
}

func mapDAOToResponseStruct(foundWord *dao.Word) model.Word {
	var responseWord model.Word
	responseWord.RawText = (*foundWord).Noun
	responseWord.Article = (*foundWord).Gender
	return responseWord
}

func cleanQueryWord(queryWord string) string {
	return strings.ToLower(queryWord)
}

func getRequestBodyAsJSON(context *gin.Context) model.Word {
	requestBody, err := ioutil.ReadAll(context.Request.Body)

	if err != nil {
		panic("Can't read body")
	}

	var requestBodyJson model.Word
	json.Unmarshal(requestBody, &requestBodyJson)
	return requestBodyJson
}

func isValidGetWordArticleRequest(requestBodyJson model.Word) bool {
	queryWord := requestBodyJson.RawText
	if queryWord == "" {
		return false
	}

	if strings.Contains(queryWord, " ") {
		return false
	}

	return true
}