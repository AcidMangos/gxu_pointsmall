package model

import "gxu_pointsmall/dao"

// Notice 公告序列化器
type Notice struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNotice 序列化轮播图
func BuildNotice(item dao.Notice) Notice {
	return Notice{
		ID:        item.ID,
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
