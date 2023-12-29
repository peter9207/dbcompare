package queries

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type Runner interface {
	PerformRead() error
	PerformWrite() error
}

type SQLRunner struct {
	db   *sql.DB
	rand rand.Source
}

func NewRunner(url string) (db *SQLRunner, err error) {

	conn, err := sql.Open("postgres", url)

	db = &SQLRunner{
		db:   conn,
		rand: rand.New(rand.NewSource(time.Now().Unix())),
	}

	return
}

func (db *SQLRunner) PerformRead() (res []int64, err error) {

	rows, err := db.db.Query("select * from items limit 20")
	if err != nil {
		return
	}
	defer rows.Close()

	res = []int64{}
	var r int64

	for rows.Next() {
		rows.Scan(&r)

		res = append(res, r)
	}

	return
}

func (db *SQLRunner) listWarehouse() (id int64, err error) {

	rows, err := db.db.Query("select * from warehouses")
	if err != nil {
		return
	}
	defer rows.Close()

	res := []int64{}
	var r int64

	for rows.Next() {
		rows.Scan(&r)
		res = append(res, r)
	}

	id = res[rand.Intn(len(res))]
	return
}

func (db *SQLRunner) randomItemType() (id int64, err error) {

	rows, err := db.db.Query("select * from item_types")
	if err != nil {
		return
	}
	defer rows.Close()

	res := []int64{}
	var r int64

	for rows.Next() {
		rows.Scan(&r)
		res = append(res, r)
	}

	id = res[rand.Intn(len(res))]
	return
}

func (db *SQLRunner) randomName() string {
	return fmt.Sprintf("%d", db.rand.Int63())
}

func (db *SQLRunner) PerformWrite() (err error) {

	warehouseID, err := db.listWarehouse()
	if err != nil {
		return
	}
	name := db.randomName()

	itemType, err := db.randomItemType()
	if err != nil {
		return
	}

	_, err = db.db.Exec("INSERT INTO items (name, item_type_id, warehouse_id)", name, itemType, warehouseID)
	if err != nil {
		return
	}

	return
}
