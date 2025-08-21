package functions

import "time"



var historique []string


func Historique(t time.Time,format string){
	timeStr:= t.Format("2006 01 02 15:04:05")
	addTime:= timeStr + format
	mu.Lock()
	historique= append(historique, addTime)
	mu.Unlock()

}