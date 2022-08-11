/*
 * @Date: 2022-02-14 15:05:35
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-04-28 17:04:03
 * @FilePath: \server\system\model\DbModel\sys.go
 */

package DbModel

import (
	"time"

	"chunDoc/system/model/CommonModel"
)

//用户
type SysUser struct {
	CommonModel.BaseModel
	LoginName   string          `json:"login_name" gorm:"uniqueIndex;size:100;not null;comment:登录名"`
	UserName    string          `json:"user_name"  gorm:"size:100"`
	Password    string          `json:"password" gorm:"size:255"`
	Status      int             `json:"status"`
	Email       string          `json:"email"  gorm:"size:100"`
	Phone       string          `json:"phone"  gorm:"size:20"`
	Token       string          `json:"token"  gorm:"-"`
	Kick        uint            `json:"kick"  gorm:"-"` //99：强制登录,1：已经登录需要询问，0：直接可以登录
	Permissions []SysPermission `json:"permissions" gorm:"many2many:sys_user_permission;joinForeignKey:UserId;joinReferences:PermissionId;"`
	Roles       []SysRole       `json:"-" gorm:"many2many:sys_user_role;joinForeignKey:UserId;joinReferences:RoleId;"`
	//Powers      []SysPermission `json:"powers"  gorm:"-"`
}

//权限
type SysPermission struct {
	CommonModel.BaseModel
	Parent     uint            `json:"parent" gorm:"default:0;comment:父级id"`
	Name       string          `json:"name" gorm:"uniqueIndex;size:100;not null;comment:唯一标识"`
	Title      string          `json:"title" gorm:"size:100;comment:名称"`
	Path       string          `json:"path" gorm:"size:100;comment:页面路径"`
	Icon       string          `json:"icon" gorm:"size:100;comment:图标"`
	Sort       uint            `json:"sort" gorm:"default:99;comment:排序"`
	Type       string          `json:"type" gorm:"size:10;default:menu;comment:类型"`
	Status     uint            `json:"status" gorm:"default:0;comment:状态"`
	Isshowmenu uint            `json:"isshowmenu" gorm:"default:1;comment:是否菜单显示"`
	Isrecord   uint            `json:"isrecord" gorm:"default:2;comment:是否记录"`
	Children   []SysPermission `json:"children" gorm:"-"`
}

//用户权限关联表
type SysUserPermission struct {
	UserId       uint          `json:"user_id" gorm:"comment:用户id"`
	PermissionId uint          `json:"permission_id" gorm:"comment:权限id"`
	Checked      uint          `json:"checked" gorm:"comment:状态：1为选中，0为半选中"`
	Permission   SysPermission `json:"permissions" gorm:"foreignKey:PermissionId;references:ID"`
}

//角色表
type SysRole struct {
	CommonModel.BaseModel
	Name        string          `json:"name" gorm:"size:100;not null;comment:角色名称"`
	Desc        string          `json:"desc"  gorm:"size:255;comment:描述"`
	Permissions []SysPermission `json:"permissions" gorm:"many2many:sys_role_permission;joinForeignKey:RoleId;joinReferences:PermissionId;"`
}

//用户权限角色
type SysUserRole struct {
	UserId uint    `json:"user_id" gorm:"comment:用户id"`
	RoleId uint    `json:"role_id" gorm:"comment:角色id"`
	Roles  SysRole `json:"roles" gorm:"foreignKey:RoleId;references:ID"`
}

//角色权限关联表
type SysRolePermission struct {
	RoleId       uint          `json:"role_id" gorm:"comment:角色id"`
	PermissionId uint          `json:"permission_id" gorm:"comment:权限id"`
	Checked      uint          `json:"checked" gorm:"comment:状态：1为选中，0为半选中"`
	Permission   SysPermission `json:"permissions" gorm:"foreignKey:PermissionId;references:ID"`
}

//日志表
type SysRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"` // 主键ID
	Uid       uint      `json:"uid" gorm:"comment:用户id"`            // 用户id
	Username  string    `json:"username" gorm:"size:100;comment:用户名称"`
	Url       string    `json:"url" gorm:"size:100;comment:Url"`
	Desc      string    `json:"desc" gorm:"size:100;comment:描述"`
	Ip        string    `json:"ip" gorm:"size:30;comment:Ip"`
	Request   string    `json:"request" gorm:"size:1000;comment:请求"`
	CreatedAt time.Time `json:"create_time" gorm:"column:create_time;comment:创建时间"` // 创建时间
}

//系统配置表
type SysConfig struct {
	Name        string    `json:"name" gorm:"size:40;primaryKey;"` // 主键ID
	Desc        string    `json:"desc" gorm:"size:100;comment:描述"`
	Group       string    `json:"group" gorm:"size:100;comment:分组"`
	Type        uint      `json:"type" gorm:"comment:类型1：字符串 2 数字 3 数组 4 字典"`
	Value       string    `json:"value" gorm:"-"`
	ValueString string    `json:"value_string" gorm:"size:1000;comment:值字符串"`
	ValueNumber float64   `json:"value_number" gorm:"comment:值数字"`
	CreatedAt   time.Time `json:"create_time" gorm:"column:create_time;comment:创建时间"` // 创建时间
}
