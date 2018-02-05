import React from 'react';
import ChatRooms from './components/chat-rooms'
import ChatRoom from './components/chat-room'
import { Menu } from './components/menu'
import { Avatar } from './components/avatar'
import { VIEW_CHAT_ROOMS, VIEW_CHAT_ROOM } from '../index'
import { getWsConnection, send, makeRead } from './wsController'

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            chatRooms: props.appData.chatRooms,
            messages: props.appData.messages,
            unreadCount: props.appData.unreadCount,
            view: props.view,
            roomId: -1
        }
    }

    dispatcherRecieved (resp) {
        const message = JSON.parse(resp.data);
        const roomMessages = [...this.state.messages[message.chat_room_id], message];
        let newMessages = this.state.messages;
        let recieved = 0;
        newMessages[message.chat_room_id] = roomMessages;
        let chatRooms = this.state.chatRooms.map(room => {
            if (room.id == message.chat_room_id) {
                if (message.author.id != this.props.user.id) {
                    recieved = 1;
                }
                return {
                    id: room.id,
                    lastMessage: message,
                    lastMessageId: message.id,
                    status: room.status
                }
            }
            return room
        })
        this.handleUnreadMessages(this.state.ws, this.state.roomId, newMessages[this.state.roomId])
        this.setState({
            chatRooms: chatRooms,
            messages: newMessages,
            unreadCount: this.state.unreadCount + recieved,
            view: this.state.view,
            roomId: this.state.roomId,
            ws: this.state.ws,
        });
    }

    dispatcherReaded(resp) {
        const data = JSON.parse(resp.data);
        const ids = data.messageIds;
        let readed = 0;
        newMessages = this.state.messages;
        newMessages[data.roomId] = newMessages[data.roomId].map(msg => {
            if (ids.includes(msg.id)) {
                if (msg.author.id != this.props.user.id) {
                    readed = 1;
                }
                msg.status = 'read'
            }
            return msg
        });
        chatRooms = this.state.chatRooms.map(room => {
            if (room.id == data.roomId) {
                room.lastMessage = newMessages[data.roomId][newMessages[data.roomId].length - 1]
            }
            return room
        })
        this.setState({
            chatRooms: chatRooms,
            messages: newMessages,
            unreadCount: this.state.unreadCount - readed,
            view: this.state.view,
            roomId: this.state.roomId,
            ws: this.state.ws,
        });
    }

    wsMessageDispatch = resp => {
        switch (resp.type) {
            case 'MESSEGE_RECIEVED':
                this.dispatcherRecieved(resp)
                break
            case 'TYPING':
                break
            case 'MESSEGE_READED':
                this.dispatcherReaded(resp)
                break
        }
    }

    isLastMessageUnread (msg, user) {
        return msg.status == 'unread' && msg.author.id != user.id;
    }

    handleUnreadMessages (ws, roomId, messages) {
        if (!messages) {
            return
        }
        if (this.isLastMessageUnread(messages[messages.length - 1], this.props.user)) {
            const ids = messages
            .filter(msg => {
                return msg.status == 'unread'
            })
            .map(msg => {
                return msg.id
            });
            makeRead(ws, roomId, ids)
        }
    }

    clickChatRoom = id => {
        const ws = getWsConnection(id, this.wsMessageDispatch);
        this.handleUnreadMessages(ws, id, this.state.messages[id])
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            unreadCount: this.state.unreadCount,
            view: VIEW_CHAT_ROOM,
            roomId: id,
            ws: ws,
        });
    }

    clickMyMessages = e => {
        e.preventDefault();
        if (this.state.ws) {
            this.state.ws.close(1000) // normal close code
        }
        this.setState({
            chatRooms: this.state.chatRooms,
            messages: this.state.messages,
            unreadCount: this.state.unreadCount,
            view: VIEW_CHAT_ROOMS,
            roomId: -1,
            ws: undefined,
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
        return <div className='col-12'>
        <Menu user={this.props.user} clickMyMessages={this.clickMyMessages} unreadCount={this.state.unreadCount} />
        <div className="row">
            <div className="col-4">
                <Avatar user={this.props.user}/>
            </div>
            <div className="col-8">
                {this.renderAppBlock()}
            </div>
        </div>
    </div>
        
    }
}