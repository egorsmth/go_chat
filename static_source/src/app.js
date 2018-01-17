import React from 'react';
import ChatRooms from './components/chat-rooms'
import ChatRoom from './components/chat-room'
import { VIEW_CHAT_ROOMS, VIEW_CHAT_ROOM } from '../index'

export default class App extends React.Component {
    constructor(props) {
        console.log(props.appData)
        super(props);
        this.state = {
            chatRooms: props.appData.chatRooms,
            messages: props.appData.messages,
            view: props.view,
            roomId: -1
        }
    }

    clickChatRoom = id => {
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            view: VIEW_CHAT_ROOM,
            roomId: id
        })
    }

    clickMyMessages = e => {
        e.preventDefault();
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            view: VIEW_CHAT_ROOMS,
            roomId: -1
        })
    }

    renderChatRooms () {
        return <ChatRooms chatRooms={this.state.chatRooms} clickChatRoom={this.clickChatRoom}/>
    }

    renderChatRoom() {
        const chatRoomMessages = this.state.messages[this.state.roomId] || []
        return <ChatRoom messages={chatRoomMessages} />
    }

    renderAppBlock() {
        if (this.state.view == VIEW_CHAT_ROOMS) {
            return this.renderChatRooms()
        }
        return this.renderChatRoom()
    }

    render() {
        return <div className="col-12">
        <div className="col-3">
            <h3>{this.props.user.username}</h3>
            <a onClick={this.clickMyMessages}>My Messages</a>
        </div>
        <div className="col-9">
            {this.renderAppBlock()}
        </div>
    </div>
        
    }
}