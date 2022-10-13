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


CREATE TABLE `resources` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(512) NOT NULL,
  `intro` longtext NOT NULL,
  `link` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 57,    "name": "UvJpIyJJkjDDfKyCmQxdwbcwA",    "intro": "yXkmMtyeYBsJGFUtZqshRSFLB",    "link": "TIPRoVvioBUWsOXXxvWxmYIMY"}



*/

// Resources struct is a row record of the resources table in the wcs database
type Resources struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Name string `gorm:"column:name;type:varchar;size:512;" json:"name"`
	//[ 2] intro                                          text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Intro string `gorm:"column:intro;type:text;size:4294967295;" json:"intro"`
	//[ 3] link                                           varchar(128)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	Link string `gorm:"column:link;type:varchar;size:128;" json:"link"`
}

var resourcesTableInfo = &TableInfo{
	Name: "resources",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "intro",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(4294967295)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       4294967295,
			GoFieldName:        "Intro",
			GoFieldType:        "string",
			JSONFieldName:      "intro",
			ProtobufFieldName:  "intro",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "link",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Link",
			GoFieldType:        "string",
			JSONFieldName:      "link",
			ProtobufFieldName:  "link",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *Resources) TableName() string {
	return "resources"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *Resources) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *Resources) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *Resources) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *Resources) TableInfo() *TableInfo {
	return resourcesTableInfo
}
