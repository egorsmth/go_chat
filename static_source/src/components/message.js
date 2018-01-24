import ReactDOM from 'react-dom';
import React from 'react';

export default class Message extends React.Component {
    render() {
        const message = this.props.messageData
        return <div className='row'>
            <div className='col-12'>
                <div className='col-4'>
                    {message.author.avatar}
                    {message.author.username}
                    {message.created}
                </div>
                <div className='col-8'>
                    {message.text}
                </div>
            </div>
        </div>
    }
}