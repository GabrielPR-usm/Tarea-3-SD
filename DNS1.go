package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Consultas"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Globals"

)


func consultasServer(lis net.Listener) {//Broker le envia a DNS las consultas de los clientes
	s := Consulta.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterConsultaBrokerServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando a las consultas de los clientes")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consultas: %s", err)
	}
}


func modificacionesServer(lis net.Listener) {//Admin le envia la operacion CRUD a DNS
	s := ModificarZF.Server{}
	grpcServer := grpc.NewServer()
	ModificarZF.RegisterModificarZFEnDNSServiceServer(grpcServer, &s)
	ModificarZF.RegisterVerificarServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando las modificaciones del admin")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve modificaciones: %s", err)
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

	lis, err := net.Listen("tcp", ":6500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func(){
		consultasServer(lis)
	}()
/*
	go func(){
		modificacionesServer(lis)
	}()

	compartirServer(lis)
*/

	modificacionesServer(lis)
}
