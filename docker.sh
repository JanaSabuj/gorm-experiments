# mysql instance
docker run --name gorm-experiments -p 127.0.0.1:3308:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:latest --default-authentication-plugin=mysql_native_password