package main

import (
	"alarm/common"
	"alarm/job"
	"alarm/models"
	_ "alarm/routers"
	"alarm/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	initConfig()

	cache, _ := cache.NewCache("memory", `{"interval":60}`)
	common.CacheMemory = cache

	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)

	connection := beego.AppConfig.String("db.username") + ":" + beego.AppConfig.String("db.pwd") + "@tcp(" + beego.AppConfig.String("db.host") + ")/" + beego.AppConfig.String("db.name") + "?charset=utf8&loc=Asia%2FShanghai"

	logs.SetLogger(logs.AdapterFile, `{"filename":"info.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", connection, 10, 10)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(models.AlarmDbConnection), new(models.AlarmNoticeMonitorSql), new(models.AlarmNoticeMsg), new(models.AlarmNoticeListener), new(models.AlarmNoticeTopic), new(models.AlarmNoticeTopicListener))

	orm.RunSyncdb("default", false, true)

	dbList := models.GetAlarmDbConnectionList()

	for _, db := range dbList {
		conn1 := db.Username + ":" + db.Pwd + "@tcp(" + db.HostPort + ")/" + db.Name + "?charset=utf8&loc=Asia%2FShanghai"
		orm.RegisterDataBase(db.AliasName, "mysql", conn1, 10, 10)
		//orm.Debug=true
	}

	cronTab := utils.NewCrontab()

	common.AlarmCronTab = cronTab

	job.InitAlarmNoticeMonitorJob()

	job.InitAlarmNoticeMsgJob()

	cronTab.Start()

	defer cronTab.Stop()

	logs.Info("start success !")
	beego.Run()
}

func initConfig() {

	rootPath := GetAPPRootPath()

	beego.SetViewsPath(rootPath + "/views")

	beego.LoadAppConfig("ini", rootPath+"/conf/app.conf")

	beego.SetStaticPath("static", rootPath+"/static")

}

func GetAPPRootPath() string {

	file, err := exec.LookPath(os.Args[0])

	if err != nil {

		return ""

	}

	p, err := filepath.Abs(file)

	if err != nil {

		return ""

	}

	return filepath.Dir(p)

}
