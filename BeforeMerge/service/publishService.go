package service

import (
	"time"
	"tryForByte/model"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// 对外开放的返回拼装好的list的函数
func GetAllPublishedByUserId(rawId int) (videoslist VideoListResponse, err error) {
	//把int转换成int64
	id := int64(rawId)
	//	todo 和从user那边查找的函数联动
	video, err := model.GetAllPublishedByUserId(id)
	if err != nil {
		return VideoListResponse{}, err
	}
	//先定义isFavorite
	var isFavorite = false
	videos := make([]Video, len(video))
	for index, v := range video {
		isFavorite, _ = model.FavoritedByUserId(id, v.Id)
		//要用service的video，因为那个才是需要拼装起来返回的值
		videos[index] = Video{
			Id:            v.Id,
			Author:        User{Name: "testing"},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFavorite,
		}
	}
	videoslist = VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	}
	return
}

// 对外开放的上传视频方法
func AddVideo(filepath string, title string, id int64) (err error) {
	//todo  视频投稿的封面不是从post里面传的,这里暂时随便填了
	video := &model.Video{
		AuthorId:    id,
		PlayUrl:     filepath,
		CoverUrl:    "notsure",
		Title:       title,
		PublishTime: time.Now(),
	}
	if err = model.AddVideo(video); err != nil {
		//todo 这里是否可以精简，这里的err如果为空会如何，不确定，所以没删
		return err
	}
	return
}
