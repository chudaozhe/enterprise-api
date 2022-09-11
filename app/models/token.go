package models

import (
	"enterprise-api/core/cache"
	"enterprise-api/core/helper"
	"errors"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func key(id int, role string) string {
	return "gin.www." + role + "." + strconv.Itoa(id) + ".token"
}

func GetToken(id int, role string) (string, error) {
	key := key(id, role)
	token, err := cache.RedisClient.Get(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) { //key不存在
			return "", nil
		}
		// 出其他错了
		return "", err
	}
	return token, err
}

func CheckToken(id int, token string, role string) (err error) {
	if len(token) == 0 {
		err = errors.New("未传递token")
		return
	}
	_token, err0 := GetToken(id, role)
	if err0 != nil {
		err = errors.New("token过期或不存在")
		return
	}
	if _token != token {
		err = errors.New("token不匹配")
		return
	}
	if _token == token { //验证成功
		return
	}
	err = errors.New("未知错误")
	return
}

func SetToken(id int, reToken bool, expire time.Duration, role string) (token string, err error) {
	key := key(id, role)
	if reToken {
		token = helper.Random(8, "")
	} else {
		_token, err0 := GetToken(id, role)
		if err0 != nil {
			return
		}
		token = _token
	}
	if expire == 0 {
		expire = 3600 * 24 * 7
	}
	_, err = cache.RedisClient.Set(key, token, expire*time.Second).Result()
	if err != nil {
		return
	}
	return
}

func DelToken(id int, role string) bool {
	key := key(id, role)
	_, err := cache.RedisClient.Del(key).Result()
	if err != nil {
		return false
	}
	return true
}
