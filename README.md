# ProcessingWorksheet

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

Processa uma planilha e após fazer um calculo de valor, será incluido no SQL Server

<p> Framework usados: Gorm e Docker </p>

```shell
cd "diretorio de sua preferencia"
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlserver
```

Usar o Docker-Compose para subir o banco de dados
```YAML
docker-compose up
```

