package api

var BASE_URL = "https://pokeapi.co/api/v2"

func GetLocationList(url *string) (*LocationListResponse, error) {
  fullURL := BASE_URL + "/location-area/"

  if url != nil {
    fullURL = *url
  } 
  
  var locationListResponse LocationListResponse
  jsonBody, err := MakeApiRequest(fullURL, "GET", &locationListResponse)
  if err != nil {
    return nil, err
  }
  return jsonBody.(*LocationListResponse), nil
}

func GetLocationArea(areaName string) (*LocationAreaResponse, error) {
  fullURL := BASE_URL + "/location-area/" + areaName
 
  var locationAreaResponse LocationAreaResponse
  jsonBody, err := MakeApiRequest(fullURL, "GET", &locationAreaResponse)
  if err != nil {
    return nil, err
  }
  return jsonBody.(*LocationAreaResponse), nil
}

func GetPokemon(nameOrId string) (*Pokemon, error) {
  fullURL := BASE_URL + "/pokemon/" + nameOrId
 
  var pokemon Pokemon
  jsonBody, err := MakeApiRequest(fullURL, "GET", &pokemon)
  if err != nil {
    return nil, err
  }
  return jsonBody.(*Pokemon), nil
}