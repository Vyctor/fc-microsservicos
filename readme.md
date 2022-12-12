# Introdução

## Conceitos básicos

Microsserviços são:

- aplicações comuns
- independentes
- com objetivos bem definidos
- que fazem parte de um ecossistema.

## Microsserviços vs Monolíticos

### Microsserviços

- Objetivos segregados, bem definidos.
- Diversas tecnologias
- Menor risco no deploy
- Uma equipe por microsserviço
- Mais complexo para iniciar projetos que ainda não sabemos todas as regras ou implicações

### Monolíticos

- Todos os contextos dentro do mesmo sistema
- Única tecnologia
- Maior risco no deploy
- Todas as equipes trabalham no mesmo sistema
- É mais simples para iniciar um projeto

## Quando utilizar microsserviços

- Quando você quiser escalar times, separando-os por microsserviço.
- Quando eu tenho contextos/área de negócios bem definidos;
- Quando você tem maturidade nos processos de entrega
- Quando você tem maturidade técnica dos times
- Quando tenho a necessidade de escalar apenas uma parte do sistema
- Quando eu preciso de tecnologias distintas/específicas em partes do sistema.

## Quando utilizar sistemas monolíticos

- Provas de conceito
- Quando não conhecemos todo o domínio
- governança simplificada sobre as tecnologias utilizadas
- Facilidade na contratação e treinamento de desenvolvedores
- Tudo está no mesmo lugar
- Compartilhamento claro de libs (shared kernel)

## Migração de monolito para microsserviços

1. Separação de contextos (DDD)
2. Evite o excesso de granularidade
3. Verifique dependências (evitar monolito distribuído)
4. Planeje o processo de migração dos bancos de dados
5. Defina os eventos que acontecem durante a utilização do seu serviço
6. Não tenha medo de duplicação de dados
7. Pensar em consistência eventual
8. CI/CD/Testes/Ambientes maduros são obrigatórios
9. Comece pelas beiradas. Migre primeiramente os serviços menos importantes, com menos impacto no dia-a-dia.
10. Aplique Strangler Patter

# Características

## Componentização via serviços

Pode ser substituído, evoluído e o seu deploy ocorre de forma independente.

## Organização através das áreas de negócio

Lei de Conway -> O sistema é uma replica da forma que a empresa é distribuída organizacionalmente.

## Produtos e não projetos

Se você desenvolveu, você também manterá o produto. Os times são criados por produto e não por projeto.
Não existe time de sustentação, o que aumenta a responsabilidade do time, visto que ele deverá manter o que produz.

## Smart endpoints e dumb pipes

A sua aplicação deve ter endpoints para realizar operações dentro do microsserviço, independente de como a comunicação chegará ao microsserviço.
O pipe, mensageiro, deve ser burro, ou seja, sem regras de negócio.

## Governança descentralizada

Temos a opção de buscar outras formas para resolver determinado problema, que não são necessariamente as formas ditadas pela governança da empresa.
Mas não é porque eu posso fazer algo que eu devo fazer, cada decisão deve ser pautada em uma necessidade específica.

## Dados descentralizados

Cada microsserviço deve ter seu banco de dados com as informações persistentes para o seu funcionamento. Para isso se deve trabalhar com consistência eventual e sincronização entre as informações.

## Automação de infraestrutura

É necessário automatizar a infraestrutura, para administrar os processos de microsserviços e garantir maturidade.

## Desenhado para falhar

É necessário desenhar o microsserviço pensando em resiliência para que eles funcionem nas piores condições possíveis.

## Design evolutivo

O Design dos serviços são feitos para tenha a possibilidade de evoluir, substituir e desabilitar os serviços sem maiores complicações.

# Resiliência

## O que é resiliência?

É um conjunto de estratégias adotadas **intencionalmente** para a adaptação de uma sistema quando uma falha ocorre.

Ter estratégias de resiliência nos possibilita minimizar os riscos de perda de dados e transações importantes para o negócio.

## Proteger e ser protegido

