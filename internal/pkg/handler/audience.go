package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"
)


var audiences = []string{"306э", "220л", "515ю", "520"}


func (h *Handler) issueAudience(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueAudience:", input)

	if input.AccessToken != 123 {
	 c.Status(http.StatusForbidden)
	 return
	}

	c.Status(http.StatusOK)


	go func() {
		time.Sleep(3 * time.Second)
		sendAudienceRequest(input)
	}()
}

func sendAudienceRequest(request models.Request) {

	var pickedAudience string
	if rand.Intn(10) % 10 >= 2 {
	 pickedAudience = audiences[rand.Intn(len(audiences))]
	}

	answer := models.AudienceRequest{
		AccessToken: request.AccessToken,
		Audience:    pickedAudience,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/lessons/%d/update_audience/", request.LessonId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
