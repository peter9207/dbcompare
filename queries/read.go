package queries

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
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

func (db *SQLRunner) PerformRead() (err error) {

	rows, err := db.db.Query("select * from items limit 20")
	if err != nil {
		return
	}
	defer rows.Close()

	// res = []int64{}
	var r int64

	for rows.Next() {
		rows.Scan(&r)

		// res = append(res, r)
	}

	return
}

func (db *SQLRunner) listWarehouse() (id int64, err error) {

	rows, err := db.db.Query("select id from warehouses")
	if err != nil {
		return
	}
	defer rows.Close()

	res := []int64{}
	var r int64

	for rows.Next() {
		err = rows.Scan(&r)
		if err != nil {
			return
		}
		res = append(res, r)
	}

	id = res[rand.Intn(len(res))]
	return
}

func (db *SQLRunner) randomItemType() (id int64, err error) {

	rows, err := db.db.Query("select id from item_types LIMIT 50")
	if err != nil {
		return
	}
	defer rows.Close()

	res := []int64{}
	var r int64

	for rows.Next() {
		err = rows.Scan(&r)
		if err != nil {
			return
		}
		res = append(res, r)
	}

	id = res[rand.Intn(len(res))]
	return
}

func (db *SQLRunner) randomName() string {
	return fmt.Sprintf("%d", db.rand.Int63())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (db *SQLRunner) Setup() (err error) {

	for i := 0; i < 10; i++ {
		name := randStringRunes(10)
		_, err = db.db.Exec("INSERT INTO warehouses ( name) VALUES ($1)", fmt.Sprintf("warehouse-%s", name))
		if err != nil {
			return
		}
	}

	for i := 0; i < 10; i++ {
		typeName := randStringRunes(10)
		_, err = db.db.Exec("INSERT INTO item_types ( name) VALUES ($1)", fmt.Sprintf("type-%s", typeName))

	}

	return
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
	query := "INSERT INTO items (name, item_type_id, warehouse_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.db.Exec(query, name, itemType, warehouseID, time.Now(), time.Now())
	if err != nil {
		return
	}

	return
}
