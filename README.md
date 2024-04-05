# Instalacion

0. Correr los microservicios dependientes a orchestrator, utilizar la rama develop de cada microservicio

 - https://github.com/b-munar/auth

 - https://github.com/b-munar/sportmen

 - https://github.com/b-munar/subscription

 - https://github.com/b-munar/partner

1. Clonar repositorio

```bash
git clone https://github.com/b-munar/orchestrator.git
```

2. Copiar el .env.template y pegarlo en un .env

3. 

```bash
docker compose build
docker compose up
```


El servicio de orchestrator consume el servicio auth, sportmen y subscription. Este servicio ahorra el trabajo de hacer register/login en auth, register/get en sportman y get en subscription.

## 1. Register User Sportmen Subscription

Crea un usuario sportment con los datos brindados, el email del usuario debe ser único.

<table>
<tr>
<td> Método </td>
<td> POST </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6350/orchestrator/register/sportman</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
{
  "email": "sportman@email.com",
  "password": "password",
  "name": "Brahian",
  "last_name": "Munar",
  "age": 26,
  "weight": 63,
  "height": 163,
  "country_birth": "Colombia",
  "city_birth": "Cali",
  "country_residence": "Colombia",
  "city_residence": "Elvira",
  "length_residence": 26,
  "sports": [
    {
      "sport": "basketball"
    }
  ]
}
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 201 </td>
<td>En el caso que el usuario se haya creado con éxito.</td>
<td>

```json
{
  "auth": {
    "email": "sportman@email.com",
    "token": "eyJ0eXA...",
    "role": 1
  },
  "sportmen": {
    "name": "Brahian",
    "last_name": "Munar",
    "age": 26,
    "weight": 63,
    "height": 163,
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "length_residence": 26,
    "sports": [
      {
        "sport": "basketball"
      }
    ]
  },
  "subscription": [
    {
      "plan": "basico",
      "price": 0
    },
    {
      "plan": "intermedio",
      "price": 19
    },
    {
      "plan": "premium",
      "price": 39
    }
  ]
}
```
</td>
</tr>
</tbody>
</table>

## 2. Register User Partner

Crea un usuario sportment con los datos brindados, el email del usuario debe ser único.

<table>
<tr>
<td> Método </td>
<td> POST </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6350/orchestrator/register/partner</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
{
  "email": "partner@email.com",
  "password": "password",
  "name": "parnert",
  "last_name": "parnert",
  "age": 20,
  "profession": "cowboy",
  "license": "Hunter X",
  "identification_type": "CC",
  "identification": "314159",
  "country_birth": "Colombia",
  "city_birth": "Cali",
  "country_residence": "Colombia",
  "city_residence": "Elvira",
  "sports": ["basketball"],
  "companies": ["Uniandes"],
  "type_services": ["education"]
}
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 201 </td>
<td>En el caso que el usuario se haya creado con éxito.</td>
<td>

```json
{
  "auth": {
    "email": "partner@email.com",
    "token": "eyJ0eXA...",
    "role": 2
  },
  "partner": {
    "name": "parnert",
    "last_name": "parnert",
    "age": 20,
    "profession": "cowboy",
    "license": "Hunter X",
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "sports": [
      "basketball"
    ],
    "companies": [
      "Uniandes"
    ],
    "type_services": [
      "education"
    ]
  }
}
```
</td>
</tr>
</tbody>
</table>




## 3. Login Sportmen

Inicia sesion de usuario con los datos brindados

<table>
<tr>
<td> Método </td>
<td> POST </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6350/orchestrator/login</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
{
  "email": "sportman@email.com",
  "password": "password"
}
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 202 </td>
<td>En el caso que el usuario hizo un login exitoso.</td>
<td>

```json
{
  "auth": {
    "email": "sportman@email.com",
    "token": "eyJ0eXA...",
    "role": 1
  },
  "sportmen": {
    "name": "Brahian",
    "last_name": "Munar",
    "age": 26,
    "weight": 63,
    "height": 163,
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "length_residence": 26,
    "sports": [
      {
        "sport": "basketball"
      }
    ]
  },
  "subscription": [
    {
      "plan": "basico",
      "price": 0
    },
    {
      "plan": "intermedio",
      "price": 19
    },
    {
      "plan": "premium",
      "price": 39
    }
  ]
}
```
</td>
</tr>
</tbody>
</table>

## 4. Login Partners

Inicia sesion de usuario con los datos brindados

<table>
<tr>
<td> Método </td>
<td> POST </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6350/orchestrator/login</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
{
  "email": "partner@email.com",
  "password": "password"
}
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 202 </td>
<td>En el caso que el usuario hizo un login exitoso.</td>
<td>

```json
{
  "auth": {
    "email": "partner@email.com",
    "token": "eyJ0eXA...",
    "role": 2
  },
  "partner": {
    "name": "parnert",
    "last_name": "parnert",
    "age": 20,
    "profession": "cowboy",
    "license": "Hunter X",
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "sports": [
      "basketball"
    ],
    "companies": [
      "Uniandes"
    ],
    "type_services": [
      "education"
    ]
  }
}
```
</td>
</tr>
</tbody>
</table>
