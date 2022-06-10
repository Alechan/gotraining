package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Interfaz que agrupa a las respuestas de error y las successful
type httpResponder interface {
	Status() int
	Body() interface{}
}

// Successful Response (no exportada)
type HttpSuccessfulResponse struct {
	status int
	body   interface{}
}

func (e *HttpSuccessfulResponse) Status() int {
	return e.status
}

func (e *HttpSuccessfulResponse) Body() interface{} {
	return e.body
}

// Http Error Response
type ErrorCode struct {
	Literal   string
	Status    int
	Alertable bool
}

type HttpErrorResponse struct {
	error
	Code    ErrorCode         `json:"code,omitempty"`
	Cause   string            `json:"cause,omitempty"`
	Message string            `json:"message,omitempty"`
	Values  map[string]string `json:"values,omitempty"`
}

func (e *HttpErrorResponse) Status() int {
	return e.Code.Status
}

func (e *HttpErrorResponse) Body() interface{} {
	return e
}

func (e *HttpErrorResponse) Error() string {
	return fmt.Sprintf("%s - %s: %s %v", e.Code.Literal, e.Cause, e.Message, e.error)
}

func (e *HttpErrorResponse) MarshalJSON() ([]byte, error) {
	// serialize CODE as
	return json.Marshal(&struct {
		Error   string            `json:"error,omitempty"`
		Cause   string            `json:"cause,omitempty"`
		Message string            `json:"message,omitempty"`
		Values  map[string]string `json:"values,omitempty"`
	}{
		Error:   e.Code.Literal,
		Cause:   e.Cause,
		Message: e.Message,
		Values:  e.Values,
	})
}

var (
	BadRequestApiError = ErrorCode{
		Status:    http.StatusBadRequest,
		Literal:   "BadRequestApiError",
		Alertable: false,
	}

	NotFoundApiError = ErrorCode{
		Status:    http.StatusNotFound,
		Literal:   "NotFoundApiError",
		Alertable: false,
	}
)

func main() {
	r := gin.Default()

	r.GET("/200", middleware(handlerDe200))
	r.GET("/201", middleware(handlerDe201))
	r.GET("/404", middleware(handlerDe404))

	_ = r.Run()
}

func handlerDe200() httpResponder {
	type TipoUno struct {
		UnCampo string `json:"UnCampo,omitempty"`
	}
	unTipoUno := TipoUno{UnCampo: "respondiendo 200"}
	response := &HttpSuccessfulResponse{
		status: http.StatusOK,
		body:   unTipoUno,
	}
	return response
}

func handlerDe201() httpResponder {
	type TipoDos struct {
		OtroCampo int `json:"OtroCampo,omitempty"`
	}
	unTipoDos := TipoDos{OtroCampo: 201}
	response := &HttpSuccessfulResponse{
		status: http.StatusCreated,
		body:   unTipoDos,
	}
	return response
}

func handlerDe404() httpResponder {
	response := &HttpErrorResponse{
		Code:    NotFoundApiError,
		Cause:   "Tu vieja not found",
		Message: "Tu vieja en serio not found",
		Values:  map[string]string{"tu_vieja": "Hola, Ema. Sos un crack"},
	}
	return response
}

func middleware(handler func() httpResponder) gin.HandlerFunc {
	return func(c *gin.Context) {
		// COSAS ANTES DE LA LLAMADA
		// (no hay nada)

		// llamada
		response := handler()

		// COSAS DESPUÉS DE LA LLAMADA
		c.JSON(response.Status(), response.Body())

		//v, isError := response.(error)
		//if isError {
		//    logError(v)
		//}
	}
}
