import requests


def lambda_handler(event, context):
    resp = requests.get("https://checkip.amazonaws.com")
    print(resp.text)
