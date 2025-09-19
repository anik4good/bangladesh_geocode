package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
	"unicode"

	Configuration "github.com/anik4good/go-echo-apiboilerplate/config"
	"github.com/anik4good/go-echo-apiboilerplate/middleware"
	"github.com/anik4good/go-echo-apiboilerplate/models"
	"github.com/labstack/echo/v4"
)

// Function to pretty-print JSON and send response
func writePrettyJSON(c echo.Context, statusCode int, data interface{}) error {
	// Pretty-print JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	// Set headers and write response
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(statusCode)
	_, err = bytes.NewBuffer(jsonData).WriteTo(c.Response())
	return err
}

// Hello Endpoint
func Hello(c echo.Context) error {
	type KeyValuePair struct {
		Status        interface{} `json:"status"`
		Name          string      `json:"name"`
		Divisions     string      `json:"divisions"`
		Districts     string      `json:"districts"`
		District      string      `json:"district"`
		Upozilla      string      `json:"upozilla"`
		Unions        string      `json:"unions"`
		Documentation string
	}

	return writePrettyJSON(c, 200, KeyValuePair{
		Status:        true,
		Name:          "Welcome to GeoAPIBD. Powered by GO, Developed by Anik Deployed in K8s",
		Divisions:     os.Getenv("SERVER_URL") + "/api/divisions",
		Districts:     os.Getenv("SERVER_URL") + "/api/districts",
		District:      os.Getenv("SERVER_URL") + "/api/division/rangpur",
		Upozilla:      os.Getenv("SERVER_URL") + "/api/division/rangpur/panchagarh",
		Unions:        os.Getenv("SERVER_URL") + "/api/division/rangpur/panchagarh/debiganj",
		Documentation: os.Getenv("SERVER_URL") + "/docs",
	})
}

// Add this function to serve the documentation
func ServeDocs(c echo.Context) error {
	return c.File("public/docs.html")
}

// Index Endpoint
func Index(c echo.Context) error {
	var division []models.Division
	start := time.Now()
	result := Configuration.GormDBConn.Find(&division)
	middleware.RecordDBQueryDuration("FindAllDivisions", time.Since(start))

	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
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

	return writePrettyJSON(c, 200, response)
}

// Show Endpoint
func Show(c echo.Context) error {
	id := c.Param("id")

	var division models.Division
	start := time.Now()
	result := Configuration.GormDBConn.First(&division, id)
	middleware.RecordDBQueryDuration("FindDivisionByID", time.Since(start))

	log.Println(result.RowsAffected)
	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
			Status:  true,
			Message: "No Data Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return writePrettyJSON(c, 200, division)
}

// GetAllDistricts Endpoint
func GetAllDistricts(c echo.Context) error {
	var district []models.District
	start := time.Now()
	result := Configuration.GormDBConn.Find(&district)
	middleware.RecordDBQueryDuration("FindAllDistricts", time.Since(start))

	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
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

	return writePrettyJSON(c, 200, response)
}

// GetAllDistrict Endpoint
func GetAllDistrict(c echo.Context) error {
	division_name := Capitalize(c.Param("division_name"))
	log.Println(division_name)

	var division models.Division
	start := time.Now()
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+division_name+"%").First(&division)
	middleware.RecordDBQueryDuration("FindDivisionByName", time.Since(start))

	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
			Status:  true,
			Message: "No Division Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var district []models.District
	start = time.Now()
	result1 := Configuration.GormDBConn.Where("division_id = ?", division.ID).Find(&district)
	middleware.RecordDBQueryDuration("FindDistrictsByDivision", time.Since(start))
	if result1.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
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

	return writePrettyJSON(c, 200, response)
}

// GetAllUpozilla Endpoint
func GetAllUpozilla(c echo.Context) error {
	district_name := Capitalize(c.Param("district_name"))
	log.Println(district_name)

	var district models.District
	start := time.Now()
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+district_name+"%").First(&district)
	middleware.RecordDBQueryDuration("FindDistrictByName", time.Since(start))

	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
			Status:  true,
			Message: "No District Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var upazilas []models.Upazila
	start = time.Now()
	result1 := Configuration.GormDBConn.Where("district_id = ?", district.ID).Find(&upazilas)
	middleware.RecordDBQueryDuration("FindUpazilasByDistrict", time.Since(start))
	if result1.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
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

	return writePrettyJSON(c, 200, response)
}

// GetAllUnions Endpoint
func GetAllUnions(c echo.Context) error {
	upazila_name := Capitalize(c.Param("upazila_name"))
	log.Println(upazila_name)

	var upazila models.Upazila
	start := time.Now()
	result := Configuration.GormDBConn.Where("name LIKE ?", "%"+upazila_name+"%").First(&upazila)
	middleware.RecordDBQueryDuration("FindUpazilaByName", time.Since(start))

	if result.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
			Status:  true,
			Message: "No Upazila Found",
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var union []models.Union
	start = time.Now()
	result1 := Configuration.GormDBConn.Where("upazilla_id = ?", upazila.ID).Find(&union)
	middleware.RecordDBQueryDuration("FindUnionsByUpazila", time.Since(start))
	if result1.Error != nil {
		return writePrettyJSON(c, 200, models.Response{
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

	return writePrettyJSON(c, 200, response)
}

// Capitalize function to make the first letter uppercase
func Capitalize(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToTitle(runes[0])

	return string(runes)
}
