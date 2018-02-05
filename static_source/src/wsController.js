const getWsConnection = (roomId, dispatcher) => {
    const ws = new WebSocket(`ws://${window.location.host}/chat/ws?id=${roomId}`); //

    ws.addEventListener('message', function(e) {
        const resp = JSON.parse(e.data);
        if (resp.status == 'success') {
            dispatcher(resp)
            return
        }
        console.error(e)
    });

    return ws
}

const wait = (ws, cb) => {
    if (ws.readyState !== 1) {
        setTimeout(cb, 1000);
    } else {
        cb()
    }
}

const send = (ws, roomId, text, userId) => {
    wait(ws, () => {
        const now = new Date();
        const message = JSON.stringify({
            action: 'send',
            data: {
                chat_room_id: roomId,
                author_id: userId,
                text: text,
                status: 'unread',
                created: now.toISOString()
            },
        });
        ws.send(message);
    });
}

const makeRead = (ws, roomId, messageIds) => {
    wait(ws, () => {
        const message = JSON.stringify({
            action: 'read',
            data: {
                messageIds: messageIds,
                roomId: roomId,
            }
        });
        ws.send(message)
    });
}

export { getWsConnection, send, makeRead }