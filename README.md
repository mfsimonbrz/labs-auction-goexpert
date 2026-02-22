Objetivo: Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

Clone o seguinte repositório: clique para acessar o repositório.

Toda rotina de criação do leilão e lances já está desenvolvida, entretanto, o projeto clonado necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

Para essa tarefa, você utilizará o go routines e deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lançes (bid) já está implementado.

Você deverá desenvolver:

    Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente;
    Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
    Um teste para validar se o fechamento está acontecendo de forma automatizada;


Dicas:

    Concentre-se na no arquivo internal/infra/database/auction/create_auction.go, você deverá implementar a solução nesse arquivo;
    Lembre-se que estamos trabalhando com concorrência, implemente uma solução que solucione isso:
    Verifique como o cálculo de intervalo para checar se o leilão (auction) ainda é válido está sendo realizado na rotina de criação de bid;
    Para mais informações de como funciona uma goroutine, clique aqui e acesse nosso módulo de Multithreading no curso Go Expert;
     

Entrega:

    O código-fonte completo da implementação.
    Documentação explicando como rodar o projeto em ambiente dev.
    Utilize docker/docker-compose para podermos realizar os testes de sua aplicação.

## Como testar

* Baixe o repositório e rode-o com o comando `docker-compose up`.
* O controle do tempo de vida de cada leilão é controlado pela variável de ambiente `AUCTION_INTERVAL`. Consulte o arquivo `.env` para alterar caso a necessidade.
* Cadastre um leilão utilizando uma ferramenta como o Postman, ou via cURL: 

`curl --request POST \
  --url http://localhost:8080/auction/ \
  --header 'Content-Type: application/json' \
  --data '{
  "product_name": "Playstation 4",
  "category": "Eletrônicos",
  "description": "Videogame muito moderno",
  "condition": 1
}'`

* Pesquise pelos leilões abertos. O leilão criado deve aparecer: 

`curl --request GET \
  --url 'http://localhost:8080/auction?status=0' \
  --header 'Authorization: Bearer <TOKEN>'
`

* Aguarde o tempo defindio em `AUCTION_INTERVAL`e repita a pesquisa. O leilão não aparece mais. Pesquise pelos leilões fechados alterando o parâmetro `condition`.

`curl --request GET \
  --url 'http://localhost:8080/auction?status=1' \
  --header 'Authorization: Bearer <TOKEN>'
`
