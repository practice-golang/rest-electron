package dbbooks // import "dbbooks"

import (
	"math/big"
	"strconv"

	// "github.com/cznic/ql"

	"modernc.org/ql"
)

// Book : 책 정보
type Book struct {
	ID     uint
	Title  string
	Author string
}

// CreateTable : 테이블 생성
func CreateTable(table string, dbname string) error {
	// DB 준비. 없으면 생성. ql.Open은 lock 파일 삭제 방도를 못 찾아서 일단 안 씀.
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	db.Run(ql.NewRWCtx(),
		`begin transaction;
		create table if not exists `+table+` (id bigint, title string, author string);
		commit;`)

	return nil
}

// DropTable : 테이블 드랍
func DropTable(table string, dbname string) error {
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	db.Run(ql.NewRWCtx(),
		`begin transaction;
		drop table `+table+`;
		commit;`)

	return nil
}

// InsertData : Crud
func InsertData(book *Book, table string, dbname string) error {
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	s, _, err := db.Run(ql.NewRWCtx(), `begin transaction; select id() from `+table+` order by id() desc limit 1; commit;`)
	if err != nil {
		panic(err)
	}

	lastID := uint64(1)
	lastData, _ := s[0].FirstRow()
	if len(lastData) > 0 {
		lastID = uint64(lastData[0].(int64) + 1)
	}

	_, _, err = db.Run(ql.NewRWCtx(),
		`begin transaction;`+
			`insert into `+table+` values (`+strconv.FormatUint(lastID, 10)+`, "`+book.Title+`", "`+book.Author+`")`+
			`;commit;`)
	if err != nil {
		panic(err)
	}

	return nil
}

// SelectData : cRud
func SelectData(id uint, table string, dbname string) (result []Book, err error) {
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	var where string
	if id > 0 {
		where = " where id=" + strconv.FormatUint(uint64(id), 16)
	} else {
		where = ""
	}
	order := " order by id() desc"

	s, _, err := db.Run(ql.NewRWCtx(),
		`begin transaction;`+
			`select id, title, author from `+table+where+order+
			`;commit;`)
	if err != nil {
		return nil, err
	}

	data, _ := s[0].Rows(-1, 0)

	book := Book{}
	result = []Book{}
	for _, d := range data {
		book.ID = uint(d[0].(*big.Int).Uint64())
		book.Title, book.Author = d[1].(string), d[2].(string)
		result = append(result, book)
	}

	return result, nil
}

// UpdateData : crUd
func UpdateData(book *Book, table string, dbname string) error {
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	idStr := strconv.FormatUint(uint64(book.ID), 16)

	_, _, err := db.Run(ql.NewRWCtx(),
		`begin transaction;`+
			`update `+table+` set title="`+book.Title+`", author="`+book.Author+`" where id==`+idStr+
			`;commit;`)
	if err != nil {
		panic(err)
	}

	return nil
}

// DeleteData : cruD
func DeleteData(id uint, table string, dbname string) error {
	db, _ := ql.OpenFile(dbname, &ql.Options{CanCreate: true, RemoveEmptyWAL: true})
	defer db.Close()

	idStr := strconv.FormatUint(uint64(id), 16)

	_, _, err := db.Run(ql.NewRWCtx(),
		`begin transaction;`+
			`delete from books where id==`+idStr+
			`;commit;`)
	if err != nil {
		panic(err)
	}

	return nil
}
