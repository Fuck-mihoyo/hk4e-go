-- 基础信息
local base_info = {
	group_id = 133103531
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 531001, gadget_id = 70500000, pos = { x = 945.511, y = 376.453, z = 1572.436 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531002, gadget_id = 70500000, pos = { x = 958.607, y = 347.896, z = 1748.330 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531003, gadget_id = 70500000, pos = { x = 942.693, y = 376.560, z = 1577.522 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531004, gadget_id = 70500000, pos = { x = 922.100, y = 353.584, z = 1537.634 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531005, gadget_id = 70500000, pos = { x = 926.401, y = 352.793, z = 1531.740 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531006, gadget_id = 70500000, pos = { x = 950.996, y = 371.054, z = 1648.088 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531007, gadget_id = 70500000, pos = { x = 874.213, y = 357.728, z = 1670.482 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531008, gadget_id = 70500000, pos = { x = 874.177, y = 359.248, z = 1667.990 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531009, gadget_id = 70500000, pos = { x = 933.557, y = 373.988, z = 1577.506 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531010, gadget_id = 70500000, pos = { x = 959.039, y = 350.061, z = 1741.203 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531011, gadget_id = 70500000, pos = { x = 848.038, y = 352.632, z = 1767.286 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 },
	{ config_id = 531012, gadget_id = 70500000, pos = { x = 852.481, y = 351.598, z = 1762.366 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 24, point_type = 9118, area_id = 6 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = ,
		monsters = { },
		gadgets = { 531001, 531002, 531003, 531004, 531005, 531006, 531007, 531008, 531009, 531010, 531011, 531012 },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================