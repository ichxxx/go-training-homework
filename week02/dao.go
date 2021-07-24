package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

func QuerySomeData() error {
	db, err := sql.Open("foo", "bar")
	if err != nil {
		return errors.Wrap(err, "dao: open data source error")
	}

	err = db.QueryRow("baz").Scan()

	// sql.ErrNoRows错误是当QueryRow()方法返回结果为空时产生的
	//
	// 该错误的官方定义如下：
	// ErrNoRows is returned by Scan when QueryRow doesn't return a
    // row. In such a case, QueryRow returns a placeholder *Row value that
    // defers this error until a Scan.
    //
    // 所以该错误并不是真正意义上的错误，所以可以忽略此特殊情况
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return errors.Wrap(err, "dao: query row error")
}