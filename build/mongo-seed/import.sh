#!/bin/bash

sleep 2

# use defaults for MONGO_HOST
echo  "MONGO_HOST: $MONGO_HOST"
echo  "MONGO_USER: $MONGO_USER"
echo  "MONGO_PASS: #####"
echo  "MONGO_FILES: $MONGO_FILES"
echo ""


# run mongo command to import Schema and collections
mongoSeed() {
  cd "$MONGO_FILES" || exit
  for d in */; do
       local schema=${d%/}
       [ "$schema" == "*" ] && break

        echo -e "\n> - schema: $schema"
        local collections=("$schema"/*.json)
        for collection in "${collections[@]}"
        do
            execMongoImport "$schema" "$collection"
        done
  done
}

execMongoImport() {
  local schema=$1
  local fullPath=$2
  local fileName=$(basename -- "$fullPath")
  local collection=${fileName%.*}

  [ "$collection" == "*" ] && echo "  -- invalid collection...skipping" && return

  echo -e "  -- collection: $collection"
  echo -e "\t - fullPath: $fullPath"
  echo -e "\t - fileName: $fileName"

  mongoimport --drop --host "$MONGO_HOST" --username "$MONGO_USER" --password "$MONGO_PASS" --authenticationDatabase admin --db "$schema" --collection "$collection" --type json --jsonArray --file "$fullPath"

}



# shellcheck disable=SC2034
mongoSeed

echo -e "\n success"
