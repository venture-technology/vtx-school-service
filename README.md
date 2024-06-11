<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> Somos segurança, velocidade e tecnologia. Somos Venture </h1>

<p align="center">
  <img src="https://i.imgur.com/yieDOSJ.png"/>
</p>

## ⚙️ API Endpoints

Todas as rotas possuem `api/v1`. Antecendo, como prefixo da rota.

### GET /ping

Retorna uma simples mensagem de "pong" para validar funcionamento da aplicação.

**Resposta**

```json
{
    "message": "pong"
}
```
---

### POST /school

Cria uma conta de escola na plataforma

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `id`     | body  | int    | ID da escola.        |
| `name`   | body  | string | Nome da escola.      |
| `cnpj`   | body  | string | CNPJ da escola.      |
| `email`  | body  | string | E-mail da escola.    |
| `password` | body | string | Senha da escola.    |
| `street` | body  | string | Rua da escola.       |
| `number` | body  | string | Número da escola.    |
| `zip`    | body  | string | CEP da escola.       |    

**Resposta**

```json
{
  "school": {
    "id": 456,
    "name": "Escola Estadual",
    "cnpj": "12.345.678/0001-90",
    "email": "escola@exemplo.com",
    "password": "segredo123",
    "street": "Rua da Escola, 123",
    "number": "456",
    "zip": "12345-678"
  }
}
```

---

### POST /login/school

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
}
"school": {
        "id": 7,
        "name": "E.M.E.F Professor Carlos Pasquale",
        "cnpj": "64025893000102",
        "email": "gustavorodrigueslima2004@gmail.com",
        "password": "",
        "street": "",
        "number": "",
        "complement": "",
        "zip": ""
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbnBqIjoiNjQwMjU4OTMwMDAxMDIiLCJleHAiOjE3MTgxOTI1MTZ9.Yf9fSXAh_akn5M5ZQvwtMzBtCHlSEznOyujR_0XBvFM"
}
```

---

### GET /school/:cnpj

Verifique uma conta de escola.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `cnpj` | uri | string  | CNPJ da escola. |     

**Resposta**

```json
{
    "school": {
      "id": 123,
      "name": "XXXXXXXX",
      "email": "john.doe@example.com",
      "cnpj": "123456789",
      "street": "123 Main Street",
      "number": "456",
      "zip": "12345-678",
      "complement": "Apt 101"
    } 
}
```

---

### GET /school

Verifique uma conta de escola.
   

**Resposta**

```json
{
    "schools": [
        {
            "id": 5,
            "name": "EE Professor Armando Gomes Araujo",
            "cnpj": "48480362000153",
            "email": "emailx@gmail.com",
            "password": "",
            "street": "Rua Alfredo Pariense",
            "number": "1",
            "complement": "",
            "zip": "08110220"
        },
        {
            "id": 6,
            "name": "E.M.E.F Professor Carlos Pasquale",
            "cnpj": "64025893000102",
            "email": "emaily@gmail.com",
            "password": "",
            "street": "Avenida Barão de Alagoas",
            "number": "322",
            "complement": "",
            "zip": "0812000"
        }
    ]
}
```

---

### PATCH /school

Altera uma conta de escola na API

**Parâmetros**

| Nome     | Local | Tipo   | Descrição            |
|----------|-------|--------|----------------------|
| `name`   | body  | string | Nome da escola.      |
| `email`  | body  | string | E-mail da escola.    |
| `password` | body | string | Senha da escola.    |
| `street` | body  | string | Rua da escola.       |
| `number` | body  | string | Número da escola.    |
| `zip`    | body  | string | CEP da escola.       |    

**Resposta**

```json
{
    "message": "updated w successfully"
}
```

---

### DELETE /school

Deleta uma conta de escola na API

**Parâmetros**

Baseado no Middleware

**Resposta**

```json
{
    "message": "deleted w successfully"
}
```

