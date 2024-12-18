import requests

url = "https://gateway.apiserpro.serpro.gov.br/token"

body = "grant_type=client_credentials"
headers = {
    'Content-Type': "application/x-www-form-urlencoded",
    'Authorization': "Basic cnNfZjd5T3FZYTlLNXM0NWtrdW42RUJxd2NNYTp4aFZxVUNaNjlQeWp3bzZhSXAyUERXZDhweThh",
    }

response = requests.request("POST", url, data=body, headers=headers)

print(response.text)
print(response.headers)