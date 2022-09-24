package utilities

import (
	"github.com/jinzhu/gorm"
)

var orm *gorm.DB

type DBCreds struct {
	dialect string
	dbInfo  interface{}
}

func (e *DBCreds) GetDialect() string {
	return e.dialect
}

func (e *DBCreds) SetDialect(dialect string) {
	e.dialect = dialect
}

func (e *DBCreds) GetDbInfo() interface{} {
	return e.dbInfo
}

func (e *DBCreds) SetDbInfo(dbInfo interface{}) {
	e.dbInfo = dbInfo
}

func GetDBInstance(creds *DBCreds) (*gorm.DB, error) {
	var err error
	if orm == nil {
		orm, err = gorm.Open(creds.GetDialect(), creds.GetDbInfo())
		if err != nil {
			return nil, err
		}
	}

	return orm, nil
}
