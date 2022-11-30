-- 基础信息
local base_info = {
	group_id = 133312020
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 20001, monster_id = 25210403, pos = { x = -3098.249, y = 330.388, z = 4725.337 }, rot = { x = 0.000, y = 294.485, z = 0.000 }, level = 32, drop_tag = "镀金旅团", disableWander = true, pose_id = 9504, area_id = 28 },
	{ config_id = 20002, monster_id = 25210202, pos = { x = -3093.750, y = 335.416, z = 4726.470 }, rot = { x = 0.000, y = 296.479, z = 0.000 }, level = 32, drop_tag = "镀金旅团", disableWander = true, pose_id = 9001, area_id = 28 }
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
		monsters = { 20001, 20002 },
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