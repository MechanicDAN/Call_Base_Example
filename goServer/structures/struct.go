package structures

import (
	"time"
)

type User struct {
	Id        int      `json:"Id" pg:"id,pk"`
	Name      string   `json:"Name" pg:"name_p,notnull"`
	Phone     string   `json:"Phone" pg:"phone_p"`
	Login 	  string   `json:"Login" pg:"login_p"`
	Password  string   `json:"-" pg:"password_p"`
	Archive   bool     `json:"-" pg:"archive_p,default:false, notnull"`
	Language  string   `json:"Language" pg:"language_p"`
	tableName struct{} `pg:"user_t,discard_unknown_columns"`
}
type Call struct {
	Id             int       `json:"Id" pg:"id,pk:user_t.id"`
	UserId         int    	 `json:"User_id" pg:"user_p,fk:user_id"`
	UserName	   string 	 `json:"Name" pg:"name_p"`
	From           string    `json:"From" pg:"from_p"`
	To             string    `json:"To" pg:"to_p"`
	StartTimestamp time.Time `json:"Start_timestamp" pg:"start_timestamp_p"`
	EndTimestamp   time.Time `json:"End_timestamp" pg:"end_timestamp_p"`
	Duration       float32   `json:"Duration" pg:"duration_p"`
	tableName      struct{}  `pg:"calls,discard_unknown_columns"`
}
type Schemata struct {
	CatalogName string	`pg:"catalog_name" ,json:"cName"`
	SchemaName 	string 	`pg:"schema_name" ,json:"sName"`
	SchemaOwner string 	`pg:"schema_owner" ,json:"oName"`
	tableName 	struct{}`pg:",discard_unknown_columns"`
}
