package functions

func Remov(users []string, user string) []string {
	var Rusers []string
	for _,r:= range users{
		if r != user{
			mu.Lock()
			Rusers = append(Rusers, r)
			mu.Unlock()
		}
	}
	return Rusers
}
