/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
