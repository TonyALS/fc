## Network no Docker

- Tipos: Bridge, Host, Overlay, Macvlan, None

### Bridge
É o tipo padrão quando não informamos nada ao criar uma network no Docker. 
Se criarmos dois containers Ubuntu por exemplo, ambos conseguirão se comunicar via ping para o
IP um do outro, mas caso o ping seja feito por nome, NÃO funcionará.


##### Documentação oficial
[Docker Network](https://docs.docker.com/network/)