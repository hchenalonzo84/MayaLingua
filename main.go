package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "diccionario"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(localhost:3306)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	log.Println("conexion exitosa")
	return conexion

}

func CerrarConexionDB(conexion *sql.DB) {
	err := conexion.Close()
	if err != nil {
		log.Println("Error al cerrar la conexión a la base de datos:", err)
	} else {
		log.Println("Conexión a la base de datos cerrada correctamente.")
	}
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	staticFiles := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	//rutas
	http.HandleFunc("/", Index)
	http.HandleFunc("/main", Principal)
	http.HandleFunc("/consulta", Consulta)
	http.HandleFunc("/obtenerDatosJson", ObtenerTodos)
	http.HandleFunc("/obtenerDatosJson2", ObtenerTodos2)

	// http.HandleFunc("/consulta", Consulta)

	log.Println("servidor corriendo.....")
	//servidor local en puerto 3000
	http.ListenAndServe("localhost:3000", nil) //cambiar a 3000

}

type Datos struct {
	Id    int64
	Spa   string
	Kechi string
}

type DatoNuevo struct {
	Id   int64
	Espa string
	Quek string
}

type PalabraJson struct {
	Id  int64  `json:"id"`
	Spa string `json:"spa"`
	Kek string `json:"kek"`
}
type PalabraJson2 struct {
	Id  int64  `json:"id"`
	Spa string `json:"spa"`
	Kek string `json:"kek"`
}

// controlador para el index o inicio de la app
func Index(rw http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(rw, "index", nil)
}

// controlador para el main o pantalla principal

func Principal(rw http.ResponseWriter, r *http.Request) {
	// Si se envía un formulario con datos
	if r.URL.Query().Get("bandera") == "formulario" && r.URL.Query().Get("entrada") != "" {
		datoNuevo := DatoNuevo{
			Espa: r.URL.Query().Get("entrada"),
			Quek: r.URL.Query().Get("salida"),
		}
		log.Println("Info redirecionada:", datoNuevo)
		plantillas.ExecuteTemplate(rw, "main", datoNuevo)
	} else {
		plantillas.ExecuteTemplate(rw, "main", nil)

	}
}

func ObtenerTodos(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	conexionEstablecida := ConexionDB()
	entrada := r.FormValue("entrada")
	nuevaEntrada := "%" + entrada + "%"
	registros, err := conexionEstablecida.Query("SELECT * FROM dato_a WHERE spa LIKE ? ORDER BY spa ASC", nuevaEntrada)
	if err != nil {
		panic(err.Error())
	}
	defer registros.Close() // Asegúrate de cerrar los registros al finalizar la función
	arregloPalabras := []PalabraJson{}
	for registros.Next() {
		var id int64
		var spa, kek string
		err := registros.Scan(&id, &spa, &kek)
		if err != nil {
			panic(err.Error())
		}
		palabra := PalabraJson{
			Id:  id,
			Spa: spa,
			Kek: kek,
		}
		arregloPalabras = append(arregloPalabras, palabra)
	}
	arr, err := json.Marshal(arregloPalabras)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(arr))
	rw.Write(arr) // Envía la respuesta JSON al cliente
}

func ObtenerTodos2(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	conexionEstablecida := ConexionDB()
	salida := r.FormValue("salida")
	nuevasalida := "%" + salida + "%"
	registros, err := conexionEstablecida.Query("SELECT * FROM dato_a WHERE kek LIKE ? ORDER BY kek ASC", nuevasalida)
	if err != nil {
		panic(err.Error())
	}
	defer registros.Close() // Asegúrate de cerrar los registros al finalizar la función
	arregloPalabras := []PalabraJson2{}
	for registros.Next() {
		var id int64
		var spa, kek string
		err := registros.Scan(&id, &spa, &kek)
		if err != nil {
			panic(err.Error())
		}
		palabra := PalabraJson2{
			Id:  id,
			Spa: spa,
			Kek: kek,
		}
		arregloPalabras = append(arregloPalabras, palabra)
	}
	arr, err := json.Marshal(arregloPalabras)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(arr))
	rw.Write(arr) // Envía la respuesta JSON al cliente
}

// controlador para la consulta de un dato en especifico desde el formulario
func Consulta(rw http.ResponseWriter, r *http.Request) {
	radiospa := r.FormValue("btnradioSpa")
	radioKek := r.FormValue("btnradioKek")
	entrada := r.FormValue("entrada")
	salida := r.FormValue("salida")
	fmt.Println("el valor de la entrada es:" + entrada)

	if radiospa == "opcionSpa" && len(entrada) != 0 {
		fmt.Println("este es el valor de radio español: " + radiospa)
		// fmt.Println(entrada)
		id_consulta := r.FormValue("id_consulta")
		conexionEstablecida := ConexionDB()
		registro, err := conexionEstablecida.Query("SELECT id,spa, kek FROM dato_a WHERE id=?", id_consulta)
		if err != nil {
			panic(err.Error())
		}

		dato := Datos{}

		for registro.Next() {
			var id int64
			var spa, kek string
			err = registro.Scan(&id, &spa, &kek)
			if err != nil {
				panic(err.Error())
			}

			dato.Id = id
			dato.Spa = spa
			dato.Kechi = kek
		}
		bandera := "formulario"
		fmt.Println("se accedio al submit del formulario")
		// fmt.Println(arregloDato)
		http.Redirect(rw, r, "/main?entrada="+dato.Spa+"&salida="+dato.Kechi+"&bandera="+bandera, http.StatusPermanentRedirect)

	} else if radioKek == "opcionKek" && len(salida) != 0 {

		fmt.Println("este es el valor de radio Kekchi: " + radioKek)

		id_consulta2 := r.FormValue("id_consulta2")
		conexionEstablecida := ConexionDB()
		registro, err := conexionEstablecida.Query("SELECT id,spa, kek FROM dato_a WHERE id=?", id_consulta2)
		if err != nil {
			panic(err.Error())
		}

		dato := Datos{}

		for registro.Next() {
			var id int64
			var spa, kek string
			err = registro.Scan(&id, &spa, &kek)
			if err != nil {
				panic(err.Error())
			}

			dato.Id = id
			dato.Spa = spa
			dato.Kechi = kek
		}
		bandera := "formulario"
		fmt.Println("se accedio al submit del formulario")
		// fmt.Println(arregloDato)
		http.Redirect(rw, r, "/main?entrada="+dato.Spa+"&salida="+dato.Kechi+"&bandera="+bandera, http.StatusPermanentRedirect)

	} else {
		http.Redirect(rw, r, "/main", http.StatusPermanentRedirect)

	}

}
