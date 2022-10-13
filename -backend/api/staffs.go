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

func configStaffsRouter(router *httprouter.Router) {
	router.GET("/staffs", GetAllStaffs)
	router.POST("/staffs", AddStaffs)
	router.GET("/staffs/:argID", GetStaffs)
	router.PUT("/staffs/:argID", UpdateStaffs)
	router.DELETE("/staffs/:argID", DeleteStaffs)
}

func configGinStaffsRouter(router gin.IRoutes) {
	router.GET("/staffs", ConverHttprouterToGin(GetAllStaffs))
	router.POST("/staffs", ConverHttprouterToGin(AddStaffs))
	router.GET("/staffs/:argID", ConverHttprouterToGin(GetStaffs))
	router.PUT("/staffs/:argID", ConverHttprouterToGin(UpdateStaffs))
	router.DELETE("/staffs/:argID", ConverHttprouterToGin(DeleteStaffs))
}

// GetAllStaffs is a function to get a slice of record(s) from staffs table in the wcs database
// @Summary Get list of Staffs
// @Tags Staffs
// @Description GetAllStaffs is a handler to get a slice of record(s) from staffs table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Staffs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /staffs [get]
// http "http://localhost:8080/staffs?page=0&pagesize=20" X-Api-User:user123
func GetAllStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "staffs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllStaffs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetStaffs is a function to get a single record from the staffs table in the wcs database
// @Summary Get record from table Staffs by  argID
// @Tags Staffs
// @ID argID
// @Description GetStaffs is a function to get a single record from the staffs table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Staffs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /staffs/{argID} [get]
// http "http://localhost:8080/staffs/1" X-Api-User:user123
func GetStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "staffs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetStaffs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddStaffs add to add a single record to staffs table in the wcs database
// @Summary Add an record to staffs table
// @Description add to add a single record to staffs table in the wcs database
// @Tags Staffs
// @Accept  json
// @Produce  json
// @Param Staffs body model.Staffs true "Add Staffs"
// @Success 200 {object} model.Staffs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /staffs [post]
// echo '{"id": 22,"name": "iBVEQSDapwAoCwOickNOSsaZD","job": "bCgVsWmQGGBqdXeGyTsemysfU","intro": "yXdkLRhXGVoatuacYYZPImjBd","avatar": "EcFdqBBRjJKCmuekeSswVwbWx"}' | http POST "http://localhost:8080/staffs" X-Api-User:user123
func AddStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	staffs := &model.Staffs{}

	if err := readJSON(r, staffs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := staffs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	staffs.Prepare()

	if err := staffs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "staffs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	staffs, _, err = dao.AddStaffs(ctx, staffs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, staffs)
}

// UpdateStaffs Update a single record from staffs table in the wcs database
// @Summary Update an record in table staffs
// @Description Update a single record from staffs table in the wcs database
// @Tags Staffs
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Staffs body model.Staffs true "Update Staffs record"
// @Success 200 {object} model.Staffs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /staffs/{argID} [put]
// echo '{"id": 22,"name": "iBVEQSDapwAoCwOickNOSsaZD","job": "bCgVsWmQGGBqdXeGyTsemysfU","intro": "yXdkLRhXGVoatuacYYZPImjBd","avatar": "EcFdqBBRjJKCmuekeSswVwbWx"}' | http PUT "http://localhost:8080/staffs/1"  X-Api-User:user123
func UpdateStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	staffs := &model.Staffs{}
	if err := readJSON(r, staffs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := staffs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	staffs.Prepare()

	if err := staffs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "staffs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	staffs, _, err = dao.UpdateStaffs(ctx,
		argID,
		staffs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, staffs)
}

// DeleteStaffs Delete a single record from staffs table in the wcs database
// @Summary Delete a record from staffs
// @Description Delete a single record from staffs table in the wcs database
// @Tags Staffs
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Staffs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /staffs/{argID} [delete]
// http DELETE "http://localhost:8080/staffs/1" X-Api-User:user123
func DeleteStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "staffs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteStaffs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
