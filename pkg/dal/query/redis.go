package query

import "fmt"

// 各个表的枚举
const (
	TableUser  = "user"
	TableVideo = "video"
	//TableLike    = "like"
	//TableComment = "comment"
)

// 统计量的枚举
const (
	CountLike    = "like_count"
	CountVideo   = "video_count"
	CountComment = "comment_count"
	CountFans    = "fans_count"
	CountFollow  = "follow_count"
)

// GenRedisKey generate redis key
// table: enum of table
// id: id of the table
func GenRedisKey(table string, id int64) string {
	return fmt.Sprintf("%s:%d", table, id)
}
