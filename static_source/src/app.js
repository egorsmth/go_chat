import React from 'react';
import ChatRooms from './components/chat-rooms'
import ChatRoom from './components/chat-room'

export default class App extends React.Component {
    renderChatRooms () {
        return <ChatRooms />
    }

    renderChatRoom() {
        return <ChatRoom />
    }

    render() {
        if (this.state.view == 'chat-rooms') {
            return this.renderChatRooms()
        }
        return this.renderChatRoom()
    }
}