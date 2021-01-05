package Consulta

import (
	//"log"
  	"golang.org/x/net/context"
	//"io"
	"fmt"
	//"google.golang.org/grpc"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Globals"
	"io/ioutil"
	//"os"
	"strconv"
	//"time"
	"strings"
)

type Server struct{
}


func (s *Server) ConsultaCliente(ctx context.Context, consul *Consulta) (*Respuesta, error) {

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

	resp := Respuesta {
		IP: "GUARDADO EXITOSAMENTE EN LOG",
		Reloj: "",
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
			Reloj: strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]),
		}
	
		return &resp, nil
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if strings.Contains(line, consul.NombreDominio) {
			ip = strings.Split(line, " ")[3]
		}
	}

	resp := Respuesta {
		IP: ip,
		Reloj: strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][0]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][1]) + "," + strconv.Itoa(Globals.RVDNS1[NombreDominio[1]][2]),
	}

	return &resp, nil
}