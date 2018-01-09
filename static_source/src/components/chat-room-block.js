import ReactDOM from 'react-dom';
import React from 'react';
import cn from 'classnames';

export default class ChatRoomBlock extends React.Component {
    onClick = () => {
        this.props.onClick(this.props.key)
    }

    render() {
        const mesasge = this.props.chatRoomData.message
        const user = this.props.ChatRoom.user
        const messageClassName = cn({
            'col-8': true,
            'msg-unread': mesasge.status == 'unread'
        })
        return <div className='row' onClick={this.onClick}>
            <div className='col-12'>
                <div className='col-4'>
                    <img src={user.avatar}/>
                </div>
                <div className={messageClassName}>{mesasge.text}</div>
            </div>
        </div>
    }
}