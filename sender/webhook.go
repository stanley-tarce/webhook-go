package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)  

 func SendWebhook(data interface{}, url string, webhookId string) error {
	jsonBytes, err := json.Marshal(data)
	if (err != nil) {
		return err
	}

	// Prepare Request 
	res, err := http.Post(url, "application/json", bytes.NewBuffer((jsonBytes)))
	if err != nil {  
	   return err  
	}

	defer func(b io.ReadCloser) {  
		err := b.Close()
		if err != nil {  
		   log.Println("Error closing response body:", err)  
		}  
	 }(res.Body)  

	 // Default value
	 status := "failed"  
	 if res.StatusCode == http.StatusOK {  
		status = "delivered"  
	 }  
  
	 log.Println(status)  
  
	 if status == "failed" {  
		return errors.New(status)  
	 }  
  
	 return nil  
 }