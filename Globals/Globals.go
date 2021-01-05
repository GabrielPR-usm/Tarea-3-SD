package Globals

var RVDNS1 = make(map[string] [3]int)//Reloj vectorial: Diccionario con Key[dominio]

func SetterR1(key string){

	RVDNS1[key] = [3]int{RVDNS1[key][0]+1, RVDNS1[key][1], RVDNS1[key][2]}
	
}