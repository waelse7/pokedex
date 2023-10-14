package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func MakeApiRequest(urlString, method string, params map[string]string, result interface{}) (interface{}, error){
  if len(params) > 0 {
    reqURL, err := url.Parse(urlString)
    if err != nil {
      return nil, err  
    }
    query := reqURL.Query()
    for key, val := range params {
      query.Add(key, val)
    }
    reqURL.RawQuery = query.Encode()
    urlString = reqURL.String()
  }

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
