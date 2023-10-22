package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeApiRequest(urlString, method string, result interface{}) (interface{}, error){

  client := getHTTPClient()

  req, err := http.NewRequest(method, urlString, nil)
  if err != nil {
    return nil, err
  }

  response, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer response.Body.Close()

  if response.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("request failed with status code: %v", response.StatusCode)
  }

  dat, err := io.ReadAll(response.Body)
  if err != nil {
    return nil, err
  }
 
  if err := json.Unmarshal(dat, &result); err != nil {
    return nil, err
  }

  return result, nil
}
