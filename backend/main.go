package main

import (
	"os"
	"io"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"os/exec"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type NodoLista struct{
	Curso string `json:"curso"`
	Nota int `json:"nota"`
	Anio int `json:"anio"`
}


//STRUCT para registro de estudiantes
type Nodo struct{
	Carnet int `json:"carnet"`
	Nombres string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	CUI string `json:"cui"`
	Correo string `json:"correo"`
	Hizq* Nodo `json:"hizq"`
	Hder* Nodo `json:"hder"`
	ListaCursos[] NodoLista `json:"listaCursos"`
}

type Respuesta struct{
	Message string `json:"message"`
}

type listaCarnet struct{
	ListaCarnets[] int `json:"listacarnets"`
}

var raiz* Nodo
var contador int

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

//GET obtiene elarbol en json
func getArbol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(raiz)
}

//GET obtiene la lista de carnets en orden
func getListaCarnetsInorden(w http.ResponseWriter, r *http.Request) {
	var lista_carnets[] int
	lista_carnets=listaInorden(raiz,lista_carnets)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&listaCarnet{ListaCarnets:lista_carnets})
}

func createNode(w http.ResponseWriter, r *http.Request) {
	var newNode* Nodo
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &newNode)
	//fmt.Printf("%d",newNode.Carnet)
	//insertamos la raiz
	raiz=crearNodo(raiz,newNode)
	escribir,err2:=json.Marshal(raiz)
	if err2 != nil {
        log.Fatal(err2)
    }
	data := []byte(escribir)
    err = ioutil.WriteFile("persiste.json", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("----------------")
	//preorden(raiz)
	//createDot(raiz)
	//Si todo ha salido bien, devolvemos un status code 201 y el arbol
	w.Header().Set("Content-Type", "application/json")
	respuesta:= &Respuesta{Message:"Alumno creado exitosamente"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}

func crearNodo(nodo* Nodo, nodoInsertar* Nodo) *Nodo{
	//insertamos el primer nodo
	if raiz==nil{
		nuevoNodo:=&Nodo{
			Carnet:nodoInsertar.Carnet, 
			Nombres: nodoInsertar.Nombres,
			Apellidos: nodoInsertar.Apellidos,
			Correo: nodoInsertar.Correo,
			CUI: nodoInsertar.CUI,
			Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
		raiz=nuevoNodo
	} else{
		//sino insertamos en donde corresponda
		raiz=insertarNodo(raiz,nodoInsertar)
	}
	return raiz
}

func insertarNodo(nodo* Nodo, nodoInsertar *Nodo) *Nodo{
	//Si es menor que el carnet actual
	if nodoInsertar.Carnet < nodo.Carnet{
		//si es null el hizq insertamos 
		if nodo.Hizq==nil{
			nuevoNodo:=&Nodo{
				Carnet:nodoInsertar.Carnet, 
				Nombres: nodoInsertar.Nombres,
				Apellidos: nodoInsertar.Apellidos,
				Correo: nodoInsertar.Correo,
				CUI: nodoInsertar.CUI,
				Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			nodo.Hizq=nuevoNodo
		} else{
			//sino seguimos recorriendo hasta insertar
			nodo.Hizq=insertarNodo(nodo.Hizq, nodoInsertar)
		}
	} else {
		//Si es mayor que el carnet entra a este conjuntp de instrucciones
		//Si el hder es null, insertamos
		if nodo.Hder==nil{
			nuevoNodo:=&Nodo{
				Carnet:nodoInsertar.Carnet, 
				Nombres: nodoInsertar.Nombres,
				Apellidos: nodoInsertar.Apellidos,
				Correo: nodoInsertar.Correo,
				CUI: nodoInsertar.CUI,
				Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			
			nodo.Hder=nuevoNodo
		} else{
			//de lo contrario seguimos recorriendo hasta insertar
			nodo.Hder=insertarNodo(nodo.Hder, nodoInsertar)
		}
	}
	return nodo 
}

//POST obtiene los cursos de un estudiante
func getListaCursos(w http.ResponseWriter, r *http.Request){
	//var cursos[] NodoLista
	var newNode* Nodo
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &newNode)
	buscado := listaCursosEstudiante(raiz, newNode.Carnet)
	estudiante:=&Nodo{
		Carnet:buscado.Carnet, 
		Nombres: buscado.Nombres,
		Apellidos: buscado.Apellidos,
		Correo: buscado.Correo,
		CUI: buscado.CUI,
		Hizq:nil, Hder:nil, ListaCursos:buscado.ListaCursos,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&estudiante)
}

func listaCursosEstudiante(nodo* Nodo, carnet int) *Nodo{
	var estudiante* Nodo
	if nodo.Carnet==carnet{
		return nodo
	}
	if nodo.Hizq != nil{
		estudiante=listaCursosEstudiante(nodo.Hizq, carnet)
	}
	if nodo.Hder != nil{
		estudiante=listaCursosEstudiante(nodo.Hder, carnet)
	}
	return estudiante
}

//Recorrido en inorden
func listaInorden(nodo* Nodo, lista[] int) []int {
	if nodo.Hizq != nil{
		lista=listaInorden(nodo.Hizq, lista)
	}
	lista = append(lista, nodo.Carnet)
	if nodo.Hder != nil{
		lista=listaInorden(nodo.Hder, lista)
	}
	return lista
}

//Recorrido en preorden
func preorden(nodo* Nodo) {
	fmt.Println(nodo.Carnet)
	if nodo.Hizq != nil{
		preorden(nodo.Hizq)
	}
	if nodo.Hder != nil{
		preorden(nodo.Hder)
	}
}

func getImagenArbol(w http.ResponseWriter, r *http.Request){
	//Generamos el archivo dot
	data := []byte(createDot(raiz))
    err := ioutil.WriteFile("grafo.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearGrafo.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./grafo.jpg")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/jpeg")
    io.Copy(w, img)

}

func recorrerArbol(nombrePadre string, hijo* Nodo, textoActual string) string{
	if hijo.Hizq!=nil{
		nombreHijo := "Nodo"
		nombreHijo+=strconv.FormatInt(int64(contador), 10)
		contador+=1
		textoActual+=nombreHijo+"[label=\""+strconv.FormatInt(int64(hijo.Hizq.Carnet),10)+"\"];\n"
		textoActual+=nombrePadre+"->"+nombreHijo+";\n"
		textoActual=recorrerArbol(nombreHijo,hijo.Hizq, textoActual)
	}
	if hijo.Hder!=nil{
		nombreHijo := "Nodo"
		nombreHijo+=strconv.FormatInt(int64(contador), 10)
		contador+=1
		textoActual+=nombreHijo+"[label=\""+strconv.FormatInt(int64(hijo.Hder.Carnet),10)+"\"];\n"
		textoActual+=nombrePadre+"->"+nombreHijo+";\n"
		textoActual=recorrerArbol(nombreHijo,hijo.Hder, textoActual)
	}
	return textoActual
}

func createDot(nodo* Nodo) string{
	var grafo string
	grafo="digraph G{\n"
	grafo+="node[shape=\"box\"]\n"
	grafo+="Nodo0[label=\""+strconv.FormatInt(int64(nodo.Carnet), 10)+"\"];\n"
	contador=1
	grafo=recorrerArbol("Nodo0", nodo, grafo)
	grafo+="}"
	return grafo
	//fmt.Println(grafo)
}

func cargarDatos() *Nodo{
	var arbol *Nodo
	datosArchivo, err := ioutil.ReadFile("./persiste.json")
    if err != nil {
        log.Fatal(err)
    }
	err=json.Unmarshal(datosArchivo, &arbol)
	if err != nil {
        log.Fatal(err)
    }
	return arbol
}

func main() {
	raiz=cargarDatos()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/arbol", getArbol).Methods("GET")
	router.HandleFunc("/crearEstudiante", createNode).Methods("POST")
	router.HandleFunc("/arbolBinario", getImagenArbol).Methods("GET")
	router.HandleFunc("/listaEstudiantes", getListaCarnetsInorden).Methods("GET")
	router.HandleFunc("/cursosEstudiante", getListaCursos).Methods("POST")
	//Se agregaron cors
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}