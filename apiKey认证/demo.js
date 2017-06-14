// 使用到的库: https://github.com/emn178/js-sha256

// 参与签名生成的 header 字段可以考虑实际业务情况:
var header_key_to_sign = ['host', 'content-length', 'content-type', 'requestURL']

function get_header_to_sign(headers) {

    var header_to_sign = {}
    var key_containes = {}

    for(var index in header_key_to_sign) {
        var key = header_key_to_sign[index]
        key_containes[key] = headers[key]
    }

    Object.keys(key_containes).sort().forEach(
        function(key) {
            header_to_sign[key] = key_containes[key]
        }
    );

    return header_to_sign
}

function gen_sigurate(method, headers, request_url, sk) {

    var header_to_sign = get_header_to_sign(headers) 

    var x = ''
    Object.keys(header_to_sign).forEach(
        function(key) {
            x += key + header_to_sign[key]
        }
    )

    var string_to_sign = method + '\n' + x + '\n' + request_url + '\n'

    var hash = sha256.create();
    hash.update(sk)
    hash.update(string_to_sign)
    return hash.hex()
}

