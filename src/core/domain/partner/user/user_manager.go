/**
 * Copyright 2014 @ ops Inc.
 * name :
 * author : jarryliu
 * date : 2014-02-14 17:46
 * description :
 * history :
 */
package user

import (
	"go2o/src/core/domain/interface/partner/user"
)

var _ user.IUserManager = new(UserManager)

type UserManager struct {
	partnerId int
	rep       user.IUserRep
}

func NewUserManager(partnerId int, rep user.IUserRep) user.IUserManager {
	return &UserManager{
		partnerId: partnerId,
		rep:       rep,
	}
}

// 获取单个用户
func (this *UserManager) GetUser(id int) user.IUser {
	v := this.rep.GetPersonValue(id)
	if v != nil {
		return newUser(v, this.rep)
	}
	return nil
}

// 获取所有配送员
func (this *UserManager) GetDeliveryStaff() []user.IDeliveryStaff {
	list := this.rep.GetDeliveryStaffPersons(this.partnerId)
	var staffs = make([]user.IDeliveryStaff, len(list))
	for i, v := range list {
		staffs[i] = newUser(v, this.rep)
	}
	return staffs
}
