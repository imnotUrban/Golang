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