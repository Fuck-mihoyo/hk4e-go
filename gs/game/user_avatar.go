package game

import (
	gdc "hk4e/gs/config"
	"hk4e/gs/constant"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/object"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

func (g *GameManager) GetAllAvatarDataConfig() map[int32]*gdc.AvatarData {
	allAvatarDataConfig := make(map[int32]*gdc.AvatarData)
	for avatarId, avatarData := range gdc.CONF.AvatarDataMap {
		if avatarId < 10000002 || avatarId >= 11000000 {
			// 跳过无效角色
			continue
		}
		if avatarId == 10000005 || avatarId == 10000007 {
			// 跳过主角
			continue
		}
		allAvatarDataConfig[avatarId] = avatarData
	}
	return allAvatarDataConfig
}

func (g *GameManager) AddUserAvatar(userId uint32, avatarId uint32) {
	player := g.userManager.GetOnlineUser(userId)
	if player == nil {
		logger.LOG.Error("player is nil, uid: %v", userId)
		return
	}
	player.AddAvatar(avatarId)
	avatar := player.AvatarMap[avatarId]
	if avatar == nil {
		logger.LOG.Error("avatar is nil, avatarId", avatarId)
		return
	}

	// 添加初始武器
	avatarDataConfig, ok := gdc.CONF.AvatarDataMap[int32(avatarId)]
	if !ok {
		logger.LOG.Error("config is nil, itemId: %v", avatarId)
		return
	}
	weaponId := g.AddUserWeapon(player.PlayerID, uint32(avatarDataConfig.InitialWeapon))

	// 角色装上初始武器
	g.WearUserAvatarEquip(player.PlayerID, avatarId, weaponId)

	// TODO 真的有必要存在吗
	g.UpdateUserAvatarFightProp(player.PlayerID, avatarId)

	// PacketAvatarAddNotify
	avatarAddNotify := new(proto.AvatarAddNotify)
	avatarAddNotify.Avatar = g.PacketAvatarInfo(avatar)
	avatarAddNotify.IsInTeam = false
	g.SendMsg(cmd.AvatarAddNotify, userId, player.ClientSeq, avatarAddNotify)
}

func (g *GameManager) WearEquipReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user wear equip, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.WearEquipReq)
	avatarGuid := req.AvatarGuid
	equipGuid := req.EquipGuid
	avatar := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	weapon := player.GameObjectGuidMap[equipGuid].(*model.Weapon)
	g.WearUserAvatarEquip(player.PlayerID, avatar.AvatarId, weapon.WeaponId)

	// PacketWearEquipRsp
	wearEquipRsp := new(proto.WearEquipRsp)
	wearEquipRsp.AvatarGuid = avatarGuid
	wearEquipRsp.EquipGuid = equipGuid
	g.SendMsg(cmd.WearEquipRsp, player.PlayerID, player.ClientSeq, wearEquipRsp)
}

func (g *GameManager) WearUserAvatarEquip(userId uint32, avatarId uint32, weaponId uint64) {
	player := g.userManager.GetOnlineUser(userId)
	if player == nil {
		logger.LOG.Error("player is nil, uid: %v", userId)
		return
	}
	avatar := player.AvatarMap[avatarId]
	weapon := player.WeaponMap[weaponId]

	world := g.worldManager.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)
	playerTeamEntity := scene.GetPlayerTeamEntity(player.PlayerID)
	team := player.TeamConfig.GetActiveTeam()

	if weapon.AvatarId != 0 {
		// 武器在别的角色身上
		weakAvatarId := weapon.AvatarId
		weakWeaponId := weaponId
		strongAvatarId := avatarId
		strongWeaponId := avatar.EquipWeapon.WeaponId
		player.TakeOffWeapon(weakAvatarId, weakWeaponId)
		player.TakeOffWeapon(strongAvatarId, strongWeaponId)
		player.WearWeapon(weakAvatarId, strongWeaponId)
		player.WearWeapon(strongAvatarId, weakWeaponId)

		weakAvatar := player.AvatarMap[weakAvatarId]
		weakWeapon := player.WeaponMap[weakAvatar.EquipWeapon.WeaponId]

		for _, aid := range team.AvatarIdList {
			if aid == 0 {
				break
			}
			if aid == weakAvatar.AvatarId {
				playerTeamEntity.weaponEntityMap[weakWeapon.WeaponId] = scene.CreateEntityWeapon()
			}
		}

		// PacketAvatarEquipChangeNotify
		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(weakAvatar, weakWeapon, playerTeamEntity.weaponEntityMap[weakWeapon.WeaponId])
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	} else if avatar.EquipWeapon != nil {
		// 角色当前有武器
		player.TakeOffWeapon(avatarId, avatar.EquipWeapon.WeaponId)
		player.WearWeapon(avatarId, weaponId)
	} else {
		// 是新角色还没有武器
		player.WearWeapon(avatarId, weaponId)
	}

	for _, aid := range team.AvatarIdList {
		if aid == 0 {
			break
		}
		if aid == avatarId {
			playerTeamEntity.weaponEntityMap[weaponId] = scene.CreateEntityWeapon()
		}
	}

	// PacketAvatarEquipChangeNotify
	avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(avatar, weapon, playerTeamEntity.weaponEntityMap[weaponId])
	g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
}

func (g *GameManager) AvatarChangeCostumeReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user change avatar costume, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarChangeCostumeReq)
	avatarGuid := req.AvatarGuid
	costumeId := req.CostumeId

	exist := false
	for _, v := range player.CostumeList {
		if v == costumeId {
			exist = true
		}
	}
	if costumeId == 0 {
		exist = true
	}
	if !exist {
		return
	}

	avatar := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	avatar.Costume = req.CostumeId

	world := g.worldManager.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	// PacketAvatarChangeCostumeNotify
	avatarChangeCostumeNotify := new(proto.AvatarChangeCostumeNotify)
	avatarChangeCostumeNotify.EntityInfo = g.PacketSceneEntityInfoAvatar(scene, player, avatar.AvatarId)
	for _, scenePlayer := range scene.playerMap {
		g.SendMsg(cmd.AvatarChangeCostumeNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, avatarChangeCostumeNotify)
	}

	// PacketAvatarChangeCostumeRsp
	avatarChangeCostumeRsp := new(proto.AvatarChangeCostumeRsp)
	avatarChangeCostumeRsp.AvatarGuid = req.AvatarGuid
	avatarChangeCostumeRsp.CostumeId = req.CostumeId
	g.SendMsg(cmd.AvatarChangeCostumeRsp, player.PlayerID, player.ClientSeq, avatarChangeCostumeRsp)
}

func (g *GameManager) AvatarWearFlycloakReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user change avatar fly cloak, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarWearFlycloakReq)
	avatarGuid := req.AvatarGuid
	flycloakId := req.FlycloakId

	exist := false
	for _, v := range player.FlyCloakList {
		if v == flycloakId {
			exist = true
		}
	}
	if !exist {
		return
	}

	avatar := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	avatar.FlyCloak = req.FlycloakId

	world := g.worldManager.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	// PacketAvatarFlycloakChangeNotify
	avatarFlycloakChangeNotify := new(proto.AvatarFlycloakChangeNotify)
	avatarFlycloakChangeNotify.AvatarGuid = avatarGuid
	avatarFlycloakChangeNotify.FlycloakId = flycloakId
	for _, scenePlayer := range scene.playerMap {
		g.SendMsg(cmd.AvatarFlycloakChangeNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, avatarFlycloakChangeNotify)
	}

	// PacketAvatarWearFlycloakRsp
	avatarWearFlycloakRsp := new(proto.AvatarWearFlycloakRsp)
	avatarWearFlycloakRsp.AvatarGuid = req.AvatarGuid
	avatarWearFlycloakRsp.FlycloakId = req.FlycloakId
	g.SendMsg(cmd.AvatarWearFlycloakRsp, player.PlayerID, player.ClientSeq, avatarWearFlycloakRsp)
}

