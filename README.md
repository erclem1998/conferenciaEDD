# **CONFERENCIA EDD 2021: DESARROLLO DE MODERNAS APLICACIONES WEB CON ANGULAR**

## Objetivo del ejemplo
El objetivo principal de este ejemplo, es introducir, a los estudiantes, el uso del framework 
Angular, que hoy en día es uno de los mas utilizados en el desarrollo de Aplicaciones Web.

## ¿Que encontrarás?
1. Servidor web desarollado con Golang.
2. Aplicación Web desarrollada con Angular.

## Lo que tienes que saber antes de hacer uso de este ejemplo

| Herramienta | Versión  |
| :----------:| :------: |
| Node.js     | 14.15.4  |
| Angular     | 11.2.2   |  
| Go          | 1.16     |

## Acerca del Servidor Web
Este ejemplo se basa en el uso de un arbol binario de búsqueda, donde se almacenan los datos de
estudiantes y los cursos que ha aprobado.

Cabe destacar que el servidor web hace uso de los siguientes paquetes:

### **1. github.com/gorilla/mux**
Utilizado para generar el enrutador de solicitudes HTTP.
Se debe instalar de la siguiente manera.

**go get github.com/gorilla/mux**

### **2. github.com/gorilla/handlers**
Utilizado para manejar los CORS del servidor, permitiendo de esta manera
el acceso a las aplicaciones web a los recursos del mismo.

**go get github.com/gorilla/handlers**

## Rutas o endpoints del Servidor

| Ruta              | Función                                                                         | Tipo    |
| :--------------:  |:------------------------------------------------------------------------------: | :-----: |
| /arbol            | Devuelve el arbol de estudiantes en formato JSON                                |   GET   | 
| /crearEstudiante  | Crea un nuevo estudiante                                                        |  POST   |
| /arbolBinario     | Devuelve el arbol de estudiantes como una imagen en formato png                 |   GET   |
| /listaEstudiantes | Devuelve un JSON con una lista de carnets                                       |   GET   |
| /cursosEstudiante | Devuelve la información completa del estudiantes, con lista de cursos aprobados |  POST   |
| /insertarCurso    | Registra un nuevo curso aprobado por el estudiante                              |  POST   |

## Acerca de la Aplicación Web con Angular
La aplicación contenida dentro de este proyecto contiene las siguientes caracteristicas:

### **1. Página de Inicio**
Es un componente que muestra una vista de bienvenida dentro de la aplicación
![Alt text](img/inicio.png?raw=true "Inicio")

### **2. Crear Estudiante**
Es un componente que permite registrar un estudiante.
![Alt text](img/crearEstudiante.png?raw=true "Inicio")

### **3. Registro de Cursos Aprobados**
Es un componente que permite registar cursos aprobados a un estudiante con nota y año de aprobación.
Además permite visualizar los cursos que actualmente posee aprobados el estudiante.
![Alt text](img/agregarCurso.png?raw=true "Inicio")