### Requirement
- Fiber
- Go Dot Env
- JWT
- Bcrypt
- Gorm & Postgres
- Swagger
- Swagger CLI

### Struktur Folder
|
|- config/              -> koneksi DB, load env
|- controllers/         -> Fiber HTTP handler
|- services/            -> logic bisnis
|- repositories/        -> akses database
|- models/              -> entitas/data struct
|- middleware/          -> auth, role guard, dll
|- routes/              -> definisi endpoint API
|- utils/               -> helper seperti JWT, hgas
|- database/
|   |- migrations/      -> folder untuk file migrate
|   |- seed             -> seed admin, user
|
|- docs/                -> swagger docs
|- .env
|- cmd/main.go
|
