import ReactDOM from 'react-dom';
import React from 'react';
import 'bootstrap';
import './style.scss'
import App from './src/app'

export const VIEW_CHAT_ROOMS = 'chat-rooms'
export const VIEW_CHAT_ROOM = 'chat-room'

ReactDOM.render(
    <App view={VIEW_CHAT_ROOMS} appData={window.appData} user={window.user}/>,
    document.getElementById('app')
);