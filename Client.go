package main



import(
  "log"
	//"errors"
	"fmt"
	//"io/ioutil"
	"context"
  //"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Consultas"
	"google.golang.org/grpc"
	//"io"
	//"math"
	//"math/rand"
	//"os"
	//"path/filepath"
	//"strconv"
	//"time"
  "strings"
)


func main(){
  diccionario:=make(map[string]string)
  var flag bool = true
  for flag == true {

    var deseo int
    fmt.Println("Que accion quiere realizar?: ")
    fmt.Println("1. Ingresar página web")
    fmt.Println("2. Exit")
    fmt.Println("Ingrese el número: ")
    fmt.Scan(&deseo)
    if(deseo==2){
      flag=false
    }
    if(deseo==1){
      var pagina string
      fmt.Println("Ingrese página web: ")
      fmt.Scan(&pagina)
      var reloj string = "0,0,0";
      for k, _ := range diccionario {
        if(k==pagina){
          reloj = strings.Split(diccionario[pagina],"-")[0]
        }
      }
      var conn0 *grpc.ClientConn
      conn0, err := grpc.Dial("dist73:9501", grpc.WithInsecure(), grpc.WithBlock())
      if err != nil{
        log.Fatalf("could not connect: %s",err)
      }
      defer conn0.Close()

      d:= Consulta.NewConsultaRequestServiceClient(conn0)
      message := Consulta.Datos{
        Url:pagina,
        Reloj:reloj,
      }

      response0,err0 := d.ConsultaRequest(context.Background(),&message)
      if err0!= nil{
        fmt.Println("Aqui")
        log.Fatalf("Error in Operacion: %s",err0)
      }
      fmt.Println(response0)
      if(response0.IP=="Not Found"){
        fmt.Println("Url no encontrada")
        for k, _ := range diccionario {
          if(k==pagina){
            delete(diccionario,pagina)
          }
        }
      }
      if(response0.IP!="Not Found"){
        diccionario[pagina]=response0.Reloj+"-"+response0.IP
        fmt.Println("La IP de "+pagina+" es "+response0.IP)
      }
    }
  }
}
