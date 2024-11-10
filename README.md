# projeto-docker-go
Container GoLang interagindo com container mysql

Linux - Ubuntu

Instalar o chrome, vscode, linguagem de melhor afinidade (.net, java, go, python, etc...), insominia/postman e o docker.
    https://askubuntu.com/questions/935569/how-to-completely-uninstall-docker
    https://www.youtube.com/watch?v=7Iw2OgmwJG0&list=PLyGI7X8uCfHBBKlqRiS8m23XRH8xYE8Il | VIDEOS ANDRE NASSERALA
    
    apt install mysql
    apt install docker-ce docker-ce-cli container.io
    systemctl start docker | inicia daemon
    systemctl enable  docker | para o docker executar sempre ao ligar maquina
    docker run hello-world | para verificar se docker esta ok

subir um container com uma imagem mysql/mariadb, expondo a porta 3306 para o host
    
    docker ps -a  | exibe todos os containers
    docker rm [id do container] | remove container | rmi remove imagem
    docker pull mariadb:latest | baixa a ultima imagem do mariadb
    docker run -d --name meu-mariadb -v mariadb-data:/var/lib/mysql mariadb:latest | cria um container nome meu-mariadb em backgroud (-d), cria um volume     pra dados não serem perdidos.
    docker run -d --name meu-mariadb -v mariadb-data:/var/lib/mysql -e MARIADB_ROOT_PASSWORD=root -e MARIADB_USER=pedro -e MARIADB_PASSWORD=root              mariadb:latest
    docker inspect meu-mariadb | inspeciona o conteiner
    docker exec ipadress meu-mariadb -a
    docker exec meu-mariadb ip a | pega o IP do container 172.17.0.2/16
    docker start meu-mariadb | para iniciar o container quando ligar a maquina


criar um banco de dados e uma estrutura de tabela para CRUD (use sua criatividade para as entidades ... clientes, séries, carros, produtos, etc...)
    
    mysql -h 172.17.0.2 -u pedro -p | acessa o mariadb via IP com o usuário criado e pedindo senha, mas acessa o root que é sucesso
    Crie o database
    Crie a tabela
    CREATE TABLE bandas ( id INT AUTO_INCREMENT PRIMARY KEY, nome VARCHAR(255) NOT NULL,musicos TEXT,generos TEXT, status BOOLEAN DEFAULT TRUE,               data_inicio DATE, data_fim DATE);

desenvolver uma API com POST, GET, PUT e DELETE (crud), no ambiente local ( linux ), utilizando Golang
    -Aplicação criada

buildar a aplicação
    
    go build -o api-go  main.go| no VScode

Criar um novo container para rodar a aplicação buildada.
    
    cd /home/pedro/dev/go-api | vai até o diretório da aplicação buildada | No terminal
    docker build -t api -f /home/pedro/dev/go-api/Dockerfile .| builda a imagem baseada no dockerfile
    docker run --name dockerapi -it api | cria o container e roda

Conseguir fazer o container do passo 8 enxergar o binário buildado do passo 7. Dica, você pode copiar o binário no momento da criação do container ou fazer um volume compartilhado entre container e host para fazer copy/paste do arquivo da aplicação.
    
    -Feito copiando o binario, configurando o dockerfile.

Com os dois containers  rodando (app e bd) acessar a API exposta e usar ela, confirmando que a alteração no banco está ocorrendo tbm.
    
    docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dockerapi | Pegar IP do container dockerapi 
    docker start dockerapi | para iniciar novamente o container
    docker exec -it dockerapi /bin/sh | para abrir o terminal do container
    ./api-go | para executar a aplicação
