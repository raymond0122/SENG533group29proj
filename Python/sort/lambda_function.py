import json

def sort(arr):
    i = 0
    while i < len(arr) - 1:
        current_elem = arr[i]
        next_elem = arr[i + 1]
        if current_elem > next_elem:
            arr[i] = next_elem
            arr[i + 1] = current_elem
            i = 0
            continue
        i += 1
        

def lambda_handler(event, context):
    # TODO implement
    size = event['size']
    data = open(f'pi_input_{size}.json')
    arr = json.load(data)['numbers']
    sort(arr)
    return {
        'statusCode': 200,
        'body': "Sorted."
    }
