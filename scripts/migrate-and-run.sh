#!/bin/sh

echo $ENVIRONMENT_NAME

./output/migrations

if [[ $ENVIRONMENT_NAME == "docker" ]]; then
    echo "seeding"
    ./output/seeder
fi

./output/server