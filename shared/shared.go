package shared

import (
	"fmt"
	"time"
)

type AvitoPage struct {
	Hint     string
	Url      string
	Title    string
	Price    string
	PriceInt string
	Image    string
	Date     time.Time
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
}

func (page AvitoPage) Dump() {
	fmt.Println(page.Title)
	fmt.Println(page.Price)
	fmt.Println(page.PriceInt)
	fmt.Println(page.Image)
}
