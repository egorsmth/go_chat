ws = new WebSocket('ws://' + window.location.host + '/ws');


function send (newMsg) {
    if (newMsg != '') {
        ws.send(
            JSON.stringify({
                email: this.email,
                username: this.username,
                message: $('<p>').html(this.newMsg).text() // Strip out html
            })
        );
        newMsg = ''; // Reset newMsg
    }
}

function created() {
    ws.addEventListener('message', function(e) {
        var msg = JSON.parse(e.data);
        chatContent = '<div class="chip">'
            + msg.username
            + '</div>'
        var element = document.getElementById('chat-messages');
        element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
}