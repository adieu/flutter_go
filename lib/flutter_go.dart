import 'dart:async';

import 'package:flutter/services.dart';

class Go {
  static const BasicMessageChannel _channel =
      const BasicMessageChannel('go_rpc', const BinaryCodec());
  static const JSONMessageCodec _jsonCodec = const JSONMessageCodec();

  static Future sendRequest(String method, [parameters]) {
    return _send(method, parameters);
  }

  static void sendNotification(String method, [parameters]) =>
      _send(method, parameters);

  static Future _send(String method, parameters) {
    if (parameters is Iterable) parameters = parameters.toList();
    if (parameters is! Map && parameters is! List && parameters != null) {
      throw new ArgumentError('Only maps and lists may be used as JSON-RPC '
          'parameters, was "$parameters".');
    }

    var message = <String, dynamic>{"jsonrpc": "2.0", "method": method, "params":[]};
    if (parameters != null) message["params"] = [parameters];

    var completer = new Completer();

    _channel.send(_jsonCodec.encodeMessage(message)).then((response) {
      var result;
      try {
        result = _jsonCodec.decodeMessage(response);
      } catch(e) {
        completer.completeError(e);
        return;
      }
      if (result.containsKey("error") && result["error"] != null) {
        completer.completeError(result["error"]);
      } else {
        completer.complete(result["result"]);
      }
    });
    return completer.future;
  }
}
