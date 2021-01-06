package CompartirZF

import (
	//"log"
  	"golang.org/x/net/context"
	//"io"
	"fmt"
	"google.golang.org/grpc"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Globals"
	"io/ioutil"
	"os"
	"strconv"
	//"time"
	"strings"
	"path/filepath"
)

type Server struct{
}

var DNSnum int = 0

func (s *Server) Notificar(ctx context.Context, noti *Notificacion) (*Respuesta, error) {
	//Le llega a los DNS y comienzan a enviar modificacion tras modificacion a los otros dnss

	fmt.Println("Espere un momento, se estan unificando los registros")
	var files []string
/*
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println("No se puedo abrir " + dirname.Name())
	}

	dir, err := os.Open(path.Join(dirname, "../../"))
	if err != nil {
		fmt.Println("No se puedo abrir " + dir.Name())
	}
*/
    err := filepath.Walk("./Registros", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), "log") {
			files = append(files, info.Name())
		}
        return nil
    })
    if err != nil {
        panic(err)
	}

	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(noti.DNSOrchestrator, grpc.WithInsecure(),grpc.WithBlock())
	if err1 != nil{
		fmt.Println("No fue posible conectarse con el orquestador de la unificacion")
	}else{
		defer conn1.Close()
	}
	
	for _, file := range files {

		path := "./Registros/" + file

		input, erro := ioutil.ReadFile(path)
		if erro != nil {
			resp := Respuesta {
				Reloj: "",
			}
			return &resp, nil
		}

		strFile := string(input)

		mer := Merge {
			FileAsStr: strFile,
			Reloj: Globals.GetClock(strings.Split(file, "-")[1]),
			Dominio: strings.Split(file, "-")[1],
		}

		merge := NewMergearServiceClient(conn1)
		_, err01 := merge.Mergear(context.Background(), &mer)
		if err01 != nil {
			fmt.Println(noti.DNSOrchestrator + " no recibio el archivo",err01)
		}else {
			fmt.Println(noti.DNSOrchestrator + "Recibio el archivo")
		}
	}

	resp := Respuesta {
		Reloj: "logs Enviados",
	}

	return &resp, nil
}

func (s *Server) CompartirZF(ctx context.Context, comp *Compartir) (*Respuesta, error) {
	//Llegan los ZF luego del merge y se guardan localmente, se borran los log

	erro := ioutil.WriteFile("./Registros/ZF-" + comp.Dominio, []byte(comp.ZFContent), 0644)
	if erro != nil {
		fmt.Println("No se pudo guardar el archivo ZF-" + comp.Dominio)
	}else{
		currentclock := strings.Split(Globals.GetClock(comp.Dominio), ",")
		orchsClock := strings.Split(comp.Reloj, ",")
		for i, _ := range currentclock {
			oC, _ := strconv.Atoi(orchsClock[i])
			cC, _ := strconv.Atoi(currentclock[i])

			if cC < oC{
				Globals.SetValue(comp.Dominio, i, oC)
			}
		}
	}

	resp := Respuesta {
		Reloj: Globals.GetClock(comp.Dominio),
	}

	return &resp, nil
}

func (s *Server) Mergear(ctx context.Context, mer *Merge) (*Respuesta, error) {

	changes := strings.Split(mer.FileAsStr, "\n")
	var current []string

	input, erro := ioutil.ReadFile("./Registros/ZF-" + mer.Dominio)
	if erro != nil {
		fmt.Println("No se ha encontrado el archivo ZF-" + mer.Dominio)
	}else{
		current := strings.Split(string(input), "\n")
	}

	fmt.Printf(mer.Dominio + " ")
	for _, change := range changes {
		fmt.Printf("|")

		values := strings.Split(change, " ")

		if values[0] == "append" {
			current = append(current, "www." + values[1] + " IN A " + values[2])

		}else{
			for i, c := range current {
				if strings.Contains(c, values[1]) {
					if values[0] == "update" {

						if len(strings.Split(values[2], ".")) == 2 {//Se esta cambiando el nombre
							current[i] = "www." + values[2] + " IN A " + strings.Split(c, " ")[3]
						}else{//Se esta cambiando la IP
							current[i] = "www." + values[1] + " IN A " + values[2]
						}

					}else if values[0] == "delete" {

						current[i] = current[len(current)-1]
						current = current[:len(current)-1]
					}
				}
			}
		}
	}
	fmt.Println()

	output := strings.Join(current, "\n")
	erro = ioutil.WriteFile("./Registros/ZF-" + mer.Dominio, []byte(output), 0644)
	if erro != nil {
		fmt.Println("No se pudo sobreescribir el archivo")
	}else{

		orchsClock := strings.Split(Globals.GetClock(mer.Dominio), ",")
		sendersClock := strings.Split(mer.Reloj, ",")

		for i := range orchsClock {
			oC, _ := strconv.Atoi(orchsClock[i])
			sC, _ := strconv.Atoi(sendersClock[i])

			if oC < sC{
				Globals.SetValue(mer.Dominio, i, sC)
			}
		}
	}

	resp := Respuesta {
		Reloj: Globals.GetClock(mer.Dominio),
	}

	return &resp, nil
}

