import ReactDOM from 'react-dom';
import React from 'react';
import cn from 'classnames';

export default class ChatRoomBlock extends React.Component {
    onClick = () => {
        this.props.onClick(this.props.chatRoomData.id)
    }

    render() {
        console.log(this.props)
        const mesasge = this.props.chatRoomData.lastMessage
        const messageClassName = cn({
            'col-8': true,
            'msg-unread': mesasge.status == 'unread'
        })
        return <div className='row' onClick={this.onClick}>
            <div className='col-12'>
                <div className='col-4'>
                    {/* <img src={user.avatar}/> */}
                    {mesasge.created}
                </div>
                <div className={messageClassName}>{mesasge.message}</div>
            </div>
        </div>
    }
}