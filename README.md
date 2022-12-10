# tado
Docker container to automate open windows detection with tado devices.
Just run the docker container with 

homename: can be found on https://app.tado.com/de/main/home as a title
username: username used to login into tado app 
password: password used to login into tado app

```
# replace parameters in sharp clips
docker run open-window -n "<homename>" -u "<username>" -p "<password>"
```
