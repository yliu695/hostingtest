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


CREATE TABLE `events` (
  `cover` varchar(512) NOT NULL DEFAULT '',
  `tags` varchar(1024) NOT NULL DEFAULT '',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `content` longtext NOT NULL,
  `title` varchar(512) NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `event_time` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "cover": "PXvcMZAaVtykxdkaiPnFcLfhu",    "tags": "dymRLWPnwryPHEVWAKyjptSUC",    "update_time": "2208-11-03T15:07:41.739514237+08:00",    "create_time": "2057-01-07T00:22:49.758092343+08:00",    "content": "NctfhDQebYWmpAGapMOhaLiCk",    "title": "BVtvmnvgRGjXXVYoTuIjUZPqV",    "id": 35,    "event_time": 55}



*/

// Events struct is a row record of the events table in the wcs database
type Events struct {
	//[ 0] cover                                          varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Cover string `gorm:"column:cover;type:varchar;size:512;" json:"cover"`
	//[ 1] tags                                           varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	Tags string `gorm:"column:tags;type:varchar;size:1024;" json:"tags"`
	//[ 2] update_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"update_time"`
	//[ 3] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"create_time"`
	//[ 4] content                                        text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Content string `gorm:"column:content;type:text;size:4294967295;" json:"content"`
	//[ 5] title                                          varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	Title string `gorm:"column:title;type:varchar;size:512;" json:"title"`
	//[ 6] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	//[ 7] event_time                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	EventTime int64 `gorm:"column:event_time;type:bigint;default:0;" json:"event_time"`
}

var eventsTableInfo = &TableInfo{
	Name: "events",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "cover",
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
			GoFieldName:        "Cover",
			GoFieldType:        "string",
			JSONFieldName:      "cover",
			ProtobufFieldName:  "cover",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "tags",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "Tags",
			GoFieldType:        "string",
			JSONFieldName:      "tags",
			ProtobufFieldName:  "tags",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "create_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "time.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "content",
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
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "title",
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
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "event_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "EventTime",
			GoFieldType:        "int64",
			JSONFieldName:      "event_time",
			ProtobufFieldName:  "event_time",
			ProtobufType:       "int64",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Events) TableName() string {
	return "events"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Events) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Events) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Events) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Events) TableInfo() *TableInfo {
	return eventsTableInfo
}
