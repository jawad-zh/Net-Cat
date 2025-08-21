package functions 

import(
	
)

var Clients[]string

func ChekConn(Name string)(bool){
	for _,r:= range Clients{
		if string(r) == Name{
			return false
			
		}
	}
	return true
	
}