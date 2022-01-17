./kata_template -package=$1
mkdir ../$1
mv output/* ../$1
cd ../$1
go mod init okscoring.com/$1
go mod tidy
code .
