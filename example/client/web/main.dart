import 'dart:html';
import 'dart:convert';

WebSocket connectClient() {
    var ws = new WebSocket("ws://localhost:4242/client/connect");
    ws.onMessage.listen((MessageEvent m) {
      print("Message receive : ${m.data}");
      querySelector("#msg_receive").appendHtml("<p>${m.data}</p>");
    });
    return ws;
}

void setupSending(WebSocket ws) {
  querySelector("#submit").onClick.listen((MouseEvent e) {
    print("Submit click");
    var id = int.parse((querySelector("#id_to_send") as InputElement).value);
    var msg = (querySelector("#msg_to_send") as InputElement).value;

    ws.sendString(JSON.encode({"topic": "msg", "to": id, "body": msg}));
  });
}

main() async {
    var ws = connectClient();
    setupSending(ws);
}
