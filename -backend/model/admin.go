package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(128) DEFAULT NULL COMMENT 'user name',
  `password` varchar(256) DEFAULT NULL COMMENT 'password',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='admin of system'

JSON Sample
-------------------------------------
{    "id": 33,    "username": "LhvvfhYxiPROoEpSrkwbwEqIo",    "password": "KOadtAHGFhoOiEsTuEKPHDqbd"}



*/

// Admin struct is a row record of the admin table in the wcs database
type Admin struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"` // id
	//[ 1] username                                       varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	Username sql.NullString `gorm:"column:username;type:varchar;size:128;" json:"username"` // user name
	//[ 2] password                                       varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	Password sql.NullString `gorm:"column:password;type:varchar;size:256;" json:"password"` // password

}

var adminTableInfo = &TableInfo{
	Name: "admin",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "username",
			Comment:            `user name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Username",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "password",
			Comment:            `password`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Password",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Admin) TableName() string {
	return "admin"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Admin) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Admin) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Admin) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Admin) TableInfo() *TableInfo {
	return adminTableInfo
}
