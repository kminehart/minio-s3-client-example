This project only exists to demonstrate exclusively using the AWS S3 Golang client to authenticate with and create buckets / objects in [minio](https://min.io/).

`docker-compose up` will start this server and minio with the credentials `root/rootroot`.

Running `curl -X POST localhost:3000/upload` will create a random file in the bucket.

View the objects created at http://127.0.0.1:9001/browser
