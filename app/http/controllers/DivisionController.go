package controllers

import (
	"log"
	"os"
	"time"
	"unicode"

	Configuration "github.com/anik4good/go-echo-apiboilerplate/config"
	"github.com/anik4good/go-echo-apiboilerplate/models"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	type KeyValuePair struct {
		Status    interface{}
		Name      string
		Divisions string
		Districts string
		District  string
		Upozilla  string
		Unions    string
	}

	return c.JSON(200, KeyValuePair{
		Status:    true,
		Name:      "Welcome to GeoAPIBD. Powered by GO, Developed by Anik",
		Divisions: os.Getenv("SERVER_URL") + "/api/divisions",
		Districts: os.Getenv("SERVER_URL") + "/api/districts",
		District:  os.Getenv("SERVER_URL") + "/api/division/rangpur",
		Upozilla:  os.Getenv("SERVER_URL") + "/api/division/rangpur/panchagarh",
		Unions:    os.Getenv("SERVER_URL") + "/api/division/rangpur/panchagarh/debiganj",
	})
}

func Index(c echo.Context) error {
	var division []models.Division
	result := Configuration.GormDBConn.Find(&division)

	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Data Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	response := models.Response{
		Status:  true,
		Message: "ALL Divisions",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data:    division,
	}

	return c.JSON(200, response)
}

func Show(c echo.Context) error {
	id := c.Param("id")

	var division models.Division
	result := Configuration.GormDBConn.First(&division, id)

	log.Println(result.RowsAffected)
	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Data Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return c.JSON(200, division)
}

func GetAllDistricts(c echo.Context) error {
	var district []models.District
	result := Configuration.GormDBConn.Find(&district)
	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Data Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	response := models.Response{
		Status:  true,
		Message: "ALL Districts",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data:    district,
	}

	return c.JSON(200, response)
}

func GetAllDistrict(c echo.Context) error {
	division_name := Capitalize(c.Param("division_name"))
	log.Println(division_name)

	var division models.Division
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+division_name+"%").First(&division)

	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Division Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var district []models.District
	result1 := Configuration.GormDBConn.Where("division_id = ?", division.ID).Find(&district)
	if result1.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Data Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	response := models.Response{
		Status:  true,
		Message: "ALL Districts Under " + division.Name,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data:    district,
	}

	return c.JSON(200, response)
}

func GetAllUpozilla(c echo.Context) error {
	district_name := Capitalize(c.Param("district_name"))
	log.Println(district_name)

	var district models.District
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+district_name+"%").First(&district)

	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No District Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var upazilas []models.Upazila
	result1 := Configuration.GormDBConn.Where("district_id = ?", district.ID).Find(&upazilas)
	if result1.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Upazilas Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	response := models.Response{
		Status:  true,
		Message: "ALL Upazilas Under " + district.Name,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data:    upazilas,
	}

	return c.JSON(200, response)
}

func GetAllUnions(c echo.Context) error {
	upazila_name := Capitalize(c.Param("upazila_name"))
	log.Println(upazila_name)

	var upazila models.Upazila
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+upazila_name+"%").First(&upazila)

	if result.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Upazila Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var union []models.Union
	result1 := Configuration.GormDBConn.Where("upazilla_id = ?", upazila.ID).Find(&union)
	if result1.Error != nil {
		return c.JSON(200, models.Response{
			Status:  true,
			Message: "No Union Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	response := models.Response{
		Status:  true,
		Message: "ALL Unions Under " + upazila.Name,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data:    union,
	}

	return c.JSON(200, response)
}

func Capitalize(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToTitle(runes[0])

	return string(runes)
}
