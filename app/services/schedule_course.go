package services

var MapTeacher = map[string][]string{}
var MapCourse = map[string][]string{}
var vis = map[string]bool{} //记录课程是否已被访问过
var p = map[string]string{} //记录当前课程被哪位老师选中
var res = map[string]string{}

func initSchedule() {
	MapTeacher = map[string][]string{}
	MapCourse = map[string][]string{}
	vis = map[string]bool{} //记录课程是否已被访问过
	p = map[string]string{} //记录当前课程被哪位老师选中
	res = map[string]string{}
}

func match(teacherID string) bool {
	for _, courseID := range MapTeacher[teacherID] {
		if !vis[courseID] { //有边且未访问
			vis[courseID] = true //记录状态未访问过
			_, ok := p[courseID]
			if !ok || match(p[courseID]) { //如果暂无匹配，或者原来匹配的左侧元素可以找到新的匹配
				p[courseID] = teacherID //当前左侧元素成为当前右侧元素的新匹配
				return true             //返回匹配成功
			}
		}
	}
	return false
}

func KM(mapTeacher map[string][]string) map[string]string {
	initSchedule()
	MapTeacher = mapTeacher
	for teacherID, courseIDs := range MapTeacher {
		for _, courseID := range courseIDs {
			MapCourse[courseID] = append(MapCourse[courseID], teacherID)
		}
	}
	for teacherID, _ := range MapTeacher {
		vis = make(map[string]bool)
		match(teacherID)
	}
	for i, j := range p {
		res[j] = i
	}
	return p
}
