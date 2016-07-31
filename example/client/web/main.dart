import 'dart:html';
import 'dart:convert';

WebSocket connectClient() {
    var ws = new WebSocket("ws://localhost:4242/client/connect");
    ws.onMessage.listen((MessageEvent m) {
      print("Message receive : ${m.data}");
      var msg = JSON.decode(m.data);
      if (msg["topic"] == "msg") {
        querySelector("#msg_receive").appendHtml("<p>${m.data}</p>");
      } else if (msg["topic"] == "myId") {

      } else if (msg["topic"] == "clientList") {
        querySelector("#client_list").setInnerHtml("${msg['body']}");
      }
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
  querySelector("#submit_client_list").onClick.listen((MouseEvent e) {
    ws.sendString(JSON.encode({"topic": "clientList"}));
  });
}

main() async {
    var ws = connectClient();
    setupSending(ws);
    setupGetClientList(ws);
}
