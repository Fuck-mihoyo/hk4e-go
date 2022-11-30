-- 基础信息
local base_info = {
	group_id = 199002085
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 85001, monster_id = 28030101, pos = { x = 438.568, y = 162.217, z = -734.522 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, drop_tag = "鸟类", disableWander = true, area_id = 401 },
	{ config_id = 85002, monster_id = 28030101, pos = { x = 541.793, y = 191.037, z = -104.943 }, rot = { x = 0.000, y = 77.917, z = 0.000 }, level = 1, drop_tag = "鸟类", disableWander = true, pose_id = 2, area_id = 402 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 85003, gadget_id = 70500000, pos = { x = 402.173, y = 194.952, z = -638.955 }, rot = { x = 0.000, y = 13.865, z = 0.000 }, level = 1, point_type = 2001, area_id = 401 },
	{ config_id = 85004, gadget_id = 70500000, pos = { x = 515.759, y = 188.849, z = -94.679 }, rot = { x = 0.000, y = 9.660, z = 0.000 }, level = 1, point_type = 2001, area_id = 402 }
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
		monsters = { 85001, 85002 },
		gadgets = { 85003, 85004 },
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