# Instalacion

0. Correr los microservicios dependientes a orchestrator

 - https://github.com/b-munar/auth

 - https://github.com/b-munar/sportmen

 - https://github.com/b-munar/subscription

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
<td> <strong>localhost:6350/orchestrator/register</strong> </td>
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
  "email": "lacabra@email.com",
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
    "email": "lacabra@email.com",
    "token": "eyJ0eXA..."
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


## 2. Login

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
  "email": "lacabra@email.com",
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
    "email": "lacabra@email.com",
    "token": "eyJ0eXA..."
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
