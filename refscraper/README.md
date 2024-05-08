## Google APIs
https://developers.google.com/sheets/api/reference/rest
https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/get
https://github.com/googleapis/google-api-go-client/blob/master/sheets/v4/sheets-gen.go

## Spells
https://docs.google.com/spreadsheets/d/1cuwb3QSvWDD7GG5McdvyyRBpqycYuKMRsXgyrvxvLFI

## Feats
https://docs.google.com/spreadsheets/d/1XqQO21AyE2WtLwW0wSjA9ov74A9tmJmVJjrhPK54JHQ

## Magic Items
https://docs.google.com/spreadsheets/d/1NQhHjDXhvFZkMiu09epBXsVwCI6YENg-AUQhF-v5ntU

## Bestiary
https://docs.google.com/spreadsheets/d/1StTeUz_ZBU3pNlW120msjUX34p9cs7kqQbZ2Ym7cSBE


# Postgresql
``` postgres
podman cp pathfinder_ref.sql <containerid>:/import.sql
podman exec pathfinder-psql psql -Upostgres -dpostgres -c "\dt"
```
