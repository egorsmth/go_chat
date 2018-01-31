import ReactDOM from 'react-dom';
import React from 'react';
import cn from 'classnames';

export default class ChatRoomBlock extends React.Component {
    onClick = () => {
        this.props.onClick(this.props.chatRoomData.id)
    }

    render() {
        const message = this.props.chatRoomData.lastMessage
        if (!message) {
            return <div className='row' onClick={this.onClick}>
                No messages yet!
            </div>
        }
        const messageClassName = cn({
            'col-8': true,
            'msg-unread': message.status == 'unread'
        })

        return <div className='row' onClick={this.onClick}>
            <div className='col-4'>
                <img className="img-fluid" src={`/${message.author.avatar}`} />
                {message.author.username} {message.created}
            </div>
            <div className={messageClassName}>{message.text}</div>
        </div>
    }
}