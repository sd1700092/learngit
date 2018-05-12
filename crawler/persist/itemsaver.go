package persist

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"imooc.com/learngo/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetURL("http://192.168.99.100:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver got item #%d: %+v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}
	indexService := client.Index().Index(index).
		Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v\n", resp)
	return err
}
