package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// API Call for GET or POST method and get Json Marshal Byte data
func APICall(callMethod string, apiUrl string, requestData any, responseData any, c *gin.Context) []byte {
	byteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
	}
	// Create a new request using http
	req, err := http.NewRequest(callMethod, apiUrl, bytes.NewBuffer(byteData))
	if err != nil {
		fmt.Println(err)
	}

	session := sessions.Default(c)
	if session.Get("login_session") != nil {
		token := session.Get("login_session").(string)
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("X-Client-Domain", c.Request.Host)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		c.JSON(500, "Connection refused, Back to previous page and try again.")
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&responseData)
	// return resp
	data, _ := json.Marshal(responseData)
	// fmt.Println(data)
	return data

}

// APICallHRM_V1 API Call for GET or POST method and get Json Marshal Byte data
func APICallHRM_V1(callMethod string, requestData any, responseData any, c *gin.Context) []byte {

	structType := reflect.TypeOf(requestData)

	// Get the name of the struct type
	structName := structType.Name()
	if len(structName) == 0 {
		structType = reflect.TypeOf(responseData)
		structName = structType.Name()
	}
	apiUrl := app.API_URL + "/v1/hrm/" + helpers.Pluralize(structName)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id > 0 {
		apiUrl = apiUrl + "/" + fmt.Sprint(id)
	}

	byteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
	}
	// Create a new request using http
	req, err := http.NewRequest(callMethod, apiUrl, bytes.NewBuffer(byteData))
	if err != nil {
		fmt.Println(err)
	}

	session := sessions.Default(c)
	if session.Get("login_session") != nil {
		token := session.Get("login_session").(string)
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("X-Client-Domain", c.Request.Host)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		c.JSON(500, "Connection refused, Back to previous page and try again.")
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&responseData)
	// return resp
	data, _ := json.Marshal(responseData)
	// fmt.Println(data)
	return data

}
