mkdir helloworldmulti
cd helloworldmulti
go mod init helloworldmulti
go mod edit -replace localhost/myproject/greetingsmulti=../greetingsmulti
go mod tidy
# Run without compiling
go run .
# Create an exectuable
go build
# Run the executable.
./helloworldmulti
