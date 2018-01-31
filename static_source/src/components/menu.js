import ReactDOM from 'react-dom';
import React from 'react';

const Menu = (props) => {
    return <div className='row'>
        <nav className="navbar navbar-expand-md navbar-light bg-faded col-12">
            <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse"
                    data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="true"
                    aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <a className="navbar-brand" href={`/user/${props.user.id}/`}>{props.user.username}</a>
            <div className="collapse navbar-collapse" id="navbarSupportedContent">
                <ul className="navbar-nav mr-auto">
                    <li className="nav-item">
                        <a className="nav-link" href="/user/members/">Members</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" href="/user/friends/">Friends</a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link" onClick={props.clickMyMessages}>My messages <span className="badge badge-primary">{props.unreadCount}</span></a>
                    </li>
                </ul>
            </div>
            <a className="nav-link" href="/logout/">Logout</a>
        </nav>
    </div>
}

export {Menu}