//Comandos 

//Docker 

//Sobe os containers
docker-compose up -d

//Entra na maquina do mysql (bash)
docker-compose exec mysql bash

//Entra na maquina do mysql (mysql)
mysql -uroot -p nomedobanco

//Listar containers em execução
docker ps

//Listar containers incluindo os parados
docker ps -a

//Listar containers apenas pelo id
docker ps -aq

//Derruba todos os containers em execucao
docker stop $(docker ps -aq)

//Remove o container
docker rm ID_do_Container ou Nome_do_Container


//Mysql

//Mostra as tabelas
show tables;

//Conceder privilegios para o usuario(ip) versao mysql <=8
CREATE USER 'root'@'172.18.0.1' IDENTIFIED BY 'sua_senha';

GRANT ALL PRIVILEGES ON *.* TO 'root'@'172.18.0.1' WITH GRANT OPTION;

FLUSH PRIVILEGES;


//Go

//Executa o arquivo
go run main.go (nome do arquivo)

