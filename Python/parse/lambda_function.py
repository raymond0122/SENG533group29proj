import json

def parse(in_str):
    counts = [0] * 26

    for i in range(0, len(in_str)):
        val = ord(in_str[i])
        if val > 96 and val < 123:
            counts[val - 97] += 1

    return counts


def lambda_handler(event, context):
    # TODO implement
    size = event['size']
    data = open(f'string_input_{size}.json')
    return {
        'statusCode': 200,
        'body': parse(json.load(data)['data'])
    }
