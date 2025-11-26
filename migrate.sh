#!/bin/sh

set -e

# DB_URL
if [ -z "${DB_URL:-}" ]; then
    echo "Error: DB_URL environment variable not set."
    echo "Please set DB_URL to your database connection string, e.g.:"
    echo "  export DB_URL=\"postgres://user:pass@localhost:5432/dbname?sslmode=disable\""
    exit 1
fi

# Migrations directory path
MIGRATIONS_DIR=${MIGRATIONS_DIR:-"./migrations"}

# Help message
show_help() {
    echo "Usage: $0 [command]"
    echo
    echo "Commands:"
    echo "  create NAME   Create a new migration with the specified name"
    echo "  up [N]        Apply all or N up migrations"
    echo "  down [N]      Apply all or N down migrations"
    echo "  force V       Force migration version to a specific version"
    echo "  goto V        Migrate to a specific version"
    echo "  drop          Drop everything in the database"
    echo "  version       Show current migration version"
    echo
    echo "Examples:"
    echo "  $0 create add_users_table"
    echo "  $0 up"
    echo "  $0 up 1"
    echo "  $0 down"
    echo "  $0 down 1"
    echo "  $0 force 1"
    echo "  $0 goto 5"
}

# Ensure migrations directory exists
ensure_migrations_dir() {
    if [ ! -d "$MIGRATIONS_DIR" ]; then
        mkdir -p "$MIGRATIONS_DIR"
        echo "Created migrations directory: $MIGRATIONS_DIR"
    fi
}

# Check if migrate is installed
check_migrate() {
    if ! command -v migrate &> /dev/null; then
        echo "Error: migrate command not found."
        echo "Please install golang-migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate"
        echo "You can install it with:"
        echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
        exit 1
    fi
}

# Main execution
main() {
    check_migrate
    ensure_migrations_dir

    case "$1" in
        create)
            if [ -z "$2" ]; then
                echo "Error: Migration name required"
                show_help
                exit 1
            fi
            migrate create -ext sql -dir "$MIGRATIONS_DIR" -seq "$2"
            echo "Created migration files in $MIGRATIONS_DIR"
            ;;
        up)
            if [ -z "$2" ]; then
                migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up
            else
                migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up "$2"
            fi
            echo "Migration(s) applied"
            ;;
        down)
            if [ -z "$2" ]; then
                migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" down
            else
                migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" down "$2"
            fi
            echo "Migration(s) rolled back"
            ;;
        force)
            if [ -z "$2" ]; then
                echo "Error: Version number required"
                show_help
                exit 1
            fi
            migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" force "$2"
            echo "Forced migration version to $2"
            ;;
        goto)
            if [ -z "$2" ]; then
                echo "Error: Version number required"
                show_help
                exit 1
            fi
            migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" goto "$2"
            echo "Migrated to version $2"
            ;;
        drop)
            migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" drop
            echo "Database dropped"
            ;;
        version)
            migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" version
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            echo "Error: Unknown command '$1'"
            show_help
            exit 1
            ;;
    esac
}

main "$@"