func (g *GameManager) PacketAvatarEquipChangeNotify(avatar *model.Avatar, weapon *model.Weapon, entityId uint32) *proto.AvatarEquipChangeNotify {
	itemDataConfig, ok := gdc.CONF.ItemDataMap[int32(weapon.ItemId)]
	avatarEquipChangeNotify := new(proto.AvatarEquipChangeNotify)
	avatarEquipChangeNotify.AvatarGuid = avatar.Guid
	if ok {
		avatarEquipChangeNotify.EquipType = uint32(itemDataConfig.EquipEnumType)
	}
	avatarEquipChangeNotify.ItemId = weapon.ItemId
	avatarEquipChangeNotify.EquipGuid = weapon.Guid
	avatarEquipChangeNotify.Weapon = &proto.SceneWeaponInfo{
		EntityId:    entityId,
		GadgetId:    uint32(gdc.CONF.ItemDataMap[int32(weapon.ItemId)].GadgetId),
		ItemId:      weapon.ItemId,
		Guid:        weapon.Guid,
		Level:       uint32(weapon.Level),
		AbilityInfo: new(proto.AbilitySyncStateInfo),
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) PacketAvatarEquipTakeOffNotify(avatar *model.Avatar, weapon *model.Weapon) *proto.AvatarEquipChangeNotify {
	avatarEquipChangeNotify := new(proto.AvatarEquipChangeNotify)
	avatarEquipChangeNotify.AvatarGuid = avatar.Guid
	itemDataConfig, ok := gdc.CONF.ItemDataMap[int32(weapon.ItemId)]
	if ok {
		avatarEquipChangeNotify.EquipType = uint32(itemDataConfig.EquipEnumType)
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) UpdateUserAvatarFightProp(userId uint32, avatarId uint32) {
	player := g.userManager.GetOnlineUser(userId)
	if player == nil {
		logger.LOG.Error("player is nil, uid: %v", userId)
		return
	}
	avatarFightPropNotify := new(proto.AvatarFightPropNotify)
	avatar := player.AvatarMap[avatarId]
	avatarFightPropNotify.AvatarGuid = avatar.Guid
	avatarFightPropNotify.FightPropMap = avatar.FightPropMap
	g.SendMsg(cmd.AvatarFightPropNotify, userId, player.ClientSeq, avatarFightPropNotify)
}

func (g *GameManager) PacketAvatarInfo(avatar *model.Avatar) *proto.AvatarInfo {
	isFocus := false
	//if avatar.AvatarId == 10000005 || avatar.AvatarId == 10000007 {
	//	isFocus = true
	//}
	pbAvatar := &proto.AvatarInfo{
		IsFocus:  isFocus,
		AvatarId: avatar.AvatarId,
		Guid:     avatar.Guid,
		PropMap: map[uint32]*proto.PropValue{
			uint32(constant.PlayerPropertyConst.PROP_LEVEL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_LEVEL),
				Val:   int64(avatar.Level),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Level)},
			},
			uint32(constant.PlayerPropertyConst.PROP_EXP): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_EXP),
				Val:   int64(avatar.Exp),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Exp)},
			},
			uint32(constant.PlayerPropertyConst.PROP_BREAK_LEVEL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_BREAK_LEVEL),
				Val:   int64(avatar.Promote),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Promote)},
			},
			uint32(constant.PlayerPropertyConst.PROP_SATIATION_VAL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_SATIATION_VAL),
				Val:   0,
				Value: &proto.PropValue_Ival{Ival: 0},
			},
			uint32(constant.PlayerPropertyConst.PROP_SATIATION_PENALTY_TIME): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_SATIATION_PENALTY_TIME),
				Val:   0,
				Value: &proto.PropValue_Ival{Ival: 0},
			},
		},
		LifeState:     1,
		EquipGuidList: object.ConvMapToList(avatar.EquipGuidList),
		FightPropMap:  nil,
		SkillDepotId:  avatar.SkillDepotId,
		FetterInfo: &proto.AvatarFetterInfo{
			ExpLevel:  uint32(avatar.FetterLevel),
			ExpNumber: avatar.FetterExp,
			// FetterList 不知道是啥 该角色在配置表里的所有FetterId
			// TODO 资料解锁条目
			FetterList:              nil,
			RewardedFetterLevelList: []uint32{10},
		},
		SkillLevelMap:     nil,
		AvatarType:        1,
		WearingFlycloakId: avatar.FlyCloak,
		CostumeId:         avatar.Costume,
		BornTime:          uint32(avatar.BornTime),
	}
	pbAvatar.FightPropMap = avatar.FightPropMap
	for _, v := range avatar.FetterList {
		pbAvatar.FetterInfo.FetterList = append(pbAvatar.FetterInfo.FetterList, &proto.FetterData{
			FetterId:    v,
			FetterState: uint32(constant.FetterStateConst.FINISH),
		})
	}
	// 解锁全部资料
	for _, v := range gdc.CONF.AvatarFetterDataMap[int32(avatar.AvatarId)] {
		pbAvatar.FetterInfo.FetterList = append(pbAvatar.FetterInfo.FetterList, &proto.FetterData{
			FetterId:    uint32(v),
			FetterState: uint32(constant.FetterStateConst.FINISH),
		})
	}
	pbAvatar.SkillLevelMap = make(map[uint32]uint32)
	for k, v := range avatar.SkillLevelMap {
		pbAvatar.SkillLevelMap[k] = v
	}
	return pbAvatar
}
