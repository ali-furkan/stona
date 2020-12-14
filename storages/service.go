package storages

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"stona/config"
	"stona/images"
	"stona/tools/messages"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type StorageService struct{}

var storageService = new(StorageService)
var ctx = context.Background()
var (
	storageClient *storage.Client
	bucket        *storage.BucketHandle
	att           *storage.BucketAttrs
)

func Service() *StorageService {
	return storageService
}

func (s *StorageService) GetAsset(c *fiber.Ctx) error {
	var err error
	obj := bucket.Object(c.Params("path") + "/" + c.Params("id"))

	nr, err := obj.NewReader(ctx)

	if err != nil {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	defer nr.Close()

	c.Set("Content-Type", nr.ContentType())

	if strings.HasPrefix(nr.ContentType(), "image") {
		var buf bytes.Buffer
		sizeStr := c.Query("size")
		size, err := strconv.ParseUint(sizeStr, 10, 0)
		if err != nil {
			if sizeStr == "" {
				return c.SendStream(nr)
			}
			return messages.ErrorMessage(c, fiber.StatusBadRequest, "Invalid 'size' query param. Please use ascii number")
		}

		t := strings.Split(nr.ContentType(), "/")[1]

		buf, err = images.ReSize(uint(size), nr, t)

		if err != nil {
			return messages.ErrorMessage(c, fiber.StatusBadRequest, err.Error())
		}

		return c.SendStream(bytes.NewReader(buf.Bytes()))
	}

	return c.SendStream(nr)
}

func (s *StorageService) GetList(c *fiber.Ctx) error {
	sizeStr := c.Query("size", "10")
	beginStr := c.Query("begin", "0")

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, "Invalid 'size' query param. Please use ascii number")
	}
	begin, err := strconv.Atoi(beginStr)
	if err != nil || begin < 0 {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, "Invalid 'begin' query param. Please use ascii number")
	}

	var files []map[string]interface{}

	query := &storage.Query{
		Prefix: c.Params("path"),
	}

	it := bucket.Objects(ctx, query)

	for i := 0; ; i++ {
		attrs, err := it.Next()
		if err == iterator.Done || size == i+1 {
			break
		}

		if i < begin {
			continue
		}

		if err != nil {
			return messages.ErrorMessage(c, fiber.StatusInternalServerError, err.Error())
		}

		files = append(files, fiber.Map{
			"url":       c.BaseURL() + config.Config().RootPath + "/" + attrs.Name,
			"type":      attrs.ContentType,
			"size":      attrs.Size,
			"createdAt": attrs.Created,
		})
	}
	if files == nil {
		return messages.ErrorMessage(c, fiber.StatusNotFound, "File not found at "+c.Params("path"))
	}
	return c.JSON(files)
}

func (s *StorageService) Delete(c *fiber.Ctx) error {

	obj := bucket.Object(c.Params("path") + "/" + c.Params("id"))

	if err := obj.Delete(ctx); err != nil {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	return c.SendString("Delete" + c.Params("id"))
}

func (s *StorageService) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	var buf bytes.Buffer
	if err != nil {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, "Form-Data is required")
	}

	f, err := file.Open()

	if err != nil {
		return messages.ErrorMessage(c, fiber.StatusBadRequest, "File is not valid")
	}

	defer f.Close()

	t := file.Header["Content-Type"][0]

	fmt.Println(t)

	if strings.HasPrefix(t, "image") {
		m, err := images.Decode(f, strings.Split(t, "/")[1])
		if err != nil {
			return messages.ErrorMessage(c, fiber.StatusInternalServerError, err.Error())
		}
		size := m.Bounds().Size()
		if size.X > config.Config().ImgMaxResolution || size.Y > config.Config().ImgMaxResolution {
			return messages.ErrorMessage(c, fiber.StatusBadRequest, fmt.Sprintf("Image Resolution must be smaller than %d px", config.Config().ImgMaxResolution))
		}
		buf, err = images.Encode(m, strings.Split(t, "/")[1])
		if err != nil {
			return messages.ErrorMessage(c, fiber.StatusInternalServerError, err.Error())
		}
	} else {
		buf.ReadFrom(f)
	}

	name := c.Params("id")
	if name == "" {
		name = uuid.New().String()
	}

	pathname := c.Params("path") + "/" + name + filepath.Ext(file.Filename)

	obj := bucket.Object(pathname)

	if _, err := obj.NewReader(ctx); err == nil {
		return messages.ErrorMessage(c, fiber.StatusConflict, "This file already exists")
	}

	sw := obj.NewWriter(ctx)

	if _, err := sw.Write(buf.Bytes()); err != nil {
		return messages.ErrorMessage(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := sw.Close(); err != nil {
		return messages.ErrorMessage(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Successfully Uploaded",
		"metadata": fiber.Map{
			"type": t,
			"size": file.Size,
			"ext":  filepath.Ext(file.Filename),
		},
		"path": c.BaseURL() + "/" + pathname,
	})
}

func init() {
	var err error

	fireabseCfg, err := json.Marshal(config.Config().FirebaseConfig)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsJSON(fireabseCfg))
	if err != nil {
		log.Fatal(err.Error())
	}
	bucket = storageClient.Bucket(config.Config().FirebaseConfig.StorageBucket)

	att, err = bucket.Attrs(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}
