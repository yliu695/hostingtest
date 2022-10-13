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

func configProjectsRouter(router *httprouter.Router) {
	router.GET("/projects", GetAllProjects)
	router.POST("/projects", AddProjects)
	router.GET("/projects/:argID", GetProjects)
	router.PUT("/projects/:argID", UpdateProjects)
	router.DELETE("/projects/:argID", DeleteProjects)
}

func configGinProjectsRouter(router gin.IRoutes) {
	router.GET("/projects", ConverHttprouterToGin(GetAllProjects))
	router.POST("/projects", ConverHttprouterToGin(AddProjects))
	router.GET("/projects/:argID", ConverHttprouterToGin(GetProjects))
	router.PUT("/projects/:argID", ConverHttprouterToGin(UpdateProjects))
	router.DELETE("/projects/:argID", ConverHttprouterToGin(DeleteProjects))
}

// GetAllProjects is a function to get a slice of record(s) from projects table in the wcs database
// @Summary Get list of Projects
// @Tags Projects
// @Description GetAllProjects is a handler to get a slice of record(s) from projects table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Projects}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /projects [get]
// http "http://localhost:8080/projects?page=0&pagesize=20" X-Api-User:user123
func GetAllProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "projects", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllProjects(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetProjects is a function to get a single record from the projects table in the wcs database
// @Summary Get record from table Projects by  argID
// @Tags Projects
// @ID argID
// @Description GetProjects is a function to get a single record from the projects table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Projects
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /projects/{argID} [get]
// http "http://localhost:8080/projects/1" X-Api-User:user123
func GetProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "projects", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetProjects(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddProjects add to add a single record to projects table in the wcs database
// @Summary Add an record to projects table
// @Description add to add a single record to projects table in the wcs database
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param Projects body model.Projects true "Add Projects"
// @Success 200 {object} model.Projects
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /projects [post]
// echo '{"id": 19,"name": "FQQabeNJaBNUtMTGgDPyyvDIJ","intro": "DhNHNTSdVPtKtMshMfUjnoGKa","link": "AabjFhdLeHVFKapYMEPumLtUU"}' | http POST "http://localhost:8080/projects" X-Api-User:user123
func AddProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	projects := &model.Projects{}

	if err := readJSON(r, projects); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := projects.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	projects.Prepare()

	if err := projects.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "projects", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	projects, _, err = dao.AddProjects(ctx, projects)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, projects)
}

// UpdateProjects Update a single record from projects table in the wcs database
// @Summary Update an record in table projects
// @Description Update a single record from projects table in the wcs database
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Projects body model.Projects true "Update Projects record"
// @Success 200 {object} model.Projects
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /projects/{argID} [put]
// echo '{"id": 19,"name": "FQQabeNJaBNUtMTGgDPyyvDIJ","intro": "DhNHNTSdVPtKtMshMfUjnoGKa","link": "AabjFhdLeHVFKapYMEPumLtUU"}' | http PUT "http://localhost:8080/projects/1"  X-Api-User:user123
func UpdateProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	projects := &model.Projects{}
	if err := readJSON(r, projects); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := projects.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	projects.Prepare()

	if err := projects.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "projects", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	projects, _, err = dao.UpdateProjects(ctx,
		argID,
		projects)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, projects)
}

// DeleteProjects Delete a single record from projects table in the wcs database
// @Summary Delete a record from projects
// @Description Delete a single record from projects table in the wcs database
// @Tags Projects
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Projects
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /projects/{argID} [delete]
// http DELETE "http://localhost:8080/projects/1" X-Api-User:user123
func DeleteProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "projects", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProjects(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
