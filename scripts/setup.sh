# Create database
createuser -s -w -l -d notabbble
createdb -O notabbble notabbble_dev
createdb -O notabbble notabbble_test