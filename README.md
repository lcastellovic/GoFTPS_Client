# GoFTPS - Simple FTPS Client by lcastellovic

> [!WARNING]
> El proyecto está en el estado pre-release y es aún experimental. Las funcionalidades son básicas para experimentar con los protocolos FTPS y las comunicaciones que se establecen con el servidor.
> En un futuro se implementarán funcionalidades descritas más abajo y se expondrán los cambios realizados.


## Funcionalidades

Actualmente están desarrolladas las siguientes funcionalidades:
+ Los datos de la conexión del servidor están implementados en el código y la conexión se establece al servidor de pruebas para experimentar con el protocolo FTP [**Rebex**](https://test.rebex.net/).
+ Se pueden ejecutar los siguientes comandos una vez se ha establecido conexión al servidor:
  + **LIST** (Se utiliza para listar un directorio pasado por parámetro y mostrar las carpetas y ficheros que contiene. En caso de que el parámetro esté vacío, lista el _working directory_).
  + **CWD** (Se utiliza para cambiar el _working directory_ al directorio pasado por parámetro. En caso de que el parámetro esté vacío, vuelve al _parent directory_).
  + **RETR** (Se utiliza para descargar el fichero pasado por parámetro al directorio raíz donde esté ubicado el programa de GoFTPS_Client).


## Futuras implementaciones

Esta es una lista de las funcionalidades e implementaciones que están ***WIP***
- [ ] Elección de los datos del servidor.
- [ ] Elección de la ruta local.
- [ ] ***STOR*** Subir un fichero seleccionado de la ruta local al _working directory_.
- [ ] ***MKD*** Crear un directorio en el _working directory_ con el nombre pasado por parámetro.


## Futuro cambio de versión

- [ ] ***GUI*** Funcionalidad bastante importante que implementará una interfaz gráfica.


> [!NOTE]
> Existen diversos _packages_ de repositorios Go para experimentar con el FTP y sus funcionalidades.

Los repositorios que están incluidos para experimentar son los siguientes:
+ [jlaffaye ftp](github.com/jlaffaye/ftp)
+ [kardianos ftps](github.com/kardianos/ftps)
+ [dutchcoders goftp ](github.com/dutchcoders/goftp)
+ [sacloud ftps](github.com/sacloud/ftps)
