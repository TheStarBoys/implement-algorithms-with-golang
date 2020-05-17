package leetcode


/**

Definition for Employee.
type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

var hash map[int]*Employee
func getImportance(employees []*Employee, id int) int {
	hash = make(map[int]*Employee)
	for _, v := range employees {
		hash[v.Id] = v
	}
	return dfs(id)
}

func dfs(id int) int {
	cur := hash[id]
	res := cur.Importance
	for _, v := range cur.Subordinates {
		res += dfs(v)
	}

	return res
}

 */