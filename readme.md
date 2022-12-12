# Conceitos básicos

Microsserviços são:

- aplicações comuns
- independentes
- com objetivos bem definidos
- que fazem parte de um ecossistema.

# Microsserviços vs Monolíticos

## Microsserviços

- Objetivos segregados, bem definidos.
- Diversas tecnologias
- Menor risco no deploy
- Uma equipe por microsserviço
- Mais complexo para iniciar projetos que ainda não sabemos todas as regras ou implicações

## Monolíticos

- Todos os contextos dentro do mesmo sistema
- Única tecnologia
- Maior risco no deploy
- Todas as equipes trabalham no mesmo sistema
- É mais simples para iniciar um projeto

# Quando utilizar microsserviços

- Quando você quiser escalar times, separando-os por microsserviço.
- Quando eu tenho contextos/área de negócios bem definidos;
- Quando você tem maturidade nos processos de entrega
- Quando você tem maturidade técnica dos times
- Quando tenho a necessidade de escalar apenas uma parte do sistema
- Quando eu preciso de tecnologias distintas/específicas em partes do sistema.

# Quando utilizar sistemas monolíticos

- Provas de conceito
- Quando não conhecemos todo o domínio
- governança simplificada sobre as tecnologias utilizadas
- Facilidade na contratação e treinamento de desenvolvedores
- Tudo está no mesmo lugar
- Compartilhamento claro de libs (shared kernel)

# Migração de monolito para microsserviços

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
