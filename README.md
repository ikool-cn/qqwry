# golang ip address query with qqwry
    qqwry.dat version: 20170610
    
**usage:**

##### 1.buid program
```
go build qqwry.go
```
##### 2.run
    
```
./qqwry -d /path/qqwry.dat -p 8888
```
    * -d the dat data path
    * -p the server listening port


##### 3.query ip like this 
    
```
http://127.0.0.1:8888/?ip=117.136.81.171,192.168.1.188
```
