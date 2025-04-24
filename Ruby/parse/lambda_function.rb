require 'json'

def parse(in_str)
    counts = Array.new(26, 0)

    for i in 0..(in_str.length - 1) do
        val = in_str[i].ord
        if val > 96 && val < 123
            counts[val - 97] += 1
        end
    end
    return counts
end

def lambda_handler(event:, context:)
    filename = "string_input_%d.json" % [event['size']]
    file = File.read(filename)
    text = JSON.parse(file)['data']
    { statusCode: 200, body: parse(text) }
end
