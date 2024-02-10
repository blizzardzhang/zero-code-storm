// Code generated by goctl. DO NOT EDIT.
package types

type AddClientReq struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Scope        string `json:"scope"`
}

type AddClientResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteClientResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteClientrReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteUserReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeptRelations struct {
	Id       int64  `json:"id"`
	Value    string `json:"value"`
	Title    string `json:"title"`
	ParentId int64  `json:"parentId"`
}

type DetailReq struct {
	Id uint64 `form:"id" validate:"number"`
}

type DetailResp struct {
	Id           int64  `json:"id"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Scope        string `json:"scope"`
	IsDeleted    string `json:"isDeleted"`
}

type JobRelations struct {
	Id      int64  `json:"id"`
	JobName string `json:"jobName"`
}

type ListClientReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	ClientId string `json:"clientId,optional"`
}

type ListClientResp struct {
	Code     string        `json:"code"`
	Message  string        `json:"message"`
	Current  int64         `json:"current,default=1"`
	Data     []*DetailResp `json:"data"`
	PageSize int64         `json:"pageSize,default=20"`
	Success  bool          `json:"success"`
	Total    int64         `json:"total"`
}

type ListMenuTree struct {
	Id       int64  `json:"id"`       // 编号
	Path     string `json:"path"`     // 菜单路径
	Name     string `json:"name"`     // 菜单名称
	ParentId int64  `json:"parentId"` // 父菜单ID，一级菜单为0
	Icon     string `json:"icon"`     // 菜单图标
}

type ListMenuTreeVue struct {
	Id           int64        `json:"id"`
	ParentId     int64        `json:"parentId"`
	Title        string       `json:"title"`
	Path         string       `json:"path"`
	Name         string       `json:"name"`
	Icon         string       `json:"icon"`
	VueRedirect  string       `json:"redirect"`
	VueComponent string       `json:"component"`
	Meta         MenuTreeMeta `json:"meta"`
}

type ListUserData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 用户名
	NickName       string `json:"nickName"`       // 昵称
	Avatar         string `json:"avatar"`         // 头像
	Email          string `json:"email"`          // 邮箱
	Mobile         string `json:"mobile"`         // 手机号
	Status         int64  `json:"status"`         // 状态  0：禁用   1：正常
	DeptId         int64  `json:"deptId"`         // 机构ID
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  0：已删除  1：正常
	JobId          int64  `json:"jobId"`
	RoleId         int64  `json:"roleId"`
	RoleName       string `json:"roleName"`
	JobName        string `json:"jobName"`
	DeptName       string `json:"deptName"`
}

type ListUserReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	Name     string `json:"name,optional"`
	NickName string `json:"nickName,optional"`
	Mobile   string `json:"mobile,optional"`
	Email    string `json:"email,optional"`
	Status   int64  `json:"status,optional,default=2"`
	DeptId   int64  `json:"deptId,optional"`
	JobId    int64  `json:"deptId,optional"`
}

type ListUserResp struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Current  int64           `json:"current,default=1"`
	Data     []*ListUserData `json:"data"`
	PageSize int64           `json:"pageSize,default=20"`
	Success  bool            `json:"success"`
	Total    int64           `json:"total"`
}

type MenuTreeMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type QueryAllRelationsData struct {
	RoleRelations []*RoleRelations `json:"roleRelations"`
	DeptRelations []*DeptRelations `json:"deptRelations"`
	JobRelations  []*JobRelations  `json:"jobRelations"`
}

type QueryAllRelationsReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type QueryAllRelationsResp struct {
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Data    QueryAllRelationsData `json:"data"`
}

type ReSetPasswordReq struct {
	Id int64 `json:"id"`
}

type ReSetPasswordResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RoleRelations struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type UpdateUserReq struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	DeptId   int64  `json:"deptId"`
	RoleId   int64  `json:"roleId"`
	Status   int64  `json:"status"`
	JobId    int64  `json:"jobId"`
	DelFlag  int64  `json:"delFlag,default=2"` // 是否删除  0：已删除  1：正常
}

type UpdateUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UserStatusReq struct {
	Id     int64 `json:"id"`
	Status int64 `json:"status"` // 状态  0：禁用   1：正常
}

type UserStatusResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AddUserReq struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	DeptId   int64  `json:"deptId"`
	RoleId   int64  `json:"roleId"`
	JobId    int64  `json:"jobId"`
	Status   int64  `json:"status"` // 状态  0：禁用   1：正常
}

type AddUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResp struct {
	Code             string `json:"code"`
	Message          string `json:"message"`
	Status           string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Id               int64  `json:"id"`
	UserName         string `json:"userName"`
	AccessToken      string `json:"token"`
	AccessExpire     int64  `json:"accessExpire"`
	RefreshAfter     int64  `json:"refreshAfter"`
}

type UserInfoResp struct {
	Code        string             `json:"code"`
	Message     string             `json:"message"`
	Avatar      string             `json:"avatar"`
	Name        string             `json:"name"`
	MenuTree    []*ListMenuTree    `json:"menuTree"`
	MenuTreeVue []*ListMenuTreeVue `json:"menuTreeVue"`
}
