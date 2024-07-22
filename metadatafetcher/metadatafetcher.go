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

func checkTagPresent(metadata map[string]string, tag string) (string, string) {
	if tagName, exists := metadata[tag]; exists {
		return tag, tagName
	}
	return tag, ""
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

	// fmt.Println("==========================================")
	// fmt.Println(url)
	_, titleTag := checkTagPresent(metadata, "og:title")
	_, urlTag := checkTagPresent(metadata, "og:url")
	_, imageTag := checkTagPresent(metadata, "og:image")
	_, typeTag := checkTagPresent(metadata, "og:type")
	_, siteNameTag := checkTagPresent(metadata, "og:site_name")
	_, descriptionTag := checkTagPresent(metadata, "og:description")
	_, videoTag := checkTagPresent(metadata, "og:video")
	_, audioTag := checkTagPresent(metadata, "og:audio")
	// fmt.Println("==========================================")

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