func (s *Server) MergeDone(ctx context.Context, pears *Pears) (*Respuesta, error) {
	//Orquestador recibe esta notificacion y envia los archivos ZF a los otros DNS	

	fmt.Println("Enviando archivos ZF a los otros DNS")
	var files []string

    err := filepath.Walk("./Registros", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), "ZF") {
			files = append(files, info.Name())
		}
		if strings.Contains(info.Name(), "log") {
			e := os.Remove("./Registros" + info.Name()) 
			if e != nil { 
				fmt.Println("No se pudo remover " + info.Name())
			} 
		}
        return nil
    })
    if err != nil {
        panic(err)
	}

	var conn1 *grpc.ClientConn
	var conn2 *grpc.ClientConn
	var conn3 *grpc.ClientConn

	s1 := "DEAD"
	s2 := "DEAD"
	s3 := "DEAD"

	if (pears.Status1 != "DEAD") && (pears.Status1 != pears.Orchestrator) {

		conn1, err1 := grpc.Dial(pears.Status1, grpc.WithInsecure(),grpc.WithBlock())
		if err1 != nil{
			fmt.Println("No fue posible conectarse con el orquestador de la unificacion")
		}else{
			s1 = "ALIVE"
			defer conn1.Close()
		}
	}

	if (pears.Status2 != "DEAD") && (pears.Status2 != pears.Orchestrator) {

		conn2, err2 := grpc.Dial(pears.Status2, grpc.WithInsecure(),grpc.WithBlock())
		if err2 != nil{
			fmt.Println("No fue posible conectarse con el orquestador de la unificacion")
		}else{
			s2 = "ALIVE"
			defer conn2.Close()
		}
	}

	if (pears.Status3 != "DEAD") && (pears.Status3 != pears.Orchestrator) {

		conn3, err3 := grpc.Dial(pears.Status3, grpc.WithInsecure(),grpc.WithBlock())
		if err3 != nil{
			fmt.Println("No fue posible conectarse con el orquestador de la unificacion")
		}else{
			s3 = "ALIVE"
			defer conn3.Close()
		}
	}
	
	for _, file := range files {

		path := "./Registros/" + file

		input, erro := ioutil.ReadFile(path)
		if erro != nil {
			resp := Respuesta {
				Reloj: "",
			}
			return &resp, nil
		}

		strFile := string(input)

		comp := Compartir {
			ZFContent: strFile,
			Reloj: Globals.GetClock(strings.Split(file, "-")[1]),
			Dominio: strings.Split(file, "-")[1],
		}

		if s1 != "DEAD"{

			compar := NewCompartirZFServiceClient(conn1)
			_, err01 := compar.CompartirZF(context.Background(), &comp)
			if err01 != nil {
				fmt.Println(pears.Status1 + " no recibio el archivo",err01)
			}else {
				fmt.Println(pears.Status1 + "Recibio el archivo")
			}
		}

		if s2 != "DEAD"{

			compar := NewCompartirZFServiceClient(conn2)
			_, err02 := compar.CompartirZF(context.Background(), &comp)
			if err02 != nil {
				fmt.Println(pears.Status2 + " no recibio el archivo",err02)
			}else {
				fmt.Println(pears.Status2 + "Recibio el archivo")
			}
		}

		if s3 != "DEAD"{

			compar := NewCompartirZFServiceClient(conn3)
			_, err03 := compar.CompartirZF(context.Background(), &comp)
			if err03 != nil {
				fmt.Println(pears.Status3 + " no recibio el archivo",err03)
			}else {
				fmt.Println(pears.Status3 + "Recibio el archivo")
			}
		}
		
	}

	resp := Respuesta {
		Reloj: "ZFs Enviados",
	}

	return &resp, nil
}