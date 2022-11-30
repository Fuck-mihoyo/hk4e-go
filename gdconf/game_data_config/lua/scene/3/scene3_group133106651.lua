-- 基础信息
local base_info = {
	group_id = 133106651
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 651001, monster_id = 28020102, pos = { x = -661.430, y = 155.668, z = 1668.546 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 36, drop_tag = "走兽", area_id = 19 },
	{ config_id = 651002, monster_id = 28020102, pos = { x = -653.362, y = 158.501, z = 1679.253 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 36, drop_tag = "走兽", area_id = 19 },
	{ config_id = 651003, monster_id = 28010202, pos = { x = -709.971, y = 131.994, z = 1639.514 }, rot = { x = 0.000, y = 280.229, z = 0.000 }, level = 36, drop_tag = "采集动物", area_id = 19 },
	{ config_id = 651004, monster_id = 20010401, pos = { x = -789.457, y = 173.124, z = 1638.207 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 36, drop_tag = "大史莱姆", area_id = 19 },
	{ config_id = 651005, monster_id = 28020102, pos = { x = -674.387, y = 135.363, z = 1639.304 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 36, drop_tag = "走兽", area_id = 19 },
	{ config_id = 651006, monster_id = 28020102, pos = { x = -664.925, y = 136.416, z = 1642.592 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 36, drop_tag = "走兽", area_id = 19 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
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
		monsters = { 651001, 651002, 651003, 651004, 651005, 651006 },
		gadgets = { },
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