# Go API

User
POST /signup
body

```json
{
	"email": "emmail@test.com",
	"password": "123"
}
```

POST /login
body

```json
{
	"email": "emmail@test.com",
	"password": "123"
}
```

Events

Gets all events
GET /events/

Gets event by id
GET /events/:id

Creates event
POST /events/
body

```json
{
	"Name": "pal norte",
	"Description": "Travel",
	"Location": "Nuevo Leon",
	"DateTime": "2024-03-27T06:00:00Z"
}
```

Update event
PUT /events/:id
body

```json
{
	"Name": "pal norte",
	"Description": "Travel",
	"Location": "Parque fundidora",
	"DateTime": "2024-03-27T06:00:00Z"
}
```

Delete event
DELETE /events/:id
