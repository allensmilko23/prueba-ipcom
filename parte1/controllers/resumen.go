package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"prueba-ipcom/entities"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MakeRequest(currentrequest string, ch chan<- []entities.Order) {
	url := fmt.Sprintf("https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-0%s", currentrequest)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var orders []entities.Order
	err = json.Unmarshal(body, &orders)
	if err != nil {
		panic(err)
	}
	ch <- orders
}

func GetchannelRequests(ch <-chan []entities.Order) []entities.Order {
	val := <-ch
	return val
}

func GetResumen(c echo.Context) error {
	dayNumber := c.QueryParam("days")
	date := c.Param("date")

	days, err := strconv.Atoi(dayNumber)
	if err == nil {
		fmt.Println(days)
	}

	if days > 30 {
		errorResponse := &entities.ErrorHandler{
			ErrorCode:   400,
			Description: "Days value must be < = 30",
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	validDate := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")
	isValidDate := validDate.MatchString(date)

	if !isValidDate {
		errorResponse := &entities.ErrorHandler{
			ErrorCode:   400,
			Description: "Invalid date format, must be YYY-MM-DD",
		}

		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	channel := make(chan []entities.Order)

	for currentrequest := 0; currentrequest < days; currentrequest++ {
		go MakeRequest(strconv.Itoa(currentrequest+1), channel)
	}

	var ordersStatistics entities.OrdersStatistics
	var orders []entities.Order

	for currentrequest := 0; currentrequest < days; currentrequest++ {
		channlesResponse := GetchannelRequests(channel)
		orders = append(orders, channlesResponse...)
	}

	for _, ord := range orders {
		ordersStatistics.Total = ordersStatistics.Total + ord.Monto
		if ord.Tdc == "visa gold" {
			ordersStatistics.Comprasportdc.Oro = ordersStatistics.Comprasportdc.Oro + 1
		}
		if ord.Tdc == "amex" {
			ordersStatistics.Comprasportdc.Amex = ordersStatistics.Comprasportdc.Oro + 1
		}
		if !ord.Compro {
			ordersStatistics.Nocompraron = ordersStatistics.Nocompraron + 1
		}
		if ord.Monto > ordersStatistics.Compramasalta {
			ordersStatistics.Compramasalta = ord.Monto
		}
	}

	fmt.Print(ordersStatistics)
	return c.JSON(http.StatusOK, ordersStatistics)

}
