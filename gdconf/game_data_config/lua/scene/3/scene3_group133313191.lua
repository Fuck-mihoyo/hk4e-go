-- 基础信息
local base_info = {
	group_id = 133313191
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 191002, monster_id = 28050106, pos = { x = -342.871, y = -187.140, z = 5447.349 }, rot = { x = 0.000, y = 189.081, z = 0.000 }, level = 32, drop_tag = "魔法生物", area_id = 32 },
	{ config_id = 191003, monster_id = 28050106, pos = { x = -325.737, y = -164.848, z = 5428.330 }, rot = { x = 0.000, y = 88.266, z = 0.000 }, level = 32, drop_tag = "魔法生物", area_id = 32 }
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
		monsters = { 191002, 191003 },
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