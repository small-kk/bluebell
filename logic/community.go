package logic

import (
	"app/dao/mysql"
	"app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查询communities表中所有数据
	dataList, err := mysql.GetCommunityList()
	return dataList, err
}

func GetCommunityDetailByID(communityId int) (*models.CommunityDetailRes, error) {
	return mysql.GetCommunityDetailByID(communityId)
}
