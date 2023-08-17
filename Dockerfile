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

# docker build -t sasmeka/coffeeback .


