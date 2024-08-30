package dto

type AnimalResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Class string `json:"class"`
    Legs  int    `json:"legs"`
}
