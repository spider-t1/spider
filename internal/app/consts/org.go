package consts

const (
	RoleSlugSuperAdmin = 1 // 超级管理员
	RoleSlugAdmin      = 2 // 管理员
	RoleSlugDeptHead   = 3 // 部门主管
	RoleSlugPartner    = 4 // 伙伴
)

const (
	RoleDataScopeAll      = 1 // 全部数据权限
	RoleDataScopeDeptAll  = 2 // 本部门及以下数据权限
	RoleDataScopeDeptSelf = 3 // 本部门数据权限
	RoleDataScopeSelf     = 4 // 自己数据权限
	RoleDataScopeCustom   = 5 // 自定义数据权限
)
