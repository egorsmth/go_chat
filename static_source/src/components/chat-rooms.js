import ReactDOM from 'react-dom';
import React from 'react';
import ChatRoomsBlock from './chat-room-block';

export default class ChatRooms extends React.Component {
    clickChatRoom = (chatRoomId) => {
        this.props.clickChatRoom(chatRoomId)
    }

    renderChatRooms () {
        if (this.props.chatRooms.length > 0) {
            return this.props.chatRooms.map(chatRoom => {
                return <ChatRoomsBlock key={chatRoom.id} chatRoomData={chatRoom} onClick={this.clickChatRoom}/>
            })
        }
        return <p>No messages yet!</p>
    }

    render () {
        return <div id='chatRooms'>
            {this.renderChatRooms()}
        </div>
    }
}