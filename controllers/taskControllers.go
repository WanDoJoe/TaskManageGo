package controllers

import (
	"github.com/astaxie/beego"
	"taskmanage/models"
	"encoding/json"
	"fmt"
)

type TaskController struct {
	beego.Controller
}

var creatTaskMsg=""
var creatTaskSuccess=""
func (this *TaskController)CreateTask(){
	params:=this.GetString("jsonParams")
	task:=new(models.Task)
	task.Completedate=GetParams(params,"completedate")
	task.Completedbyuserid=StringToInt(GetParams(params,"completedbyuserid"))
	task.Completedbyusername=GetParams(params,"completedate")
	task.Createbyuserid=StringToInt(GetParams(params,"createbyuserid"))
	task.Createdate=GetParams(params,"createdate")
	task.Remake=GetParams(params,"remake")
	task.Taskname=GetParams(params,"taskname")
	task.Taskstate=GetParams(params,"taskstate")
	task.Weight=GetParams(params,"weight")

	returnid,_:=models.CreateTask(task)

	if returnid>0{
		creatTaskMsg="创建成功"
		creatTaskSuccess="true"
	}else{
		creatTaskMsg="创建失败"
		creatTaskSuccess="false"
	}

	this.Data["json"]= map[string]interface{}{"success":creatTaskSuccess,"message":creatTaskMsg,"data":returnid}
	this.ServeJSON(false)
}

func (this *TaskController)MyTaskList(){

}

func (this *TaskController)TaskDetise(){
	params:=this.GetString("jsonParams")
	taskId:=GetParams(params,"taskid")

	task,err:=models.FindTaskById(StringToInt(taskId))
	if err!=nil{
		panic(err)
	}
	taskJson,er:=json.Marshal(task)
	if er!=nil{
		panic(er)
	}
	fmt.Println(taskJson)
	this.Data["json"] = map[string]interface{}{"success":"true","message":"成功","data":taskJson}
	this.ServeJSON(false)

}
func (this *TaskController) EditTask(){

}