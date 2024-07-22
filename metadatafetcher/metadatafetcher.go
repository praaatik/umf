package metadatafetcher

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type OpenGraphMetaData struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	Type        string `json:"type"`
	SiteName    string `json:"site_name"`
	Description string `json:"description"`
	Audio       string `json:"audio"`
	Video       string `json:"video"`
}

func checkTagPresent(metadata map[string]string, tag string) string {
	if tagName, exists := metadata[tag]; exists {
		return tagName
	}
	return ""
}

func FetchHTML(url string) (OpenGraphMetaData, error) {
	respon, err := http.Get(url)
	if err != nil {
		return OpenGraphMetaData{}, err
	}

	doc, err := goquery.NewDocumentFromReader(respon.Body)
	if err != nil {
		return OpenGraphMetaData{}, err
	}

	metadata := make(map[string]string)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		content, _ := s.Attr("content")
		property, _ := s.Attr("property")
		if name != "" {
			metadata[name] = content
		} else if property != "" {
			metadata[property] = content
		}
	})

	titleTag := checkTagPresent(metadata, "og:title")
	urlTag := checkTagPresent(metadata, "og:url")
	imageTag := checkTagPresent(metadata, "og:image")
	typeTag := checkTagPresent(metadata, "og:type")
	siteNameTag := checkTagPresent(metadata, "og:site_name")
	descriptionTag := checkTagPresent(metadata, "og:description")
	videoTag := checkTagPresent(metadata, "og:video")
	audioTag := checkTagPresent(metadata, "og:audio")

	return OpenGraphMetaData{
		Title:       titleTag,
		URL:         urlTag,
		Image:       imageTag,
		Type:        typeTag,
		SiteName:    siteNameTag,
		Description: descriptionTag,
		Audio:       audioTag,
		Video:       videoTag,
	}, nil
}
