package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Consultas"


)


func consultasServer(lis net.Listener) {//Broker le envia a DNS las consultas de los clientes
	c := Consulta.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterConsultaRequestServiceServer(grpcServer, &c)
	fmt.Println("Broker Escuchando a las consultas de los clientes")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consultas: %s", err)
	}
}


func modificacionesServer(lis net.Listener) {//Admin le envia la operacion CRUD a DNS
	s := ModificarZF.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	ModificarZF.RegisterModificarZFRequestServiceServer(grpcServer, &s)
	//Consultas.RegisterConsultaClienteServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("BROKER ESCUCHANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}



/*
func compartirServer(lis net.Listener) {//DNSs comparten sus ZF para llevar a cabo la consistencia
	s := CompartirZF.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterCompartirZFServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando otros DNS para hacer consistencia")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consistencia: %s", err)
	}
}
*/

func main() {

	lis1, err1 := net.Listen("tcp", ":9500")
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}

	lis2, err2 := net.Listen("tcp", ":9501")
	if err2 != nil {
		log.Fatalf("failed to listen: %v", err2)
	}


	go func(){
		consultasServer(lis2)
	}()
	modificacionesServer(lis1)

}
