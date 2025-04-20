# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands
- **Build**: `go build -o techbloghub-server ./cmd/main.go`
- **Run Tests**: `go test ./...`
- **Single Test**: `go test -v ./path/to/package -run TestName`
- **Start Dev DB**: `./script/dev-db.sh`
- **Start Test DB**: `./script/test-db.sh`
- **Create Migration**: `./script/create_migration.sh migration_name`
- **Stop Test Containers**: `docker stop test-tbh-postgres`

## Code Style
- **Imports**: Standard library first, third-party next, internal packages last
- **Error Handling**: Always check errors with `if err != nil`, return errors up the call stack
- **Naming**: camelCase for private, PascalCase for exported; lowercase package names
- **Testing**: Use testify for assertions, TransactionalTest for DB tests
- **Structure**: HTTP handlers in internal/http/handler/, DB operations with EntGo ORM
- **Comments**: Document public functions, Korean comments acceptable
- **Formatting**: Always run `go fmt ./...` before committing

## Architecture
- HTTP routing via Gin framework
- Database access via EntGo ORM
- Soft delete via SoftDeleteMixin
- Pagination support via TechbloghubPaging
- RSS feed processing with gofeed parser

## Testing
1. **Database Test Setup**:
   - Start test database: `./script/test-db.sh`
   - Clean up after tests: `docker stop test-tbh-postgres`

2. **Running Tests**:
   - Run specific package tests: `go test -v ./internal/package/...`
   - Run all tests: `go test ./...`

3. **Test Structure**:
   - Use `testutils.TransactionalTest` for database tests
   - Create test data using the EntGo client
   - When testing entities, always include all required fields
   - Remember to set all required fields when creating test data (e.g., Company requires name, logo_url, blog_url)

4. **Test Assertions**:
   - Use `assert.NoError()` to check for successful operations
   - Use `assert.Error()` to verify expected errors
   - Use `assert.Equal()` to compare expected vs. actual values
   - Use `assert.Contains()` to check error message contents
