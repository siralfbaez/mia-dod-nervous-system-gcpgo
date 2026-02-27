#!/bin/bash
# MIA-DoD Database Initialization Script

echo "🚀 Initializing AlloyDB/Postgres environment..."

# Check if psql is installed
if ! command -v psql &> /dev/null
then
    echo "❌ Error: psql is not installed. Please install postgresql-client."
    exit 1
fi

# Run migrations (assuming DATABASE_URL is set)
echo "📂 Running migrations from ../migrations/..."
psql "$DATABASE_URL" -f ../migrations/000001_init_schema.up.sql

# Seed the database
echo "🌱 Seeding test data..."
psql "$DATABASE_URL" -f ../seeds/test_menu_items.sql

echo "✅ Database setup complete."