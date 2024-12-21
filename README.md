# YaCalc
It's an API for servis calculator

How to install:
----------------
1. Install [golang](https://go.dev/doc/install)
2. Enter the command:
``` bash
git clone github.com/StepanShel/YaCalc
```
How to use
----------
1. Enter the command
```bash
go run ./cmd/main.go
```
2. Send cURl requests
```bash
curl --location 'localhost:8080/api/v1/calculate' \         
--header 'Content-Type: application/json' \         
--data '
{
  "expression": "your_expression"
}'
```
>[!TIP]
>
>1. IF you've got Windows, it's better to send requests to **Git Bash** (or **Postman**)
>2. Default localhost port is :8080. If you want to change it, add PORT value(in bash, not in shell) like this:
>```bash
>export PORT=8020 && go run ./cmd/main.go
>```

Example Responses
---------------
1. Status Ok response:
>[!TIP]
>
>``` shell
>curl --location 'localhost:8080/api/v1/calculate' \
>--header 'Content-Type: application/json' \
>--data '{
>  "expression": "2+2*2"
>}'
>```
2. Status 422 response's (entity unprocessable):
> [!CAUTION]
>
> ``` shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2+2*2--"
> }'
> ```

> [!CAUTION]
>
> ``` shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2/0"
> }'
> ```

> [!CAUTION]
>
> ``` shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2+2)"
> }'
> ```

3. Status 405 resronse (method other than POST):
> [!CAUTION]
>
> ``` shell
> curl -l 'localhost:8080/api/v1/calculate' \
> ```
