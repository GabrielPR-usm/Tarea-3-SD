package main

import (
	//"bufio"
	//"log"
	//"errors"
	"fmt"
	//"io/ioutil"
	//"math"
	//"os"
	//"strconv"
	//"bytes"
	//"path/filepath"
	//"math/rand"
	"time"
	"golang.org/x/net/context"
  	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	)

func main() {

	var t time.Duration = 5000000000

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":6500",grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err != nil{
		fmt.Println("Servidor :6500 no disponible: ", err)
	}else{
		fmt.Println("Servidor :6500 esta disponible")
		defer conn.Close()
	}

	/*
	modificacion1 := ModificarZF.Modificacion {
		Accion : "append",
		ValorAfectado : "google.com",
		NuevoValor : "12.10.110.1",
	}
	*/

	modificacion2 := ModificarZF.Modificacion {
		Accion : "update",
		ValorAfectado : "gugli.com",
		NuevoValor : "guglision.com",
	}

	modificacion3 := ModificarZF.Modificacion {
		Accion : "delete",
		ValorAfectado : "guglision.com",
		NuevoValor : "guglision.compadre",
	}

	prop1 := ModificarZF.NewModificarZFEnDNSServiceClient(conn)
	/*
	resp1, err1 := prop1.ModificarZFEnDNS(context.Background(), &modificacion1)
	if err1 != nil {
		fmt.Println("No se puedo realizar la modificacion",err1)
	}else {
		fmt.Println("Cambio realizado")
		fmt.Println("reloj " + resp1.Reloj)
	}
	*/

	resp2, err2 := prop1.ModificarZFEnDNS(context.Background(), &modificacion2)
	if err2 != nil {
		fmt.Println("No se puedo realizar la modificacion",err2)
	}else {
		fmt.Println("Cambio realizado")
		fmt.Println("reloj " + resp2.Reloj)
	}

	resp3, err3 := prop1.ModificarZFEnDNS(context.Background(), &modificacion3)
	if err3 != nil {
		fmt.Println("No se puedo realizar la modificacion",err3)
	}else {
		fmt.Println("Cambio realizado")
		fmt.Println("reloj " + resp3.Reloj)
	}


	

}
