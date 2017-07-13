# Tool that displays the contents of our DB release table written in golang
Output on frontend is shown in JSON format

Required tpkgs:
go-1.5.1-5-x86_64.tpkg
mysql-5.6.24-3.1-x86_64.tpkg
mysqltuner-1.2.0-redhat-x86_64.tpkg

vpdb.go is the main go script that runs the http server.
# I used vpdb2.go for testing changes. 
vpdb.sh is a simple shell script that runs the go script via "go run" to run it in the background.

# contact magaling.markizm@gmail.com 
