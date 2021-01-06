package main

import (
	"log"
	"net"
	"fmt"
	"time"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	//"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"

)

var t time.Duration = 5000000000

func mergeen() {

	s1 := ":6502"
	s2 := ":7502"
	s3 := ":8502"

	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(s1, grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err1 != nil{
		s1 = "DEAD"
	}else{
		defer conn1.Close()
	}

	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(s2, grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err2 != nil{
		s2 = "DEAD"
	}else{
		defer conn2.Close()
	}

	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial(s3, grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err3 != nil{
		s3 = "DEAD"
	}else{
		defer conn3.Close()
	}

	orch := "NOT SET"

	if s1 != "DEAD" {
		orch = s1
	}else if s2 != "DEAD" {
		orch = s2
	}else if s3 != "DEAD"{
		orch = s3
	}

	noti := CompartirZF.Notificacion {
		DNSOrchestrator: orch,
	}

	if (s1 != orch) && (s1 != "DEAD") {

		noti1 := CompartirZF.NewNotificacionServiceClient(conn1)
		_, err01 := noti1.Notificar(context.Background(), &noti)
		if err01 != nil {
			fmt.Println(s1 + " no recibio la notificacion ",err01)
		}else {
			fmt.Println(s1 + "Recibio la notificacion")
		}
	}

	if (s2 != orch) && (s2 != "DEAD") {

		noti2 := CompartirZF.NewNotificacionServiceClient(conn2)
		_, err02 := noti2.Notificar(context.Background(), &noti)
		if err02 != nil {
			fmt.Println(s2 + " no recibio la notificacion ",err02)
		}else {
			fmt.Println(s2 + "Recibio la notificacion")
		}
	}

	if (s3 != orch) && (s3 != "DEAD") {

		noti3 := CompartirZF.NewNotificacionServiceClient(conn3)
		_, err03 := noti3.Notificar(context.Background(), &noti)
		if err03 != nil {
			fmt.Println(s3 + " no recibio la notificacion ",err03)
		}else {
			fmt.Println(s3 + "Recibio la notificacion")
		}
	}

	time.Sleep(10 * time.Second)

	//Le digo al orchestrator que mande de vuelta los ZF unificados
	mergeRdy := CompartirZF.Pears {
		Status1: s1,
    	Status2: s2,
    	Status3: s3,
		Orchestrator: orch,
	}

	if s1 == orch {

		mergeD1 := CompartirZF.NewMergeDoneServiceClient(conn1)
		_, err01 := mergeD1.MergeDone(context.Background(), &mergeRdy)
		if err01 != nil {
			fmt.Println(s1 + " No recibio la orden de enviar el consolidado ",err01)
		}else {
			fmt.Println(s1 + "Recibio la orden de enviar el consolidado")
		}
	}

	if s2 == orch {

		mergeD2 := CompartirZF.NewMergeDoneServiceClient(conn2)
		_, err02 := mergeD2.MergeDone(context.Background(), &mergeRdy)
		if err02 != nil {
			fmt.Println(s2 + " No recibio la orden de enviar el consolidado ",err02)
		}else {
			fmt.Println(s2 + "Recibio la orden de enviar el consolidado")
		}
	}

	if s3 == orch {

		mergeD3 := CompartirZF.NewMergeDoneServiceClient(conn3)
		_, err03 := mergeD3.MergeDone(context.Background(), &mergeRdy)
		if err03 != nil {
			fmt.Println(s3 + " No recibio la orden de enviar el consolidado ",err03)
		}else {
			fmt.Println(s3 + "Recibio la orden de enviar el consolidado")
		}
	}
}

func main() {

	lis, err := net.Listen("tcp", ":9500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func () {
		for true {
			mergeen()
			time.Sleep(5 * 60 * time.Second)
		}
	}()

	s := ModificarZF.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	ModificarZF.RegisterModificarZFRequestServiceServer(grpcServer, &s)
	
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("BROKER ESCUCHANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
