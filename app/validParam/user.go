package validParam

import (
	"course-system/app/common"
	"github.com/go-playground/validator/v10"
)

func diyValidParam(request common.CreateMemberRequest) bool {
	flag := true
	test1 := request.Username
	for i := 0; i < len(test1); i++ {
		if test1[i] < 65 || test1[i] > 90 && test1[i] < 97 || test1[i] > 122 {
			flag = false
			break
		}
	}
	test2 := request.Password
	a := [3]int{0, 0, 0}
	for i := 0; i < len(test2); i++ {
		if test2[i] >= 65 && test2[i] <= 90 {
			a[0] = 1
		}
		if test2[i] >= 97 && test2[i] <= 122 {
			a[1] = 1
		}
		if test2[i] >= 48 && test2[i] <= 57 {
			a[2] = 1
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] != 1 {
			flag = false
			break
		}
	}
	return flag
}

func CreateUserValidParam(request common.CreateMemberRequest) bool {
	validate := validator.New()
	err := validate.Struct(request)
	return diyValidParam(request) && (err == nil)
}

func UpdateUserValidParam(request common.UpdateMemberRequest) bool {
	validate := validator.New()
	err := validate.Struct(request)
	return err == nil
}
