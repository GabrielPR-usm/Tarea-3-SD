package ModificarZF

import (
	"log"
  	"golang.org/x/net/context"
	//"io"
	"fmt"
	"google.golang.org/grpc"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	"github.com/GabrielPR-usm/Tarea-3-SD/Globals"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"strings"
	"math/rand"
)

type Server struct{
}
func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

var t time.Duration = 1500000000

func (s *Server) ModificarZFRequest(ctx context.Context, solicitud *Solicitud) (*RespuestaBroker, error) {
	fmt.Println("Analizando Request")
	servers := make([]string, 3)
	servers[0] = ":8500"
	servers[1] = ":7500"
	servers[2] = ":6500"
	var puerto string="";
	var aleatorio int=-1;
	for flag:=false;flag==false;{
		aleatorio=random(0,3)
		var conn *grpc.ClientConn
		fmt.Println("Aleatorio " + servers[aleatorio])
		conn, err := grpc.Dial(servers[aleatorio], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
		if err != nil{
			fmt.Println("could not connect: %s",err)
			continue
		}
		defer conn.Close()

		c := NewVerificarServiceClient(conn)
		message0 := Enviodom{
			Dominio:solicitud.Dominio,
		}
		response,err := c.Verificar(context.Background(),&message0)
		if err!= nil{
			log.Fatalf("Error in Operacion: %s",err)
		}
		var t1 []string = strings.Split(solicitud.Reloj, ",");
		var t2 []string = strings.Split(response.Reloj, ",");
		var count int =0
		var indice int = 0
		for;indice<3;indice++{
			value1,err := strconv.Atoi(t1[indice])
			if err!= nil{
				log.Fatalf("Error in Operacion: %s",err)
			}
			value2,err1 := strconv.Atoi(t2[indice])
			if err1!= nil{
				log.Fatalf("Error in Operacion: %s",err1)
			}
			if(value2>=value1){
				count++
			}
		}
		fmt.Println(count)
		if count==3{
			flag=true
			puerto=servers[aleatorio]
		}
	}
	retorno:=RespuestaBroker{
		IPserver:puerto,
	}
	return &retorno,nil
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

		fiz.WriteString("\nwww." + modif.ValorAfectado + " IN A " + modif.NuevoValor)

		fiz.Close()

	}else{

		input, erro := ioutil.ReadFile("./Registros/ZF-" + NombreDominio[1])
		if erro != nil {
			fmt.Println("No se ha encontrado el archivo ZF-" + NombreDominio[1])
			resp := RespuestaDNS {
				Reloj: strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]) ,
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
		Reloj: strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]),
	}

	return &resp, nil
}

func (s *Server) Verificar(ctx context.Context, dominio *Enviodom) (*RespuestaDNS, error) {

	fmt.Println("Verificando")

	resp := RespuestaDNS {
		Reloj: strconv.Itoa(Globals.RVDNS1[dominio.Dominio][0]) + "," + strconv.Itoa(Globals.RVDNS1[dominio.Dominio][1]) + "," + strconv.Itoa(Globals.RVDNS1[dominio.Dominio][2]),
	}

	return &resp, nil
}