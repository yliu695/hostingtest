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

func configNewsRouter(router *httprouter.Router) {
	router.GET("/news", GetAllNews)
	router.POST("/news", AddNews)
	router.GET("/news/:argID", GetNews)
	router.PUT("/news/:argID", UpdateNews)
	router.DELETE("/news/:argID", DeleteNews)
}

func configGinNewsRouter(router gin.IRoutes) {
	router.GET("/news", ConverHttprouterToGin(GetAllNews))
	router.POST("/news", ConverHttprouterToGin(AddNews))
	router.GET("/news/:argID", ConverHttprouterToGin(GetNews))
	router.PUT("/news/:argID", ConverHttprouterToGin(UpdateNews))
	router.DELETE("/news/:argID", ConverHttprouterToGin(DeleteNews))
}

// GetAllNews is a function to get a slice of record(s) from news table in the wcs database
// @Summary Get list of News
// @Tags News
// @Description GetAllNews is a handler to get a slice of record(s) from news table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.News}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /news [get]
// http "http://localhost:8080/wcs?page=0&pagesize=20" X-Api-User:user123
func GetAllNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	//might havr to h=change this
	if err := ValidateRequest(ctx, r, "news", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllNews(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetNews is a function to get a single record from the news table in the wcs database
// @Summary Get record from table News by  argID
// @Tags News
// @ID argID
// @Description GetNews is a function to get a single record from the news table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.News
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /wcs/{argID} [get]
// http "http://localhost:8080/wcs/1" X-Api-User:user123
func GetNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}
	//might have to change this
	if err := ValidateRequest(ctx, r, "news", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetNews(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddNews add to add a single record to news table in the wcs database
// @Summary Add an record to news table
// @Description add to add a single record to news table in the wcs database
// @Tags News
// @Accept  json
// @Produce  json
// @Param News body model.News true "Add News"
// @Success 200 {object} model.News
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /news [post]
// echo '{"id": 76,"title": "cGSfstycEikZjCWZYJhEuWwWC","content": "YkQrALQfQfpqwiSLOctWUkrOs","create_time": "2311-07-11T12:25:43.563373812+08:00","update_time": "2177-04-07T01:41:28.623684615+08:00","tags": "erGBbmpFXmOvpNmVsAOvLkCuF","cover": "kljHXlIKVdfpvdiQDEksfgyqH"}' | http POST "http://localhost:8080/news" X-Api-User:user123
func AddNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	news := &model.News{}
	//might need to change this
	if err := readJSON(r, news); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := news.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	news.Prepare()

	if err := news.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "news", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	news, _, err = dao.AddNews(ctx, news)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, news)
}

// UpdateNews Update a single record from news table in the wcs database
// @Summary Update an record in table news
// @Description Update a single record from news table in the wcs database
// @Tags News
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  News body model.News true "Update News record"
// @Success 200 {object} model.News
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /news/{argID} [put]
// echo '{"id": 76,"title": "cGSfstycEikZjCWZYJhEuWwWC","content": "YkQrALQfQfpqwiSLOctWUkrOs","create_time": "2311-07-11T12:25:43.563373812+08:00","update_time": "2177-04-07T01:41:28.623684615+08:00","tags": "erGBbmpFXmOvpNmVsAOvLkCuF","cover": "kljHXlIKVdfpvdiQDEksfgyqH"}' | http PUT "http://localhost:8080/news/1"  X-Api-User:user123
func UpdateNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	news := &model.News{}
	if err := readJSON(r, news); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := news.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	news.Prepare()

	if err := news.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "news", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	news, _, err = dao.UpdateNews(ctx,
		argID,
		news)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, news)
}

// DeleteNews Delete a single record from news table in the wcs database
// @Summary Delete a record from news
// @Description Delete a single record from news table in the wcs database
// @Tags News
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.News
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /news/{argID} [delete]
// http DELETE "http://localhost:8080/news/1" X-Api-User:user123
func DeleteNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "news", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteNews(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
