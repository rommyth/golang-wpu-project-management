### Requirement
- Fiber
- Go Dot Env
- JWT
- Bcrypt
- Gorm & Postgres
- Swagger
- Swagger CLI

### Struktur Folder
```
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
```

Flow Struktur Folder

- Dari Routes -> Controller -> Service -> Repository -> Model

- Routes ini sebagai api request dari frontend
- Controller ini sebagai handler dari routes
- Service ini sebagai logic bisnis
- Repository ini sebagai akses database
- Model ini sebagai entitas/data struct
- Middleware ini sebagai auth, role guard, dll
- Utils ini sebagai helper seperti JWT, hash password, dll
- Config ini sebagai koneksi DB, load env
- Database ini sebagai folder untuk file migrate dan seed
- Docs ini sebagai swagger docs
- .env ini sebagai file environment
- cmd/main.go ini sebagai entry point aplikasi
