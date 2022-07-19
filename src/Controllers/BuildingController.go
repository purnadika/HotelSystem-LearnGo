package controller

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"HotelSystem-LearnGo/Services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BuildingController struct {
	BuildingService Services.IBuildingService
}

func NewBuildingController(buildingService Services.IBuildingService) IBuildingController {
	return &BuildingController{
		BuildingService: buildingService,
	}
}

func (controller *BuildingController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingCreateRequest := Requests.BuildingCreateRequest{}
	helper.ReadFromRequestBody(request, &buildingCreateRequest)
	helper.LogDebug("CreateBuilding", buildingCreateRequest)
	buildingResponse := controller.BuildingService.Create(buildingCreateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *BuildingController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingUpdateRequest := Requests.BuildingUpdateRequest{}
	helper.ReadFromRequestBody(request, &buildingUpdateRequest)
	helper.LogDebug("UpdateBuilding", request)
	buildingResponse := controller.BuildingService.Update(buildingUpdateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *BuildingController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingIdString := params.ByName("buildingId")
	buildingId := helper.StringToUint(buildingIdString)
	helper.LogDebug("DeleteBuilding", params)
	buildingResponse := controller.BuildingService.Delete(uint(buildingId))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *BuildingController) FindByBuildingName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingName := params.ByName("buildingName")
	buildingResponse := controller.BuildingService.FindByName(buildingName)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *BuildingController) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	getAllRequest := Requests.GetAllRequest{}
	helper.ReadFromRequestBody(request, &getAllRequest)
	helper.LogDebug("GetBuildings", getAllRequest)
	buildingsResponse := controller.BuildingService.GetAll(getAllRequest)
	jsontext, err := json.Marshal(buildingsResponse)
	helper.PanicIfError(err)
	log.Println(string(jsontext))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   buildingsResponse,
	}
	helper.WriteToResponseBody(writer, response)
}
