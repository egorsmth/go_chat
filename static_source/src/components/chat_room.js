import ReactDOM from 'react-dom';
import React from 'react';
import Message from './message';
import SendMessageForm from './send-message-form'

export default class ChatRoom extends React.Component {
    renderMessages() {
        if (this.state.messages.length > 0) {
            return this.state.messages.map(message => {
                return <Message key={message.id} messageData={message} />
            })
        }
        return <div className='row'>
            There is no messages yet!
        </div>
    }

    render () {
        return <div id='chatRoom'>
            {this.renderMessages()}
            <SendMessageForm />
        </div>
    }
}