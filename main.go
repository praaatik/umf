package main

import (
	"log"

	"github.com/praaatik/umf/api"
	// "github.com/praaatik/umf/metadatafetcher"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(len(listOfUrls))
	//
	// metadataChannel := make(chan metadatafetcher.OpenGraphMetaData, len(listOfUrls))
	//
	// for _, url := range listOfUrls {
	// 	go func(url string) {
	// 		defer wg.Done()
	// 		urlMetadata := metadatafetcher.FetchHTML(url)
	// 		metadataChannel <- urlMetadata
	// 	}(url)
	// }
	//
	// go func() {
	// 	wg.Wait()
	// 	close(metadataChannel)
	// }()
	//
	// for metadata := range metadataChannel {
	// 	fmt.Println("title -> ", metadata.Title)
	// 	fmt.Println("url -> ", metadata.URL)
	// 	fmt.Println("descripion -> ", metadata.Description)
	// 	fmt.Println("audio -> ", metadata.Audio)
	// 	fmt.Println("video -> ", metadata.Video)
	// 	fmt.Println("sitename -> ", metadata.SiteName)
	// 	fmt.Println("type -> ", metadata.Type)
	// 	fmt.Println("image -> ", metadata.Image)
	// }

	// server code goes in here
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("unable to create new server")
	}

	server.Start()
}
