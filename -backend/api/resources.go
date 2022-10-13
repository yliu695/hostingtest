package api

import (
	"net/http"

	"wcs/dao"
	"wcs/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configResourcesRouter(router *httprouter.Router) {
	router.GET("/resources", GetAllResources)
	router.POST("/resources", AddResources)
	router.GET("/resources/:argID", GetResources)
	router.PUT("/resources/:argID", UpdateResources)
	router.DELETE("/resources/:argID", DeleteResources)
}

func configGinResourcesRouter(router gin.IRoutes) {
	router.GET("/resources", ConverHttprouterToGin(GetAllResources))
	router.POST("/resources", ConverHttprouterToGin(AddResources))
	router.GET("/resources/:argID", ConverHttprouterToGin(GetResources))
	router.PUT("/resources/:argID", ConverHttprouterToGin(UpdateResources))
	router.DELETE("/resources/:argID", ConverHttprouterToGin(DeleteResources))
}

// GetAllResources is a function to get a slice of record(s) from resources table in the wcs database
// @Summary Get list of Resources
// @Tags Resources
// @Description GetAllResources is a handler to get a slice of record(s) from resources table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Resources}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resources [get]
// http "http://localhost:8080/resources?page=0&pagesize=20" X-Api-User:user123
func GetAllResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "resources", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllResources(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetResources is a function to get a single record from the resources table in the wcs database
// @Summary Get record from table Resources by  argID
// @Tags Resources
// @ID argID
// @Description GetResources is a function to get a single record from the resources table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Resources
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /resources/{argID} [get]
// http "http://localhost:8080/resources/1" X-Api-User:user123
func GetResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resources", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetResources(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddResources add to add a single record to resources table in the wcs database
// @Summary Add an record to resources table
// @Description add to add a single record to resources table in the wcs database
// @Tags Resources
// @Accept  json
// @Produce  json
// @Param Resources body model.Resources true "Add Resources"
// @Success 200 {object} model.Resources
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resources [post]
// echo '{"id": 57,"name": "UvJpIyJJkjDDfKyCmQxdwbcwA","intro": "yXkmMtyeYBsJGFUtZqshRSFLB","link": "TIPRoVvioBUWsOXXxvWxmYIMY"}' | http POST "http://localhost:8080/resources" X-Api-User:user123
func AddResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	resources := &model.Resources{}

	if err := readJSON(r, resources); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resources.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resources.Prepare()

	if err := resources.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resources", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	resources, _, err = dao.AddResources(ctx, resources)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resources)
}

// UpdateResources Update a single record from resources table in the wcs database
// @Summary Update an record in table resources
// @Description Update a single record from resources table in the wcs database
// @Tags Resources
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Resources body model.Resources true "Update Resources record"
// @Success 200 {object} model.Resources
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /resources/{argID} [put]
// echo '{"id": 57,"name": "UvJpIyJJkjDDfKyCmQxdwbcwA","intro": "yXkmMtyeYBsJGFUtZqshRSFLB","link": "TIPRoVvioBUWsOXXxvWxmYIMY"}' | http PUT "http://localhost:8080/resources/1"  X-Api-User:user123
func UpdateResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resources := &model.Resources{}
	if err := readJSON(r, resources); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := resources.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	resources.Prepare()

	if err := resources.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "resources", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	resources, _, err = dao.UpdateResources(ctx,
		argID,
		resources)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, resources)
}

// DeleteResources Delete a single record from resources table in the wcs database
// @Summary Delete a record from resources
// @Description Delete a single record from resources table in the wcs database
// @Tags Resources
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Resources
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /resources/{argID} [delete]
// http DELETE "http://localhost:8080/resources/1" X-Api-User:user123
func DeleteResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "resources", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteResources(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
