# Reqstrict

Reqstrict is rule based forward proxy server. It restricts only HTTP (and not HTTPs) requests based on certain rules as configured by the user.

## Motivation
Reqstrict is created to ease down the process of configuring proxy server. The user is expected to modify script in order to manage the proxy server.

## Installation
```
$ git clone https://github.com/gophergala2016/reqstrict.git
$ cd reqstrict
$ go get github.com/yuin/gopher-lua
$ go build
$ ./reqstrict
```

## Usage
In order to use reqstrict as your proxy server, you first need to configure proxy settings of your browser. You can do it by following steps for [Chrome](https://support.google.com/chrome/answer/96815?hl=en) and [Firefox](http://www.wikihow.com/Enter-Proxy-Settings-in-Firefox)

```
./reqstrict -b [bindaddress] -p [port] -f [scriptfile] -m [maxconnection]
  -b string
    	The proxy server's bind address (default "0.0.0.0")
  -f string
    	The path of the lua script file (default "rule.lua")
  -m int
    	Maximum connections to proxy server (default 20)
  -p int
    	The proxy server's port (default 8000)

```
You can configure rules in rule.lua

for example, if you want to block http://www.google.com, then simply put function as :
```
  function filter (request)
    if string.find(request,"google") then
      return false
    else
      return true
    end
  end
```
Similarly you can block requests coming from certain User-Agents also.

Note : Do not change the function name or return type (true or false).

## License
Code released under [the MIT license](https://github.com/gophergala2016/reqstrict/blob/master/LICENSE).
