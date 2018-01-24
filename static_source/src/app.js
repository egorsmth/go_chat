import React from 'react';
import ChatRooms from './components/chat-rooms'
import ChatRoom from './components/chat-room'
import { VIEW_CHAT_ROOMS, VIEW_CHAT_ROOM } from '../index'
import { getWsConnection, send } from './wsController'

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

    wsMessageDispatch = resp => {
        switch (resp.type){
            case 'MESSEGE_RECIEVED':
                const message = JSON.parse(resp.data)
                const roomMessages = [...this.state.messages[message.chat_room_id], message];
                const newMessages = Object.assign({}, this.state.messages, {
                    [message.chat_room_id]: roomMessages
                });

                const chatRooms = this.state.chatRooms.map(room => {
                    if (room.id == message.chat_room_id) {
                        return {
                            id: room.id,
                            lastMessage: message,
                            lastMessageId: message.id,
                            status: room.status
                        }
                    }
                    return room
                })
                this.setState({
                    chatRooms: chatRooms,
                    messages: newMessages,
                    view: this.state.view,
                    roomId: this.state.roomId
                });
                break
            case 'TYPING':
                break
            case 'READ':
                break
        }
        // var mesasgeHtml = document.createElement('div');
        // mesasgeHtml.innerHTML = `<p>${msg.username}</p><p>${msg.message}</p><p>${msg.created}</p>`;
        // var element = document.getElementById('messages-block');
        // element.appendChild(mesasgeHtml)
    }

    clickChatRoom = id => {
        const ws = getWsConnection(id, this.wsMessageDispatch);
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            view: VIEW_CHAT_ROOM,
            roomId: id,
            ws: ws
        });
    }

    clickMyMessages = e => {
        e.preventDefault();
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            view: VIEW_CHAT_ROOMS,
            roomId: -1
        });
    }

    submitMessage = text => {
        send(this.state.ws, this.state.roomId, text, window.user.id)
    }

    renderChatRooms () {
        return <ChatRooms chatRooms={this.state.chatRooms} clickChatRoom={this.clickChatRoom}/>
    }

    renderChatRoom() {
        const chatRoomMessages = this.state.messages[this.state.roomId] || []
        return <ChatRoom messages={chatRoomMessages} submitMessage={this.submitMessage}/>
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