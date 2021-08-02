package topics

import (
	"fmt"
	"libria/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	topicService Service
}

func NewDelivery(topicService Service) Delivery {
	return Delivery{
		topicService: topicService,
	}
}

func (d *Delivery) GetAll(c echo.Context) error {
	topics, err := d.topicService.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, topics)
}
func (d *Delivery) GetById(c echo.Context) error {
	id := c.Param("id")
	topic, err := d.topicService.GetById(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, topic)
}

func (d *Delivery) GetRandom(c echo.Context) error {
	topics, err := d.topicService.GetRandom()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, topics)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.Topic)
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	topic, err := d.topicService.Post(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, topic.ID)
}

func (d *Delivery) Update(c echo.Context) (err error) {
	id := c.Param("id")

	requestBody := new(models.Topic)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	topic, err := d.topicService.Update(id, requestBody)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, topic)
}

func (d *Delivery) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	topic, err := d.topicService.Delete(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, topic)
}
