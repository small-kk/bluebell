package mysql

import (
	"app/models"
	"errors"
	"gorm.io/gorm"
)

// GetCommunityList 获取community中所有记录
func GetCommunityList() ([]*models.Community, error) {
	var dataList = make([]*models.Community, 0, 20)
	err := db.Find(&dataList).Error
	return dataList, err
}

// GetCommunityDetailByID 根据id查询社区分类详情
func GetCommunityDetailByID(communityId int) (*models.CommunityDetailRes, error) {
	communityDetailRes := new(models.CommunityDetailRes)
	err := db.Table("communities").Where("id=?", communityId).Find(&communityDetailRes).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.New("invalid community id")
	}
	return communityDetailRes, err
}
