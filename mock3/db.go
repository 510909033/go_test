package mock3

import (
	"fmt"
)

/*
第一步：使用 mockgen 生成 db_mock.go。
一般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package。

mockgen -source=db.go -destination=db_mock.go -package=mock3

*/

// db.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	fmt.Println("db.go GetFromDB, ", db, key)
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
