# CLIma  

Programa em CLI para verificar o clima local ou de qualquer lugar do mundo em poucas teclas.

## Para usar:

Você pode clonar o repositório e rodar: 

```
go build
```

Após isso, mova o arquivo para a pasta `/bin` para ter acesso ao comando no terminal

```
sudo mv ./clima /usr/local/bin/
```

Para testar se a instalação foi feita de maneira correta:
```
clima
```
----------

## Exemplo de resposta:

```cli
clima
Fetching data for São Paulo...
Timezone: America/Sao Paulo

Day: 04/07/2023
02:00h - 11.0°C

Day   - Min   - Max   - Sunrise - Sunset
04/07 - 9.3°C - 19.7°C - 06:49 - 17:32
05/07 - 8.9°C - 22.2°C - 06:49 - 17:32
06/07 - 11.2°C - 22.8°C - 06:49 - 17:33
07/07 - 11.8°C - 23.5°C - 06:49 - 17:33
08/07 - 12.5°C - 25.4°C - 06:49 - 17:34
09/07 - 16.0°C - 19.9°C - 06:49 - 17:34
10/07 - 15.8°C - 21.2°C - 06:49 - 17:34
```

```
clima belo horizonte
Fetching data for Belo Horizonte...
Timezone: America/Sao Paulo

Day: 04/07/2023
02:00h - 12.2°C

Day   - Min   - Max   - Sunrise - Sunset
04/07 - 10.1°C - 21.2°C - 06:31 - 17:28
05/07 - 10.0°C - 20.3°C - 06:31 - 17:29
06/07 - 9.4°C - 21.0°C - 06:31 - 17:29
07/07 - 9.0°C - 22.9°C - 06:31 - 17:29
08/07 - 9.7°C - 25.6°C - 06:31 - 17:30
09/07 - 13.4°C - 28.5°C - 06:31 - 17:30
10/07 - 16.7°C - 28.8°C - 06:31 - 17:30
```
