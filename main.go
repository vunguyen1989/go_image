package main

import (
	"io"
	"net/http"
	"strconv"

	"github.com/h2non/bimg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/resize", handleResize)

	e.Logger.Fatal(e.Start(":8080"))
}

func handleResize(c echo.Context) error {
	// Lấy width từ header
	widthStr := c.Request().Header.Get("X-Width")
	width, err := strconv.Atoi(widthStr)
	if err != nil || width <= 0 {
		return c.String(http.StatusBadRequest, "Invalid or missing X-Width header")
	}

	// Lấy height từ header (optional)
	heightStr := c.Request().Header.Get("X-Height")
	height := 0
	if heightStr != "" {
		height, err = strconv.Atoi(heightStr)
		if err != nil || height < 0 {
			return c.String(http.StatusBadRequest, "Invalid X-Height header")
		}
	}

	if height == 0 {
		height = width // Keep aspect ratio if only width provided
	}

	// Lấy file từ request
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Missing image file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot open file")
	}
	defer src.Close()

	// Đọc nội dung file
	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot read file")
	}

	// Xử lý bằng bimg
	img := bimg.NewImage(imageBytes)

	_, err = img.Metadata()
	if err != nil {
		return c.String(http.StatusInternalServerError, "bimg: Failed to load image")
	}

	options := bimg.Options{
		Width:   width,
		Height:  height,
		Quality: 75,
	}

	result, err := img.Process(options)
	if err != nil {
		return c.String(http.StatusInternalServerError, "bimg: Error during processing")
	}

	return c.Blob(http.StatusOK, "image/jpeg", result)
}
