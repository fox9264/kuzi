package models

import (
	"demo/dao"
)


type Weibo struct {
	Phone uint64 `json:"phone"`
	Uid   uint64 `json:"uid"`
}


func FindByUid(uid string) (WeiboList []*Weibo, err error){
	if err = dao.DB.Where("uid=?", uid).Find(&WeiboList).Error; err != nil{
		return nil, err
	}
	return
}

