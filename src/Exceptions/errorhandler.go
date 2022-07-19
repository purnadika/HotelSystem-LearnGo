package exceptions

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Responses"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, err) {
		return
	}

	if validationErrors(writer, err) {
		return
	}

	internalServerError(writer, err)
}

func validationErrors(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := Responses.GeneralResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := Responses.GeneralResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := Responses.GeneralResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
