package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func InitProgreSQLDBPool(path string) error {
	//普通的db方式
	DbConfMap := &PostgrogreSQLMapConf{}
	err := ParseConfig(path, DbConfMap)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if len(DbConfMap.List) == 0 {
		fmt.Printf("[INFO] %s%s\n", time.Now().Format(TimeFormat), " empty mysql config.")
	}

	//DBMapPool = map[string]*sql.DB{}
	//GORMMapPool = map[string]*gorm.DB{}
	fmt.Println(DbConfMap)
	for confName, DbConf := range DbConfMap.List {
		fmt.Println(":::::::", confName)

		dbpool, err := sql.Open("postgres", DbConf.DataSourceName)
		if err != nil {
			fmt.Printf("链接数据失败：%v\n", err.Error())
			return err
		}

		dbpool.SetMaxOpenConns(DbConf.MaxOpenConn)
		dbpool.SetMaxIdleConns(DbConf.MaxIdleConn)
		dbpool.SetConnMaxLifetime(time.Duration(DbConf.MaxConnLifeTime) * time.Second)
		err = dbpool.Ping()
		if err != nil {
			return err
		}

		//gorm连接方式

		dbGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: dbpool}), &gorm.Config{
			//Logger: &DefaultMysqlGormLogger,
		})
		if err != nil {
			return err
		}
		DBMapPool[confName] = dbpool
		GORMMapPool[confName] = dbGorm
	}

	//手动配置连接
	//if dbpool, err := GetDBPool("postgres"); err == nil {
	//	DBDefaultPool = dbpool
	//}
	//if dbpool, err := GetGormPool("postgres"); err == nil {
	//	GORMDefaultPool = dbpool
	//}
	return nil
}
