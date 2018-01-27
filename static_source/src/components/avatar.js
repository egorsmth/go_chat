import ReactDOM from 'react-dom';
import React from 'react';

const Avatar = props => {
    return <img className="img-fluid" src={props.user.avatar} alt=""/>
}

export { Avatar }