package shared

import (
	"fmt"
	"time"
)

const MainHost = "avitowww.ru"

type AvitoPage struct {
	Hint     string
	Number   string
	Url      string
	Title    string
	Price    string
	PriceInt string
	Image    string
	Date     time.Time
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
}

type AvitoLitePage struct {
	Hint string
	Url  string
}

func (page AvitoPage) Dump() {
	fmt.Println(page.Title)
	fmt.Println(page.Price)
	fmt.Println(page.PriceInt)
	fmt.Println(page.Image)
}
