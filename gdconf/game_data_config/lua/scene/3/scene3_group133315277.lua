-- 基础信息
local base_info = {
	group_id = 133315277
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 277001, monster_id = 26090801, pos = { x = 394.457, y = 213.140, z = 2226.885 }, rot = { x = 0.000, y = 9.416, z = 0.000 }, level = 27, drop_tag = "蕈兽", pose_id = 102, area_id = 20 },
	{ config_id = 277002, monster_id = 26090801, pos = { x = 380.565, y = 215.679, z = 2220.532 }, rot = { x = 0.000, y = 7.368, z = 0.000 }, level = 27, drop_tag = "蕈兽", pose_id = 102, area_id = 20 },
	{ config_id = 277003, monster_id = 26090801, pos = { x = 384.176, y = 214.162, z = 2232.306 }, rot = { x = 0.000, y = 36.894, z = 0.000 }, level = 27, drop_tag = "蕈兽", pose_id = 102, area_id = 20 }
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
		monsters = { 277001, 277002, 277003 },
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