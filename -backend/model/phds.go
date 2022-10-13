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


CREATE TABLE `phds` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `job` varchar(512) NOT NULL,
  `intro` longtext NOT NULL,
  `avatar` varchar(512) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 51,    "name": "iptrUeYYumwGPNtqxbwXxhAXn",    "job": "TmSLqCHJTseAlgRrANZBjBvHw",    "intro": "FnfwXHBltWDjbebJfRnbCWXns",    "avatar": "OuiNsicSMHnSSfyaHXkVcbBTC"}



*/

// Phds struct is a row record of the phds table in the wcs database
type Phds struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 1] name                                           varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	Name string `gorm:"column:name;type:varchar;size:256;" json:"name"`
	//[ 2] job                                            varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Job string `gorm:"column:job;type:varchar;size:512;" json:"job"`
	//[ 3] intro                                          text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Intro string `gorm:"column:intro;type:text;size:4294967295;" json:"intro"`
	//[ 4] avatar                                         varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Avatar string `gorm:"column:avatar;type:varchar;size:512;" json:"avatar"`
}

var phdsTableInfo = &TableInfo{
	Name: "phds",
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
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "job",
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
			GoFieldName:        "Job",
			GoFieldType:        "string",
			JSONFieldName:      "job",
			ProtobufFieldName:  "job",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
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
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "avatar",
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
			GoFieldName:        "Avatar",
			GoFieldType:        "string",
			JSONFieldName:      "avatar",
			ProtobufFieldName:  "avatar",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Phds) TableName() string {
	return "phds"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Phds) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Phds) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Phds) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Phds) TableInfo() *TableInfo {
	return phdsTableInfo
}
