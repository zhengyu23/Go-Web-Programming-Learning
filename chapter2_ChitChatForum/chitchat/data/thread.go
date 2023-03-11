package data

/* 帖子 - Thread */
import (
	"time"
)

// Thread 存储所有帖子的相关代码
type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// Threads 从数据库里取出所有帖子并返回给index处理器函数
func Threads() (threads []Thread, err error) {
	// ORDER BY created_at DESC: 按照 created_at 逆序排序
	// ASC: 正序
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at" +
		" FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}
