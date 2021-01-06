package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Consultas"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
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


func compartirServer(lis net.Listener) {//DNSs comparten sus ZF para llevar a cabo la consistencia
	s := CompartirZF.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterCompartirZFServiceServer(grpcServer, &s)
	Consulta.RegisterNotificacionServiceServer(grpcServer, &s)
	Consulta.RegisterMergearServiceServer(grpcServer, &s)
	Consulta.RegisterMergeDoneService(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando otros DNS para hacer consistencia")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consistencia: %s", err)
	}
}


func main() {

	lis, err := net.Listen("tcp", ":6500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	lis1, err1 := net.Listen("tcp", ":6501")
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}

	lis2, err2 := net.Listen("tcp", ":6502")
	if err2 != nil {
		log.Fatalf("failed to listen: %v", err2)
	}

	go func(){
		consultasServer(lis)
	}()

	go func(){
		modificacionesServer(lis1)
	}()

	compartirServer(lis2)

}
