package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SaiHLu/simple-rest/models"
	"github.com/SaiHLu/simple-rest/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	if err := context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed.",
		})

		return err
	}

	if err := r.DB.Create(&book).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not create the book",
		})

		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Book has been created.",
	})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	booksModels := &[]models.Book{}

	if err := r.DB.Find(booksModels).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get the books.",
		})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get the books",
		"data":    booksModels,
	})

	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	booksModel := models.Book{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Id  is required.",
		})

		return nil
	}

	if err := r.DB.Delete(booksModel, id).Error; err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Something went wrong",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book Deleted.",
	})

	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Book{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Id  is required.",
		})

		return nil
	}

	if err := r.DB.Where("id = ?", id).First(bookModel).Error; err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Book not found.",
		})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Found Book.",
		"data":    &bookModel,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-books", r.CreateBook)
	api.Delete("/delete-books/:id", r.DeleteBook)
	api.Get("/get-books/:id", r.GetBookById)
	api.Get("/get-books", r.GetBooks)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		// SSLMode:  os.Getenv("DB_SSLMode"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database.")
	}

	if err = models.MigrateBooks(db); err != nil {
		log.Fatal("Could migrate db.")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":8080")
}
