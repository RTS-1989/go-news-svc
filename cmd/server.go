package main

import (
	"GoNews/pkg/config"
	"GoNews/pkg/pb"
	"GoNews/pkg/rss"
	"GoNews/pkg/services"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Start GoNews")

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	connString := c.DBUrl

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postsChan := make(chan storage.Post)
	errChan := make(chan error)
	runErrorsCheck(errChan)
	db, err := postgres.New(ctx, connString, postsChan, errChan)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	configFile := "config.json"
	sourceRss, err := rss.NewSourceRss(configFile, postsChan, errChan)

	if err != nil {
		log.Fatal(err)
	}

	// Получаем список последних публикаций из базы данных (если они есть). Чтобы не дублировать данные в бд.
	lastPubTimeMap, _ := db.GetLastPubDateForSources(sourceRss.SourceRssList)
	sourceRss.LastPubTimeFromDB = lastPubTimeMap

	sourceRss.RunGetSourcesInfo()
	db.RunInsertPosts()

	s := services.Server{
		H: db,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterGoNewsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

// runErrorsCheck читает ошибки из канало ошибок
func runErrorsCheck(errChan chan error) {
	go func() {
		for {
			select {
			case err := <-errChan:
				log.Println(err)
			}
		}
	}()
}