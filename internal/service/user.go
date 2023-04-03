package service

import (
	"Metart/internal/dao"
	"Metart/internal/model"
)

// CreateUser 新建User信息
func CreateUser(user *model.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

// GetAllUser 查询所有user信息
func GetAllUser() (userList []*model.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

// DeleteUserById 根据id删除对应的user信息
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&model.User{}).Error
	return
}

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}
