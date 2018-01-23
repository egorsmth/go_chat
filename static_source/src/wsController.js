const getWsConnection = (roomId) => {
    const ws = new WebSocket(`ws://${window.location.host}/chat/ws?id=${roomId}`); //

    ws.addEventListener('message', function(e) {
        console.log('message comming!!!')
        console.log(e)
        // dispatchWsResponse(e)
    });

    // const btn = document.getElementById('btn-send')
    // btn.addEventListener('click', (e) => {
    //     console.log('clicked')
    //     const inp = document.getElementById('inp-msg').value
    //     console.log(inp)
    //     if (inp != '') {
    //         const now = new Date();
    //         console.log(now.toISOString())
    //         ws.send(
    //             JSON.stringify({
    //                 chat_room_id: window.chat_room_id,
    //                 user_id: window.user_id,
    //                 message: inp,
    //                 created: now.toISOString()
    //             })
    //         );
    //         document.getElementById('inp-msg').value = ''; // Reset newMsg
    //     }
    // });

    // dispatchWsResponse = (msgResponse) => {
    //     var msg = JSON.parse(e.data);
    //     var mesasgeHtml = document.createElement('div');
    //     mesasgeHtml.innerHTML = `<p>${msg.username}</p><p>${msg.message}</p><p>${msg.created}</p>`;
    //     var element = document.getElementById('messages-block');
    //     element.appendChild(mesasgeHtml)
    // }
}

const send = (ws, roomId, text, userId) => {
    const now = new Date();
    const message = JSON.stringify({
        chat_room_id: roomId,
        user_id: userId,
        text: text,
        created: now.toISOString()
    });
    console.log('send!!!!')
    console.log(message)
    ws.send(message);
}

export { getWsConnection, send }