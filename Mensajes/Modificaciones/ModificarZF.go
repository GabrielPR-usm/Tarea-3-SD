package ModificarZF

import (
	"log"
  	"golang.org/x/net/context"
	//"io"
	"fmt"
	//"google.golang.org/grpc"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	"github.com/GabrielPR-usm/Tarea-3-SD/Globals"
	"io/ioutil"
	"os"
	"strconv"
	//"time"
	"strings"
)

type Server struct{
}

func (s *Server) ModificarZFRequest(ctx context.Context, modif *Modificacion) (*RespuestaBroker, error) {

	fmt.Println("Guardando distribucion de chunks en LOG")
	//start := time.Now()

	/* GUARDAR EN LOG */
	/*
	fi, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	fi.WriteString(noti.BookName + " " + strconv.FormatUint(uint64(noti.NChunks1+noti.NChunks2+noti.NChunks3), 10) + "\n")
	
	var i int64 = 0
	for ;i<noti.NChunks1 ;i++{
		fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto1+"\n")
	}
	for ;i<(noti.NChunks1 +noti.NChunks2);i++{
	 	fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto2+"\n")
	}
	for ;i<(noti.NChunks1 +noti.NChunks2 +noti.NChunks3);i++{
	   	fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto3+"\n")
	}
	fi.Close()
	
	elapsed := time.Since(start)
	fmt.Println("Archivo LOG creado satisfactoriamente")
	fmt.Println("EXECUTION TIME: ", elapsed)
	
	*/
	resp := RespuestaBroker {
		IP: "GUARDADO EXITOSAMENTE EN LOG",
	}

	return &resp, nil
}

func (s *Server) ModificarZFEnDNS(ctx context.Context, modif *Modificacion) (*RespuestaDNS, error) {

	fmt.Println("Realizando " + modif.Accion + " de " + modif.ValorAfectado)
	
	NombreDominio := strings.Split(modif.ValorAfectado, ".")
	flag := false

	//Modifico Archivo ZF
	if modif.Accion == "append" {

		fiz, errz := os.OpenFile("./Registros/ZF-" + NombreDominio[1] , os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if errz != nil {
			panic(errz)
		}

		fiz.WriteString("www." + modif.ValorAfectado + " IN A " + modif.NuevoValor + "\n")

		fiz.Close()

	}else{

		input, erro := ioutil.ReadFile("./Registros/ZF-" + NombreDominio[1])
		if erro != nil {
			fmt.Println("No se ha encontrado el archivo ZF-" + NombreDominio[1])
			resp := RespuestaDNS {
				Reloj: "[" + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]) + "]",
			}
		
			return &resp, nil
		}

		lines := strings.Split(string(input), "\n")

		if modif.Accion == "update" {

			for i, line := range lines {
				if strings.Contains(line, modif.ValorAfectado) {
					flag = true
					if len(strings.Split(modif.NuevoValor, ".")) == 2 {//Se esta cambiando el nombre
						lines[i] = "www." + modif.NuevoValor + " IN A " + strings.Split(line, " ")[3]
					}else{//Se esta cambiando la IP
						lines[i] = "www." + modif.ValorAfectado + " IN A " + modif.NuevoValor
					}
				}
			}
			if !flag {
				fmt.Println("No se ha encontrado el dominio a modificar")
			}

		}
		if modif.Accion == "delete" {

			for i, line := range lines {
				if strings.Contains(line, modif.ValorAfectado) {
					if flag {
						lines[i-1] = lines[i]
					}
					flag = true
				}
			}
			if !flag {
				fmt.Println("No se ha encontrado el dominio a borrar")
			}else{
				lines = lines[:len(lines)-1]
			}
		}

		output := strings.Join(lines, "\n")
		erro = ioutil.WriteFile("./Registros/ZF-" + NombreDominio[1], []byte(output), 0644)
		if erro != nil {
			log.Fatalln(erro)
		}
		
	}

	//Escribo el cambio en el LOG del archivo ZF correspondiente
	fi, err := os.OpenFile("./Registros/log-" + NombreDominio[1] , os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	fi.WriteString(modif.Accion + " " + modif.ValorAfectado + " " + modif.NuevoValor + "\n")

	fi.Close()

	if flag {
		Globals.SetterR1(NombreDominio[1])
	}

	resp := RespuestaDNS {
		Reloj: "[" + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]) + "]",
	}

	return &resp, nil
}