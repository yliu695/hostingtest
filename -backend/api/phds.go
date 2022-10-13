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

func configPhdsRouter(router *httprouter.Router) {
	router.GET("/phds", GetAllPhds)
	router.POST("/phds", AddPhds)
	router.GET("/phds/:argID", GetPhds)
	router.PUT("/phds/:argID", UpdatePhds)
	router.DELETE("/phds/:argID", DeletePhds)
}

func configGinPhdsRouter(router gin.IRoutes) {
	router.GET("/phds", ConverHttprouterToGin(GetAllPhds))
	router.POST("/phds", ConverHttprouterToGin(AddPhds))
	router.GET("/phds/:argID", ConverHttprouterToGin(GetPhds))
	router.PUT("/phds/:argID", ConverHttprouterToGin(UpdatePhds))
	router.DELETE("/phds/:argID", ConverHttprouterToGin(DeletePhds))
}

// GetAllPhds is a function to get a slice of record(s) from phds table in the wcs database
// @Summary Get list of Phds
// @Tags Phds
// @Description GetAllPhds is a handler to get a slice of record(s) from phds table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Phds}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /phds [get]
// http "http://localhost:8080/phds?page=0&pagesize=20" X-Api-User:user123
func GetAllPhds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "phds", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPhds(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPhds is a function to get a single record from the phds table in the wcs database
// @Summary Get record from table Phds by  argID
// @Tags Phds
// @ID argID
// @Description GetPhds is a function to get a single record from the phds table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Phds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /phds/{argID} [get]
// http "http://localhost:8080/phds/1" X-Api-User:user123
func GetPhds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "phds", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPhds(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPhds add to add a single record to phds table in the wcs database
// @Summary Add an record to phds table
// @Description add to add a single record to phds table in the wcs database
// @Tags Phds
// @Accept  json
// @Produce  json
// @Param Phds body model.Phds true "Add Phds"
// @Success 200 {object} model.Phds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /phds [post]
// echo '{"id": 51,"name": "iptrUeYYumwGPNtqxbwXxhAXn","job": "TmSLqCHJTseAlgRrANZBjBvHw","intro": "FnfwXHBltWDjbebJfRnbCWXns","avatar": "OuiNsicSMHnSSfyaHXkVcbBTC"}' | http POST "http://localhost:8080/phds" X-Api-User:user123
func AddPhds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	phds := &model.Phds{}

	if err := readJSON(r, phds); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := phds.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	phds.Prepare()

	if err := phds.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "phds", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	phds, _, err = dao.AddPhds(ctx, phds)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, phds)
}

// UpdatePhds Update a single record from phds table in the wcs database
// @Summary Update an record in table phds
// @Description Update a single record from phds table in the wcs database
// @Tags Phds
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Phds body model.Phds true "Update Phds record"
// @Success 200 {object} model.Phds
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /phds/{argID} [put]
// echo '{"id": 51,"name": "iptrUeYYumwGPNtqxbwXxhAXn","job": "TmSLqCHJTseAlgRrANZBjBvHw","intro": "FnfwXHBltWDjbebJfRnbCWXns","avatar": "OuiNsicSMHnSSfyaHXkVcbBTC"}' | http PUT "http://localhost:8080/phds/1"  X-Api-User:user123
func UpdatePhds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	phds := &model.Phds{}
	if err := readJSON(r, phds); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := phds.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	phds.Prepare()

	if err := phds.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "phds", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	phds, _, err = dao.UpdatePhds(ctx,
		argID,
		phds)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, phds)
}

// DeletePhds Delete a single record from phds table in the wcs database
// @Summary Delete a record from phds
// @Description Delete a single record from phds table in the wcs database
// @Tags Phds
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Phds
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /phds/{argID} [delete]
// http DELETE "http://localhost:8080/phds/1" X-Api-User:user123
func DeletePhds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "phds", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePhds(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
