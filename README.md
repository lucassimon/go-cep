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
