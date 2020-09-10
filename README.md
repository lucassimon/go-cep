# GO CEP
API em Golang para consulta de endereço através de CEP

## Compilar
Além de ter o Go instalado no sistema operacional é necessário executar o comando ```go build```
e um binário com nome de *go-cep* será criado na raiz do projeto.

## Executar
Basta executar, execute o comando```go run main.go``` 
e a seguinte mensagem irá aparecer no console 
informando onde a aplicação estará sendo executada: 
```Servidor executando no endereço: http://127.0.0.1:3000```

## Testando
Para rodar os tests, execute o comando: ```go test```

## Uso
Basta acessar a URL como no exemplo abaixo ```http://localhost:3000/cep/89201405```
O retorno será um JSON com o conteúdo 

## Instalação da aplicação em sistema linux baseado em debian
Para colocar a aplicação para ser executada em produção em um servidor ubuntu server, 
basta executar o arquivo *install.sh* como root, exemplo: ```sudo ./install```

Feito isso a seguinte mensagem deve aparecer:
```
GO-CEP - Serviço de consulta de endereço por CEP

[√] Arquivos da aplicação copiado para /usr/local/gocep/
[√] Permissões setadas
[√] Arquivo go-cep.service copiado para /etc/systemd/system/go-cep.service
[√] Atualizando lista de serviços do sistema operacional
[√] Executando serviço


● go-cep.service
   Loaded: loaded (/etc/systemd/system/go-cep.service; disabled; vendor preset: enabled)
   Active: active (running) since Fri 2020-09-04 10:48:55 -03; 10ms ago
 Main PID: 3296 (go-cep)
    Tasks: 3 (limit: 4915)
   CGroup: /system.slice/go-cep.service
           └─3296 /usr/local/gocep/go-cep

set 04 10:48:55 tayron-IPMH110G-DDR3 systemd[1]: Started go-cep.service.

Fim da instalação, em caso de problema execute o comando: [ journalctl -u go-cep -f ] para mais detalhes
```
Ao executar o comando *journalctl -u go-cep -f* verá em qual endereço e porta a aplicação está acessível

```
set 04 10:48:55 tayron-IPMH110G-DDR3 systemd[1]: Started go-cep.service.
set 04 10:48:55 tayron-IPMH110G-DDR3 go-cep[3296]: Arquivo de configuração: /usr/local/gocep/.env
set 04 10:48:55 tayron-IPMH110G-DDR3 go-cep[3296]: Servidor executando no endereço: http://127.0.0.1:3003
```

Conforme o log a aplicação está sendo executada no endereço: ```http://127.0.0.1:3003```

## Desinstalação da aplicação em sistema linux baseado em debian
Para desinstalar basta executar o comando: ```sudo ./uninstall.sh```

Feito isso a seguinte mensagem deve aparecer:
```
Desinstalando GO-CEP - Serviço de consulta de endereço por CEP

[√] Parando serviço
[√] Removendo instalação antiga
[√] Removendo serviço antigo
[√] Atualizando lista de serviços do sistema operacional

Desinstalação concluída
```
