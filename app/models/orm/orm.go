package orm

/*
 * 	Custom ORM package
 * 	Based on gorm ORM.
 * 	Make sure the every database actions within the transaction.
 * 	Do begin, rollback and commit the db transaction automatically.
 * 	Return proper error
 *
*/

import (
	"github.com/jinzhu/gorm"
	cgorm "mio/db/gorm"
	"database/sql"
	"fmt"
)

type (
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error
)

type (
	DBFunc2 func(tx *gorm.DB) (*sql.Rows, error) // func type which accept *gorm.DB and return error
)

// Create
// Helper function to insert gorm model to database by using 'WithinTransaction'
func Create(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !cgorm.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback()                                                                                                           // rollback
			return err
		}
		return err
	})
}

// Save
// Helper function to save gorm model to database by using 'WithinTransaction'
func Save(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if cgorm.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Save(v).Error; err != nil {
			tx.Rollback()                                                                                                       // rollback
			return err
		}
		return err
	})
}

// Delete
// Helper function to save gorm model to database by using 'WithinTransaction'
func Delete(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			tx.Rollback()                                                                                                           // rollback
			return err
		}
		return err
	})
}

//m, err := tx.Table("usuarios").Where("userid = ? and pass = ?", userid, pass).Select("firstname, pass2").Rows()

// FindOneByID
// Helper function to find a record by using 'WithinTransaction'
func FindOneByID(v interface{}, id uint64) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Last(v, id).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

func Mierda(userid, pass string) (err error) {
	fmt.Println()


	db:= cgorm.DBManager()
	rows:= &sql.Rows{}
	rows, err = db.Table("usuarios").Where("userid = ? and pass = ?", userid, pass).Select("firstname, pass2").Rows()
	

	for rows.Next() {
		var firstname string
		var pass2 string
		
		rows.Scan(&firstname, &pass2)
		fmt.Println(firstname, pass2)
	}



	return err;

	// return WithinTransaction(func(tx *gorm.DB) (error) {
	// 	_, err := tx.Table("usuarios").Where("userid = ? and pass = ?", userid, pass).Select("firstname, pass2").Rows()
		
	// 	if err != nil {
	// 		tx.Rollback() // rollback db transaction
	// 		return err
	// 	}
		
	// 	fmt.Println(err)

	// 	return nil
	// })
}

func FindOneByFields(t, v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(t).Last(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindAll
// Helper function to find records by using 'WithinTransaction'
func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindOneByQuery
// Helper function to find a record by using 'WithinTransaction'
func FindOneByQuery(v interface{}, params map[string]interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Last(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindByQuery
// Helper function to find records by using 'WithinTransaction'
func FindByQuery(v interface{}, params map[string]interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// WithinTransaction
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func WithinTransaction(fn DBFunc) (err error) {
	tx := cgorm.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

func WithinTransaction2(fn DBFunc2) (err error) {
	tx := cgorm.DBManager().Begin() // start db transaction
	defer tx.Commit()
	_, err = fn(tx)
	// close db transaction
	return err
}

