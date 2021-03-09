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
	"github.com/gorilla/mux"
)

type NodoLista struct{
	Curso string `json:"curso"`
	Nota int `json:"nota"`
	Anio int `json:"anio"`
}

type Nodo struct{
	Carnet int `json:"carnet"`
	Hizq* Nodo `json:"hizq"`
	Hder* Nodo `json:"hder"`
	ListaCursos[] NodoLista `json:"listaCursos"`
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

func createNode(w http.ResponseWriter, r *http.Request) {
	var newNode Nodo
	//leemos el body de la petición
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//tomamos los valores del body y los colocamos en una variable de struct de Nodo
	json.Unmarshal(reqBody, &newNode)
	//fmt.Printf("%d",newNode.Carnet)
	//insertamos la raiz
	raiz=crearNodo(raiz,newNode.Carnet)
	fmt.Println("----------------")
	//preorden(raiz)
	//createDot(raiz)
	//Si todo ha salido bien, devolvemos un status code 201 y el arbol
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(raiz)

}

func crearNodo(nodo* Nodo, valor int) *Nodo{
	//insertamos el primer nodo
	if raiz==nil{
		nuevoNodo:=&Nodo{Carnet:valor, Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
		raiz=nuevoNodo
	} else{
		//sino insertamos en donde corresponda
		raiz=insertarNodo(raiz,valor)
	}
	return raiz
}

func insertarNodo(nodo* Nodo, valor int) *Nodo{
	//Si es menor que el carnet actual
	if valor < nodo.Carnet{
		//si es null el hizq insertamos 
		if nodo.Hizq==nil{
			nuevoNodo:=&Nodo{Carnet:valor, Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			nodo.Hizq=nuevoNodo
		} else{
			//sino seguimos recorriendo hasta insertar
			nodo.Hizq=insertarNodo(nodo.Hizq, valor)
		}
	} else {
		//Si es mayor que el carnet entra a este conjuntp de instrucciones
		//Si el hder es null, insertamos
		if nodo.Hder==nil{
			nuevoNodo:=&Nodo{Carnet:valor, Hizq:nil, Hder:nil, ListaCursos:[]NodoLista{}}
			nodo.Hder=nuevoNodo
		} else{
			//de lo contrario seguimos recorriendo hasta insertar
			nodo.Hder=insertarNodo(nodo.Hder, valor)
		}
	}
	return nodo 
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
	router.HandleFunc("/crearNodo", createNode).Methods("POST")
	router.HandleFunc("/arbolBinario", getImagenArbol).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}