package services

import (
	"context"
	"goddamnnoob/RabbitMQ-ProductAPI/integrations"
	"goddamnnoob/RabbitMQ-ProductAPI/repositories"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func ConsumerService(ctx context.Context, wg *sync.WaitGroup) {

	connection, err := integrations.GetNewRabbitMQConnection()
	if err != nil {
		log.Println("Error while Getting Connection")
		log.Println(err.Error())
		return
	}
	defer func() {
		connection.Close()
		wg.Done()
	}()

	ch, err := connection.Channel()
	if err != nil {
		log.Println("Error while connecting to consumer channel")
		log.Println(err.Error())
	}

	productids, err := ch.Consume(queuename, "", true, false, false, false, nil)

	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "NOT_FOUND") {
			_, err = ch.QueueDeclare(
				queuename, // Queue name
				false,     // Durable
				false,     // Delete when unused
				false,     // Exclusive
				false,     // No-wait
				nil,       // Arguments
			)
			if err != nil {
				log.Println("Failed to declare queue:")
				log.Println(err.Error())
				return
			}
		} else {
			log.Println(err.Error())
			return
		}
	}
	go func() {
		for productid := range productids {

			uuidObj, err := uuid.Parse(string(productid.Body))
			log.Println(uuidObj.String())
			if err != nil {
				log.Println("Error while getting productids")
				log.Println(err.Error())
			}
			go getImageURLSDownloadUpdateDB(uuidObj)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting Down consumer service")
}

func getImageURLSDownloadUpdateDB(productid uuid.UUID) (err error) {

	product, err := repositories.GetProductURLsFromID(productid)
	log.Println(product)

	if err != nil {
		return err
	}

	for _, productimageurl := range product.Productimages {
		compressedimagepath, err := downloadImageCompressAndStore(productimageurl)
		if err != nil {
			continue
		}
		product.Compressed_product_images = append(product.Compressed_product_images, compressedimagepath)
	}
	err = repositories.UpdateCompressedImagePaths(product)

	if err != nil {
		return err
	}
	return nil
}

func downloadImageCompressAndStore(imageurl string) (path string, err error) {

	response, err := http.Get(imageurl)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	imagefilename := filepath.Base(imageurl)

	imagefile, err := os.Create(imagefilename)

	if err != nil {
		return "", err
	}
	defer imagefile.Close()

	_, err = io.Copy(imagefile, response.Body)

	if err != nil {
		return "", err
	}

	imagefile, err = os.Open(imagefilename)

	if err != nil {
		return "", err
	}
	defer imagefile.Close()

	imageData, _, err := image.Decode(imagefile)

	if err != nil {
		return "", err
	}

	compressedImage := resize.Resize(800, 0, imageData, resize.Lanczos3)

	compressedFileName := "c_" + imagefilename
	compressedFile, err := os.Create(compressedFileName)
	if err != nil {
		return "", err
	}
	defer compressedFile.Close()

	err = jpeg.Encode(compressedFile, compressedImage, nil)
	if err != nil {
		return "", err
	}

	cfstat, err := compressedFile.Stat()

	if err != nil {
		return "", nil
	}

	return cfstat.Name(), nil

}
