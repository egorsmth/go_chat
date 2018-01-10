import React from 'react';
import ChatRooms from './components/chat-rooms'
import ChatRoom from './components/chat-room'
import { VIEW_CHAT_ROOMS } from '../index'

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            chatRooms: props.appData.chatRooms,
            messages: props.appData.messages,
            view: props.view
        }
    }

    renderChatRooms () {
        return <ChatRooms chatRooms={this.state.chatRooms} />
    }

    renderChatRoom() {
        return <ChatRoom messages={this.state.messages} />
    }

    render() {
        if (this.state.view == VIEW_CHAT_ROOMS) {
            return this.renderChatRooms()
        }
        return this.renderChatRoom()
    }
}