package model

type Favorite struct {
	Id      int64
	UserId  int64
	VideoId int64
}

func (f Favorite) TableName() string {
	return "favorite"
}

func FavoritedByUserId(uid int64, vid int64) (favorited bool, err error) {
	favorite := make([]*Favorite, 0)
	if err = DB.Where("?=?", uid, vid).Find(&favorite).Error; err != nil {
		return false, err
	}

	//todo 这里可否优化一下
	if len(favorite) != 0 {
		return true, nil
	}
	return
}
