package main

import (
	"database/sql"
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

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	staticFiles := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	http.HandleFunc("/", Index)
	http.HandleFunc("/main", Principal)
	http.HandleFunc("/consulta", Consulta)
	// http.HandleFunc("/consulta", Consulta)

	log.Println("servidor corriendo.....")
	//servidor local en puerto 3000
	http.ListenAndServe("localhost:3000", nil)

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

// controlador para el index o inicio de la app
func Index(rw http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(rw, "index", nil)
}

// controlador para el main o pantalla principal
func Principal(rw http.ResponseWriter, r *http.Request) {
	datoNuevo := DatoNuevo{}
	bandera := r.URL.Query().Get("bandera")
	spa := r.URL.Query().Get("entrada")
	kek := r.URL.Query().Get("salida")
	if bandera == "formulario" && spa != "" {
		datoNuevo.Espa = spa
		datoNuevo.Quek = kek
		log.Println("info redirecionada")
		fmt.Println(datoNuevo)
		plantillas.ExecuteTemplate(rw, "main", datoNuevo)

	} else {
		plantillas.ExecuteTemplate(rw, "main", nil)
	}
}

// controlador para la consulta de un dato en especifico desde el formulario
func Consulta(rw http.ResponseWriter, r *http.Request) {
	radiospa := r.FormValue("btnradioSpa")
	radioKek := r.FormValue("btnradioKek")
	entrada := r.FormValue("entrada")
	fmt.Println("el valor de la entrada es:" + entrada)

	if radiospa == "opcionSpa" && len(entrada) != 0 {
		fmt.Println("este es el valor de radio español: " + radiospa)
		// fmt.Println(entrada)
		entrada := r.FormValue("entrada")
		conexionEstablecida := ConexionDB()
		registro, err := conexionEstablecida.Query("SELECT id,spa, kek FROM dato_a WHERE id=?", entrada)
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

	} else if radioKek == "opcionKek" && len(entrada) != 0 {
		fmt.Println("este es el valor de radio Kekchi: " + radioKek)

		fmt.Println("condidional de radio de kekchi")
		http.Redirect(rw, r, "/main", http.StatusPermanentRedirect)

	} else {
		http.Redirect(rw, r, "/main", http.StatusPermanentRedirect)

	}

}
