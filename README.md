# Postbin

Stupid simple API that echoes what has been recieved from a request as is.
I needed to test some webhooks and came up with this since http://postb.in/ wasn't working for me.

### Installation
go get github.com/mauleyzaola/postbin

### How it works

When you run `./postbin` a key is generated. This is a simple way to ignore requests from unknown clients.

The key can also be provided as parameter.

Append this key to the url `<server>/key`. Look below for curl examples on this.
```bash
postbin$ ./postbin
key:  4a407d07-04fb-4cb9-b839-e95ed351820d
port:  8080
========================================
```

The port can be passed as parameter, default is `8080`.
```bash
Usage of ./postbin:
  -key string
        if passed it will use to filter request. if not provided a random key will be used
  -port int
        specifies the port number where to serve requests from (default 8080)
```

### Examples
```bash
curl "http://localhost:8080/4a407d07-04fb-4cb9-b839-e95ed351820d"
```
Response
```bash
Received at:
2018-07-29 14:29:09.367973766 -0500 CDT m=+109.693599790
Remote Address:
[::1]:51380
Method:
GET
Headers:
User-Agent:curl/7.54.0
Accept:*/*
Query String Values:
Body:
```

```bash
curl "http://localhost:8080/4a407d07-04fb-4cb9-b839-e95ed351820d?par1=somthing&par2=another1&par3=true"
```
Response
```bash
Received at:
2018-07-29 14:29:35.040244716 -0500 CDT m=+135.365874807
Remote Address:
[::1]:51381
Method:
GET
Headers:
User-Agent:curl/7.54.0
Accept:*/*
Query String Values:
par1:somthing
par2:another1
par3:true
Body:
```

```bash
curl "http://localhost:8080/4a407d07-04fb-4cb9-b839-e95ed351820d" -d'{"name":"Mauricio","ok":true}'
```
Response
```bash
Received at:
2018-07-29 14:29:52.748265931 -0500 CDT m=+153.073898827
Remote Address:
[::1]:51382
Method:
POST
Headers:
Content-Length:29
Content-Type:application/x-www-form-urlencoded
User-Agent:curl/7.54.0
Accept:*/*
Query String Values:
Body:
{"name":"Mauricio","ok":true}
```

```bash
curl --header "token:xxx33" --header "Content-Type:application/json" "http://localhost:8080/31bf3b67-d25b-4744-8bf3-dd396aa849fe" -d'{"name":"Mauricio","ok":true}'
```
Response
```bash
Received at:
2018-07-29 14:37:07.457274749 -0500 CDT m=+33.664058948
Remote Address:
[::1]:51430
Method:
POST
Headers:
User-Agent:curl/7.54.0
Accept:*/*
Token:xxx33
Content-Type:application/json
Content-Length:29
Query String Values:
Body:
{"name":"Mauricio","ok":true}
```