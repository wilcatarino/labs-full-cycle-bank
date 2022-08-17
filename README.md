# Sobre

- Banco fictício com um checkout, criado durante as aulas da [imersão Full Cycle](https://imersao.fullcycle.com.br/)

# Desafio

- Performance entre as transações que ocorrem no banco fictício e no checkout
- Utilização de gRPC (HTTP/2)
- Utilização de Apache Kafka
- Utilização de Elasticsearch e Kibana com Kafka Connect
- Utilização de Istio, Kiali, Prometheus e Grafana
- Utilização de Docker e Kubernetes

# Arquitetura

- Banco (backend) - utilizando Golang
- Extrato (backend) - 
- Extrato (frontend) - 
- Checkout (backend) - utilizando NestJS
- Checkout (frontend) - utilizando Next.js

De forma bem resumida:
- Banco (backend) atualiza banco de dados e publica as transações no tópico do Apache Kafka
- Extrato (backend) consome as transações do tópico do Apache Kafka
- Extrato (frontend) exibe as transações já consumidas do tópico do Apache Kafka
- Kafka Connect alimenta o Elasticsearch através do tópico do Apache Kafka 
- Kibana (frontend) acessa o Elasticsearch

# Comandos

- `cd src/catarinobank`
- `sudo docker-compose up -d`
- `sudo docker ps --all`
- `sudo docker exec -it catarinobank bash`
- `sudo docker-compose down`
