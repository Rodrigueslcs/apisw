# apisw
para leventar a aplicação  

docker-compose up -d 




rotas

metodo GET na rota http://127.0.0.1:3456/v1/planet
retorna todos planetas cadastrados

metodo POST na rota http://127.0.0.1:3456/v1/planet
insere um novo planeta 
{
    "name": "Tatooine",
    "climate":"quente",
    "terrain": "agua"

}

Metodo DELETE rota http://127.0.0.1:3456/v1/planet?id=...
passando o ID como paramentro
exclui um planeta atravez do ID

metodo GET na rota http://127.0.0.1:3456/v1/planet/?id=...
busca o pleneta pelo id


metodo GET na rota http://127.0.0.1:3456/v1/planet/?name=...
busca os planete pelo nome 
