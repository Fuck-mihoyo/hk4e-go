package constant

var FightPropertyConst *FightProperty

type FightProperty struct {
	FIGHT_PROP_NONE                             uint16
	FIGHT_PROP_BASE_HP                          uint16
	FIGHT_PROP_HP                               uint16
	FIGHT_PROP_HP_PERCENT                       uint16
	FIGHT_PROP_BASE_ATTACK                      uint16
	FIGHT_PROP_ATTACK                           uint16
	FIGHT_PROP_ATTACK_PERCENT                   uint16
	FIGHT_PROP_BASE_DEFENSE                     uint16
	FIGHT_PROP_DEFENSE                          uint16
	FIGHT_PROP_DEFENSE_PERCENT                  uint16
	FIGHT_PROP_BASE_SPEED                       uint16
	FIGHT_PROP_SPEED_PERCENT                    uint16
	FIGHT_PROP_HP_MP_PERCENT                    uint16
	FIGHT_PROP_ATTACK_MP_PERCENT                uint16
	FIGHT_PROP_CRITICAL                         uint16
	FIGHT_PROP_ANTI_CRITICAL                    uint16
	FIGHT_PROP_CRITICAL_HURT                    uint16
	FIGHT_PROP_CHARGE_EFFICIENCY                uint16
	FIGHT_PROP_ADD_HURT                         uint16
	FIGHT_PROP_SUB_HURT                         uint16
	FIGHT_PROP_HEAL_ADD                         uint16
	FIGHT_PROP_HEALED_ADD                       uint16
	FIGHT_PROP_ELEMENT_MASTERY                  uint16
	FIGHT_PROP_PHYSICAL_SUB_HURT                uint16
	FIGHT_PROP_PHYSICAL_ADD_HURT                uint16
	FIGHT_PROP_DEFENCE_IGNORE_RATIO             uint16
	FIGHT_PROP_DEFENCE_IGNORE_DELTA             uint16
	FIGHT_PROP_FIRE_ADD_HURT                    uint16
	FIGHT_PROP_ELEC_ADD_HURT                    uint16
	FIGHT_PROP_WATER_ADD_HURT                   uint16
	FIGHT_PROP_GRASS_ADD_HURT                   uint16
	FIGHT_PROP_WIND_ADD_HURT                    uint16
	FIGHT_PROP_ROCK_ADD_HURT                    uint16
	FIGHT_PROP_ICE_ADD_HURT                     uint16
	FIGHT_PROP_HIT_HEAD_ADD_HURT                uint16
	FIGHT_PROP_FIRE_SUB_HURT                    uint16
	FIGHT_PROP_ELEC_SUB_HURT                    uint16
	FIGHT_PROP_WATER_SUB_HURT                   uint16
	FIGHT_PROP_GRASS_SUB_HURT                   uint16
	FIGHT_PROP_WIND_SUB_HURT                    uint16
	FIGHT_PROP_ROCK_SUB_HURT                    uint16
	FIGHT_PROP_ICE_SUB_HURT                     uint16
	FIGHT_PROP_EFFECT_HIT                       uint16
	FIGHT_PROP_EFFECT_RESIST                    uint16
	FIGHT_PROP_FREEZE_RESIST                    uint16
	FIGHT_PROP_TORPOR_RESIST                    uint16
	FIGHT_PROP_DIZZY_RESIST                     uint16
	FIGHT_PROP_FREEZE_SHORTEN                   uint16
	FIGHT_PROP_TORPOR_SHORTEN                   uint16
	FIGHT_PROP_DIZZY_SHORTEN                    uint16
	FIGHT_PROP_MAX_FIRE_ENERGY                  uint16
	FIGHT_PROP_MAX_ELEC_ENERGY                  uint16
	FIGHT_PROP_MAX_WATER_ENERGY                 uint16
	FIGHT_PROP_MAX_GRASS_ENERGY                 uint16
	FIGHT_PROP_MAX_WIND_ENERGY                  uint16
	FIGHT_PROP_MAX_ICE_ENERGY                   uint16
	FIGHT_PROP_MAX_ROCK_ENERGY                  uint16
	FIGHT_PROP_SKILL_CD_MINUS_RATIO             uint16
	FIGHT_PROP_SHIELD_COST_MINUS_RATIO          uint16
	FIGHT_PROP_CUR_FIRE_ENERGY                  uint16
	FIGHT_PROP_CUR_ELEC_ENERGY                  uint16
	FIGHT_PROP_CUR_WATER_ENERGY                 uint16
	FIGHT_PROP_CUR_GRASS_ENERGY                 uint16
	FIGHT_PROP_CUR_WIND_ENERGY                  uint16
	FIGHT_PROP_CUR_ICE_ENERGY                   uint16
	FIGHT_PROP_CUR_ROCK_ENERGY                  uint16
	FIGHT_PROP_CUR_HP                           uint16
	FIGHT_PROP_MAX_HP                           uint16
	FIGHT_PROP_CUR_ATTACK                       uint16
	FIGHT_PROP_CUR_DEFENSE                      uint16
	FIGHT_PROP_CUR_SPEED                        uint16
	FIGHT_PROP_NONEXTRA_ATTACK                  uint16
	FIGHT_PROP_NONEXTRA_DEFENSE                 uint16
	FIGHT_PROP_NONEXTRA_CRITICAL                uint16
	FIGHT_PROP_NONEXTRA_ANTI_CRITICAL           uint16
	FIGHT_PROP_NONEXTRA_CRITICAL_HURT           uint16
	FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY       uint16
	FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY         uint16
	FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT       uint16
	FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT           uint16
	FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT           uint16
	FIGHT_PROP_NONEXTRA_WATER_ADD_HURT          uint16
	FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT          uint16
	FIGHT_PROP_NONEXTRA_WIND_ADD_HURT           uint16
	FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT           uint16
	FIGHT_PROP_NONEXTRA_ICE_ADD_HURT            uint16
	FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT           uint16
	FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT           uint16
	FIGHT_PROP_NONEXTRA_WATER_SUB_HURT          uint16
	FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT          uint16
	FIGHT_PROP_NONEXTRA_WIND_SUB_HURT           uint16
	FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT           uint16
	FIGHT_PROP_NONEXTRA_ICE_SUB_HURT            uint16
	FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO    uint16
	FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO uint16
	FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT       uint16
	STRING_MAP                                  map[string]uint16
}

