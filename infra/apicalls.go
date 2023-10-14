package infra


func GetLocation(url *string) *LocationResponse {
  BASE_URL := "https://pokeapi.co/api/v2"
  fullURL := BASE_URL + "/location/"


  if url != nil {
    fullURL = *url
  } 
  
  param := map[string]string{}
  var locationResponse LocationResponse
  jsonBody, err := MakeApiRequest(fullURL, "GET", param, &locationResponse)
  if err != nil {
    return nil
  }
  return jsonBody.(*LocationResponse)
}
