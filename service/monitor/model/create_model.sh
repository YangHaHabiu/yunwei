#!/bin/bash
goctl model mysql ddl --src="./*.sql" --dir="./"
rm -rf ./*.sql