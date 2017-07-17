package models

import "github.com/astaxie/beego/orm"


const TABLENAME_TASK="task"

type Task struct {
	Id int
	Taskname string
	Taskstate string
	Createbyuserid int
	Createbyusername string
	Completedbyuserid int
	Completedbyusername string
	Createdate string
	Completedate string
	Weight string
	Remake string
	Introduction string
}


func (t *Task) TableName() string {
	return TableName(TABLENAME_USER)
}

func CreateTask(task *Task) (int64, error) {
	return orm.NewOrm().Insert(task)
}
func (t *Task) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}
func FindTaskById(id int) (*Task, error){
	t := new(Task)

	err := orm.NewOrm().QueryTable(TableName(TABLENAME_TASK)).Filter("id", id).One(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}