func InitFightPropertyConst() {
	FightPropertyConst = new(FightProperty)

	FightPropertyConst.FIGHT_PROP_NONE = 0
	FightPropertyConst.FIGHT_PROP_BASE_HP = 1
	FightPropertyConst.FIGHT_PROP_HP = 2
	FightPropertyConst.FIGHT_PROP_HP_PERCENT = 3
	FightPropertyConst.FIGHT_PROP_BASE_ATTACK = 4
	FightPropertyConst.FIGHT_PROP_ATTACK = 5
	FightPropertyConst.FIGHT_PROP_ATTACK_PERCENT = 6
	FightPropertyConst.FIGHT_PROP_BASE_DEFENSE = 7
	FightPropertyConst.FIGHT_PROP_DEFENSE = 8
	FightPropertyConst.FIGHT_PROP_DEFENSE_PERCENT = 9
	FightPropertyConst.FIGHT_PROP_BASE_SPEED = 10
	FightPropertyConst.FIGHT_PROP_SPEED_PERCENT = 11
	FightPropertyConst.FIGHT_PROP_HP_MP_PERCENT = 12
	FightPropertyConst.FIGHT_PROP_ATTACK_MP_PERCENT = 13
	FightPropertyConst.FIGHT_PROP_CRITICAL = 20
	FightPropertyConst.FIGHT_PROP_ANTI_CRITICAL = 21
	FightPropertyConst.FIGHT_PROP_CRITICAL_HURT = 22
	FightPropertyConst.FIGHT_PROP_CHARGE_EFFICIENCY = 23
	FightPropertyConst.FIGHT_PROP_ADD_HURT = 24
	FightPropertyConst.FIGHT_PROP_SUB_HURT = 25
	FightPropertyConst.FIGHT_PROP_HEAL_ADD = 26
	FightPropertyConst.FIGHT_PROP_HEALED_ADD = 27
	FightPropertyConst.FIGHT_PROP_ELEMENT_MASTERY = 28
	FightPropertyConst.FIGHT_PROP_PHYSICAL_SUB_HURT = 29
	FightPropertyConst.FIGHT_PROP_PHYSICAL_ADD_HURT = 30
	FightPropertyConst.FIGHT_PROP_DEFENCE_IGNORE_RATIO = 31
	FightPropertyConst.FIGHT_PROP_DEFENCE_IGNORE_DELTA = 32
	FightPropertyConst.FIGHT_PROP_FIRE_ADD_HURT = 40
	FightPropertyConst.FIGHT_PROP_ELEC_ADD_HURT = 41
	FightPropertyConst.FIGHT_PROP_WATER_ADD_HURT = 42
	FightPropertyConst.FIGHT_PROP_GRASS_ADD_HURT = 43
	FightPropertyConst.FIGHT_PROP_WIND_ADD_HURT = 44
	FightPropertyConst.FIGHT_PROP_ROCK_ADD_HURT = 45
	FightPropertyConst.FIGHT_PROP_ICE_ADD_HURT = 46
	FightPropertyConst.FIGHT_PROP_HIT_HEAD_ADD_HURT = 47
	FightPropertyConst.FIGHT_PROP_FIRE_SUB_HURT = 50
	FightPropertyConst.FIGHT_PROP_ELEC_SUB_HURT = 51
	FightPropertyConst.FIGHT_PROP_WATER_SUB_HURT = 52
	FightPropertyConst.FIGHT_PROP_GRASS_SUB_HURT = 53
	FightPropertyConst.FIGHT_PROP_WIND_SUB_HURT = 54
	FightPropertyConst.FIGHT_PROP_ROCK_SUB_HURT = 55
	FightPropertyConst.FIGHT_PROP_ICE_SUB_HURT = 56
	FightPropertyConst.FIGHT_PROP_EFFECT_HIT = 60
	FightPropertyConst.FIGHT_PROP_EFFECT_RESIST = 61
	FightPropertyConst.FIGHT_PROP_FREEZE_RESIST = 62
	FightPropertyConst.FIGHT_PROP_TORPOR_RESIST = 63
	FightPropertyConst.FIGHT_PROP_DIZZY_RESIST = 64
	FightPropertyConst.FIGHT_PROP_FREEZE_SHORTEN = 65
	FightPropertyConst.FIGHT_PROP_TORPOR_SHORTEN = 66
	FightPropertyConst.FIGHT_PROP_DIZZY_SHORTEN = 67
	FightPropertyConst.FIGHT_PROP_MAX_FIRE_ENERGY = 70
	FightPropertyConst.FIGHT_PROP_MAX_ELEC_ENERGY = 71
	FightPropertyConst.FIGHT_PROP_MAX_WATER_ENERGY = 72
	FightPropertyConst.FIGHT_PROP_MAX_GRASS_ENERGY = 73
	FightPropertyConst.FIGHT_PROP_MAX_WIND_ENERGY = 74
	FightPropertyConst.FIGHT_PROP_MAX_ICE_ENERGY = 75
	FightPropertyConst.FIGHT_PROP_MAX_ROCK_ENERGY = 76
	FightPropertyConst.FIGHT_PROP_SKILL_CD_MINUS_RATIO = 80
	FightPropertyConst.FIGHT_PROP_SHIELD_COST_MINUS_RATIO = 81
	FightPropertyConst.FIGHT_PROP_CUR_FIRE_ENERGY = 1000
	FightPropertyConst.FIGHT_PROP_CUR_ELEC_ENERGY = 1001
	FightPropertyConst.FIGHT_PROP_CUR_WATER_ENERGY = 1002
	FightPropertyConst.FIGHT_PROP_CUR_GRASS_ENERGY = 1003
	FightPropertyConst.FIGHT_PROP_CUR_WIND_ENERGY = 1004
	FightPropertyConst.FIGHT_PROP_CUR_ICE_ENERGY = 1005
	FightPropertyConst.FIGHT_PROP_CUR_ROCK_ENERGY = 1006
	FightPropertyConst.FIGHT_PROP_CUR_HP = 1010
	FightPropertyConst.FIGHT_PROP_MAX_HP = 2000
	FightPropertyConst.FIGHT_PROP_CUR_ATTACK = 2001
	FightPropertyConst.FIGHT_PROP_CUR_DEFENSE = 2002
	FightPropertyConst.FIGHT_PROP_CUR_SPEED = 2003
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ATTACK = 3000
	FightPropertyConst.FIGHT_PROP_NONEXTRA_DEFENSE = 3001
	FightPropertyConst.FIGHT_PROP_NONEXTRA_CRITICAL = 3002
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ANTI_CRITICAL = 3003
	FightPropertyConst.FIGHT_PROP_NONEXTRA_CRITICAL_HURT = 3004
	FightPropertyConst.FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY = 3005
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY = 3006
	FightPropertyConst.FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT = 3007
	FightPropertyConst.FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT = 3008
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT = 3009
	FightPropertyConst.FIGHT_PROP_NONEXTRA_WATER_ADD_HURT = 3010
	FightPropertyConst.FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT = 3011
	FightPropertyConst.FIGHT_PROP_NONEXTRA_WIND_ADD_HURT = 3012
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT = 3013
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ICE_ADD_HURT = 3014
	FightPropertyConst.FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT = 3015
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT = 3016
	FightPropertyConst.FIGHT_PROP_NONEXTRA_WATER_SUB_HURT = 3017
	FightPropertyConst.FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT = 3018
	FightPropertyConst.FIGHT_PROP_NONEXTRA_WIND_SUB_HURT = 3019
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT = 3020
	FightPropertyConst.FIGHT_PROP_NONEXTRA_ICE_SUB_HURT = 3021
	FightPropertyConst.FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO = 3022
	FightPropertyConst.FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO = 3023
	FightPropertyConst.FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT = 3024

	FightPropertyConst.STRING_MAP = make(map[string]uint16)

	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONE"] = 0
	FightPropertyConst.STRING_MAP["FIGHT_PROP_BASE_HP"] = 1
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HP"] = 2
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HP_PERCENT"] = 3
	FightPropertyConst.STRING_MAP["FIGHT_PROP_BASE_ATTACK"] = 4
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ATTACK"] = 5
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ATTACK_PERCENT"] = 6
	FightPropertyConst.STRING_MAP["FIGHT_PROP_BASE_DEFENSE"] = 7
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DEFENSE"] = 8
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DEFENSE_PERCENT"] = 9
	FightPropertyConst.STRING_MAP["FIGHT_PROP_BASE_SPEED"] = 10
	FightPropertyConst.STRING_MAP["FIGHT_PROP_SPEED_PERCENT"] = 11
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HP_MP_PERCENT"] = 12
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ATTACK_MP_PERCENT"] = 13
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CRITICAL"] = 20
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ANTI_CRITICAL"] = 21
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CRITICAL_HURT"] = 22
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CHARGE_EFFICIENCY"] = 23
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ADD_HURT"] = 24
	FightPropertyConst.STRING_MAP["FIGHT_PROP_SUB_HURT"] = 25
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HEAL_ADD"] = 26
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HEALED_ADD"] = 27
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ELEMENT_MASTERY"] = 28
	FightPropertyConst.STRING_MAP["FIGHT_PROP_PHYSICAL_SUB_HURT"] = 29
	FightPropertyConst.STRING_MAP["FIGHT_PROP_PHYSICAL_ADD_HURT"] = 30
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DEFENCE_IGNORE_RATIO"] = 31
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DEFENCE_IGNORE_DELTA"] = 32
	FightPropertyConst.STRING_MAP["FIGHT_PROP_FIRE_ADD_HURT"] = 40
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ELEC_ADD_HURT"] = 41
	FightPropertyConst.STRING_MAP["FIGHT_PROP_WATER_ADD_HURT"] = 42
	FightPropertyConst.STRING_MAP["FIGHT_PROP_GRASS_ADD_HURT"] = 43
	FightPropertyConst.STRING_MAP["FIGHT_PROP_WIND_ADD_HURT"] = 44
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ROCK_ADD_HURT"] = 45
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ICE_ADD_HURT"] = 46
	FightPropertyConst.STRING_MAP["FIGHT_PROP_HIT_HEAD_ADD_HURT"] = 47
	FightPropertyConst.STRING_MAP["FIGHT_PROP_FIRE_SUB_HURT"] = 50
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ELEC_SUB_HURT"] = 51
	FightPropertyConst.STRING_MAP["FIGHT_PROP_WATER_SUB_HURT"] = 52
	FightPropertyConst.STRING_MAP["FIGHT_PROP_GRASS_SUB_HURT"] = 53
	FightPropertyConst.STRING_MAP["FIGHT_PROP_WIND_SUB_HURT"] = 54
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ROCK_SUB_HURT"] = 55
	FightPropertyConst.STRING_MAP["FIGHT_PROP_ICE_SUB_HURT"] = 56
	FightPropertyConst.STRING_MAP["FIGHT_PROP_EFFECT_HIT"] = 60
	FightPropertyConst.STRING_MAP["FIGHT_PROP_EFFECT_RESIST"] = 61
	FightPropertyConst.STRING_MAP["FIGHT_PROP_FREEZE_RESIST"] = 62
	FightPropertyConst.STRING_MAP["FIGHT_PROP_TORPOR_RESIST"] = 63
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DIZZY_RESIST"] = 64
	FightPropertyConst.STRING_MAP["FIGHT_PROP_FREEZE_SHORTEN"] = 65
	FightPropertyConst.STRING_MAP["FIGHT_PROP_TORPOR_SHORTEN"] = 66
	FightPropertyConst.STRING_MAP["FIGHT_PROP_DIZZY_SHORTEN"] = 67
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_FIRE_ENERGY"] = 70
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_ELEC_ENERGY"] = 71
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_WATER_ENERGY"] = 72
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_GRASS_ENERGY"] = 73
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_WIND_ENERGY"] = 74
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_ICE_ENERGY"] = 75
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_ROCK_ENERGY"] = 76
	FightPropertyConst.STRING_MAP["FIGHT_PROP_SKILL_CD_MINUS_RATIO"] = 80
	FightPropertyConst.STRING_MAP["FIGHT_PROP_SHIELD_COST_MINUS_RATIO"] = 81
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_FIRE_ENERGY"] = 1000
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_ELEC_ENERGY"] = 1001
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_WATER_ENERGY"] = 1002
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_GRASS_ENERGY"] = 1003
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_WIND_ENERGY"] = 1004
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_ICE_ENERGY"] = 1005
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_ROCK_ENERGY"] = 1006
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_HP"] = 1010
	FightPropertyConst.STRING_MAP["FIGHT_PROP_MAX_HP"] = 2000
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_ATTACK"] = 2001
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_DEFENSE"] = 2002
	FightPropertyConst.STRING_MAP["FIGHT_PROP_CUR_SPEED"] = 2003
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ATTACK"] = 3000
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_DEFENSE"] = 3001
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_CRITICAL"] = 3002
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ANTI_CRITICAL"] = 3003
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_CRITICAL_HURT"] = 3004
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY"] = 3005
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY"] = 3006
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT"] = 3007
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT"] = 3008
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT"] = 3009
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_WATER_ADD_HURT"] = 3010
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT"] = 3011
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_WIND_ADD_HURT"] = 3012
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT"] = 3013
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ICE_ADD_HURT"] = 3014
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT"] = 3015
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT"] = 3016
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_WATER_SUB_HURT"] = 3017
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT"] = 3018
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_WIND_SUB_HURT"] = 3019
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT"] = 3020
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_ICE_SUB_HURT"] = 3021
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO"] = 3022
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO"] = 3023
	FightPropertyConst.STRING_MAP["FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT"] = 3024
}
