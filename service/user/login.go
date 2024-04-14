package user

import (
	receive "fast-gin/interaction/receive/user"
)

func Register(data *receive.RegisterReceive) (results interface{}, err error) {
	return data, nil
}

func Login(data *receive.LoginReceive) (results interface{}, err error) {
	return data, nil
}
