package models

import (
	orm "vAdmin/database"
	"fmt"
	//"vAdmin/tools"
	_ "time"
)

type Releasecicd struct {
	Id          int    `json:"id" gorm:"type:int(11) unsigned;primary_key"` // 自增ID
	Productname string `json:"productname" gorm:"type:varchar(250);"`
	Rid         string `sql:"index" json:"rid" gorm:"type:varchar(250);"`                                           // 计划ID
	Fid         string `sql:"index" json:"fid" gorm:"type:varchar(25);"`                                            // 需求ID
	Bid         string `sql:"index" json:"bid" gorm:"type:varchar(25);"`                                            // bugID
	Rnotes      string `json:"rnotes" gorm:"type:text;"` // Release Notes
	Openedby   string `json:"openedby" gorm:"type:varchar(25);"`  // 产品发起人
	Openeddate string `json:"openeddate" gorm:"type:datetime;"`   // 发起时间

	Appname     string `sql:"index" json:"appname" gorm:"type:varchar(25);"` // 应用名
	Devtime   string `json:"devtime" gorm:"type:varchar(5);"`   // 开发耗时
	Commitid    string `json:"commitid" gorm:"type:varchar(50);"` // CommitID
	Cicdtime   string `json:"cicdtime" gorm:"type:datetime;"`     // 提测时间
	Civersion   string `json:"civersion" gorm:"type:text;"` // CICD版本
	Num        int    `json:"num" gorm:"type:int(3);"`  // 测试发布次数
	Stgtime   string `json:"stgtime" gorm:"type:varchar(5);"`   // 测试耗时

	Prdversion string `json:"prdversion" form:"prdversion" gorm:"column:prdversion;comment:'上线版本';type:varchar(70)"`
	Deploytime string `json:"deploytime" gorm:"type:datetime;"`   // 发布时间
	Fidstatus string `sql:"index" json:"fidstatus" gorm:"type:varchar(5);"` // 发布状态
	CreateBy   string `json:"createBy" gorm:"type:varchar(50);"` // 创建人
	UpdateBy   string `json:"updateBy" gorm:"type:varchar(50);"` // 更新者
	Daylabel string `json:"daylabel " gorm:"type:varchar(50);"`

	Sure_status  string `json:"sure_status" gorm:"type:enum('需求中'，'待开发排期','开发中','开发完成','待测试排期','测试中','测试完成','已上线','');default:''"` //认领状态
	Sure_time    string `json:"sure_time" gorm:"type:timestamp;"` // 开始处理时间
	Sure_reasion string `json:"sure_reasion" gorm:"type:varchar(500);"`
	Username   string `json:"username" gorm:"type:varchar(25);"`  // 指派给谁

	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (Releasecicd) TableName() string {
	return "releasecicd"
}

type Releaserid struct {
	Id int `json:"id" gorm:"type:int(11) unsigned;primary_key"` // 自增ID
	Rid string `json:"rid" gorm:"type:varchar(250);"` // 计划ID
	Openeddate string `json:"openeddate" gorm:"type:datetime;"` // 发起时间
	//CreateBy string `json:"createBy" gorm:"type:varchar(128);"` // 创建人
	//UpdateBy string `json:"updateBy" gorm:"type:varchar(128);"` // 更新者
	Rnotesadd string `json:"rnotesadd" gorm:"size:2048;"` // Rid视图
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
}

func (Releaserid) RidTableName() string {
	return "releaserid"
}

// 条件查询获取Releasecicd
func (e *Releasecicd) Get() (Releasecicd, error) {
	var doc Releasecicd
	table := orm.Eloquent.Table(e.TableName())

	// 此if条件很关键，提供查看、修改、删除数据。
	if e.Id != 0  {
		table = table.Where("id = ?", e.Id)
	}

	if e.Rid != ""  {
		table = table.Where("rid = ?", e.Rid)
	}


	if e.Fid != ""  {
		table = table.Where("fid = ?", e.Fid)
	}


	if e.Bid != ""  {
		table = table.Where("bid = ?", e.Bid)
	}

	if e.Appname != ""  {
		table = table.Where("appname = ?", e.Appname)
	}

	if e.Openedby != ""  {
		table = table.Where("openedby = ?", e.Openedby)
	}

	if e.Openeddate != ""  {
		table = table.Where("openeddate >=  ?", e.Openeddate)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取已上线带分页
func (e *Releasecicd) GetPage(pageSize int, pageIndex int) ([]Releasecicd, int, error) {
	var doc []Releasecicd

	fmt.Println(pageSize,pageIndex)
	table := orm.Eloquent.Select("*," +
		"case " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)=0 then '0.5' " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)>=1 and DATEDIFF(cicdtime,openeddate)<180 then DATEDIFF(cicdtime,openeddate)"+
		"else '半年以上' end as devtime," +
		"case " +
		"when deploytime is not null and cicdtime is not null and DATEDIFF(deploytime,cicdtime)=0 then '0.5' " +
		"when deploytime is not null and cicdtime is not null and DATEDIFF(deploytime,cicdtime)>=1 and DATEDIFF(deploytime,cicdtime)<180 then DATEDIFF(deploytime,cicdtime)"+
		"else '半年以上' end as stgtime," +
		"case " +
		"when DATEDIFF(deploytime,openeddate)=0 then '1' " +
		"when DATEDIFF(deploytime,openeddate)>=1 and DATEDIFF(deploytime,openeddate)<180 then DATEDIFF(deploytime,openeddate)"+
		"else '半年以上' end as totaltime," +
		"case " +
		" when deploytime is not null and DATEDIFF(NOW(),deploytime)=0 then 0" +
		" when deploytime is not null and DATEDIFF(NOW(),deploytime)=1 then 1" +
		" when deploytime is not null and DATEDIFF(NOW(),deploytime)=2 then 2" +
		" when deploytime is not null and DATEDIFF(NOW(),deploytime)>2 and DATEDIFF(NOW(),deploytime)<=7 then 3" +
		" else DATEDIFF(NOW(),deploytime) end daylabel").Table(e.TableName())

	if e.Id != 0  {
		table = table.Where("id = ?", e.Id)
	}

	if e.Rid != ""  {
		table = table.Where("rid = ?", e.Rid)
	}

	if e.Fid != ""  {
		table = table.Where("fid = ?", e.Fid)
	}

	if e.Bid != ""  {
		table = table.Where("bid = ?", e.Bid)
	}

	if e.Appname != ""  {
		table = table.Where("appname = ?", e.Appname)
	}

	if e.Openedby != ""  {
		table = table.Where("openedby = ?", e.Openedby)
	}

	if e.Openeddate != ""  {
		table = table.Where("openeddate >=  ?", e.Openeddate)
	}

	if e.Fidstatus != "" {
		table = table.Where("fidstatus = ?", e.Fidstatus)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	/*dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table,err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	*/
	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and (fidstatus='已发生产' or sure_status='已上线')").Order("DATE_FORMAT(deploytime,'%Y-%m-%d') desc").Order("fid desc").Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	if err := table.Where("deleted_at is NULL and unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and (fidstatus='已发生产' or sure_status='已上线')").Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

// 获取带分页合并的rid列表页
func (e *Releaserid) GetPageAdd(pageSize int, pageIndex int) ([]Releaserid, int, error) {

	//table := orm.Eloquent.Raw("rid,group_concat(CONCAT_WS(';',fid,rnotes,openedby,openeddate,appx.appdesc) " +
	//	"Separator '\\n') as rnotesadd from ( select rid,fid,rnotes,openedby,openeddate," +
	//	"group_concat(CONCAT_WS(';',appname,cicdtime,num,deploytime) Separator '\\n') as appdesc " +
	//	"from releasecicd group by rid) as appx").Table(e.RidTableName()).Group("rid")
	//
	//if e.Rid != ""  {
	//	table = table.Where("rid = ?", e.Rid)
	//}
	//
	//var count int
	//if err := table.Offset((pageIndex - 1) * pageSize).Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00')").Limit(pageSize).Order("deploytime desc").Find(&doc).Error; err != nil {
	//	return nil, 0, err
	//}
	//table.Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00')").Count(&count)

	var doc []Releaserid
	var page []Releaserid

	//sql:="SELECT rid,group_concat(CONCAT_WS('|',fid,rnotes,openedby,openeddate,appx.appdesc) Separator '\n\b') as rnotesadd from ( select rid,fid,rnotes,openedby,openeddate,group_concat(CONCAT_WS(';',appname,cicdtime,num,deploytime) Separator '\n\b') as appdesc from releasecicd WHERE unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') group by rid) as appx GROUP BY rid"
	//sql:="SELECT rid,group_concat(CONCAT_WS('|',fid,rnotes,openedby,openeddate,appx.appdesc) Separator '\n\b' ) as rnotesadd\n\nFROM\n\n(select rid,fid,rnotes,openedby,openeddate,group_concat(CONCAT_WS('; ',appname,cicdtime,num,deploytime,\n(case \nwhen DATEDIFF(cicdtime,openeddate)=0 then '1天内' \nwhen DATEDIFF(cicdtime,openeddate)>=1 and DATEDIFF(cicdtime,openeddate)<30 then CONCAT(DATEDIFF(cicdtime,openeddate),'天')\nelse '1个月以上' end),\n(case \nwhen DATEDIFF(deploytime,cicdtime)=0 then '1天内'\nwhen DATEDIFF(deploytime,cicdtime)>=1 and DATEDIFF(deploytime,cicdtime)<30 then CONCAT(DATEDIFF(deploytime,cicdtime),'天')\nelse '1个月以上' end ),\n(case \nwhen DATEDIFF(deploytime,openeddate)=0 then '1天内' \nwhen DATEDIFF(deploytime,openeddate)>=1 and DATEDIFF(deploytime,openeddate)<30 then CONCAT(DATEDIFF(deploytime,openeddate),'天')\nelse '1个月以上' end),\n(case \nwhen DATEDIFF(NOW(),deploytime)=0 then 0\nwhen DATEDIFF(NOW(),deploytime)=1 then 1\nwhen DATEDIFF(NOW(),deploytime)=2 then 2\nwhen DATEDIFF(NOW(),deploytime)>2 and DATEDIFF(NOW(),deploytime)<7 then 3\nelse 99 end)\n\n) Separator '\n\b') as appdesc\nfrom releasecicd\nWHERE unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') group by rid) as appx \n\nGROUP BY rid"
	sql:="SELECT rid,group_concat(CONCAT_WS('|',fid,rnotes,openedby,openeddate,appx.appdesc) Separator '\n\b' ) as rnotesadd\n\nFROM\n\n(select rid,fid,rnotes,openedby,openeddate,group_concat(CONCAT_WS('; ',appname,cicdtime,num,deploytime,\n(case \nwhen DATEDIFF(cicdtime,openeddate)=0 then 'Dev耗时1天内' \nwhen DATEDIFF(cicdtime,openeddate)>=1 and DATEDIFF(cicdtime,openeddate)<30 then CONCAT('Dev耗时',DATEDIFF(cicdtime,openeddate),'天')\nelse 'Dev耗时1个月以上' end),\n(case \nwhen DATEDIFF(deploytime,cicdtime)=0 then 'STG耗时1天内'\nwhen DATEDIFF(deploytime,cicdtime)>=1 and DATEDIFF(deploytime,cicdtime)<30 then CONCAT('STG耗时',DATEDIFF(deploytime,cicdtime),'天')\nelse 'STG耗时1个月以上' end ),\n(case \nwhen DATEDIFF(deploytime,openeddate)=0 then 'Total耗时1天内' \nwhen DATEDIFF(deploytime,openeddate)>=1 and DATEDIFF(deploytime,openeddate)<30 then CONCAT('Total耗时',DATEDIFF(deploytime,openeddate),'天')\nelse 'Total耗时1个月以上' end),\n(case \nwhen DATEDIFF(NOW(),deploytime)=0 then '今日上线的'\nwhen DATEDIFF(NOW(),deploytime)=1 then '昨日上线的'\nwhen DATEDIFF(NOW(),deploytime)=2 then '前天上线的'\nwhen DATEDIFF(NOW(),deploytime)>2 and DATEDIFF(NOW(),deploytime)<7 then '本周上线的'\nelse CONCAT('上线后过去',DATEDIFF(NOW(),deploytime),'天') end)\n\n) Separator '\n\b') as appdesc\nfrom releasecicd\nWHERE unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') group by rid) as appx \n\nGROUP BY rid"
	orm.Eloquent.Raw(sql).Scan(&page)
	orm.Eloquent.Raw(sql+" limit ?,?",(pageIndex - 1) * pageSize,10).Scan(&doc)

	/*
	rows,err:=orm.Eloquent.Raw(sql).Rows()
    tools.HasError(err,"数据查询错误",500)
	defer rows.Close()
	for rows.Next(){
		var releaserid Releaserid
		orm.Eloquent.ScanRows(rows,&releaserid)
		doc=append(doc, releaserid)
	}
	 */

	return doc, len(page), nil

}

// 获取带分页的未提测列表
func (e *Releasecicd) GetPageTodo(pageSize int, pageIndex int) ([]Releasecicd, int, error) {
	var doc []Releasecicd

	/*
	table := orm.Eloquent.Select("*," +
		"case " +
		"when openeddate is not null and DATEDIFF(NOW(),openeddate)>=1 and DATEDIFF(NOW(),openeddate)<180 then DATEDIFF(NOW(),openeddate)"+
		"when openeddate is not null and DATEDIFF(NOW(),openeddate)=0 then '0.5'" +
		"else '半年以上' end as totaltime," +
		"case " +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=0 then 0" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=1 then 1" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=2 then 2" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)>2 and DATEDIFF(NOW(),openeddate)<=7 then 3" +
		" else DATEDIFF(NOW(),openeddate) end daylabel").Table(e.TableName())
	*/

	table := orm.Eloquent.Select("releasecicd.id,productname,rid,fid,bid,rnotes,sure_status,sure_time,sure_reasion,appname,civersion,commitid,prdversion,openedby,openeddate,cicdtime,num,deploytime,fidstatus,releasecicd.updated_at,sys_user.username," +
		"case " +
		"when openeddate is not null and DATEDIFF(NOW(),openeddate)>=1 and DATEDIFF(NOW(),openeddate)<180 then DATEDIFF(NOW(),openeddate)"+
		"when openeddate is not null and DATEDIFF(NOW(),openeddate)=0 then '0.5'" +
		"else '半年以上' end as totaltime," +
		"case " +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=0 then 0" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=1 then 1" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)=2 then 2" +
		" when openeddate is not null and DATEDIFF(NOW(),openeddate)>2 and DATEDIFF(NOW(),openeddate)<=7 then 3" +
		" else DATEDIFF(NOW(),openeddate) end daylabel").Table(e.TableName()).Joins("LEFT JOIN sys_user on releasecicd.update_by=sys_user.user_id")

	if e.Id != 0  {
		table = table.Where("id = ?", e.Id)
	}

	if e.Rid != ""  {
		table = table.Where("rid = ?", e.Rid)
	}

	if e.Fid != ""  {
		table = table.Where("fid = ?", e.Fid)
	}

	if e.Bid != ""  {
		table = table.Where("bid = ?", e.Bid)
	}

	if e.Openedby != ""  {
		table = table.Where("openedby = ?", e.Openedby)
	}

	if e.Openeddate != ""  {
		table = table.Where("openeddate >=  ?", e.Openeddate)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	/*dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table,err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	*/
	var count int

	//if err := table.Offset((pageIndex - 1) * pageSize).Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and status='未提测'").Limit(pageSize).Order("openeddate").Order("CAST(fid AS UNSIGNED)").Find(&doc).Error; err != nil {
	if err := table.Offset((pageIndex - 1) * pageSize).Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and fidstatus='未提测' and sure_status!='已上线'").Order("DATE_FORMAT(openeddate,'%Y-%m-%d')").Order("fid desc").Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("releasecicd.deleted_at is NULL and unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and fidstatus='未提测' and sure_status!='已上线'").Count(&count)
	return doc, count, nil
}

// 获取带分页的测试中列表
func (e *Releasecicd) GetPageStg(pageSize int, pageIndex int) ([]Releasecicd, int, error) {
	var doc []Releasecicd

	table := orm.Eloquent.Select("*," +
		"case " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)=0 then '0.5' " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)>=1 and DATEDIFF(cicdtime,openeddate)<180 then DATEDIFF(cicdtime,openeddate)"+
		"else '半年以上' end as devtime," +
		"case " +
		"when cicdtime is not null and DATEDIFF(NOW(),cicdtime)=0 then '0.5' " +
		"when cicdtime is not null and DATEDIFF(NOW(),cicdtime)>=1 and DATEDIFF(NOW(),cicdtime)<180 then DATEDIFF(NOW(),cicdtime)"+
		"else '半年以上' end as stgtime," +
		"case " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)=0 and DATEDIFF(NOW(),cicdtime)=0 then '1' " +
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)>0 and DATEDIFF(NOW(),cicdtime)>0 and DATEDIFF(NOW(),openeddate)>=1 and DATEDIFF(NOW(),openeddate)<180 then DATEDIFF(NOW(),openeddate)"+
		"when cicdtime is not null and DATEDIFF(cicdtime,openeddate)=0 or DATEDIFF(NOW(),cicdtime)=0 then DATEDIFF(NOW(),openeddate)+0.5 " +
		"when cicdtime is null then DATEDIFF(DATE_ADD(NOW(),INTERVAL 1 DAY),openeddate)" +
		"when cicdtime is not null then DATEDIFF(NOW(),openeddate)" +
		"when cicdtime is null and DATEDIFF(NOW(),openeddate)>0 then DATEDIFF(NOW(),openeddate)" +
		"when cicdtime is null and DATEDIFF(NOW(),openeddate)= 0 then '0.5'" +
		"else '半年以上' end as totaltime," +
		"case " +
		" when cicdtime is not null and DATEDIFF(NOW(),cicdtime)=0 then 0" +
		" when cicdtime is not null and DATEDIFF(NOW(),cicdtime)=1 then 1" +
		" when cicdtime is not null and DATEDIFF(NOW(),cicdtime)=2 then 2" +
		" when cicdtime is not null and DATEDIFF(NOW(),cicdtime)>2 and DATEDIFF(NOW(),cicdtime)<=7 then 3" +
		" else DATEDIFF(NOW(),cicdtime) end daylabel").Table(e.TableName())

	if e.Id != 0  {
		table = table.Where("id = ?", e.Id)
	}

	if e.Rid != ""  {
		table = table.Where("rid = ?", e.Rid)
	}

	if e.Fid != ""  {
		table = table.Where("fid = ?", e.Fid)
	}

	if e.Bid != ""  {
		table = table.Where("bid = ?", e.Bid)
	}

	if e.Appname != ""  {
		table = table.Where("appname = ?", e.Appname)
	}

	if e.Openedby != ""  {
		table = table.Where("openedby = ?", e.Openedby)
	}

	if e.Openeddate != ""  {
		table = table.Where("openeddate >=  ?", e.Openeddate)
	}

	if e.Fidstatus != "" {
		table = table.Where("fidstatus = ?", e.Fidstatus)
	}

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	/*dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table,err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	*/
	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Where("unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and fidstatus='已提测'").Order("fid desc").Order("DATE_FORMAT(cicdtime,'%Y-%m-%d') desc").Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("deleted_at is NULL and unix_timestamp(`openeddate`) >= unix_timestamp('2020-07-01 00:00:00') and fidstatus='已提测'").Count(&count)
	return doc, count, nil
}

// 创建Releasecicd
func (e *Releasecicd) Create() (Releasecicd, error) {
	var doc Releasecicd
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 更新Releasecicd
func (e *Releasecicd) Update(id int) (update Releasecicd, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Releasecicd
func (e *Releasecicd) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&Releasecicd{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Releasecicd) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Releasecicd{}).Error; err != nil {
		return
	}
	Result = true
	return
}

type chartModel struct {
	Dtime string
	Dcount int
}

//获取每天上线发布数趋势
func (e *Releasecicd) DeployByDay() (interface{}, error){
	//var deployct map[string]string
	//再使用make函数创建一个非nil的map，nil map不能赋值
	//deployct = make(map[string]string)
	//deployct := make(map[string]string)
	list :=[]chartModel{}
	sql:="SELECT time as dtime,COUNT(time) as dcount FROM (\nSELECT date_format(deploytime,'%Y-%m-%d') time from  releasecicd WHERE deploytime is not null) as T\nGROUP BY time"
	orm.Eloquent.Raw(sql).Scan(&list)
	return list,nil
}

//获取每小时上线发布数趋势
func (e *Releasecicd) DeploybyHour() (interface{}, error) {
	list :=[]chartModel{}
	sql:="SELECT time as dtime,COUNT(time) as dcount FROM (SELECT date_format(deploytime,'%Y-%m-%d %d') time from  releasecicd WHERE deploytime is not null) as T GROUP BY time"
	orm.Eloquent.Raw(sql).Scan(&list)

	return list,nil
}