import ReactDOM from 'react-dom';
import React from 'react';

export default class Message extends React.Component {
    render() {
        const message = this.props.messageData
        return <div className='row'>
            <div className='col-12'>
                <div className='col-4'>
                    {this.props.messageData.user.avatar}
                </div>
                <div className='col-8'>
                    {this.props.messageData.text}
                </div>
            </div>
        </div>
    }
}