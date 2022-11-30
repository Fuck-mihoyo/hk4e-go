-- 基础信息
local base_info = {
	group_id = 133106551
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
	{ config_id = 551001, gadget_id = 70210101, pos = { x = -677.652, y = 109.329, z = 1799.750 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜通用璃月", persistent = true, area_id = 19 },
	{ config_id = 551002, gadget_id = 70210101, pos = { x = -677.059, y = 109.426, z = 1796.836 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜通用璃月", persistent = true, area_id = 19 },
	{ config_id = 551003, gadget_id = 70210101, pos = { x = -676.872, y = 108.927, z = 1797.941 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜通用璃月", persistent = true, area_id = 19 },
	{ config_id = 551004, gadget_id = 70220025, pos = { x = -690.192, y = 109.688, z = 1792.958 }, rot = { x = 90.000, y = 0.000, z = 0.000 }, level = 36, area_id = 19 },
	{ config_id = 551005, gadget_id = 70220025, pos = { x = -692.298, y = 109.688, z = 1796.077 }, rot = { x = 90.000, y = 192.981, z = 0.000 }, level = 36, area_id = 19 }
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

-- 废弃数据
garbages = {
	monsters = {
		{ config_id = 551018, monster_id = 28030101, pos = { x = -669.468, y = 130.201, z = 1840.253 }, rot = { x = 0.000, y = 182.794, z = 0.000 }, level = 36, drop_tag = "鸟类", pose_id = 901, area_id = 19 },
		{ config_id = 551019, monster_id = 28030101, pos = { x = -669.609, y = 130.226, z = 1832.407 }, rot = { x = 0.000, y = 182.794, z = 0.000 }, level = 36, drop_tag = "鸟类", pose_id = 901, area_id = 19 }
	}
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
		gadgets = { 551001, 551002, 551003, 551004, 551005 },
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