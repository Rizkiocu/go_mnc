package controller

import (
	"fmt"
	"net/http"
	"test_mnc/model"
	"test_mnc/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentUC usecase.PaymentUseCase
	rg        *gin.RouterGroup
}

func (s *PaymentController) createHandler(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := s.paymentUC.CreateNew(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (u *PaymentController) listHandler(c *gin.Context) {
	payments, err := u.paymentUC.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, payments)
}

func (u *PaymentController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")
	payment, err := u.paymentUC.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, payment)
}

func (u *PaymentController) updateHandler(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := u.paymentUC.Update(payment)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update payment",
	})
}

func (u *PaymentController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.paymentUC.Delete(id); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfully delete payment with id %s", id)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (p *PaymentController) Route() {
	p.rg.POST("/payments", p.createHandler)
	p.rg.GET("/payments", p.listHandler)
	p.rg.GET("/payments/:id", p.getByIdHandler)
	p.rg.PUT("/payments", p.updateHandler)
	p.rg.DELETE("/payments/:id", p.deleteHandler)
}

func NewPaymentController(paymentUC usecase.PaymentUseCase, rg *gin.RouterGroup) *PaymentController {
	return &PaymentController{
		paymentUC: paymentUC,
		rg:        rg,
	}
}
