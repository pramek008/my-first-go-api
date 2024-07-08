package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pramek008/first-golang/database"
	"github.com/pramek008/first-golang/models"
)

func Home(c *fiber.Ctx)error{
	return c.SendString("First Go Rest API")
}

func CreateFact(c *fiber.Ctx)error{
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		}) 
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
  var fact models.Fact
  id := c.Params("id")

  // Find the existing fact by ID
  result := database.DB.Db.Where("id = ?", id).First(&fact)
  if result.Error != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "message": "Fact not found",
    })
  }

  // Parse the request body to populate the fact object
  var updatedFact models.Fact
  if err := c.BodyParser(&updatedFact); err != nil {
    return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
  }

  // Update the fields of the existing fact
  fact.Question = updatedFact.Question
  fact.Answer = updatedFact.Answer

  // Save the updated fact to the database
  result = database.DB.Db.Save(&fact)
  if result.Error != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message": "Error updating fact",
    })
  }

  // Return the updated fact
  return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
    var fact models.Fact
    id := c.Params("id")

    // Find the existing fact by ID
    result := database.DB.Db.Where("id = ?", id).First(&fact)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Fact not found",
        })
    }

    // Delete the fact
    if err := database.DB.Db.Delete(&fact).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error deleting fact",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Fact deleted successfully",
    })
}
