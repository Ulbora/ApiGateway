cd cache
go test -coverprofile=coverage.out
sleep 15
cd ..

cd cluster
go test -coverprofile=coverage.out
sleep 15
cd ..

cd common
go test -coverprofile=coverage.out
sleep 15
cd ..

cd errors
go test -coverprofile=coverage.out
sleep 15
cd ..

cd handlers
go test -coverprofile=coverage.out
sleep 15
cd ..

cd managers
go test -coverprofile=coverage.out
sleep 15
cd ..

cd performance
go test -coverprofile=coverage.out
sleep 15
cd ..