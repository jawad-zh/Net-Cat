package functions



var historique []string


func Historique(format string){
	mu.Lock()
	historique= append(historique, format)
	mu.Unlock()

}