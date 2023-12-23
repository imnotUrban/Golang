# Contexto del proyecto

Hacer una api CRUD que utilice un ORM llamado GORM

## Como iniciar una API con GO


Primero, se debe iniciar el go mod, de esta manera se hará un seguimiento de las dependencias del código,, a medida que se le vayan sumando dependencias, el archivo irá mutando

```
go mod init API_MUX_GORM
```


Luego, se debe instalar la dependencia de gorilla mux

```
go get github.com/gorilla/mux
```

Luego es importante instalar fresh para que después de hacer cambios se actualice sola la página

```
go get github.com/pilu/fresh
```


luego, para ejecutar la aplicacion, en vez de go run ., es usaremos 

```
go run github.com/pilu/fresh
```

#### Endpoints

-> GET con parámetros
```
http://localhost:8084/api/ejemplo/87
```

-> GET con querystring
```
http://localhost:8084/api/query-string?id=1
```


Más de goroutines y channels 

https://www.developerro.com/2023/06/14/goroutines-channels/



#### DTO

La carpeta dto sirve para tener la estructura de los objetos que nos llega en las peticiones por ejemplo
transferir información entre las capas de una aplicación 


##### Endpoint para subir un archivo

Upload_file



#### Endpoint para ver un archivo
 -> Es un método de renderización buffer

http://localhost:8084/api/view?folder=fotos&file=20231222192809.png


### Activar CORS

go get github.com/rs/cors

Luego se pone en el main


## Instalar GORM

``
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
``

## instalar variables de entorno

go get github.com/joho/godotenv


## Lib para hacer slugs

go get github.com/gosimple/slug
