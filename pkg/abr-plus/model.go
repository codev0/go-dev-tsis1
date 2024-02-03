package abrplus

import (
	"errors"
)

type Restaurant struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

var restaurants = []Restaurant{
	{
		Id:      "1",
		Title:   "The Cinnamon Club",
		Address: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		Type:    "Indian",
	},
	{
		Id:      "2",
		Title:   "Nobu",
		Address: "19 Old Park Ln, Mayfair, London W1K 1LB",
		Type:    "Japanese",
	},
	{
		Id:      "3",
		Title:   "Gordon Ramsay",
		Address: "68 Royal Hospital Rd, Chelsea, London SW3 4HP",
		Type:    "French",
	},
	{
		Id:      "4",
		Title:   "The Ledbury",
		Address: "127 Ledbury Rd, Notting Hill, London W11 2AQ",
		Type:    "Modern European",
	},
	{
		Id:      "5",
		Title:   "Wahaca",
		Address: "66 Chandos Pl, Covent Garden, London WC2N 4HG",
		Type:    "Mexican",
	},
}

func GetRestaurants() []Restaurant {
	return restaurants
}

func GetRestaurant(id string) (*Restaurant, error) {
	for _, r := range restaurants {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Restaurant not found")
}
