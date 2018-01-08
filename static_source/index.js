// const ws = new WebSocket(`ws://${window.location.host}/ws?id=${window.chat_room_id}`);

// ws.addEventListener('message', function(e) {
//     console.log('message comming!!!')
//     console.log(e)
//     var msg = JSON.parse(e.data);
//     var mesasgeHtml = document.createElement('div');
//     mesasgeHtml.innerHTML = `<p>${msg.username}</p><p>${msg.message}</p><p>${msg.created}</p>`;
//     var element = document.getElementById('messages-block');
//     element.appendChild(mesasgeHtml)
// });

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
import ReactDOM from 'react-dom';

ReactDOM.render(
    <App />,
    document.getElementById('container')
);