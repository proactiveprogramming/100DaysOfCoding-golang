mkdir helloworld2
cd helloworld2
go mod init helloworld2
go mod edit -replace localhost/myproject/greetings=../greetings
go mod tidy
# Run without compiling
go run .
# Create an exectuable
go build
# Run the executable.
./helloworld2
