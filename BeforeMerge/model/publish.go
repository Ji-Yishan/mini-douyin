package model

import (
	"time"
)

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	PublishTime   time.Time
}

func (v Video) TableName() string {
	return "video"
}

// 数据库查询相关的返回数据库内list的函数
func GetAllPublishedByUserId(id int64) (videos []*Video, err error) {
	//uid, err := strconv.Atoi(id)
	//if err != nil {
	//	return nil, err
	//}
	//todo 查询不到数据返回错误（？ 需要返回吗？
	videos = make([]*Video, 0)
	if err = DB.Where("author_id =?", id).Find(&videos).Error; err != nil {
		return nil, err
	}
	return
}

// 在数据库内添加视频信息
func AddVideo(video *Video) (err error) {
	//todo cover_url没找到在哪里弄哎
	if err := DB.Create(&video).Error; err != nil {
		return err
	}
	return nil
}
