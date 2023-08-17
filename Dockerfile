#buat container
FROM golang:1.20.3-alpine

#buat folder untuk menyimpan code
WORKDIR /coffee_back

#copy semua file ke workdir
COPY . .

#install dependencys
RUN go mod download

#build
RUN go build -v -o /coffee_back/coffeeback ./cmd/main.go

#open port
EXPOSE 8080

#run app
ENTRYPOINT [ "/coffee_back/coffeeback" ]

# docker build -t coffeeback .
# docker run --name coffeebackapp --net=bridge -e PGHOST=host.docker.internal -e PGPORT=5433 -e PGUSER=fazztrack -e PGDATABASE=postgres -e PGPASSWORD=root -p 8081:8080 sasmeka/coffeeback



