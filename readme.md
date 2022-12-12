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
