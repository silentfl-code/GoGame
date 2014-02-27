import 'dart:html';
import 'dart:async';

void main() {
	new Client();
}

class Client {
	var messages;	//лог сообщений textarea
	var message;	//элемент input с отправляемым сообщением
	WebSocket ws;
	WebSocket needReload;
	
	Client() {
		//кнопка "Отправить"
		var send = querySelector('#send');
		messages = querySelector('#messages');
		message = querySelector('#message');
		
		print('Hi');
		send.onClick.listen(SendMessage);

		
		//лог сообщений
		messages = querySelector('#messages');
		send.value = 'Hi';
		
		//вебсокет
		ws = new WebSocket('ws://localhost:8080/ws');
		ws.onOpen.listen((e) => AddMessage('Connect'));
		ws.onError.listen((e) => AddMessage('Error connecting'));
		ws.onMessage.listen(MessageReceive);
		
		/*
		//перезагрузка страницы
		needReload = new WebSocket('ws://localhost:8080/needReload');
		needReload.onMessage.listen((e) => AddMessage('Reloading')); //(e) => window.location.reload(true));
		*/
	}
	
	//отправляет сообщение на удаленный вебсокет
	void SendMessage(MouseEvent e) {
		if (message.value != '') {
			ws.send(message.value);
			message.value = '';
		}
	}
	
	//добавляет сообщение в лог сообщений
	void AddMessage(String msg) {
		messages.value = messages.value + msg + '\n';
	}
	
	//получение сообщения от удаленного узла по протоколу websocket
	void MessageReceive(e) {
		//e.data
	}
}