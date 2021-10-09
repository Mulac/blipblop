This is where all our executable stuff goes.

Try:
- `DEBUG=true go run scraper/main.go` to test scraper output
- `DB_DRIVER=sqlite3 DB_NAME=prototype.db SCRAPER=apify go run scraper/main.go` to fill the prototype database
- `DB_DRIVER=sqlite3 DB_NAME=prototype.db go run backend/main.go` to start the server with prototype database