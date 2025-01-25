# Email Searcher - Enron Corp

Email Searcher es una aplicacion web que permite la busquedad de correos electronicos del conjunto de datos Enron Mail. Este proyecto esta compuesto por el indexador de datos, el api y el cliente Web.

El proyecto tuvo las siguientes etapas:

- **Indexar data**: Mediante el indexer se realizo el proceso de indexacion utilizando Go Rutines para una mayor velocidad, haciendo uso de varias WorkerPools. El indexer Toma el directorio, procesa cada uno de los archivos, estructurando los Emails y enviando los al inidice "Emails" y de cada Email va almacenando todas las personas involucrados, relacionando las con su email y nombre para luego indexarlas en el indice "Persons"

- **Profiling**: Se relizo profiling de CPU, memoria y trace, que permitio observar las funcionalidades que tenian un costo alto de procesamiento, consulmo de memoria y el comportamiento de las GoRutines.

- **Visualizador**: Interfaz web que permite, realizar busquedas especificas y filtrado de los emails, permitiendo interactuar con estos, viendo la relacion entre los diferentes correos electronicos y las personas, por ejemplo permite visualizar todos los correos enviados, recibidos y copiados de una persona.

- **Optimización**: Se realizao mejoras en el indexer en base al analisis de estos profilings. Se realizaron pruebas a ambas versiones del indexer en condiciones similares de conexion a internet y evitando ejecutar otros progrmas durante el proceso y en el indexer_optimized se redujo el tiempo de indexado.

- **Deployment**: Se creo mediante IaC utilizando terraform y levantando la infraestructura en AWS.

## Estructura del Proyecto

```plaintext
.
├── api/                # Contiene la lógica backend, recuperando los datos de ZincSearch
├── client /            # Aplicación frontend para interactuar con los datos
├── indexer /           # Hace un proceso de ETL de los correos y los indexa en ZincSearch
├── indexer_optimized/  # Es una version optimizada del indexer, despues del analisis de profilings
└── terraform  /        # Infraestructura de los recuros para el despliegue de la aplicacion

```

## Tecnologias utilizadas

- ZincSearch
- Go 1.21.1
- Go-Chi v5
- Vue
- TypeScript
- Tailwind CSS
- Terraform

## Requisitos

1- Docker

2- Go v1.21.1

3- Enron Email Dataset

Puede descargar y extraer el Enron Email Dataset con el siguiente comando

```bash
curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o enron_mail_20110402.tgz && tar -xf enron_mail_20110402.tgz
```

4- Servicio de ZincSearch

- Download [ZincSearch](https://github.com/zincsearch/zincsearch) and follow the [ZincSearch Quick Start Guide](https://zincsearch-docs.zinc.dev/quickstart/).
- Alternatively, you can run ZincSearch with Docker using the following commands:

```bash
mkdir data

docker run -v /full/path/of/volume/of/data:/data -e ZINC_DATA_PATH="/data" -p 4080:4080 \
-e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 \
--name zincsearch public.ecr.aws/zinclabs/zincsearch:latest`
```

### Getting Started

## Indexer

Execute el comando:

```bash
go run . ./full/path/of/dataset
```

## API

- Agrega un archivo .env con las environment variables

- Ejecute el comando:

```bash
go run .
```

## Client

Ejecute los comandos:

```bash
npm install
npm run dev
```

## Contacto

Para cualquier consulta, puedes realizarla a melisadavila2001@gmail.com.
