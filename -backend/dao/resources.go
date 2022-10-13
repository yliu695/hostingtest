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

// GetAllResources is a function to get a slice of record(s) from resources table in the wcs database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllResources(ctx context.Context, page, pagesize int64, order string) (results []*model.Resources, totalRows int, err error) {

	resultOrm := DB.Model(&model.Resources{})
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

// GetResources is a function to get a single record from the resources table in the wcs database
// error - ErrNotFound, db Find error
func GetResources(ctx context.Context, argID int32) (record *model.Resources, err error) {
	record = &model.Resources{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddResources is a function to add a single record to resources table in the wcs database
// error - ErrInsertFailed, db save call failed
func AddResources(ctx context.Context, record *model.Resources) (result *model.Resources, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateResources is a function to update a single record from resources table in the wcs database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateResources(ctx context.Context, argID int32, updated *model.Resources) (result *model.Resources, RowsAffected int64, err error) {

	result = &model.Resources{}
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

// DeleteResources is a function to delete a single record from resources table in the wcs database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteResources(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Resources{}
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
