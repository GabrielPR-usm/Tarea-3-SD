package main



import(
  "log"
	//"errors"
	"fmt"
	//"io/ioutil"
	"context"
  //"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"google.golang.org/grpc"
	//"io"
	//"math"
	"math/rand"
	//"os"
	//"path/filepath"
	"strconv"
	"time"
  "strings"
)


func main(){
  diccionario:=make(map[string]string)
  var flag bool = true
  for flag == true {

    var deseo int
    fmt.Println("Que accion quiere realizar?: ")
    fmt.Println("1. Create/Agregar nombre.dominio")
    fmt.Println("2. Update nombre.dominio")
    fmt.Println("3. Delete nombre.dominio")
    fmt.Println("4. Exit")
    fmt.Println("Ingrese el número: ")
    fmt.Scan(&deseo)
    if(deseo==4){
      flag=false
    }
    if(deseo==1 || deseo==2||deseo==3){
      var nomdom string;
      fmt.Println("Ingrese el nombre.dominio: ")
      fmt.Scan(&nomdom)
      var nombre string = strings.Split(nomdom,".")[0]
      var dominio string = strings.Split(nomdom,".")[1]
      //fmt.Println(nombre)
      //fmt.Println(dominio)
      if(deseo==1){
        //var Newip string = Ipaleatoria()
        var Newip string;
        fmt.Println("Asigne una IP Ej:127.0.255.9 --> ")
        fmt.Scan(&Newip)
        //fmt.Println(Newip)
        //fmt.Println(diccionario)
        var reloj string = "000";
        for k := range diccionario {
          if(k==dominio){
            reloj = strings.Split(diccionario[dominio],"-")[0]
          }
        }
        var conn0 *grpc.ClientConn
        conn0, err := grpc.Dial(":9500", grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn0.Close()

        d := ModificarZF.NewModificarZFRequestServiceClient(conn0)
        message0 := ModificarZF.Solicitud{
          Dominio:dominio,
          Reloj:reloj,
        }

        response0,err := d.ModificarZFRequest(context.Background(),&message0)
        if err0!= nil{
          log.Fatalf("Error in Operacion: %s",err0)
        }
        //fmt.Println(response0.IPserver)
        var conn *grpc.ClientConn
        conn, err = grpc.Dial(response0.IPserver, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn.Close()

        c := ModificarZF.NewModificarZFEnDNSServiceClient(conn)
        message := ModificarZF.Modificacion{
          Accion:"append",
          ValorAfectado:nomdom,
          NuevoValor:Newip,
        }
        response,err := c.ModificarZFEnDNS(context.Background(),&message)
        if err!= nil{
          log.Fatalf("Error in Operacion: %s",err)
        }
        diccionario[dominio]= response.Reloj+"-"+response0.IPserver
      }
      if(deseo==2){
        var Newvalue string;
        var deseo1 int
        fmt.Println("Que tipo de modificación quiere hacer?: ")
        fmt.Println("1. Update nombre")
        fmt.Println("2. Update IP")
        fmt.Println("Ingrese el número: ")
        fmt.Scan(&deseo1)
        if(deseo1==1){
          fmt.Println("Asigne un nuevo nombre --> ")
          fmt.Scan(&Newvalue)
        }
        if(deseo1==2){
          fmt.Println("Asigne una nueva IP Ej:127.0.255.9 --> ")
          fmt.Scan(&Newvalue)
        }
        //fmt.Println(Newip)
        //fmt.Println(diccionario)
        var reloj string = "000";
        for k := range diccionario {
          if(k==dominio){
            reloj = strings.Split(diccionario[dominio],"-")[0]
          }
        }
        var conn0 *grpc.ClientConn
        conn0, err := grpc.Dial(":9500", grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn0.Close()

        d := ModificarZF.NewModificarZFRequestServiceClient(conn0)
        message0 := ModificarZF.Solicitud{
          Dominio:dominio,
          Reloj:reloj,
        }

        response0,err := d.ModificarZFRequest(context.Background(),&message0)
        if err!= nil{
          log.Fatalf("Error in Operacion: %s",err)
        }

        var conn *grpc.ClientConn
        conn, err = grpc.Dial(response0.IPserver, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn.Close()

        c := ModificarZF.NewModificarZFEnDNSServiceClient(conn)
        message := ModificarZF.Modificacion{
          Accion:"update",
          ValorAfectado:nomdom,
          NuevoValor:Newvalue,
        }
        response,err := c.ModificarZFEnDNS(context.Background(),&message)
        if err!= nil{
          log.Fatalf("Error in Operacion: %s",err)
        }
        diccionario[dominio]= response.Reloj+"-"+response0.IPserver


      }
      if(deseo==3){
        var reloj string = "000";
        var Newip string = "0.0.0.0"
        for k := range diccionario {
          if(k==dominio){
            reloj = strings.Split(diccionario[dominio],"-")[0]
          }
        }
        var conn0 *grpc.ClientConn
        conn0, err := grpc.Dial(":9500", grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn0.Close()

        d := ModificarZF.NewModificarZFRequestServiceClient(conn0)
        message0 := ModificarZF.Solicitud{
          Dominio:dominio,
          Reloj:reloj,
        }

        response0,err := d.ModificarZFRequest(context.Background(),&message0)
        if err!= nil{
          log.Fatalf("Error in Operacion: %s",err)
        }

        var conn *grpc.ClientConn
        conn, err = grpc.Dial(response0.IPserver, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil{
          log.Fatalf("could not connect: %s",err)
        }
        defer conn.Close()

        c := ModificarZF.NewModificarZFEnDNSServiceClient(conn)
        message := ModificarZF.Modificacion{
          Accion:"delete",
          ValorAfectado:nomdom,
          NuevoValor:"",
        }
        response,err := c.ModificarZFEnDNS(context.Background(),&message)
        if err!= nil{
          log.Fatalf("Error in Operacion: %s",err)

        }
        diccionario[dominio]= response.Reloj+"-"+response0.IPserver

      }
    }
  }

}
