# pwshed (PassWord haSHED)

### Hashing standards

https://security.blogoverflow.com/2013/09/about-secure-password-hashing/
While SHA-3 is emerging as the standard for hashing, it is not necessarily the best standard for password hashing. The document above details different hashing algorithms and their properties.

For the purpose of this project, we allow the user to use an algorithm of your choice. But, we fall back to bcrypt as the standard.

### API

API version of the tool allows you to make API calls to a http server with the password and optinally the algorithm to hash it with. The server responds after 5 seconds.

To start the server:
```
go build

./pwshed
```

Usage:
```
curl -d "password=procore" http://localhost:8080/hash
# 5B3milHYYsskm+5N+QPnny3vPVOJe9yaMyfwoLjWh/dWUghzv5YOCmSSEDPstX2wfMZk9b39d/j+i0A3/rTarA==

curl -d "password=angryMonkey" http://localhost:8080/hash?alg=SHA512
# ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==
```

### CLI

CLI version of the tool allows you to pass in a password and an optional algorithm of your choice to hash the password using that algorithm

Usage:
```
go build

./pwshed -cli=true -password=securestring
# chBQTKUoivgDzB3H9zDrIjYsVJvFhwGZ1ZwI1ZsQecttcTcoOWk07K1SyPfhfzsNf6XmBys0stnbQhHGku8qgw==

./pwshed -cli=true -password=procore -alg=SHA512
# VWx+vwK0xCGmazm68Bs7grHIXJv7Nl0W3vwR2DZ79dLGgoG0L+/9O3zc1xRmM28ltCujLRUb1/nEqJU3fQJMRw==
```


