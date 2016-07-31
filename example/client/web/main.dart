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

void setupGetClientList(WebSocket ws) {
  querySelector("#client_list").onClick.listen((MouseEvent e) {
    ws.sendString(JSON.encode({"topic": "clientList"}));
  });
}

main() async {
    var ws = connectClient();
    setupSending(ws);
    setupGetClientList(ws);
}
