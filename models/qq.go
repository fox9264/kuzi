package models

import (
	"demo/dao"
)


type Qq struct {
	Qq uint64 `json:"qq"`
	Phone   uint64 `json:"phone"`
}

func FindByQq(qq string) (QqList []*Qq, err error){
	if err = dao.DB.Where("qq=?", qq).Find(&QqList).Error; err != nil{
		return nil, err
	}
	return
}

