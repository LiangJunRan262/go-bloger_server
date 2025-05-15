package redis_jwt

import (
	"bloger_server/global"
	"bloger_server/utils/jwts"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type BlackType int8

const (
	UserBlackToken  BlackType = iota
	AdminBlackType            // 管理员手动下线
	DeviceBlackType           // 设备把自己挤掉了
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
	//switch b {
	//case UserBlackToken:
	//	return "用户token"
	//case AdminBlackType:
	//	return "管理员手动下线"
	//case DeviceBlackType:
	//	return "设备把自己挤掉了"
	//default:
	//	return "未知"
	//}
}

func ParseBlackTyoe(val string) BlackType {
	switch val {
	case "0":
		return UserBlackToken
	case "1":
		return AdminBlackType
	case "2":
		return DeviceBlackType
	default:
		return UserBlackToken
	}
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)

	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token：【%s】解析失败", token)
		return
	}

	seconds := claims.ExpiresAt - claims.IssuedAt
	if seconds <= 0 {
		logrus.Errorf("token：【%s】过期时间小于等于0", token)
		return
	}

	res, err := global.Redis.Set(global.CommonContext, key, value.String(), time.Duration(seconds)*time.Hour).Result()
	if err != nil {
		logrus.Errorf("redis添加黑名单失败")
		return
	}
	logrus.Infof("token<UNK>%s<UNK> 已加入黑名单，原因：%s", token, res)
}

func HasTokenBlack(token string) (blk BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(global.CommonContext, key).Result()
	if err != nil {
		logrus.Errorf("token：【%s】不存在", token)
		return blk, false
	}
	blk = ParseBlackTyoe(value)
	return blk, true
}
