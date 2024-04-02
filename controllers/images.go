package controllers

import (
	"github.com/lemadane/admin_backend_gofiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

// UploadImage handles the HTTP POST request for uploading images.
// It expects a multipart form with one or more image files.
// The function saves the uploaded files to the "./uploads" directory.
// It returns a JSON response with a success message if the upload is successful.
func UploadImage(context *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(context, "images"); err != nil {
		return err
	}
	form, err := context.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	var filename = ""
	for _, file := range files {
		filename = file.Filename
		if err := context.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}
	return context.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
