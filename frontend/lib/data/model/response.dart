import 'dart:convert';

class Result {
  int? code;
  Tx? data;
  String? msg;

  Result({this.code, this.data, this.msg});

  Result.fromJson(Map<String, dynamic> json) {
    code = json['code'];
    data = json['data'] != null ? Tx.fromJson(json['data']) : null;
    msg = json['msg'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['code'] = code;
    if (this.data != null) {
      data['data'] = this.data!.toJson();
    }
    data['msg'] = msg;
    return data;
  }
}

class Tx {
  int? nonce;
  int? gasPrice;
  int? gasLimit;
  String? to;
  int? value;
  String? data;
  String? from;
  int? chainId;
  String? r;
  String? v;
  String? s;

  Tx(
      {this.nonce,
      this.gasPrice,
      this.gasLimit,
      this.to,
      this.value,
      this.data,
      this.from,
      this.chainId,
      this.r,
      this.v,
      this.s});

  Tx.fromJson(Map<String, dynamic> json) {
    nonce = json['nonce'];
    gasPrice = json['gasPrice'];
    gasLimit = json['gasLimit'];
    to = json['to'];
    value = json['value'];
    data = json['data'];
    from = json['from'];
    chainId = json['chainId'];
    r = json['r'];
    v = json['v'];
    s = json['s'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['nonce'] = nonce;
    data['gasPrice'] = gasPrice;
    data['gasLimit'] = gasLimit;
    data['to'] = to;
    data['value'] = value;
    data['data'] = this.data;
    data['from'] = from;
    data['chainId'] = chainId;
    data['r'] = r;
    data['v'] = v;
    data['s'] = s;
    return data;
  }
}

String getPrettyJSONString(jsonObject) {
  var encoder = const JsonEncoder.withIndent('  ');
  return encoder.convert(jsonObject);
}
