# pwshed

### CLI

CLI version of the tool allows you to pass in a password and an optional algorithm of your choice to hash the password using that algorithm

Usage:
```
go build

./pwshed -password=securestring
# chBQTKUoivgDzB3H9zDrIjYsVJvFhwGZ1ZwI1ZsQecttcTcoOWk07K1SyPfhfzsNf6XmBys0stnbQhHGku8qgw==

./pwshed -password=procore -alg=SHA512
# VWx+vwK0xCGmazm68Bs7grHIXJv7Nl0W3vwR2DZ79dLGgoG0L+/9O3zc1xRmM28ltCujLRUb1/nEqJU3fQJMRw==
```
