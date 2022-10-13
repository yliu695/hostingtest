package dao

import (
	"context"
	"time"

	"wcs/model"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllAdmin is a function to get a slice of record(s) from admin table in the wcs database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAdmin(ctx context.Context, page, pagesize int64, order string) (results []*model.Admin, totalRows int, err error) {

	resultOrm := DB.Model(&model.Admin{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetAdmin is a function to get a single record from the admin table in the wcs database
// error - ErrNotFound, db Find error
func GetAdmin(ctx context.Context, argID int32) (record *model.Admin, err error) {
	record = &model.Admin{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// GetAdminByName is a function to get a single record by name from the admin table in the wcs database
// error - ErrNotFound, db Find error
func GetAdminByName(ctx context.Context, userName string) (record *model.Admin, err error) {
	record = &model.Admin{}
	if err = DB.First(record, "username = ?", userName).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAdmin is a function to add a single record to admin table in the wcs database
// error - ErrInsertFailed, db save call failed
func AddAdmin(ctx context.Context, record *model.Admin) (result *model.Admin, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAdmin is a function to update a single record from admin table in the wcs database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAdmin(ctx context.Context, argID int32, updated *model.Admin) (result *model.Admin, RowsAffected int64, err error) {

	result = &model.Admin{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteAdmin is a function to delete a single record from admin table in the wcs database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAdmin(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Admin{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
