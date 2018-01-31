import ReactDOM from 'react-dom';
import React from 'react';
import cn from 'classnames';


export default class Message extends React.Component {
    render() {
        const message = this.props.messageData
        const messageClassName = cn({
            'row': true,
            'msg-unread': message.status == 'unread'
        })
        return <div className={messageClassName} >
            <div className='col-4'>
            <img className='img-fluid' src={`/${message.author.avatar}`} />
                {message.author.username}
                {message.created}
            </div>
            <div className='col-8'>
                {message.text}
            </div>
        </div>

    }
}