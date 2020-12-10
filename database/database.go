package database

import (
	"fmt"
	"github.com/iwandede/go-via/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

/* -------------------------------------------------------------------------- */
/*                              INIT CONFIG TO DB                             */
/* -------------------------------------------------------------------------- */
func DataStore(conf *config.Config) (*sqlx.DB, error) {
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Database, conf.DB.Password, conf.DB.SSL)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	go func(connection string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second}
		for {
			time.Sleep(60 * time.Second)
			if e := db.Ping(); e != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					e2 := RetryHandler(3, func() (bool, error) {
						var e error
						db, e = sqlx.Open(conf.DB.Driver, connection)
						if e != nil {
							return false, e
						}
						return true, nil
					})
					if e2 != nil {
						fmt.Println(e.Error())
						time.Sleep(intervals[i])
						if i == len(intervals)-1 {
							i--
						}
						continue
					}
					break L
				}

			}
		}
	}(connection)
	return db, nil
}

/*func DataStore(conf *config.Config) (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Database, conf.DB.Password, conf.DB.SSL)
	db, err := gorm.Open(conf.DB.Driver, connection)
	if err != nil {
		return nil, fmt.Errorf("Error : %v", err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(30)
	//db.AutoMigrate(&models.Service{})
	db.AutoMigrate(&models.ThridParty{})
	go func(connection string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second}
		for {
			time.Sleep(60 * time.Second)
			if e := db.DB().Ping(); e != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					e2 := RetryHandler(3, func() (bool, error) {
						var e error
						db, e = gorm.Open(conf.DB.Driver, connection)
						if e != nil {
							return false, e
						}
						return true, nil
					})
					if e2 != nil {
						fmt.Println(e.Error())
						time.Sleep(intervals[i])
						if i == len(intervals)-1 {
							i--
						}
						continue
					}
					break L
				}

			}
		}
	}(connection)
	//defer db.Close()
	return db, nil
}*/

func RetryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	}
	if n-1 > 0 {
		return RetryHandler(n-1, f)
	}
	return er
}
