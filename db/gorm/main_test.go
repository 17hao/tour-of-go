package main

import (
	"testing"

	"github.com/bytedance/mockey"
)

func TestQuery(t *testing.T) {
	mockey.PatchConvey("", t, func() {
		// db, mock, err := sqlmock.New()
		// if err != nil {
		// 	t.Fatal(err)
		// }
		//
		// rows := mock.NewRows([]string{"id", "name", "city"}).AddRow(1, "name1", "city1").AddRow(2, "name2", "city2")
		// //mock.ExpectQuery("select id, name, city from employees where id > ?").WillReturnRows(rows).WithArgs(1)
		// mock.ExpectQuery("select id, name, city from employees").WillReturnRows(rows)
		//
		// gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}))
		// if err != nil {
		// 	t.Fatal(err)
		// }
		//
		// mockey.Mock(gorm.Open).Return(gormDB, nil).Build()
		// res := query()
		// assert.Equal(t, 2, len(res))
	})
}
