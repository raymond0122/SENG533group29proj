require 'json'

def sort(arr)
    i = 0
    while i < arr.length - 1
    current_elem = arr[i]
    next_elem = arr[i+1]

    if current_elem > next_elem
        arr[i] = next_elem
        arr[i+1] = current_elem
        i = 0
    else
        i += 1
    end

    end
end

def lambda_handler(event:, context:)
    # TODO implement
    arr = event['body']
    sort(arr)
    { statusCode: 200, body: arr }
end
