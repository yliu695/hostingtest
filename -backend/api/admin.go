package api

import (
	"net/http"

	"wcs/dao"
	"wcs/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configAdminRouter(router *httprouter.Router) {
	router.GET("/admin", GetAllAdmin)
	router.POST("/admin", AddAdmin)
	router.GET("/admin/:argID", GetAdmin)
	router.PUT("/admin/:argID", UpdateAdmin)
	router.DELETE("/admin/:argID", DeleteAdmin)
}

func configGinAdminRouter(router gin.IRoutes) {
	router.GET("/isAdminLogin", IsAdminLogin)
	router.POST("/adminLogin", AdminLogin)
	router.POST("/adminLogout", AdminLogout)
	router.GET("/admin", ConverHttprouterToGin(GetAllAdmin))
	router.POST("/admin", ConverHttprouterToGin(AddAdmin))
	router.GET("/admin/:argID", ConverHttprouterToGin(GetAdmin))
	router.PUT("/admin/:argID", ConverHttprouterToGin(UpdateAdmin))
	router.DELETE("/admin/:argID", ConverHttprouterToGin(DeleteAdmin))
}

// GetAllAdmin is a function to get a slice of record(s) from admin table in the wcs database
// @Summary Get list of Admin
// @Tags Admin
// @Description GetAllAdmin is a handler to get a slice of record(s) from admin table in the wcs database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Admin}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /admin [get]
// http "http://localhost:8080/admin?page=0&pagesize=20" X-Api-User:user123
func GetAllAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "admin", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAdmin(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetAdmin is a function to get a single record from the admin table in the wcs database
// @Summary Get record from table Admin by  argID
// @Tags Admin
// @ID argID
// @Description GetAdmin is a function to get a single record from the admin table in the wcs database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} model.Admin
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /admin/{argID} [get]
// http "http://localhost:8080/admin/1" X-Api-User:user123
func GetAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAdmin(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

func IsAdminLogin(c *gin.Context) {
	session := sessions.Default(c)
	isLogin := session.Get("currentAdmin") != nil
	c.JSON(200, gin.H{
		"isLogin": isLogin,
	})
}

func AdminLogin(c *gin.Context) {
	userName, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")

	admin, err := dao.GetAdminByName(c, userName)
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}

	if admin.Password.Valid && admin.Password.String == password {
		session := sessions.Default(c)
		session.Set("currentAdmin", admin.ID)
		session.Options(sessions.Options{
			MaxAge: 3600 * 12, // 12hrs
		})
		session.Save()
		c.JSON(200, gin.H{
			"isLogin": true,
		})
		return
	}

	c.JSON(200, gin.H{
		"isLogin": false,
	})
}

func AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("currentAdmin")
	session.Options(sessions.Options{
		MaxAge: 3600 * 12, // 12hrs
	})
	session.Save()

	c.JSON(200, gin.H{})
}

// AddAdmin add to add a single record to admin table in the wcs database
// @Summary Add an record to admin table
// @Description add to add a single record to admin table in the wcs database
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param Admin body model.Admin true "Add Admin"
// @Success 200 {object} model.Admin
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /admin [post]
// echo '{"id": 33,"username": "LhvvfhYxiPROoEpSrkwbwEqIo","password": "KOadtAHGFhoOiEsTuEKPHDqbd"}' | http POST "http://localhost:8080/admin" X-Api-User:user123
func AddAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	admin := &model.Admin{}

	if err := readJSON(r, admin); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := admin.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	admin.Prepare()

	if err := admin.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	admin, _, err = dao.AddAdmin(ctx, admin)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, admin)
}

// UpdateAdmin Update a single record from admin table in the wcs database
// @Summary Update an record in table admin
// @Description Update a single record from admin table in the wcs database
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Param  Admin body model.Admin true "Update Admin record"
// @Success 200 {object} model.Admin
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /admin/{argID} [put]
// echo '{"id": 33,"username": "LhvvfhYxiPROoEpSrkwbwEqIo","password": "KOadtAHGFhoOiEsTuEKPHDqbd"}' | http PUT "http://localhost:8080/admin/1"  X-Api-User:user123
func UpdateAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	admin := &model.Admin{}
	if err := readJSON(r, admin); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := admin.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	admin.Prepare()

	if err := admin.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "admin", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	admin, _, err = dao.UpdateAdmin(ctx,
		argID,
		admin)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, admin)
}

// DeleteAdmin Delete a single record from admin table in the wcs database
// @Summary Delete a record from admin
// @Description Delete a single record from admin table in the wcs database
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 204 {object} model.Admin
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /admin/{argID} [delete]
// http DELETE "http://localhost:8080/admin/1" X-Api-User:user123
func DeleteAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt32(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "admin", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAdmin(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
