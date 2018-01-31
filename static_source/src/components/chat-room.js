import ReactDOM from 'react-dom';
import React from 'react';
import Message from './message';
import SendMessageForm from './send-message-form'

export default class ChatRoom extends React.Component {
    renderMessages() {
        if (this.props.messages.length > 0) {
            return this.props.messages.map(message => {
                return <Message key={message.id} messageData={message} />
            })
        }
        return <div className='row'>
            There is no messages yet!
        </div>
    }

    render () {
        return <div id='chatRoom' className='row'>
            <div className="messages-scrollable col-12">
                {this.renderMessages()}
            </div>
            
            <SendMessageForm submitMessage={this.props.submitMessage}/>
        </div>
    }
}