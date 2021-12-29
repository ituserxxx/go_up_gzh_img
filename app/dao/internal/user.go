// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserDao is the manager for logic model data accessing and custom defined data operations functions management.
type UserDao struct {
	gmvc.M                                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C       userColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.	
	DB      gdb.DB                              // DB is the raw underlying database management object.
	Table   string                              // Table is the underlying table name of the DAO.
}

// UserColumns defines and stores column names for table user.
type userColumns struct {
	Id          string //   
    Nickname    string //   
    Openid      string //   
    SessionKey  string //   
    AvatarUrl   string //   
    CreateTime  string //
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
    columns := userColumns{
		Id:         "id",           
            Nickname:   "nickname",     
            Openid:     "openid",       
            SessionKey: "session_key",  
            AvatarUrl:  "avatar_url",   
            CreateTime: "create_time",
	}
	return &UserDao{
		C: 	   columns,
		M:     g.DB("default").Model("user").Safe(),
		DB:    g.DB("default"),
		Table: "user",
	}
}