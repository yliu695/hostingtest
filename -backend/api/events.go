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

func configEventsRouter(router *httprouter.Router) {
	router.GET("/events", GetAllEvents)
	router.POST("/events", AddEvents)
	router.GET("/events/:argID", GetEvents)
	router.PUT("/events/:argID", UpdateEvents)
	router.DELETE("/events/:argID", DeleteEvents)
}

func configGinEventsRouter(router gin.IRoutes) {
	router.GET("/events", ConverHttprouterToGin(GetAllEvents))
	router.POST("/events", ConverHttprouterToGin(AddEvents))
	router.GET("/events/:argID", ConverHttprouterToGin(GetEvents))
	router.PUT("/events/:argID", ConverHttprouterToGin(UpdateEvents))
	router.DELETE("/events/:argID", ConverHttprouterToGin(DeleteEvents))
}

// GetAllEvents is a function to get a slice of record(s) from events table in the wcs database
// @Summary Get list of Events
// @Tags Events
// @Description GetAllEvents is a handler to get a slice of record(s) from events table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Events}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /events [get]
// http "http://localhost:8080/events?page=0&pagesize=20" X-Api-User:user123
func GetAllEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "events", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllEvents(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetEvents is a function to get a single record from the events table in the wcs database
// @Summary Get record from table Events by  argID
// @Tags Events
// @ID argID
// @Description GetEvents is a function to get a single record from the events table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Events
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /events/{argID} [get]
// http "http://localhost:8080/events/1" X-Api-User:user123
func GetEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "events", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetEvents(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddEvents add to add a single record to events table in the wcs database
// @Summary Add an record to events table
// @Description add to add a single record to events table in the wcs database
// @Tags Events
// @Accept  json
// @Produce  json
// @Param Events body model.Events true "Add Events"
// @Success 200 {object} model.Events
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /events [post]
// echo '{"cover": "PXvcMZAaVtykxdkaiPnFcLfhu","tags": "dymRLWPnwryPHEVWAKyjptSUC","update_time": "2208-11-03T15:07:41.739514237+08:00","create_time": "2057-01-07T00:22:49.758092343+08:00","content": "NctfhDQebYWmpAGapMOhaLiCk","title": "BVtvmnvgRGjXXVYoTuIjUZPqV","id": 35,"event_time": 55}' | http POST "http://localhost:8080/events" X-Api-User:user123
func AddEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	events := &model.Events{}

	if err := readJSON(r, events); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := events.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	events.Prepare()

	if err := events.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "events", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	events, _, err = dao.AddEvents(ctx, events)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, events)
}

// UpdateEvents Update a single record from events table in the wcs database
// @Summary Update an record in table events
// @Description Update a single record from events table in the wcs database
// @Tags Events
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Events body model.Events true "Update Events record"
// @Success 200 {object} model.Events
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /events/{argID} [put]
// echo '{"cover": "PXvcMZAaVtykxdkaiPnFcLfhu","tags": "dymRLWPnwryPHEVWAKyjptSUC","update_time": "2208-11-03T15:07:41.739514237+08:00","create_time": "2057-01-07T00:22:49.758092343+08:00","content": "NctfhDQebYWmpAGapMOhaLiCk","title": "BVtvmnvgRGjXXVYoTuIjUZPqV","id": 35,"event_time": 55}' | http PUT "http://localhost:8080/events/1"  X-Api-User:user123
func UpdateEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	events := &model.Events{}
	if err := readJSON(r, events); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := events.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	events.Prepare()

	if err := events.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "events", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	events, _, err = dao.UpdateEvents(ctx,
		argID,
		events)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, events)
}

// DeleteEvents Delete a single record from events table in the wcs database
// @Summary Delete a record from events
// @Description Delete a single record from events table in the wcs database
// @Tags Events
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Events
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /events/{argID} [delete]
// http DELETE "http://localhost:8080/events/1" X-Api-User:user123
func DeleteEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "events", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteEvents(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
