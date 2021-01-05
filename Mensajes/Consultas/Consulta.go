package Consulta

import (
	"log"
  	"golang.org/x/net/context"
	"io"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Server struct{
}


func (s *Server) ConsultaCliente(ctx context.Context, consul *Consulta) (*Respuesta, error) {

	fmt.Println("Guardando distribucion de chunks en LOG")
	start := time.Now()

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
	resp := Respuesta {
		Resultado: "GUARDADO EXITOSAMENTE EN LOG",
	}

	return &resp, nil
}

func (s *Server) ConsultaBroker(ctx context.Context, consul *Consulta) (*Respuesta, error) {

	fmt.Println("Buscando " + consul.NombreDominio + " en los archivos ZF")
	
	NombreDominio := strings.Split(consul.NombreDominio, ".")
	ip := "Not Found"

	input, erro := ioutil.ReadFile("./Registros/ZF-" + NombreDominio[1])
	if erro != nil {
		fmt.Println("No se ha encontrado el archivo ZF-" + NombreDominio[1])
		resp := Respuesta {
			IP: ip,
			Reloj: "[" + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]) + "]",
		}
	
		return &resp, nil
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, consul.NombreDominio) {
			ip = strings.Split(line, " ")[3]
		}
	}

	resp := Respuesta {
		IP: ip,
		Reloj: "[" + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]) + "]",
	}

	return &resp, nil
}