package Globals

import (
	"strconv"
)

var RVDNS1 = make(map[string] [3]int)//Reloj vectorial: Diccionario con Key[dominio]

func Setter(key string, num int){

	if num == 0{
		RVDNS1[key] = [3]int{RVDNS1[key][0]+1, RVDNS1[key][1], RVDNS1[key][2]}
	}else if num == 1{
		RVDNS1[key] = [3]int{RVDNS1[key][0], RVDNS1[key][1]+1, RVDNS1[key][2]}
	}else if num == 2{
		RVDNS1[key] = [3]int{RVDNS1[key][0], RVDNS1[key][1], RVDNS1[key][2]+1}
	}
}

func SetValue(key string, num int, value int){
	
	if num == 0{
		RVDNS1[key] = [3]int{value, RVDNS1[key][1], RVDNS1[key][2]}
	}else if num == 1{
		RVDNS1[key] = [3]int{RVDNS1[key][0], value, RVDNS1[key][2]}
	}else if num == 2{
		RVDNS1[key] = [3]int{RVDNS1[key][0], RVDNS1[key][1], value}
	}
}

func GetClock(key string) string {
	clock := strconv.Itoa(RVDNS1[key][0]) + "," + strconv.Itoa(RVDNS1[key][1]) + "," + strconv.Itoa(RVDNS1[key][2])
	return clock
}