Um sistema em uma arquitetura distribuída precisa adotar mecanismos de autopreservação para garantir ao máximo sua operação com **qualidade**

Um sistema precisa sempre se comportar da mesma forma, independente do cenário enfrentado.

Um sistema não pode ser egoísta ao ponto de realizar mais requisições em um sistema que está falhando.

Um sistema lento no ar muitas vezes é pior que um sistema fora do ar.

## Health Check

- Sem sinais vitais não é possível saber a saúde de um sistema
- Deve bater em todas as dependências do sistema
- Um sistema que não está saudável possui uma chance de se recuperar caso o tráfego para de ser direcionado a ele temporariamente

## Rate Limiting

- Protege o sistema baseado no que ele foi projetado para suportar
- Limita a quantidade de requisições que o sistema pode lidar
- Pode ser programada por tipo de client

## Circuit Breaker

- Protege o sistema fazendo com que as requisições feitas para ele sejam negadas
- Circuito fechado = Requisições chegam normalmente
- Circuito aberto = Requisições não chegam ao sistema. Erro instantâneo ao client
- Meio aberto = Permite uma quantidade limitada de requisições para verificação se o sistema tem condições de voltar ao ar integralmente

## API Gateway

- Garante que as requisições inapropriadas não cheguem ao sistema.
- Implementa políticas de Rate Limiting, Health Check
- Ajuda a organizar microsserviços em contexto - Estrela da morte

## Service Mesh

- Controla o tráfego de rede através de proxies
- Evita implementações de proteção pelo próprio sistema
- mTLS
- Implementa Circuit Breaker, retry, timeout, fault injection

## Trabalhar de forma assíncrona

- Evita perda de dados
- Não há perda de dados no envio de uma transação se o servidor estiver fora do ar
- Servidor pode processar a transação em seu tempo, quando estiver online
- Entender com profundidade o message broker/sistema de stream

## Retry

- Exponential backoff (o tempo de re-tentativa é aumentado exponencialmente)
- Exponential backoff com Jitter (o tento de re-tentativa é aumentado exponencialmente, porém com o Jitter eu tenho os tempos de chamada embaralhados, pois é mandado o tempo da chamada mais um valor aleatório para se somar ao tempo total de espera)

## Garantias de entrega

- **Ack 0**. Fire and Forget -> Não importa se a mensagem foi recebida. Não precisa de ack, não garante a entrega.
- **Ack 1**. O líder recebe e confirma a mensagem. Garante mais ou menos a entrega.
- **Ack -1**. O líder recebe a mensagem e encaminha para todos os followers. Garante totalmente a entrega.

## Situações complexas

- O que acontece se o message broker cair?
- Haverá perda de mensagens?
- Seu sistema ficará fora do ar?
- Como garantir resiliência

## Transaction outbox

Criar uma tabela com registros temporários para armazenas as mensagens, e após mandá-las para o broker.
Quando o broker receber a mensagem eu posso deletar esse item da tabela

## Garantia de recebimento

- Auto Ack = false e commit manual (preciso processar e fazer o commit do recebimento)
- Prefetch alinhado a volumetria (receber um batch de mensagens que será processado pelo consumer)
- Saber quantas mensagens consigo consumir

## Idempotência e políticas de fallback

- O ato de conseguir lidar com duplicidade de informação
- Ter a condição de identificar a duplicidade de mensagens e descartar-las
- Políticas claras de fallback

## Observabilidade

- APM
  - Consegue monitor a aplicação
- Tracing Distribuído
  - Monitora o caminho que uma requisição percorreu e onde eventualmente ocorreram erros
- Métricas personalizadas
  - Métricas que garantem informações de negócio e de aplicação
- Spans personalizados
  - Mostrar tudo que acontece dentro do software, de forma clara

# Referências

- Exponential backoff and Jitter: https://aws.amazon.com/pt/blogs/architecture/exponential-backoff-and-jitter/
- Remédio ou Veneno - https://www.youtube.com/watch?v=1MkPpKPyBps
- OTEL - https://opentelemetry.io/
