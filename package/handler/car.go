package handler

import (
	"autokatolog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create a car
// @Tags cars
// @Description Create a new car with an owner
// @ID create-car
// @Accept  json
// @Produce  json
// @Param input body autokatolog.Car true "Car and Owner Info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cars [post]
func (h *Handler) createCar(c *gin.Context) {

	input := autokatolog.Car{
		RegNums: "Default",
		Mark:    "Default",
		Model:   "Default",
		Year:    0,
		Owner: autokatolog.People{
			ID:         0,
			Name:       "default",
			Surname:    "default",
			Patronymic: "default",
		},
	}

	//fmt.Println(input)
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Println(input)

	var owner autokatolog.People

	id, err := h.services.CreateOwner(owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	input.Owner.ID = id
	fmt.Println(input)

	if err := h.services.CreateCar(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, statusResponse{"Ok"})
}

// @Summary Get all cars
// @Tags cars
// @Description Get all cars with pagination and optional filtering by any car attribute
// @ID get-all-cars
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param reg_num query string false "Filter by registration number"
// @Param mark query string false "Filter by car mark"
// @Param model query string false "Filter by car model"
// @Param year query int false "Filter by manufacturing year"
// @Success 200 {array} autokatolog.Car
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cars [get]
func (h *Handler) getAllCars(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		newErrorResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		newErrorResponse(c, http.StatusBadRequest, "invalid page size")
		return
	}

	// Handling year as integer
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		year = 0 // Assuming 0 means "not set" and will not filter by year
	}

	filter := autokatolog.Car{
		RegNums: c.Query("reg_num"),
		Mark:    c.Query("mark"),
		Model:   c.Query("model"),
		Year:    year,
	}

	cars, err := h.services.GetAllCars(filter, page, pageSize)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cars)
}

// @Summary Get car by registration number
// @Tags cars
// @Description Get a single car by its registration number
// @ID get-car-by-regnum
// @Accept  json
// @Produce  json
// @Param reg_num path string true "Car Registration Number"
// @Success 200 {object} autokatolog.Car
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cars/{reg_num} [get]
func (h *Handler) getCarByRegNum(c *gin.Context) {

	regNum := c.Param("reg_num")
	if regNum == "" {
		newErrorResponse(c, http.StatusBadRequest, "reg_num parameter is required")
		return
	}
	//fmt.Println(regNum)

	car, err := h.services.GetCarByRegNum(regNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"not found": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// @Summary Update car
// @Tags cars
// @Description Update car attributes by registration number
// @ID update-car
// @Accept  json
// @Produce  json
// @Param reg_num path string true "Car Registration Number"
// @Param input body autokatolog.Car true "Car Update Data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cars/{reg_num} [put]
func (h *Handler) updateCar(c *gin.Context) {
	var input autokatolog.Car
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	regNum := c.Param("reg_num")
	if regNum == "" {
		newErrorResponse(c, http.StatusBadRequest, "reg_num parameter is required")
		return
	}
	car, err := h.services.GetCarByRegNum(regNum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Owner.ID = car.Owner.ID
	input.RegNums = regNum
	fmt.Println(input)
	_, err = h.services.UpdateCar(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete car
// @Tags cars
// @Description Delete a car by its registration number
// @ID delete-car
// @Accept  json
// @Produce  json
// @Param reg_num path string true "Car Registration Number"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cars/{reg_num} [delete]
func (h *Handler) deleteCar(c *gin.Context) {
	regNum := c.Param("reg_num")
	if regNum == "" {
		newErrorResponse(c, http.StatusBadRequest, "reg_num parameter is required")
		return
	}

	if err := h.services.DeleteCar(regNum); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
