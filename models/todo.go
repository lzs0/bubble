package models

import "bubble/dao"

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (todo *Todo) CreateTodo() (err error) {
	// 写进数据库
	err = dao.DB.Create(&todo).Error
	return
}

func (todo *Todo) UpdateTodo() (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func (todo *Todo) GetATodo(id string) (err error) {
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return
	}
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
