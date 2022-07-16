package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
	"xyz.frankity/gosanime/main/config"
	"xyz.frankity/gosanime/main/models"
)

func ovas() ([]models.Anime, error) {
	animes := []models.Anime{}
	resp, err := soup.Get(fmt.Sprintf("%v%s", config.Rooturl, config.Ovasurl))
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	main := doc.Find("div", "class", "anime__page__content").FindAll("div", "class", "anime__item")

	for _, p := range main {
		anime := models.Anime{
			ID:     strings.Split(p.Find("h5").Find("a").Attrs()["href"], "/")[3],
			Name:   p.Find("h5").Find("a").Text(),
			Poster: p.Find("a").Find("div", "class", "anime__item__pic").Attrs()["data-setbg"],
			State:  p.Find("div", "class", "anime__item__text").Find("ul").FindAll("li")[0].Text(),
			Type:   strings.TrimSpace(p.Find("div", "class", "anime__item__text").Find("ul").FindAll("li")[1].Text()),
		}
		animes = append(animes, anime)
	}

	return animes, nil
}
