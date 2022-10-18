# Image base:
FROM nginx:latest
# Diretorio que vai ser trabalhado dentro do container (como se criasse 
# uma pasta e fosse setada como contexto) - opcional
WORKDIR /app
# Atualiza os pacotes da imagem, normal nas distros Linux
# Pode-se usar tb apt-get update && apt-get install vim -y para quebrar linha usamos \
RUN apt-get update
# Instala o vim. A flag -y eh necessaria para dar "sim" nas interacoes do terminal
RUN apt-get install vim -y
# Copia arquivos da pasta html no host para /usr/share/nginx no container
COPY html /usr/share/nginx/html

ENTRYPOINT [ "/docker-entrypoint.sh" ]

CMD [ "nginx", "-g", "daemon off;" ]